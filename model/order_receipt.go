package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// OrderReceipt 订单发票信息
type OrderReceipt struct {
	g.TENANCY_MODEL

	ReceiptInfo  string    `gorm:"column:receipt_info;type:varchar(500);default:''" json:"receiptInfo"` // 发票类型：1.普通发票，2.增值税发票
	Status       int8      `gorm:"column:status;type:tinyint;default:0" json:"status"`                  // 开票状态：1.已出票,10.已寄出
	ReceiptSn    string    `gorm:"column:receipt_sn;type:varchar(255);default:''" json:"receiptSn"`     // 发票单号
	ReceiptNo    string    `gorm:"column:receipt_no;type:varchar(255)" json:"receiptNo"`                // 发票编号
	DeliveryInfo string    `gorm:"column:delivery_info;type:varchar(255)" json:"deliveryInfo"`          // 收票联系信息
	Mark         string    `gorm:"column:mark;type:varchar(255)" json:"mark"`                           // 用户备注
	ReceiptPrice float64   `gorm:"column:receipt_price;type:decimal(10,2)" json:"receiptPrice"`         // 开票金额
	OrderPrice   float64   `gorm:"column:order_price;type:decimal(10,2)" json:"orderPrice"`             // 订单金额
	StatusTime   time.Time `gorm:"column:status_time;type:datetime;not null" json:"statusTime"`         // 状态变更时间
	MerMark      string    `gorm:"column:mer_mark;type:varchar(255)" json:"merMark"`                    // 备注

	OrderID      string `gorm:"column:order_id;type:varchar(255);not null;default:0" json:"orderId"` // 订单ID
	SysUserID    uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID uint   `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
