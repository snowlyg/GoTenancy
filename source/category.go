package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Category = new(category)

type category struct{}

var categories = []model.TenancyCategory{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CateName: "计生保健", Pid: "0", Path: "", Sort: 0, Level: 0, IsShow: true, SysTenancyID: 1, Pic: "http://qmplusimg.henrongyi.top/head.png"},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CateName: "大保健", Pid: "1", Path: "", Sort: 1, Level: 1, IsShow: true, SysTenancyID: 1, Pic: "http://qmplusimg.henrongyi.top/head.png"},
}

func (m *category) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.TenancyCategory{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> categories 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&categories).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> categories 表初始数据成功!")
		return nil
	})
}
