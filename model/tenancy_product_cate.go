package model

// TenancyProductCate 商品商户分类关联表
type TenancyProductCate struct {
	TenancyProductID  int  `gorm:"column:tenancy_product_id;type:int" json:"tenancyProductId"`
	TenancyCategoryID uint `gorm:"index:tenancy_category_id;column:tenancy_category_id;type:int;not null" json:"tenancyCategoryId"` // 分类id
	SysTenancyID      uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`                // 商户 id
}
