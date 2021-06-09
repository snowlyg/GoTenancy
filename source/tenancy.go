package source

import (
	"time"

	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Tenancy = new(tenancy)

type tenancy struct{}

var tenancies = []model.SysTenancy{
	{BaseTenancy: model.BaseTenancy{UUID: uuid.NewV4(), Name: "宝安中心人民医院", Tele: "0755-23568911", Address: "xxx街道888号", BusinessTime: "08:30-17:30", Status: g.StatusTrue, SysRegionCode: 1}},
	{TENANCY_MODEL: g.TENANCY_MODEL{CreatedAt: time.Now().AddDate(0, 0, -1)}, BaseTenancy: model.BaseTenancy{UUID: uuid.NewV4(), Name: "广东东莞第三人民医院", Tele: "0755-23568922", Address: "xxx街道777号", BusinessTime: "08:35-17:35", Status: g.StatusTrue, SysRegionCode: 2}},
	{TENANCY_MODEL: g.TENANCY_MODEL{CreatedAt: time.Now().AddDate(0, 0, -2)}, BaseTenancy: model.BaseTenancy{UUID: uuid.NewV4(), Name: "湖南长沙湘雅医院", Tele: "0755-23568922", Address: "xxx街道777号", BusinessTime: "08:35-17:35", Status: g.StatusTrue, SysRegionCode: 3}},
	{TENANCY_MODEL: g.TENANCY_MODEL{CreatedAt: time.Now().AddDate(0, 0, -15)}, BaseTenancy: model.BaseTenancy{UUID: uuid.NewV4(), Name: "湖北武汉第三人民医院", Tele: "0755-23568922", Address: "xxx街道777号", BusinessTime: "08:35-17:35", Status: g.StatusTrue, SysRegionCode: 4}},
	{BaseTenancy: model.BaseTenancy{UUID: uuid.NewV4(), Name: "上海第三人民医院", Tele: "0755-23568922", Address: "xxx街道777号", BusinessTime: "08:35-17:35", Status: g.StatusFalse, SysRegionCode: 5}},
}

//@description: sys_tenancies 表数据初始化
func (a *tenancy) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysTenancy{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_tenancies 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&tenancies).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_tenancies 表初始数据成功!")
		return nil
	})
}
