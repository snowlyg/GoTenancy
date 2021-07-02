package model

import "github.com/snowlyg/go-tenancy/g"

// ShippingTemplateFree 指定包邮信息表
type ShippingTemplateFree struct {
	g.TENANCY_MODEL
	Number uint    `gorm:"column:number;type:int unsigned;not null;default:0" json:"number"`            // 包邮件数
	Price  float64 `gorm:"column:price;type:decimal(10,2) unsigned;not null;default:0.00" json:"price"` // 包邮金额
	Code   int     `json:"code" gorm:""`                                                                // 城市ID

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}
