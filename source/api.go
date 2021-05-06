package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var Api = new(api)

type api struct{}

var apis = []model.SysApi{
	{g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/base/login", "用户登录", "base", "POST"},
	{g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/register", "用户注册", "user", "POST"},
	{g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/createApi", "创建api", "api", "POST"},
	{g.TENANCY_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiList", "获取api列表", "api", "POST"},
	{g.TENANCY_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getApiById", "获取api详细信息", "api", "POST"},
	{g.TENANCY_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApi", "删除Api", "api", "POST"},
	{g.TENANCY_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/updateApi", "更新Api", "api", "POST"},
	{g.TENANCY_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/getAllApis", "获取所有api", "api", "POST"},
	{g.TENANCY_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/createAuthority", "创建角色", "authority", "POST"},
	{g.TENANCY_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/deleteAuthority", "删除角色", "authority", "POST"},
	{g.TENANCY_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/getAuthorityList", "获取角色列表", "authority", "POST"},
	{g.TENANCY_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenu", "获取菜单树", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuList", "分页获取基础menu列表", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addBaseMenu", "新增菜单", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuTree", "获取用户动态路由", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/addMenuAuthority", "增加menu和角色关联关系", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getMenuAuthority", "获取指定角色menu", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/deleteBaseMenu", "删除菜单", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/updateBaseMenu", "更新菜单", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/menu/getBaseMenuById", "根据id获取菜单", "menu", "POST"},
	{g.TENANCY_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/changePassword", "修改密码", "user", "POST"},
	{g.TENANCY_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/getUserList", "获取用户列表", "user", "POST"},
	{g.TENANCY_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserAuthority", "修改用户角色", "user", "POST"},
	{g.TENANCY_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/logout", "退出", "user", "GET"},
	{g.TENANCY_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/clean", "清空", "user", "GET"},
	{g.TENANCY_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/updateCasbin", "更改角色api权限", "casbin", "POST"},
	{g.TENANCY_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/getPolicyPathByAuthorityId", "获取权限列表", "casbin", "POST"},
	{g.TENANCY_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/setDataAuthority", "设置角色资源权限", "authority", "POST"},
	{g.TENANCY_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getSystemConfig", "获取配置文件内容", "system", "POST"},
	{g.TENANCY_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/setSystemConfig", "设置配置文件内容", "system", "POST"},
	{g.TENANCY_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "创建客户", "customer", "POST"},
	{g.TENANCY_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "更新客户", "customer", "PUT"},
	{g.TENANCY_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "删除客户", "customer", "DELETE"},
	{g.TENANCY_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customer", "获取单一客户", "customer", "GET"},
	{g.TENANCY_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/customer/customerList", "获取客户列表", "customer", "GET"},
	{g.TENANCY_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/casbin/casbinTest/:pathParam", "RESTFUL模式测试", "casbin", "GET"},
	{g.TENANCY_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/updateAuthority", "更新角色信息", "authority", "PUT"},
	{g.TENANCY_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/authority/copyAuthority", "拷贝角色", "authority", "POST"},
	{g.TENANCY_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/deleteUser", "删除用户", "user", "DELETE"},
	{g.TENANCY_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/createSysOperationRecord", "新增操作记录", "sysOperationRecord", "POST"},
	{g.TENANCY_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecord", "删除操作记录", "sysOperationRecord", "DELETE"},
	{g.TENANCY_MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/findSysOperationRecord", "根据ID获取操作记录", "sysOperationRecord", "GET"},
	{g.TENANCY_MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/getSysOperationRecordList", "获取操作记录列表", "sysOperationRecord", "GET"},
	{g.TENANCY_MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/sysOperationRecord/deleteSysOperationRecordByIds", "批量删除操作历史", "sysOperationRecord", "DELETE"},
	{g.TENANCY_MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/user/setUserInfo", "设置用户信息", "user", "PUT"},
	{g.TENANCY_MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/system/getServerInfo", "获取服务器信息", "system", "POST"},
	{g.TENANCY_MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/email/emailTest", "发送测试邮件", "email", "POST"},
	{g.TENANCY_MODEL{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, "/api/deleteApisByIds", "批量删除api", "api", "DELETE"},
}

//@author: [SliverHorn](https://github.com/SliverHorn)
//@description: sys_apis 表数据初始化
func (a *api) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 67}).Find(&[]model.SysApi{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_apis 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&apis).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_apis 表初始数据成功!")
		return nil
	})
}
