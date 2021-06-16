package request

import "github.com/snowlyg/go-tenancy/model"

type TenancyProductPageInfo struct {
	Page              int    `json:"page" form:"page" binding:"required"`
	PageSize          int    `json:"pageSize" form:"pageSize" binding:"required"`
	TenancyCategoryId int    `json:"tenancyCategoryId" form:"tenancyCategoryId"`
	Type              string `json:"type" form:"type"  binding:"required"`
	Keyword           string `json:"keyword" form:"keyword"`
}

type UpdateTenancyProduct struct {
	Id uint `json:"id"`
	model.BaseTenancyProduct
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
