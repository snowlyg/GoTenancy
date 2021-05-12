package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"

	"gorm.io/gorm"
)

var BaseMenu = new(menu)

type menu struct{}

var menus = []model.SysBaseMenu{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "dashboard", Name: "dashboard", Hidden: false, Component: "view/dashboard/index.vue", Sort: 1, Meta: model.Meta{Title: "仪表盘", Icon: "setting"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "超级管理员", Icon: "user-solid"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "adminUser", Name: "adminUser", Component: "view/superAdmin/user/admin.vue", Sort: 4, Meta: model.Meta{Title: "员工管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "tenancyUser", Name: "tenancyUser", Component: "view/superAdmin/user/tenancy.vue", Sort: 5, Meta: model.Meta{Title: "商户用户管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "generalUser", Name: "generalUser", Component: "view/superAdmin/user/general.vue", Sort: 6, Meta: model.Meta{Title: "普通用户管理", Icon: "coordinate"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 7, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 17, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "14", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 3, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 18, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "dictionary", Name: "dictionary", Component: "view/superAdmin/dictionary/sysDictionary.vue", Sort: 5, Meta: model.Meta{Title: "字典管理", Icon: "notebook-2"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 19, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "3", Path: "dictionaryDetail/:id", Name: "dictionaryDetail", Component: "view/superAdmin/dictionary/sysDictionaryDetail.vue", Sort: 1, Meta: model.Meta{Title: "字典详情", Icon: "s-order"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 20, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "3", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 6, Meta: model.Meta{Title: "操作历史", Icon: "time"}},

	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "0", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},
}

//Init sys_base_menus 表数据初始化
func (m *menu) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 29}).Find(&[]model.SysBaseMenu{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_base_menus 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&menus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_base_menus 表初始数据成功!")
		return nil
	})
}
