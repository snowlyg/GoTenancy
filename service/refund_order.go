package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func GetRefundOrder(orderIds []uint) (float64, error) {
	var refundPayPrice request.Result
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Select("sum(refund_price) as count").
		Where("order_id in ?", orderIds).
		Where("status = ?", model.RefundStatusEnd).
		Where("is_system_del", g.StatusFalse).
		First(&refundPayPrice).Error
	if err != nil {
		return 0, err
	}
	return refundPayPrice.Count, nil
}

func getRefundOrderSearch(info request.RefundOrderPageInfo, ctx *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	if info.Status != "" {
		db = db.Where("refund_orders.status = ?", info.Status)
	}

	if multi.IsTenancy(ctx) {
		db = db.Where("refund_orders.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}

	if info.Date != "" {
		db = filterDate(db, info.Date, "refund_orders")
	}

	if info.IsTrader != "" {
		db = db.Where("sys_tenancies.is_trader = ?", info.IsTrader)
	}
	if info.OrderSn != "" {
		db = db.Where("orders.order_sn like ?", info.OrderSn+"%")
	}

	if info.RefundOrderSn != "" {
		db = db.Where("refund_orders.refund_order_sn like ?", info.RefundOrderSn+"%")
	}

	return db.Where("is_system_del", g.StatusFalse), nil
}

// GetRefundOrderInfoList
func GetRefundOrderInfoList(info request.RefundOrderPageInfo, ctx *gin.Context) ([]response.RefundOrderList, map[string]int64, int64, error) {
	stat := map[string]int64{
		"agree":    0,
		"all":      0,
		"audit":    0,
		"backgood": 0,
		"end":      0,
		"refuse":   0,
	}
	var refundOrderList []response.RefundOrderList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Select("refund_orders.*,sys_tenancies.name as tenancy_name,sys_tenancies.is_trader as is_trader,sys_general_infos.nick_name as user_nick_name,orders.order_sn as order_sn,orders.activity_type as activity_type").
		Joins("left join orders on refund_orders.order_id = orders.id").
		Joins("left join sys_tenancies on refund_orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_users on refund_orders.sys_user_id = sys_users.id").
		Joins("left join sys_general_infos on sys_general_infos.sys_user_id = sys_users.id")

	db, err := getRefundOrderSearch(info, ctx, db)
	if err != nil {
		return refundOrderList, stat, total, err
	}
	stat, err = getRefundStat(stat)
	if err != nil {
		return refundOrderList, stat, total, err
	}

	err = db.Count(&total).Error
	if err != nil {
		return refundOrderList, stat, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&refundOrderList).Error
	if err != nil {
		return refundOrderList, stat, total, err
	}

	if len(refundOrderList) > 0 {
		var refundOrderIds []uint
		for _, refundOrder := range refundOrderList {
			refundOrderIds = append(refundOrderIds, refundOrder.ID)
		}

		var refundProducts []response.RefundProduct
		err = g.TENANCY_DB.Model(&model.RefundProduct{}).
			Select("refund_products.*,order_products.*").
			Joins("left join order_products on refund_products.order_product_id = order_products.id").
			Where("refund_products.refund_order_id in ?", refundOrderIds).Find(&refundProducts).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return refundOrderList, stat, total, err
		}

		for i := 0; i < len(refundOrderList); i++ {
			for _, refundProduct := range refundProducts {
				if refundOrderList[i].ID == refundProduct.RefundOrderID {
					refundOrderList[i].RefundProduct = append(refundOrderList[i].RefundProduct, refundProduct)
				}
			}
		}
	}

	return refundOrderList, stat, total, nil
}

func getRefundStat(stat map[string]int64) (map[string]int64, error) {
	// 已支付订单数量
	var all int64
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Count(&all).Error
	if err != nil {
		return nil, err
	}
	stat["all"] = all

	var agree int64
	err = g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("status = ?", model.RefundStatusAgree).Count(&agree).Error
	if err != nil {
		return nil, err
	}
	stat["agree"] = agree

	var audit int64
	err = g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("status = ?", model.RefundStatusAudit).Count(&audit).Error
	if err != nil {
		return nil, err
	}
	stat["audit"] = audit

	var backgood int64
	err = g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("status = ?", model.RefundStatusBackgood).Count(&backgood).Error
	if err != nil {
		return nil, err
	}
	stat["backgood"] = backgood

	var end int64
	err = g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("status = ?", model.RefundStatusEnd).Count(&end).Error
	if err != nil {
		return nil, err
	}
	stat["end"] = end

	var refuse int64
	err = g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("status = ?", model.RefundStatusRefuse).Count(&refuse).Error
	if err != nil {
		return nil, err
	}
	stat["refuse"] = refuse

	return stat, nil
}

func GetRefundOrderRecord(id uint, info request.PageInfo) ([]model.RefundStatus, int64, error) {
	var returnRecord []model.RefundStatus
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.RefundStatus{}).Where("refund_order_id = ?", id)
	err := db.Count(&total).Error
	if err != nil {
		return returnRecord, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&returnRecord).Error
	if err != nil {
		return returnRecord, total, err
	}
	return returnRecord, total, nil
}

func GetRefundOrderRemarkMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	remark, err := GetRefundOrderRemarkByID(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"merMark","value":"%s","title":"备注","props":{"type":"text","placeholder":"请输入备注"}}],"action":"","method":"POST","title":"备注信息","config":{}}`, remark)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/refundOrder/remarkRefundOrder", id), ctx)
	return form, err
}

func GetRefundOrderRemarkByID(id uint) (string, error) {
	var merMark string
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("mer_mark").Where("id = ?", id).Where("is_system_del", g.StatusFalse).First(&merMark).Error
	return merMark, err
}

func RemarkRefundOrder(id uint, merMark map[string]interface{}) error {
	return g.TENANCY_DB.Model(&model.RefundOrder{}).Where("is_system_del", g.StatusFalse).Where("id = ?", id).Updates(merMark).Error
}
