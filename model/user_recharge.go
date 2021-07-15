package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// UserRecharge 用户充值表
type UserRecharge struct {
	g.TENANCY_MODEL

	Price        float64   `gorm:"column:price;type:decimal(8,2) unsigned;not null;default:0.00" json:"price"`      // 充值金额
	GivePrice    float64   `gorm:"column:give_price;type:decimal(8,2);not null;default:0.00" json:"givePrice"`      // 购买赠送金额
	RechargeType string    `gorm:"column:recharge_type;type:varchar(32);not null" json:"rechargeType"`              // 充值类型
	Paid         int       `gorm:"column:paid;type:tinyint unsigned;not null;default:0" json:"paid"`                // 是否充值
	PayTime      time.Time `gorm:"column:pay_time;type:timestamp" json:"payTime"`                                   // 充值支付时间
	RefundPrice  float64   `gorm:"column:refund_price;type:decimal(10,2) unsigned;default:0.00" json:"refundPrice"` // 退款金额

	SysUserID uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	OrderID   string `gorm:"unique;column:order_id;type:varchar(32);not null" json:"orderId"` // 订单号
}
