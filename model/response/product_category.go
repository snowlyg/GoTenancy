package response

import "github.com/snowlyg/go-tenancy/model"

type ProductCategory struct {
	TenancyResponse
	model.BaseProductCategory
	Children []ProductCategory `json:"children" gorm:"-"`
}
