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

	// 系统管理
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "admin", Name: "superAdmin", Component: "view/superAdmin/index.vue", Sort: 3, Meta: model.Meta{Title: "系统管理", Icon: "user-solid"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "tenancy", Name: "tenancy", Component: "view/superAdmin/tenancy/index.vue", Sort: 7, Meta: model.Meta{Title: "商户管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "brand", Name: "brand", Component: "view/superAdmin/brand/index.vue", Sort: 9, Meta: model.Meta{Title: "品牌管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "brandCategory", Name: "brandCategory", Component: "view/superAdmin/brandCategory/index.vue", Sort: 10, Meta: model.Meta{Title: "品牌分类管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "authority", Name: "authority", Component: "view/superAdmin/authority/authority.vue", Sort: 1, Meta: model.Meta{Title: "角色管理", Icon: "s-custom"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "menu", Name: "menu", Component: "view/superAdmin/menu/menu.vue", Sort: 2, Meta: model.Meta{Title: "菜单管理", Icon: "s-order", KeepAlive: true}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "api", Name: "api", Component: "view/superAdmin/api/api.vue", Sort: 3, Meta: model.Meta{Title: "api管理", Icon: "s-platform", KeepAlive: true}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "adminUser", Name: "adminUser", Component: "view/superAdmin/user/admin.vue", Sort: 4, Meta: model.Meta{Title: "员工管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "config", Name: "config", Component: "view/superAdmin/config/index.vue", Sort: 15, Meta: model.Meta{Title: "系统配置管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "mini", Name: "mini", Component: "view/superAdmin/mini/index.vue", Sort: 8, Meta: model.Meta{Title: "小程序管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "2", Path: "operation", Name: "operation", Component: "view/superAdmin/operation/sysOperationRecord.vue", Sort: 20, Meta: model.Meta{Title: "操作历史", Icon: "time"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, ParentId: "2", Path: "state", Name: "state", Hidden: false, Component: "view/system/state.vue", Sort: 6, Meta: model.Meta{Title: "服务器状态", Icon: "cloudy"}},

	// 商铺管理
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 21, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "tenancy", Name: "tenancy", Component: "view/tenancy/index.vue", Sort: 7, Meta: model.Meta{Title: "商铺管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 22, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "21", Path: "media", Name: "media", Component: "view/tenancy/media.vue", Sort: 11, Meta: model.Meta{Title: "媒体库", Icon: "upload"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 23, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "21", Path: "category", Name: "category", Component: "view/tenancy/category.vue", Sort: 12, Meta: model.Meta{Title: "分类管理", Icon: "cloudy"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 24, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "21", Path: "product", Name: "product", Component: "view/tenancy/product/index.vue", Sort: 13, Meta: model.Meta{Title: "商品管理", Icon: "cloudy"}},

	// 用户管理
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 41, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "user", Name: "user", Component: "view/user/index.vue", Sort: 3, Meta: model.Meta{Title: "用户管理", Icon: "user-solid"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 42, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "41", Path: "tenancyUser", Name: "tenancyUser", Component: "view/user/tenancy.vue", Sort: 5, Meta: model.Meta{Title: "商户用户管理", Icon: "coordinate"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 43, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "41", Path: "generalUser", Name: "generalUser", Component: "view/user/general.vue", Sort: 6, Meta: model.Meta{Title: "普通用户管理", Icon: "coordinate"}},

	// 系统工具
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 51, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "0", Path: "systemTools", Name: "systemTools", Component: "view/systemTools/index.vue", Sort: 5, Meta: model.Meta{Title: "系统工具", Icon: "s-cooperation"}},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 52, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: false, ParentId: "51", Path: "system", Name: "system", Component: "view/systemTools/system/system.vue", Sort: 17, Meta: model.Meta{Title: "系统配置", Icon: "s-operation"}},

	// 个人信息
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 81, CreatedAt: time.Now(), UpdatedAt: time.Now()}, MenuLevel: 0, Hidden: true, ParentId: "0", Path: "person", Name: "person", Component: "view/person/person.vue", Sort: 16, Meta: model.Meta{Title: "个人信息", Icon: "message-solid"}},
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
