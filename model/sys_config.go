package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// SysConfig 系统配置
type SysConfig struct {
	g.TENANCY_MODEL
	ConfigName string `gorm:"column:config_name;type:varchar(64);not null" json:"configName"  binding:"required"`               // 字段名称
	ConfigKey  string `gorm:"unique;column:config_key;type:varchar(64);not null" json:"configKey"  binding:"required"`          // 字段 key
	ConfigType string `gorm:"column:config_type;type:varchar(20);not null;default:input" json:"configType"  binding:"required"` // 配置类型
	ConfigRule string `gorm:"column:config_rule;type:varchar(255)" json:"configRule"`                                           // 规则
	Required   int    `gorm:"column:required;type:tinyint unsigned;not null;default:1" json:"required"`                         // 必填
	Info       string `gorm:"column:info;type:varchar(128);default:''" json:"info"`                                             // 配置说明
	Sort       uint16 `gorm:"column:sort;type:smallint unsigned;not null;default:0" json:"sort"`                                // 排序
	UserType   uint8  `gorm:"column:user_type;type:tinyint unsigned;not null;default:1" json:"userType"  binding:"required"`    // 2=总后台配置 1=商户后台配置
	Status     int    `gorm:"column:status;type:tinyint unsigned;not null;default:1" json:"status"  binding:"required"`         // 是否显示

	SysConfigCategoryID int `gorm:"index:sys_config_category_id;column:sys_config_category_id;type:int unsigned;not null" json:"sysConfigCategoryId"  binding:"required"` // 商户 id

	TypeName string `gorm:"-" json:"typeName"` // 配置类型
	Value    string `gorm:"-" json:"value"`    // 配置类型
}
