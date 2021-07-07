package model

import "github.com/snowlyg/go-tenancy/g"

// RefundProduct 退款单产品表
type RefundProduct struct {
	g.TENANCY_MODEL

	RefundOrderID  uint `gorm:"index:refund_order_id;column:refund_order_id;type:int unsigned;not null" json:"refundOrderId"` // 退款单
	OrderProductID uint `gorm:"column:order_product_id;type:int unsigned;not null" json:"orderProductId"`                     // 订单产品id
	RefundNum      uint `gorm:"column:refund_num;type:int unsigned;not null;default:0" json:"refundNum"`                      // 退货数
}
