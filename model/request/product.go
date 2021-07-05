package request

import "github.com/snowlyg/go-tenancy/model"

type ProductPageInfo struct {
	Page              int    `json:"page" form:"page" binding:"required"`
	PageSize          int    `json:"pageSize" form:"pageSize" binding:"required"`
	ProductCategoryId uint   `json:"tenancyCategoryId" form:"tenancyCategoryId"`
	Type              string `json:"type" form:"type"  binding:"required"`
	CateId            int    `json:"cateId" form:"cateId"`
	Keyword           string `json:"keyword" form:"keyword"`
	IsGiftBag         string `json:"isGiftBag" form:"isGiftBag"`
}

type UpdateProduct struct {
	Id uint `json:"id"`
	model.BaseProduct
	Content string `json:"content"`
}
type SetProductFicti struct {
	Ficti  int32  `json:"ficti"`
	Number string `json:"number"`
	Type   int    `json:"type" binding:"required"` // 1:+ ,2:-
}

type ChangeProductStatus struct {
	Id      []uint `json:"id" form:"id" binding:"required,gt=0"`
	Status  int    `json:"status" binding:"required"`
	Refusal string `json:"refusal" `
}

type ChangeProductIsShow struct {
	Id     uint `json:"id" form:"id" binding:"required,gt=0"`
	IsShow int  `json:"isShow" binding:"required"`
}

type CreateProduct struct {
	model.BaseProduct
	SliderImages []string           `json:"sliderImages"`      // 轮播图
	CateId       uint               `json:"cateId"`            // 平台分类id
	CategoryIds  []uint             `json:"tenancyCategoryId"` // 平台分类id
	Content      string             `json:"content"`           // 商品内容
	Attr         []Value            `json:"attr"`              //规格参数
	AttrValue    []ProductAttrValue `json:"attrValue"`         // 商品内容
}

type ProductAttrValue struct {
	model.BaseProductAttrValue
	Detail map[string]interface{} `json:"detail"`
	Value0 string                 `json:"value0"`
}
