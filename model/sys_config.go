package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// SysConfig 系统配置
type SysConfig struct {
	g.TENANCY_MODEL
	Type  string `json:"type" form:"type" gorm:"column:type;comment:类型"`
	Name  string `json:"name" form:"name" gorm:"column:name;comment:名称"`
	Value string `json:"value" form:"value" gorm:"column:value;comment:设置值"`

	SysTenancyID int `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
