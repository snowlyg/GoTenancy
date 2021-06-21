package model

// ProductCate 商品商户分类关联表
type ProductCate struct {
	ProductID         int  `gorm:"column:product_id;type:int" json:"productId"`
	ProductCategoryID uint `gorm:"index:product_category_id;column:product_category_id;type:int;not null" json:"productCategoryId"` // 分类id
	SysTenancyID      uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`                // 商户 id
}
