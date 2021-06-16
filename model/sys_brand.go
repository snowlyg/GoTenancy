package model

import "github.com/snowlyg/go-tenancy/g"

type SysBrand struct {
	g.TENANCY_MODEL
	BrandName string `gorm:"column:brand_name;type:varchar(100);not null" json:"brandName" binding:"required"` // 品牌名称
	Sort      int32  `gorm:"column:sort;type:mediumint;not null" json:"sort"`                                  // 排序
	Pic       string `gorm:"column:pic;type:varchar(128);not null;default:''" json:"pic"`                      // 图标
	Status    int    `gorm:"column:status;type:tinyint(1);not null" json:"status"`                             // 是否显示

	BrandCategoryID uint `gorm:"column:brand_category_id;type:mediumint;not null" json:"brandCategoryId"` // 分类id
}
