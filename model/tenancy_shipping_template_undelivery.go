package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyShippingTemplateUndelivery 指定不配送区域表
type TenancyShippingTemplateUndelivery struct {
	g.TENANCY_MODEL
	TempID uint   `gorm:"index:temp_id;column:temp_id;type:int unsigned;not null;default:0" json:"tempId"` // 模板ID
	CityID string `gorm:"column:city_id;type:text;not null" json:"cityId"`                                 // 城市ID /id/id/id/
}
