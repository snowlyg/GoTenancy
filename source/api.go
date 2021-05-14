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
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/register", Description: "用户注册", ApiGroup: "user", Method: "POST"},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/createApi", Description: "创建api", ApiGroup: "api", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/getApiList", Description: "获取api列表", ApiGroup: "api", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/getApiById", Description: "获取api详细信息", ApiGroup: "api", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/deleteApi", Description: "删除Api", ApiGroup: "api", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/updateApi", Description: "更新Api", ApiGroup: "api", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/getAllApis", Description: "获取所有api", ApiGroup: "api", Method: "POST"},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/createAuthority", Description: "创建角色", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/deleteAuthority", Description: "删除角色", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/getAuthorityList", Description: "获取角色列表", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/getMenu", Description: "获取菜单树", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/getMenuList", Description: "分页获取基础menu列表", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/addBaseMenu", Description: "新增菜单", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/getBaseMenuTree", Description: "获取用户动态路由", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 16, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/addMenuAuthority", Description: "增加menu和角色关联关系", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/getMenuAuthority", Description: "获取指定角色menu", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/deleteBaseMenu", Description: "删除菜单", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/updateBaseMenu", Description: "更新菜单", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/menu/getBaseMenuById", Description: "根据id获取菜单", ApiGroup: "menu", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/changePassword", Description: "修改密码", ApiGroup: "user", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/getAdminList", Description: "获取管理员列表", ApiGroup: "user", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/getTenancyList", Description: "获取商户员工列表", ApiGroup: "user", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 25, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/getGeneralList", Description: "获取普通用户列表", ApiGroup: "user", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 26, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/setUserAuthority", Description: "修改用户角色", ApiGroup: "user", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 27, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/logout", Description: "退出", ApiGroup: "user", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 28, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/clean", Description: "清空", ApiGroup: "user", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 29, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/casbin/updateCasbin", Description: "更改角色api权限", ApiGroup: "casbin", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 30, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/casbin/getPolicyPathByAuthorityId", Description: "获取权限列表", ApiGroup: "casbin", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 31, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/setDataAuthority", Description: "设置角色资源权限", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 32, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/system/getSystemConfig", Description: "获取配置文件内容", ApiGroup: "system", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 33, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/system/setSystemConfig", Description: "设置配置文件内容", ApiGroup: "system", Method: "POST"},

	// 商户
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 34, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/getTenancies/{code:int}", Description: "根据地区获取商户", ApiGroup: "tenancy", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 35, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/getTenancyList", Description: "获取商户列表", ApiGroup: "tenancy", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 36, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/createTenancy", Description: "添加商户", ApiGroup: "tenancy", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 37, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/setTenancyRegion", Description: "设置商户地区", ApiGroup: "tenancy", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 38, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/getTenancyById", Description: "获取商户详细信息", ApiGroup: "tenancy", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 39, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/updateTenancy", Description: "更新商户", ApiGroup: "tenancy", Method: "PUT"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 40, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/tenancy/deleteTenancy", Description: "删除商户", ApiGroup: "tenancy", Method: "DELETE"},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/casbin/casbinTest/:pathParam", Description: "RESTFUL模式测试", ApiGroup: "casbin", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/updateAuthority", Description: "更新角色信息", ApiGroup: "authority", Method: "PUT"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/copyAuthority", Description: "拷贝角色", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 44, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/deleteUser", Description: "删除用户", ApiGroup: "user", Method: "DELETE"},

	// 小程序
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 45, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/mini/getMiniList", Description: "获取小程序列表", ApiGroup: "mini", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 46, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/mini/createMini", Description: "添加小程序", ApiGroup: "mini", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 47, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/mini/getMiniById", Description: "获取小程序详细信息", ApiGroup: "mini", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 48, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/mini/updateMini", Description: "更新小程序", ApiGroup: "mini", Method: "PUT"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 49, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/mini/deleteMini", Description: "删除小程序", ApiGroup: "mini", Method: "DELETE"},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 50, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/getAdminAuthorityList", Description: "获取员工角色列表", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/getTenancyAuthorityList", Description: "获取商户角色列表", ApiGroup: "authority", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/authority/getGeneralAuthorityList", Description: "获取普通用户角色列表", ApiGroup: "authority", Method: "POST"},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 54, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/sysOperationRecord/createSysOperationRecord", Description: "新增操作记录", ApiGroup: "sysOperationRecord", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 55, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/sysOperationRecord/deleteSysOperationRecord", Description: "删除操作记录", ApiGroup: "sysOperationRecord", Method: "DELETE"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 56, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/sysOperationRecord/findSysOperationRecord", Description: "根据ID获取操作记录", ApiGroup: "sysOperationRecord", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 57, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/sysOperationRecord/getSysOperationRecordList", Description: "获取操作记录列表", ApiGroup: "sysOperationRecord", Method: "GET"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 61, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/sysOperationRecord/deleteSysOperationRecordByIds", Description: "批量删除操作历史", ApiGroup: "sysOperationRecord", Method: "DELETE"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 65, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/user/setUserInfo", Description: "设置用户信息", ApiGroup: "user", Method: "PUT"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 66, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/system/getServerInfo", Description: "获取服务器信息", ApiGroup: "system", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 67, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/email/emailTest", Description: "发送测试邮件", ApiGroup: "email", Method: "POST"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 85, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Path: "/v1/admin/api/deleteApisByIds", Description: "批量删除api", ApiGroup: "api", Method: "DELETE"},
}

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
