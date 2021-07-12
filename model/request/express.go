package request

// Paging common input parameter structure
type ExpressPageInfo struct {
	Name     string `json:"name" form:"name"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
}
