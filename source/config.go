package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Config = new(config)

type config struct{}

var configs = []model.SysConfig{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "token", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "appid", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "appsecret", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "encodingaeskey", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 5, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_id", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 6, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_key", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 7, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_ssl_type", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 8, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_ssl_p12", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 9, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_ssl_key", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "wechat", Name: "mch_ssl_cer", Value: "", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 11, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "system", Name: "sys_login_logo", Value: "http:\\\\mer.crmeb.net\\uploads\\def\\20200816\\6c50374375d5fd6b2a8c40d49884daf6.png", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 12, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "system", Name: "sys_menu_logo", Value: "http:\\\\mer.crmeb.net\\uploads\\def\\20200816\\6c50374375d5fd6b2a8c40d49884daf6.png", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 13, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "system", Name: "sys_menu_slogo", Value: "http:\\\\mer.crmeb.net\\uploads\\def\\20200816\\e11a0f712ca67edff15e900858d690fa.png", SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 14, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Type: "system", Name: "sys_login_title", Value: "", SysTenancyID: 0},
}

func (m *config) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysConfig{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_configs 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configs).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_configs 表初始数据成功!")
		return nil
	})
}
