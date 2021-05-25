package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Brand = new(brand)

type brand struct{}

var brands = []model.SysBrand{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, BrandName: "杜蕾斯", Sort: 0, Pic: "http://qmplusimg.henrongyi.top/head.png", IsShow: true, BrandCategoryID: 1},
}

func (m *brand) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysBrand{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_brands 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&brands).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_brands 表初始数据成功!")
		return nil
	})
}
