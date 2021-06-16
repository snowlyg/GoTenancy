package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type SysConfigValue struct {
	g.TENANCY_MODEL
	ConfigKey string `gorm:"index:config_name;column:config_key;type:varchar(32);not null" json:"configKey"` // 配置分类key
	Value     string `gorm:"column:value;type:varchar(2000);not null" json:"value"`                          // 值

	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int unsigned;not null" json:"sysTenancyId"` // 商户 id
}
