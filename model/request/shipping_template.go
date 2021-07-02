package request

type ShippingTemplatePageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	Name     string `json:"name" form:"name"`
}

type UpdateShippingTemplate struct {
}
