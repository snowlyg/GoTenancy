package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" binding:"required"`
	PageSize int `json:"pageSize" form:"pageSize" binding:"required"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id" binding:"required"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authorityId" form:"authorityId" binding:"required"`
}

type Empty struct{}
