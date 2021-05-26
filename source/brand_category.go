package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var BrandCategory = new(brand_category)

type brand_category struct{}

var brand_categories = []model.SysBrandCategory{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CateName: "计生保健", Pid: "0", Path: "", Sort: 0, Level: 0, IsShow: true},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()}, CateName: "大保健", Pid: "1", Path: "", Sort: 1, Level: 1, IsShow: true},
}

func (m *brand_category) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysBrandCategory{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> sys_brand_categories 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&brand_categories).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_brand_categories 表初始数据成功!")
		return nil
	})
}
