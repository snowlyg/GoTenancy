package response

import "github.com/snowlyg/go-tenancy/model"

type TenancyProductList struct {
	TenancyResponse
	model.BaseTenancyProduct
	SysTenancyName string        `json:"sysTenancyName"`        // 商户名称
	CateName       string        `json:"cateName"`              // 分类名称
	BrandName      string        `json:"brandName"`             // 商户名称
	ProductCates   []ProductCate `gorm:"-" json:"productCates"` // 商户分类
}

type TenancyProductFicti struct {
	Ficti int32 `json:"ficti"`
}

type TenancyProductDetail struct {
	TenancyResponse
	model.BaseTenancyProduct
	SysTenancyName string        `json:"sysTenancyName"` // 商户名称
	CateName       string        `json:"cateName"`       // 分类名称
	BrandName      string        `json:"brandName"`      // 商户名称
	Content        string        `json:"content"`
	ProductCates   []ProductCate `gorm:"-" json:"productCates"` // 商户分类
}

type TenancyProductFilter struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type TenancyProductCondition struct {
	Type       int                    `json:"type"`
	Name       string                 `json:"name"`
	Conditions map[string]interface{} `json:"conditions"`
	IsDeleted  bool                   `json:"is_deleted"`
}

type ProductCate struct {
	ID       uint   `json:"id"`
	CateName string `json:"cateName"` // 分类名称
}
