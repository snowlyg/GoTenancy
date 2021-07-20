package model

import "github.com/snowlyg/go-tenancy/g"

// UserUserLabel 用户标签关系表
type UserUserLabel struct {
	g.TENANCY_MODEL

	SysUserID    uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	UserLabelID  uint `gorm:"column:user_label_id;" json:"userLabelId"`
	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
