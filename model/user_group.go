package model

import "github.com/snowlyg/go-tenancy/g"

// UserGroup 用户分组表
type UserGroup struct {
	g.TENANCY_MODEL
	GroupName string `gorm:"column:group_name;type:varchar(64);not null" json:"groupName" binding:"required"` // 用户分组名称
}
