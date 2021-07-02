package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// ShippingTemplate 运费表
type ShippingTemplate struct {
	g.TENANCY_MODEL
	BaseShippingTemplate

	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"`
}

type BaseShippingTemplate struct {
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`                           // 模板名称
	Type       uint8  `gorm:"column:type;type:tinyint unsigned;not null;default:2" json:"type"`             // 计费方式 1=数量 2=重量 3=体积
	Appoint    uint8  `gorm:"column:appoint;type:tinyint unsigned;not null;default:2" json:"appoint"`       // 开启指定包邮
	Undelivery uint8  `gorm:"column:undelivery;type:tinyint unsigned;not null;default:2" json:"undelivery"` // 开启指定区域不配送
	IsDefault  int    `gorm:"column:is_default;type:tinyint;default:2" json:"isDefault"`                    // 默认模板
	Sort       int    `gorm:"index:mer_id;column:sort;type:int;not null;default:0" json:"sort"`             // 排序
}
