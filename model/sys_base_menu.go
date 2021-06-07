package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysBaseMenu struct {
	g.TENANCY_MODEL
	Pid           uint           `gorm:"index:pid;column:pid;type:int unsigned;not null;default:0" json:"pid"`         // 父级id
	Path          string         `gorm:"column:path;type:varchar(512);not null" json:"path"`                           // 路径
	Icon          string         `gorm:"column:icon;type:varchar(32);default:''" json:"icon"`                          // 图标
	MenuName      string         `gorm:"column:menu_name;type:varchar(128);not null;default:''" json:"menu_name"`      // 按钮名
	Route         string         `gorm:"column:route;type:varchar(64);not null" json:"route"`                          // 路由名称
	Params        string         `gorm:"column:params;type:varchar(128);not null;default:''" json:"params"`            // 参数
	Sort          int8           `gorm:"column:sort;type:tinyint;not null;default:1" json:"sort"`                      // 排序
	Hidden        uint8          `gorm:"column:hidden;type:tinyint unsigned;not null;default:1" json:"hidden"`         // 是否显示
	IsTenancy     uint8          `gorm:"column:is_tenancy;type:tinyint unsigned;not null;default:1" json:"is_tenancy"` // 模块，1 平台， 2商户
	IsMenu        uint8          `gorm:"column:is_menu;type:tinyint unsigned;not null;default:1" json:"is_menu"`       // 类型，1菜单 2 权限
	SysAuthoritys []SysAuthority `json:"authoritys" gorm:"many2many:sys_authority_menus;"`
	Children      []SysBaseMenu  `json:"children" gorm:"-"`
}
