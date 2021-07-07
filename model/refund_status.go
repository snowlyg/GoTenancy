package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// RefundStatus 订单操作记录表
type RefundStatus struct {
	g.TENANCY_MODEL

	RefundOrderID uint      `gorm:"index:refund_order_id;column:refund_order_id;type:int unsigned;not null" json:"refundOrderId"` // 退款单订单id
	ChangeType    string    `gorm:"index:change_type;column:change_type;type:varchar(32);not null" json:"changeType"`             // 操作类型
	ChangeMessage string    `gorm:"column:change_message;type:varchar(256);not null" json:"changeMessage"`                        // 操作备注
	ChangeTime    time.Time `gorm:"column:change_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"changeTime"`       // 操作时间
}
