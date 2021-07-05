package response

import (
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

type ProductList struct {
	TenancyResponse
	model.BaseProduct
	SysTenancyName string        `json:"sysTenancyName"`        // 商户名称
	CateName       string        `json:"cateName"`              // 分类名称
	BrandName      string        `json:"brandName"`             // 商户名称
	ProductCates   []ProductCate `gorm:"-" json:"productCates"` // 商户分类

}

type ProductFicti struct {
	Ficti int32 `json:"ficti"`
}

type ProductDetail struct {
	TenancyResponse
	model.BaseProduct
	SysTenancyName string                     `json:"sysTenancyName"` // 商户名称
	CateName       string                     `json:"cateName"`       // 分类名称
	BrandName      string                     `json:"brandName"`      // 商户名称
	TempName       string                     `json:"tempName"`       // 运费模板名称
	Content        string                     `json:"content"`
	SliderImage    string                     `json:"sliderImage"`           // 轮播图
	SliderImages   []string                   `gorm:"-" json:"sliderImages"` // 轮播图
	Attr           []request.Value            `gorm:"-" json:"attr"`
	AttrValue      []request.ProductAttrValue `gorm:"-" json:"attrValue"`
	CateId         uint                       `gorm:"-" json:"cateId"`            // 平台分类id
	CategoryIds    []uint                     `gorm:"-" json:"tenancyCategoryId"` // 商户分类
	ProductCates   []ProductCate              `gorm:"-" json:"productCates"`      // 商户分类
}

type ProductFilter struct {
	Type  int    `json:"type"`
	Name  string `json:"name"`
	Count int64  `json:"count"`
}

type ProductCondition struct {
	Type       int                    `json:"type"`
	Name       string                 `json:"name"`
	Conditions map[string]interface{} `json:"conditions"`
	IsDeleted  bool                   `json:"is_deleted"`
}

type ProductCate struct {
	ID       uint   `json:"id"`
	CateName string `json:"cateName"` // 分类名称
}
