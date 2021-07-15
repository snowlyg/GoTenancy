package model

import "github.com/snowlyg/go-tenancy/g"

// UserVisit 商品浏览分析表
type UserVisit struct {
	g.TENANCY_MODEL

	Type    string `gorm:"index:type;column:type;type:varchar(32);not null" json:"type"`        // 记录类型 page,product
	TypeID  int    `gorm:"index:type;column:type_id;type:int;not null;default:0" json:"typeId"` // 商品ID
	Content string `gorm:"column:content;type:varchar(255)" json:"content"`                     // 备注描述

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
