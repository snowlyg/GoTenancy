package model

import "github.com/snowlyg/go-tenancy/g"

const (
	UserLabelTypeUnknown int = iota
	UserLabelTypeSD          //手动标签
	UserLabelTypeZD          //自动标签
)

// UserLabel 用户标签表
type UserLabel struct {
	g.TENANCY_MODEL

	LabelName string `gorm:"column:label_name;type:varchar(255);not null;default:''" json:"labelName" ` // 标签名称
	Type      int    `gorm:"column:type;type:tinyint unsigned;not null;default:1" json:"type"`          // 1=手动标签 2=自动标签

	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
