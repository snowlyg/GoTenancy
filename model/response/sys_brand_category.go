package response

import "github.com/snowlyg/go-tenancy/model"

type SysBrandCategory struct {
	TenancyResponse
	model.BaseBrandCategory
	Children []SysBrandCategory `json:"children" gorm:"-"`
}
