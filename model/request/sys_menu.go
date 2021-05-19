package request

import "github.com/snowlyg/go-tenancy/model"

// Add menu authority info structure
type AddMenuAuthorityInfo struct {
	Menus       []model.SysBaseMenu
	AuthorityId string `json:"authorityId"  validate:"required"`
}
