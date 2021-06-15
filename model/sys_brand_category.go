package model

import "github.com/snowlyg/go-tenancy/g"

type SysBrandCategory struct {
	g.TENANCY_MODEL
	Pid      string             `gorm:"index:pid;column:pid;type:mediumint;not null" json:"pid"`        // 父id
	CateName string             `gorm:"column:cate_name;type:varchar(100);not null" json:"cateName"`    // 分类名称
	Path     string             `gorm:"column:path;type:varchar(255);not null;default:''" json:"path"`  // 路径
	Sort     int32              `gorm:"index;column:sort;type:mediumint;not null" json:"sort"`          // 排序
	Status   int                `gorm:"column:status;type:tinyint(1);not null" json:"status"`           // 是否显示
	Level    uint               `gorm:"column:level;type:int unsigned;not null;default:0" json:"level"` // 等级
	Children []SysBrandCategory `json:"children" gorm:"-"`
}
