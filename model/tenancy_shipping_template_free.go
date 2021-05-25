package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyShippingTemplateFree 指定包邮信息表
type TenancyShippingTemplateFree struct {
	g.TENANCY_MODEL
	Number uint    `gorm:"column:number;type:int unsigned;not null;default:0" json:"number"`                // 包邮件数
	Price  float64 `gorm:"column:price;type:decimal(10,2) unsigned;not null;default:0.00" json:"price"`     // 包邮金额
	TempID uint    `gorm:"index:temp_id;column:temp_id;type:int unsigned;not null;default:0" json:"tempId"` // 模板ID
	CityID string  `gorm:"column:city_id;type:text;not null" json:"cityId"`                                 // 城市ID
}
