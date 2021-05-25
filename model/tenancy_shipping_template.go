package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// TenancyShippingTemplate 运费表
type TenancyShippingTemplate struct {
	g.TENANCY_MODEL
	Name       string `gorm:"column:name;type:varchar(255);not null" json:"name"`                           // 模板名称
	Type       uint8  `gorm:"column:type;type:tinyint unsigned;not null;default:1" json:"type"`             // 计费方式 0=数量 1=重量 2=体积
	Appoint    uint8  `gorm:"column:appoint;type:tinyint unsigned;not null;default:0" json:"appoint"`       // 开启指定包邮
	Undelivery uint8  `gorm:"column:undelivery;type:tinyint unsigned;not null;default:0" json:"undelivery"` // 开启指定区域不配送
	IsDefault  bool   `gorm:"column:is_default;type:bool;default:false" json:"isDefault"`                   // 默认模板
	Sort       int    `gorm:"index:mer_id;column:sort;type:int;not null;default:0" json:"sort"`             // 排序

	TenancyID int `gorm:"index:tenancy_id;column:tenancy_id;type:int;not null" json:"tenancyId"` // 商户 id
}
