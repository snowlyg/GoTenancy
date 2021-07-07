package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

func GetRefundOrder(orderIds []uint) (float64, error) {
	var refundPayPrice request.Result
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("sum(refund_price) as count").Where("order_id in ?", orderIds).Where("status = ?", model.RefundStatusEnd).First(&refundPayPrice).Error
	if err != nil {
		return 0, err
	}
	return refundPayPrice.Count, nil
}
