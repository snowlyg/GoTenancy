package source

import (
	"github.com/snowlyg/go-tenancy/g"

	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gookit/color"
	"gorm.io/gorm"
)

var Casbin = new(casbin)

type casbin struct{}

var carbines = []gormadapter.CasbinRule{
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/register", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/logout", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/clean", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/createApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/getApiList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/getApiById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/deleteApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/updateApi", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/getAllApis", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/createAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/deleteAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/getAuthorityList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/setDataAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/updateAuthority", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/authority/copyAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/getMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/getMenuList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/addBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/getBaseMenuTree", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/addMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/getMenuAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/deleteBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/updateBaseMenu", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/menu/getBaseMenuById", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/changePassword", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/getAdminList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/getTenancyList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/getGeneralList", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/setUserAuthority", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/deleteUser", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/casbin/updateCasbin", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/casbin/getPolicyPathByAuthorityId", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/casbin/casbinTest/:pathParam", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/system/getSystemConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/system/setSystemConfig", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/system/getServerInfo", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionaryDetail/createSysDictionaryDetail", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionaryDetail/deleteSysDictionaryDetail", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionaryDetail/updateSysDictionaryDetail", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionaryDetail/findSysDictionaryDetail", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionaryDetail/getSysDictionaryDetailList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionary/createSysDictionary", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionary/deleteSysDictionary", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionary/updateSysDictionary", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionary/findSysDictionary", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysDictionary/getSysDictionaryList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/createSysOperationRecord", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/deleteSysOperationRecord", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/updateSysOperationRecord", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/findSysOperationRecord", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/getSysOperationRecordList", V2: "GET"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/sysOperationRecord/deleteSysOperationRecordByIds", V2: "DELETE"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/user/setUserInfo/{user_id:int}", V2: "PUT"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/email/emailTest", V2: "POST"},
	{Ptype: "p", V0: AdminAuthorityId, V1: "/v1/api/deleteApisByIds", V2: "DELETE"},
}

//Init casbin_rule 表数据初始化
func (c *casbin) Init() error {
	g.TENANCY_DB.AutoMigrate(gormadapter.CasbinRule{})
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Find(&[]gormadapter.CasbinRule{}).RowsAffected == 154 {
			color.Danger.Println("\n[Mysql] --> casbin_rule 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&carbines).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> casbin_rule 表初始数据成功!")
		return nil
	})
}
