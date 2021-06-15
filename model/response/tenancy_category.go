package response

import "github.com/snowlyg/go-tenancy/model"

type TenancyCategory struct {
	TenancyResponse
	model.BaseTenancyCategory
	Children []TenancyCategory `json:"children" gorm:"-"`
}
