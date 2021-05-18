package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Address = new(address)

type address struct{}

var addresses = []model.TenancyAddress{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, Name: "余思琳", Sex: model.Female, Country: "中国", Province: "广东省", City: "东莞市", IsDefault: 1, DetailAddress: "松山湖阿里产业园", SysUserID: 3},
}

func (m *address) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.TenancyAddress{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_addresses 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&addresses).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_addresses 表初始数据成功!")
		return nil
	})
}
