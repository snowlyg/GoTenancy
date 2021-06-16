package model

// TenancyProductContent 商品详情表
type TenancyProductContent struct {
	Content string `gorm:"column:content;type:longtext;not null" json:"content"`       // 商品详情
	Type    int32  `gorm:"column:type;type:tinyint(1);not null;default:1" json:"type"` // 商品类型 1=普通

	TenancyProductID uint `gorm:"index:tenancy_product_id;column:tenancy_product_id;type:int;not null" json:"tenancyProductId"` // 商品id
}
