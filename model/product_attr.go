package model

import "github.com/snowlyg/go-tenancy/g"

// ProductAttr 商品属性表
type ProductAttr struct {
	g.TENANCY_MODEL

	AttrName   string `gorm:"column:attr_name;type:varchar(32);not null" json:"attrName"`       // 属性名
	AttrValues string `gorm:"column:attr_values;type:varchar(2000);not null" json:"attrValues"` // 属性值
	Type       int    `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                // 活动类型 1=商品

	ProductID uint `gorm:"index:product_id;column:product_id;type:int;not null" json:"productId"` // 商品id // 商品id
}
