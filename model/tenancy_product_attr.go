package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyProductAttr 商品属性表
type TenancyProductAttr struct {
	g.TENANCY_MODEL

	AttrName         string `gorm:"column:attr_name;type:varchar(32);not null" json:"attrName"`                                   // 属性名
	AttrValues       string `gorm:"column:attr_values;type:varchar(2000);not null" json:"attrValues"`                             // 属性值
	Type             int    `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                                            // 活动类型 1=商品
	TenancyProductID int    `gorm:"index:tenancy_product_id;column:tenancy_product_id;type:int;not null" json:"tenancyProductId"` // 商品id
}
