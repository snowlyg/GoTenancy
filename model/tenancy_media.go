package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type TenancyMedia struct {
	g.TENANCY_MODEL
	Name string `json:"name" gorm:"comment:文件名"` // 文件名
	Url  string `json:"url" gorm:"comment:文件地址"` // 文件地址
	Tag  string `json:"tag" gorm:"comment:文件标签"` // 文件标签
	Key  string `json:"key" gorm:"comment:编号"`   // 编号

	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
