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
