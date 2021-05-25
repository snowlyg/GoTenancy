package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyShippingTemplateRegion 配送区域表
type TenancyShippingTemplateRegion struct {
	g.TENANCY_MODEL
	First         float64 `gorm:"column:first;type:decimal(10,2) unsigned;not null;default:0.00" json:"first"`                  // 首件
	FirstPrice    float64 `gorm:"column:first_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"firstPrice"`       // 首件运费
	Continue      float64 `gorm:"column:continue;type:decimal(10,2) unsigned;not null;default:0.00" json:"continue"`            // 续件
	ContinuePrice float64 `gorm:"column:continue_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"continuePrice"` // 续件运费
	TempID        uint    `gorm:"index:temp_id;column:temp_id;type:int unsigned;not null;default:0" json:"tempId"`              // 模板ID
	CityID        string  `gorm:"column:city_id;type:text;not null" json:"cityId"`                                              // 城市ID /id/id/id/
}
