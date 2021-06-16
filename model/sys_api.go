package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysApi struct {
	g.TENANCY_MODEL
	Path        string `json:"path" gorm:"comment:api路径" binding:"required"`
	Description string `json:"description" gorm:"comment:api中文描述" binding:"required"`
	ApiGroup    string `json:"apiGroup" gorm:"comment:api组" binding:"required"`
	Method      string `json:"method" gorm:"default:POST;comment:方法" binding:"required"`
}
