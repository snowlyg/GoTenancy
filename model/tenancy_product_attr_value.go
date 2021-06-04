package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyProductAttrValue 商品属性值表
type TenancyProductAttrValue struct {
	g.TENANCY_MODEL

	Detail       string  `gorm:"column:detail;type:varchar(1000);not null;default:''" json:"detail"`
	Sku          string  `gorm:"index:sku;column:sku;type:varchar(128);not null" json:"sku"`              // 商品属性索引值 (attr_value|attr_value[|....])
	Stock        uint    `gorm:"column:stock;type:int unsigned;not null" json:"stock"`                    // 属性对应的库存
	Sales        uint    `gorm:"column:sales;type:int unsigned;not null;default:0" json:"sales"`          // 销量
	Image        string  `gorm:"column:image;type:varchar(128)" json:"image"`                             // 图片
	BarCode      string  `gorm:"column:bar_code;type:varchar(50);not null;default:''" json:"barCode"`     // 产品条码
	Cost         float64 `gorm:"column:cost;type:decimal(8,2) unsigned;not null" json:"cost"`             // 成本价
	OtPrice      float64 `gorm:"column:ot_price;type:decimal(8,2);not null;default:0.00" json:"otPrice"`  // 原价
	Price        float64 `gorm:"column:price;type:decimal(8,2) unsigned;not null" json:"price"`           // 价格
	Volume       float64 `gorm:"column:volume;type:decimal(8,2);not null;default:0.00" json:"volume"`     // 体积
	Weight       float64 `gorm:"column:weight;type:decimal(8,2);not null;default:0.00" json:"weight"`     // 重量
	Type         int     `gorm:"column:type;type:tinyint(1);default:1" json:"type"`                       // 活动类型 1=商品
	ExtensionOne float64 `gorm:"column:extension_one;type:decimal(8,2);default:0.00" json:"extensionOne"` // 一级佣金
	ExtensionTwo float64 `gorm:"column:extension_two;type:decimal(8,2);default:0.00" json:"extensionTwo"` // 二级佣金
	Unique       string  `gorm:"index;column:unique;type:char(12);not null;default:''" json:"unique"`     // 唯一值

	TenancyProductID int `gorm:"index:tenancy_product_id;column:tenancy_product_id;type:int;not null" json:"tenancyProductId"` // 商品id
}
