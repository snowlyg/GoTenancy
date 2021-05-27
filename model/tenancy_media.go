package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type TenancyMedia struct {
	g.TENANCY_MODEL
	Name      string `json:"name" gorm:"comment:文件名"`                                       // 文件名
	Url       string `json:"url" gorm:"comment:文件地址"`                                       // 文件地址
	Tag       string `json:"tag" gorm:"comment:文件标签"`                                       // 文件标签
	Key       string `json:"key" gorm:"comment:编号"`                                         // 编号
	TenancyID uint   `gorm:"column:tenancy_id;type:int unsigned;not null" json:"tenancyId"` // 商户Id
}
