package request

import "github.com/snowlyg/go-tenancy/model"

type SetDataAuthority struct {
	AuthorityId     string               `json:"authorityId" binding:"required"`
	DataAuthorityId []model.SysAuthority `json:"dataAuthorityId" binding:"required"`
}

type DeleteAuthority struct {
	AuthorityId  string              `json:"authorityId"  binding:"required"`
	SysBaseMenus []model.SysBaseMenu `json:"menus" gorm:"many2many:sys_authority_menus;"`
}
