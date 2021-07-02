package model

import "github.com/snowlyg/go-tenancy/g"

// ShippingTemplateRegion 配送区域表
type ShippingTemplateRegion struct {
	g.TENANCY_MODEL
	First         float64 `gorm:"column:first;type:decimal(10,2) unsigned;not null;default:0.00" json:"first"`                  // 首件
	FirstPrice    float64 `gorm:"column:first_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"firstPrice"`       // 首件运费
	Continue      float64 `gorm:"column:continue;type:decimal(10,2) unsigned;not null;default:0.00" json:"continue"`            // 续件
	ContinuePrice float64 `gorm:"column:continue_price;type:decimal(10,2) unsigned;not null;default:0.00" json:"continuePrice"` // 续件运费
	Code          int     `json:"code" gorm:""`

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}
