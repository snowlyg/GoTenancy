package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
)

type SysTenancy struct {
	g.TENANCY_MODEL
	BaseTenancy
	Region SysRegion `json:"region" gorm:"foreignKey:SysRegionCode;references:code;comment:所属区域"`
}

type BaseTenancy struct {
	UUID         uuid.UUID `json:"uuid" gorm:"comment:UUID"`
	Name         string    `json:"name" form:"name" gorm:"column:name;comment:商户名称" binding:"required"`
	Tele         string    `json:"tele" form:"tele" gorm:"column:tele;comment:商户联系电话"`
	Address      string    `json:"address" form:"address" gorm:"column:address;comment:商户详细地址"`
	BusinessTime string    `json:"businessTime" form:"businessTime" gorm:"column:business_time;comment:商户营业时间"`
	Status       int       `gorm:"column:status;type:tinyint(1);not null;default:2" json:"status" binding:"required"` // 商户是否禁用2锁定,1正常

	SysRegionCode uint `json:"sysRegionCode" form:"sysRegionCode" gorm:"column:sys_region_code;comment:商户所属区域code" binding:"required"`
}
