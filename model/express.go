package model

import "github.com/snowlyg/go-tenancy/g"

// Express 快递公司表
type Express struct {
	g.TENANCY_MODEL
	Code   string `gorm:"unique;column:code;type:varchar(50);not null" json:"code"`          // 快递公司简称
	Name   string `gorm:"column:name;type:varchar(50);not null" json:"name"`                 // 快递公司全称
	Sort   int    `gorm:"column:sort;type:int;not null" json:"sort"`                         // 排序
	Status int    `gorm:"index:status;column:status;type:tinyint(1);not null" json:"status"` // 是否显示
}
