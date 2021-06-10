package request

type DeleteMedia struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type UpdateMediaName struct {
	Name string `json:"name" form:"name" binding:"required"`
}

type MediaPageInfo struct {
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pageSize" form:"pageSize" binding:"required"`
	Name     string `json:"name" form:"name"`
}
