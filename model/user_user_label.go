package model

import "github.com/snowlyg/go-tenancy/g"

// UserUserLabel 用户标签关系表
type UserUserLabel struct {
	g.TENANCY_MODEL

	SysUserID   uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	UserLabelID uint `gorm:"column:user_label_id;" json:"userLabelId"`
}
