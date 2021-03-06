package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"required"`
}

type ChangeStatus struct {
	Id     uint `json:"id" form:"id" binding:"required,gt=0"`
	Status int  `json:"status" binding:"required"`
}

// Find by id structure
type GetById struct {
	Id uint `json:"id" uri:"id" form:"id" binding:"required"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" binding:"required"`
}

type Result struct {
	Count float64
}

type Empty struct{}
