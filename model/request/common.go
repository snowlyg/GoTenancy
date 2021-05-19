package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page" validate:"required"`
	PageSize int `json:"pageSize" form:"pageSize" validate:"required"`
}

// Find by id structure
type GetById struct {
	Id float64 `json:"id" form:"id" validate:"required"`
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids" validate:"required"`
}

// Get role by id structure
type GetAuthorityId struct {
	AuthorityId string `json:"authority_id" form:"authority_id" validate:"required"`
}

type Empty struct{}
