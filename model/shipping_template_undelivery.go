package model

import "github.com/snowlyg/go-tenancy/g"

// ShippingTemplateUndelivery 指定不配送区域表
type ShippingTemplateUndelivery struct {
	g.TENANCY_MODEL
	Code int `json:"code" gorm:""`

	ShippingTemplateID uint `gorm:"index:shipping_template_id;column:shipping_template_id;type:int unsigned;not null;default:0" json:"shippingTemplateId"` // 模板ID
}
