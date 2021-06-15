package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Category = new(category)

type category struct{}

var categories = []model.TenancyCategory{
	{BaseTenancyCategory: model.BaseTenancyCategory{CateName: "计生保健", Pid: 0, Path: "", Sort: 0, Level: 0, Status: g.StatusTrue, Pic: "http://qmplusimg.henrongyi.top/head.png"}, SysTenancyID: 1},
	{BaseTenancyCategory: model.BaseTenancyCategory{CateName: "大保健", Pid: 1, Path: "", Sort: 1, Level: 1, Status: g.StatusTrue, Pic: "http://qmplusimg.henrongyi.top/head.png"}, SysTenancyID: 1},
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
