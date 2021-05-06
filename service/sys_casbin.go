package service

import (
	"errors"
	"fmt"
	"strings"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	_ "github.com/go-sql-driver/mysql"
)

// UpdateCasbin 更新casbin权限
func UpdateCasbin(authorityId string, casbinInfos []request.CasbinInfo) error {
	ClearCasbin(0, authorityId)
	rules := [][]string{}
	for _, v := range casbinInfos {
		cm := model.CasbinModel{
			Ptype:       "p",
			AuthorityId: authorityId,
			Path:        v.Path,
			Method:      v.Method,
		}
		rules = append(rules, []string{cm.AuthorityId, cm.Path, cm.Method})
	}
	e, _ := Casbin()
	success, _ := e.AddPolicies(rules)
	if success == false {
		return errors.New("存在相同api,添加失败,请联系管理员")
	}
	return nil
}

// UpdateCasbinApi API更新随动
func UpdateCasbinApi(oldPath string, newPath string, oldMethod string, newMethod string) error {
	err := g.TENANCY_DB.Table("casbin_rule").Model(&model.CasbinModel{}).Where("v1 = ? AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	return err
}

// GetPolicyPathByAuthorityId 获取权限列表
func GetPolicyPathByAuthorityId(authorityId string) (pathMaps []request.CasbinInfo) {
	e, _ := Casbin()
	list := e.GetFilteredPolicy(0, authorityId)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

//ClearCasbin 清除匹配的权限
func ClearCasbin(v int, p ...string) bool {
	e, _ := Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

// Casbin 持久化到数据库  引入自定义规则
func Casbin() (*casbin.Enforcer, error) {
	a, err := gormadapter.NewAdapterByDB(g.TENANCY_DB)
	if err != nil {
		return nil, fmt.Errorf("casbin new adapter by db err %w", err)
	}
	e, err := casbin.NewEnforcer(g.TENANCY_CONFIG.Casbin.ModelPath, a)
	if err != nil {
		return nil, fmt.Errorf("casbin new enforcer err %w", err)
	}
	e.AddFunction("ParamsMatch", ParamsMatchFunc)
	err = e.LoadPolicy()
	if err != nil {
		return nil, fmt.Errorf("casbin load policy err %w", err)
	}
	return e, nil
}

//ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, key2)
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)

	return ParamsMatch(name1, name2), nil
}
