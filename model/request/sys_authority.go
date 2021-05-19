package request

import "github.com/snowlyg/go-tenancy/model"

type SetDataAuthority struct {
	AuthorityId     string               `json:"authorityId" validate:"required"`
	DataAuthorityId []model.SysAuthority `json:"dataAuthorityId" validate:"required"`
}

type DeleteAuthority struct {
	AuthorityId  string              `json:"authorityId"  validate:"required"`
	SysBaseMenus []model.SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
