package model

import "github.com/snowlyg/go-tenancy/g"

// UserBill 用户账单表
type UserBill struct {
	g.TENANCY_MODEL

	Pm       uint8   `gorm:"column:pm;type:tinyint unsigned;not null;default:0" json:"pm"`                   // 0 = 支出 1 = 获得
	Title    string  `gorm:"column:title;type:varchar(64);not null" json:"title"`                            // 账单标题
	Category string  `gorm:"index:type;column:category;type:varchar(64);not null" json:"category"`           // 明细种类
	Type     string  `gorm:"index:type;column:type;type:varchar(64);not null;default:''" json:"type"`        // 明细类型
	Number   float64 `gorm:"column:number;type:decimal(8,2) unsigned;not null;default:0.00" json:"number"`   // 明细数字
	Balance  float64 `gorm:"column:balance;type:decimal(8,2) unsigned;not null;default:0.00" json:"balance"` // 剩余
	Mark     string  `gorm:"column:mark;type:varchar(512);not null" json:"mark"`                             // 备注
	Status   int     `gorm:"column:status;type:tinyint(1);not null;default:2" json:"status"`                 // 1 = 待确定 2 = 有效 3 = 无效

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	LinkID    uint `gorm:"index:type;column:link_id;type:varchar(32);not null;default:0" json:"linkId"` // 关联id
}
