package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysApi struct {
	g.TENANCY_MODEL
	Path        string `json:"path" gorm:"comment:api路径" validate:"required"`
	Description string `json:"description" gorm:"comment:api中文描述" validate:"required"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组" validate:"required"`
	Method      string `json:"method" gorm:"default:POST" gorm:"comment:方法" validate:"required"`
}
