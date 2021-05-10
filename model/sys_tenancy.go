package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
)

type SysTenancy struct {
	g.TENANCY_MODEL
	UUID          uuid.UUID `json:"uuid" gorm:"comment:UUID"`
	Name          string    `json:"name" form:"name" gorm:"column:name;comment:商户名称"`
	Tele          string    `json:"tele" form:"tele" gorm:"column:tele;comment:商户联系电话"`
	Address       string    `json:"address" form:"address" gorm:"column:address;comment:商户详细地址"`
	BusinessTime  string    `json:"business_hours" form:"business_hours" gorm:"column:business_hours;comment:商户营业时间"`
	Region        SysRegion `json:"region" gorm:"foreignKey:SysRegionCode;references:code;comment:所属区域"`
	SysRegionCode int       `json:"sys_region_code" form:"sys_region_code" gorm:"column:code;comment:商户所属区域code"`
}
