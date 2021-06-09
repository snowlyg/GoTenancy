package model

import "github.com/snowlyg/go-tenancy/g"

type SysConfigCategory struct {
	g.TENANCY_MODEL
	Name   string `gorm:"column:name;type:varchar(255);not null" json:"name" binding:"required"`      // 配置分类名称
	Key    string `gorm:"unique;column:key;type:varchar(255);not null" json:"key" binding:"required"` // 配置分类英文名称
	Info   string `gorm:"column:info;type:varchar(30)" json:"info"`                                   // 配置分类说明
	Sort   uint16 `gorm:"column:sort;type:smallint unsigned;not null;default:0" json:"sort"`          // 排序
	Icon   string `gorm:"column:icon;type:varchar(30)" json:"icon"`                                   // 图标
	Status int    `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"`       // 配置分类状态
}
