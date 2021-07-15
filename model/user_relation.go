package model

import "github.com/snowlyg/go-tenancy/g"

// UserRelation 用户记录表
type UserRelation struct {
	g.TENANCY_MODEL

	Type   int  `gorm:"uniqueIndex:uid;column:type;type:tinyint;not null;default:0" json:"type"` // 关联类型(1= 普通商品、10 = 店铺、12=购买过)
	TypeID uint `gorm:"uniqueIndex:uid;column:type_id;type:int unsigned;not null" json:"typeId"` // 类型的 id

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
