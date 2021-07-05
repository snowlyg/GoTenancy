package model

// ProductContent 商品详情表
type ProductContent struct {
	Content string `gorm:"column:content;type:longtext;not null" json:"content"`       // 商品详情
	Type    int32  `gorm:"column:type;type:tinyint(1);not null;default:1" json:"type"` // 商品类型 1=普通

	ProductID uint `gorm:"product_contents:product_id;column:product_id;type:int;not null" json:"productId"` // 商品id
}
