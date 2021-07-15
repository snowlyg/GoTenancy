package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func GetRefundOrder(orderIds []uint, ctx *gin.Context) (float64, error) {
	var refundPayPrice request.Result
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).
		Select("sum(refund_price) as count").
		Where("order_id in ?", orderIds).
		Where("status = ?", model.RefundStatusEnd)

	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}

	err := db.First(&refundPayPrice).Error
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

	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where("refund_orders."+isDelField, g.StatusFalse)
	}

	return db, nil
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
		Select("refund_orders.*,sys_tenancies.name as tenancy_name,sys_tenancies.is_trader as is_trader,general_infos.nick_name as user_nick_name,orders.order_sn as order_sn,orders.activity_type as activity_type").
		Joins("left join orders on refund_orders.order_id = orders.id").
		Joins("left join sys_tenancies on refund_orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_users on refund_orders.sys_user_id = sys_users.id").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id")

	db, err := getRefundOrderSearch(info, ctx, db)
	if err != nil {
		return refundOrderList, stat, total, err
	}
	stat, err = getRefundStat(stat, ctx)
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

		refundProducts, err := getRefundProducts(refundOrderIds)
		if err != nil {
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

func getRefundProducts(refundOrderIds []uint) ([]response.RefundProduct, error) {
	var refundProducts []response.RefundProduct
	err := g.TENANCY_DB.Model(&model.RefundProduct{}).
		Select("refund_products.*,order_products.*").
		Joins("left join order_products on refund_products.order_product_id = order_products.id").
		Where("refund_products.refund_order_id in ?", refundOrderIds).Find(&refundProducts).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return refundProducts, fmt.Errorf("get refund products %w", err)
	}
	return refundProducts, nil
}

func getRefundStat(stat map[string]int64, ctx *gin.Context) (map[string]int64, error) {
	isDelField := GetIsDelField(ctx)

	// 已支付订单数量
	{
		var all int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Count(&all).Error
		if err != nil {
			return nil, err
		}
		stat["all"] = all
	}

	{
		var agree int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("status = ?", model.RefundStatusAgree).Count(&agree).Error
		if err != nil {
			return nil, err
		}
		stat["agree"] = agree
	}

	{
		var audit int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("status = ?", model.RefundStatusAudit).Count(&audit).Error
		if err != nil {
			return nil, err
		}
		stat["audit"] = audit
	}

	{
		var backgood int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("status = ?", model.RefundStatusBackgood).Count(&backgood).Error
		if err != nil {
			return nil, err
		}
		stat["backgood"] = backgood
	}

	{
		var end int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("status = ?", model.RefundStatusEnd).Count(&end).Error
		if err != nil {
			return nil, err
		}
		stat["end"] = end
	}

	{
		var refuse int64
		db := g.TENANCY_DB.Model(&model.RefundOrder{})
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("status = ?", model.RefundStatusRefuse).Count(&refuse).Error
		if err != nil {
			return nil, err
		}
		stat["refuse"] = refuse
	}

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
	refundOrder, err := GetRefundOrderByID(id, ctx)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"mer_mark","value":"%s","title":"备注","props":{"type":"text","placeholder":"请输入备注"}}],"action":"","method":"POST","title":"备注信息","config":{}}`, refundOrder.MerMark)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/refundOrder/remarkRefundOrder", id), ctx)
	return form, err
}

func GetRefundOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := `{"rule":[{"type":"radio","field":"status","value":5,"title":"审核","props":{},"control":[{"value":5,"rule":[{"type":"input","field":"failMessage","value":"","title":"拒绝原因","props":{"type":"text","placeholder":"请输入拒绝原因"},"validate":[{"message":"请输入拒绝原因","required":true,"type":"string","trigger":"change"}]}]}],"options":[{"value":2,"label":"同意"},{"value":5,"label":"拒绝"}]}],"action":"","method":"POST","title":"退款审核","config":{}}`

	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/refundOrder/auditRefundOrder", id), ctx)
	return form, err
}

func GetRefundOrderByID(id uint, ctx *gin.Context) (model.RefundOrder, error) {
	var refundOrder model.RefundOrder
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", id)
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	err := db.First(&refundOrder).Error
	return refundOrder, err
}

func RemarkRefundOrder(id uint, merMark map[string]interface{}, ctx *gin.Context) error {
	db := g.TENANCY_DB.Model(&model.RefundOrder{})
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	return db.Where("id = ?", id).Updates(merMark).Error
}

// GetRefundPriceByOrderIds 获取已退金额
func GetRefundPriceByOrderIds(ids []uint, isDelField string) (float64, error) {
	var wxPayPrice request.Result
	db := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("sum(refund_price) as count")

	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	err := db.Where("status = ?", model.RefundStatusEnd).Where("order_id in ?", ids).First(&wxPayPrice).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return wxPayPrice.Count, err
	}
	return wxPayPrice.Count, nil
}

func checkRefundPrice(refundOrder model.RefundOrder, ctx *gin.Context) (float64, error) {
	order, err := GetOrderById(refundOrder.OrderID, ctx)
	if err != nil {
		return 0, fmt.Errorf("get order %w", err)
	}
	refundPrice, err := GetRefundPriceByOrderIds([]uint{refundOrder.OrderID}, GetIsDelField(ctx))
	if err != nil {
		return 0, fmt.Errorf("get refund order price %w", err)
	}
	payPrice := decimal.NewFromFloat(order.PayPrice)
	refundPriceD := decimal.NewFromFloat(refundPrice)
	pefundPrice := decimal.NewFromFloat(refundOrder.RefundPrice)
	if payPrice.Sub(refundPriceD).LessThanOrEqual(pefundPrice) {
		return 0, fmt.Errorf("退款金额超出订单可退金额")
	}
	return refundPrice, nil
}

func GetOtherRefundOrderIds(orderId, refundOrderId uint) ([]uint, error) {
	var ids []uint
	err := g.TENANCY_DB.Model(&model.RefundOrder{}).Select("id").Where("order_id = ?", orderId).
		Where("status in ?", []int{model.RefundStatusAudit, model.RefundStatusAgree, model.RefundStatusBackgood}).
		Where("id != ?", refundOrderId).Find(&ids).Error
	if err != nil {
		return ids, fmt.Errorf("get other refund order ids %w", err)
	}
	return ids, nil
}

func AuditRefundOrder(id uint, audit request.OrderAudit, ctx *gin.Context) error {
	refundOrder, err := GetRefundOrderByID(id, ctx)
	if err != nil {
		return fmt.Errorf("get refund order %w", err)
	}

	if audit.Status == model.RefundStatusAgree {
		err := agreeRefundOrder(refundOrder, ctx)
		if err != nil {
			return err
		}
	} else if audit.Status == model.RefundStatusRefuse {
		err := refuseRefundOrder(refundOrder, audit.FailMessage, ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

//  2. 拒绝退款
//    2.1 如果退款数量 等于 购买数量 返还可退款数 is_refund = 0
//    2.2 商品总数小于可退数量 返还可退数 以商品数为准
//    2.3 是否存在其他图款单,是 ,退款中 ,否, 部分退款
func refuseRefundOrder(refundOrder model.RefundOrder, failMessage string, ctx *gin.Context) error {
	status := refundOrder.Status
	refundOrderIds, err := GetOtherRefundOrderIds(refundOrder.OrderID, refundOrder.ID)
	if err != nil {
		return err
	}
	refundProducts, err := getRefundProducts([]uint{refundOrder.ID})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单商品状态
		for _, refundProduct := range refundProducts {
			var isRefund uint8
			refundNum := refundProduct.RefundNum + refundProduct.OrderProduct.RefundNum //返还可退款数
			if refundProduct.OrderProduct.ProductNum == refundNum {
				isRefund = 0
			}
			if refundProduct.OrderProduct.ProductNum < refundNum {
				refundNum = refundProduct.OrderProduct.ProductNum
			}

			if len(refundOrderIds) > 0 {
				var count int64
				err := g.TENANCY_DB.Model(&model.RefundProduct{}).Where("refund_order_id in ?", refundOrderIds).Where("order_product_id = ?", refundProduct.ProductID).Count(&count).Error
				if err != nil {
					return fmt.Errorf("get check refund product %w", err)
				}
				if count > 0 {
					isRefund = 1
				}
			}
			refundProduct.OrderProduct.IsRefund = isRefund
			err := tx.Model(&model.OrderProduct{}).Where("id = ?", refundProduct.OrderProduct.ID).Updates(map[string]interface{}{"is_refund": isRefund, "refund_num": refundNum}).Error
			if err != nil {
				return fmt.Errorf("update refund product is_refund %w", err)
			}
		}
		status = model.RefundStatusRefuse
		err := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", refundOrder.ID).Updates(map[string]interface{}{"status": status, "status_time": time.Now(), "fail_message": failMessage}).Error
		if err != nil {
			return fmt.Errorf("update refund order status %w", err)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// agreeRefundOrder
//  1.同意退款
//    1.1 仅退款
//       1.1.1 是 , 如果退款数量 等于 购买数量 is_refund = 3 全退退款 不等于 is_refund = 2 部分退款
//       1.1.2 否, is_refund = 1 退款中
//    1.2 退款退货 is_refund = 1
func agreeRefundOrder(refundOrder model.RefundOrder, ctx *gin.Context) error {
	status := refundOrder.Status
	refundOrderIds, err := GetOtherRefundOrderIds(refundOrder.OrderID, refundOrder.ID)
	if err != nil {
		return err
	}
	refundPrice, err := checkRefundPrice(refundOrder, ctx)
	if err != nil {
		return err
	}
	refundProducts, err := getRefundProducts([]uint{refundOrder.ID})
	if err != nil {
		return err
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		// 更新订单商品状态
		for _, refundProduct := range refundProducts {
			var isRefund uint8
			if refundOrder.RefundType == model.RefundTypeTK {
				if refundProduct.RefundNum == refundProduct.BaseOrderProduct.ProductNum {
					isRefund = 3
				} else {
					isRefund = 2
				}
			}
			if len(refundOrderIds) > 0 {
				var count int64
				err := g.TENANCY_DB.Model(&model.RefundProduct{}).Where("refund_order_id in ?", refundOrderIds).Where("order_product_id = ?", refundProduct.ProductID).Count(&count).Error
				if err != nil {
					return fmt.Errorf("get check refund product %w", err)
				}
				if count > 0 {
					isRefund = 1
				}
			}
			refundProduct.OrderProduct.IsRefund = isRefund
			err := tx.Model(&model.OrderProduct{}).Where("id = ?", refundProduct.OrderProduct.ID).Update("is_refund", isRefund).Error
			if err != nil {
				return fmt.Errorf("update refund product is_refund %w", err)
			}
		}

		// 更新退款单状态
		if refundOrder.RefundType == model.RefundTypeTK {
			status = model.RefundStatusEnd
			// TODO:: 退款操作 func actionRefundPrice(id uint, refundPrice float64){}
			fmt.Printf("refundPrice %f", refundPrice)
		} else if refundOrder.RefundType == model.RefundTypeAll {
			status = model.RefundStatusAgree
			err := AddRefundOrderStatus(refundOrder.ID, "refund_agree", "退款申请已通过，请将商品寄回")
			if err != nil {
				return fmt.Errorf("add refund order status %w", err)
			}
		}
		err := g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", refundOrder.ID).Updates(map[string]interface{}{"status": status, "status_time": time.Now()}).Error
		if err != nil {
			return fmt.Errorf("update refund order status %w", err)
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

// DeleteRefundOrder
func DeleteRefundOrder(id uint) error {
	return g.TENANCY_DB.Model(&model.RefundOrder{}).Where("id = ?", id).Update("is_system_del", g.StatusTrue).Error
}

func AddRefundOrderStatus(id uint, cahngeType, changeMessage string) error {
	status := model.RefundStatus{
		RefundOrderID: id,
		ChangeType:    cahngeType,
		ChangeMessage: changeMessage,
		ChangeTime:    time.Now(),
	}
	return g.TENANCY_DB.Model(&model.RefundStatus{}).Create(&status).Error
}
