package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func DeliveryOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	formStr := `{"rule":[{"type":"radio","field":"deliveryType","value":1,"title":"发货类型","props":{},"control":[{"value":1,"rule":[{"type":"select","field":"deliveryName","value":"","title":"快递名称","props":{"multiple":false,"placeholder":"请选择快递名称"},"options":[]},{"type":"input","field":"deliveryId","value":"","title":"快递单号","props":{"type":"text","placeholder":"请输入快递单号"},"validate":[{"message":"请输入快递单号","required":true,"type":"string","trigger":"change"}]}]},{"value":2,"rule":[{"type":"input","field":"deliveryName","value":"","title":"送货人姓名","props":{"type":"text","placeholder":"请输入送货人姓名"},"validate":[{"message":"请输入送货人姓名","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"deliveryId","value":"","title":"手机号","props":{"type":"text","placeholder":"请输入手机号"},"validate":[{"message":"请输入手机号","required":true,"type":"string","trigger":"change"}]}]},{"value":3,"rule":[]}],"options":[{"value":1,"label":"发货"},{"value":2,"label":"送货"},{"value":3,"label":"无需物流"}]}],"action":"","method":"POST","title":"添加发货信息","config":{}}`
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/deliveryOrder", id), ctx)

	opts, err := GetExpressOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Control[0].Rule[0].Options = opts
	return form, err
}

func GetOrderRemarkMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	order, err := GetOrderRemarkAndUpdateByID(id, ctx)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"remark","value":"%s","title":"备注","props":{"type":"text","placeholder":"请输入备注"},"validate":[{"message":"请输入备注","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"修改备注","config":{}}`, order.Remark)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/remarkOrder", id), ctx)
	return form, err
}

func GetEditOrderMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	order, err := GetOrderRemarkAndUpdateByID(id, ctx)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"inputNumber","field":"totalPrice","value":%02f,"title":"订单总价","props":{"placeholder":"请输入订单总价"},"validate":[{"message":"请输入订单总价","required":true,"type":"number","trigger":"change"}]},{"type":"inputNumber","field":"payPrice","value":%02f,"title":"实际支付金额","props":{"placeholder":"请输入实际支付金额"},"validate":[{"message":"请输入实际支付金额","required":true,"type":"number","trigger":"change"}]},{"type":"inputNumber","field":"totalPostage","value":%02f,"title":"订单邮费","props":{"placeholder":"请输入订单邮费"},"validate":[{"message":"请输入订单邮费","required":true,"type":"number","trigger":"change"}]}],"action":"","method":"POST","title":"修改订单","config":{}}`, order.TotalPrice, order.PayPrice, order.TotalPostage)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/order/updateOrder", id), ctx)
	return form, err
}

func GetOrderRemarkAndUpdateByID(id uint, ctx *gin.Context) (request.OrderRemarkAndUpdate, error) {
	var order request.OrderRemarkAndUpdate
	db := g.TENANCY_DB.Model(&model.Order{}).Select("remark,total_price,pay_price,total_postage").Where("id = ?", id)
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	err := db.First(&order).Error
	return order, err
}

// getOrderCount
func getOrderCount(name string, ctx *gin.Context) (int64, error) {
	var count int64
	wheres := getOrderConditions()
	for _, where := range wheres {
		if where.Name == name {
			db := g.TENANCY_DB.Model(&model.Order{})
			if where.Conditions != nil && len(where.Conditions) > 0 {
				for key, cn := range where.Conditions {
					if cn == nil {
						db = db.Where(key)
					} else {
						db = db.Where(fmt.Sprintf("%s = ?", key), cn)
					}
				}
			}

			isDelField := GetIsDelField(ctx)
			if isDelField != "" {
				db = db.Where(isDelField, g.StatusFalse)
			}

			err := db.Count(&count).Error
			if err != nil {
				return count, err
			}
		}
	}

	return count, nil
}

func GetFilter(ctx *gin.Context) ([]map[string]interface{}, error) {
	charts := []map[string]interface{}{
		{"count": 0, "orderType": "", "title": "全部"},
		{"count": 0, "orderType": "1", "title": "普通订单"},
		{"count": 0, "orderType": "2", "title": "核销订单"},
	}

	for _, chart := range charts {
		if chart["orderType"] == "" {
			continue
		}
		var count int64
		db := g.TENANCY_DB.Model(&model.Order{})
		isDelField := GetIsDelField(ctx)
		if isDelField != "" {
			db = db.Where(isDelField, g.StatusFalse)
		}
		err := db.Where("order_type = ?", chart["orderType"]).Count(&count).Error
		if err != nil {
			return nil, err
		} else {
			chart["count"] = count
		}

	}

	return charts, nil
}

func GetChart(ctx *gin.Context) (map[string]interface{}, error) {
	charts := map[string]interface{}{
		"all":        0,
		"complete":   0,
		"del":        0,
		"refund":     0,
		"statusAll":  0,
		"unevaluate": 0,
		"unpaid":     0,
		"unshipped":  0,
		"untake":     0,
	}
	for name, _ := range charts {
		if cc, err := getOrderCount(name, ctx); err != nil {
			return nil, err
		} else {
			charts[name] = cc
		}

	}

	return charts, nil
}

//1: 未支付 2: 未发货 3: 待收货 4: 待评价 5: 交易完成 6: 已退款 7: 已删除
// getOrderConditions
func getOrderConditions() []response.OrderCondition {
	conditions := []response.OrderCondition{
		{Name: "all", Type: 0, Conditions: nil},
		{Name: "unpaid", Type: 1, Conditions: map[string]interface{}{"paid": 2, "is_del": 2}},
		{Name: "unshipped", Type: 2, Conditions: map[string]interface{}{"paid": 1, "status": 1, "is_del": 2}},
		{Name: "untake", Type: 3, Conditions: map[string]interface{}{"status": 2, "is_del": 2}},
		{Name: "unevaluate", Type: 4, Conditions: map[string]interface{}{"status": 3, "is_del": 2}},
		{Name: "complete", Type: 5, Conditions: map[string]interface{}{"status": 4, "is_del": 2}},
		{Name: "refund", Type: 6, Conditions: map[string]interface{}{"status": 5, "is_del": 2}},
		{Name: "del", Type: 7, Conditions: map[string]interface{}{"is_del": 1}},
	}
	return conditions
}

// getOrderConditionByStatus
func getOrderConditionByStatus(status int) response.OrderCondition {
	conditions := getOrderConditions()
	for _, condition := range conditions {
		if condition.Type == status {
			return condition
		}
	}
	return conditions[0]
}

func GetOrderById(id uint, ctx *gin.Context) (response.OrderDetail, error) {
	var order response.OrderDetail
	db := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,sys_general_infos.nick_name as user_nick_name").
		Joins("left join sys_users on orders.sys_user_id = sys_users.id").
		Joins("left join sys_general_infos on sys_general_infos.sys_user_id = sys_users.id").
		Joins(fmt.Sprintf("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id and sys_authorities.authority_type = %d", multi.GeneralAuthority))

	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where("orders."+isDelField, g.StatusFalse)
	}

	err := db.Where("orders.id = ?", id).
		First(&order).Error
	if err != nil {
		return order, err
	}
	return order, nil
}

func GetOrderRecord(id uint, info request.PageInfo) ([]model.OrderStatus, int64, error) {
	var orderRecord []model.OrderStatus
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.OrderStatus{}).Where("order_id = ?", id)
	err := db.Count(&total).Error
	if err != nil {
		return orderRecord, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&orderRecord).Error
	if err != nil {
		return orderRecord, total, err
	}
	return orderRecord, total, nil
}

func RemarkOrder(id uint, remark map[string]interface{}, ctx *gin.Context) error {
	db := g.TENANCY_DB.Model(&model.Order{})
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	return db.Where("id = ?", id).Updates(remark).Error
}

func DeliveryOrder(id uint, delivery request.DeliveryOrder, ctx *gin.Context) error {

	var changeMessage string
	var deliveryName string
	switch delivery.DeliveryType {
	case model.DeliverTypeFH:
		express, err := GetExpressByCode(delivery.DeliveryName)
		if err != nil {
			return fmt.Errorf("get express err %w", err)
		}
		deliveryName = express.Name
		changeMessage = fmt.Sprintf("订单已配送【快递名称】:%s; 【快递单号】：%s", deliveryName, delivery.DeliveryId)
	case model.DeliverTypeSH:
		deliveryName = delivery.DeliveryName
		regexp := regexp.MustCompile(`^1[3456789]{1}\d{9}$`)
		if !regexp.MatchString(delivery.DeliveryId) {
			return fmt.Errorf("手机号格式错误")
		}
		changeMessage = fmt.Sprintf("订单已配送【送货人姓名】:%s; 【手机号】：%s", deliveryName, delivery.DeliveryId)
	case model.DeliverTypeXN:
		changeMessage = "订单已配送【虚拟发货】"
	default:
		return fmt.Errorf("error deliver type %d", delivery.DeliveryType)
	}

	orderDelivery := map[string]interface{}{
		"delivery_id":   delivery.DeliveryId,
		"delivery_name": deliveryName,
		"delivery_type": delivery.DeliveryType,
		"status":        model.OrderStatusNoReceive,
	}
	db := g.TENANCY_DB.Model(&model.Order{})
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	err := db.Where("id = ?", id).Updates(orderDelivery).Error
	if err != nil {
		return fmt.Errorf("update order info %w", err)
	}

	orderStatus := model.OrderStatus{
		ChangeType:    fmt.Sprintf("delivery_%d", delivery.DeliveryType),
		ChangeMessage: changeMessage,
		ChangeTime:    time.Now(),
		OrderID:       id,
	}
	err = g.TENANCY_DB.Model(&model.OrderStatus{}).Create(&orderStatus).Error
	if err != nil {
		return fmt.Errorf("update order status info %d", err)
	}
	return nil
}

// GetOrderInfoList
func GetOrderInfoList(info request.OrderPageInfo, ctx *gin.Context) ([]response.OrderList, []map[string]interface{}, int64, error) {
	stat := []map[string]interface{}{
		{"className": "el-icon-s-goods", "count": 0, "field": "件", "name": "已支付订单数量"},
		{"className": "el-icon-s-order", "count": 0, "field": "元", "name": "实际支付金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "已退款金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "微信支付金额"},
		{"className": "el-icon-s-finance", "count": 0, "field": "元", "name": "余额支付金额"},
		{"className": "el-icon-s-cooperation", "count": 0, "field": "元", "name": "支付宝支付金额"},
	}
	var orderList []response.OrderList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Order{}).
		Select("orders.*,sys_tenancies.name as tenancy_name,sys_tenancies.is_trader as is_trader,group_orders.group_order_sn as group_order_sn").
		Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")

	db, err := getOrderSearch(info, ctx, db)
	if err != nil {
		return orderList, stat, total, err
	}

	stat, err = getStat(info, ctx, stat)
	if err != nil {
		return orderList, stat, total, err
	}

	err = db.Count(&total).Error
	if err != nil {
		return orderList, stat, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&orderList).Error
	if err != nil {
		return orderList, stat, total, err
	}

	if len(orderList) > 0 {
		var orderIds []uint
		for _, order := range orderList {
			orderIds = append(orderIds, order.ID)
		}

		var orderProducts []response.OrderProduct
		err = g.TENANCY_DB.Model(&model.OrderProduct{}).Where("order_id in ?", orderIds).Find(&orderProducts).Error
		if err != nil {
			return orderList, stat, total, err
		}

		for i := 0; i < len(orderList); i++ {
			for _, orderProduct := range orderProducts {
				if orderList[i].ID == orderProduct.OrderID {
					orderList[i].OrderProduct = append(orderList[i].OrderProduct, orderProduct)
				}
			}
		}
	}

	return orderList, stat, total, nil
}

func getOrderSearch(info request.OrderPageInfo, ctx *gin.Context, db *gorm.DB) (*gorm.DB, error) {
	if info.Status != "" {
		status, err := strconv.Atoi(info.Status)
		if err != nil {
			return db, err
		}
		cond := getOrderConditionByStatus(status)
		if cond.IsDeleted {
			db = db.Unscoped()
		}

		for key, cn := range cond.Conditions {
			if cn == nil {
				db = db.Where(fmt.Sprintf("%s%s", "orders.", key))
			} else {
				db = db.Where(fmt.Sprintf("%s%s = ?", "orders.", key), cn)
			}
		}
	}

	if multi.IsTenancy(ctx) {
		db = db.Where("orders.sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	} else {
		if info.SysTenancyId > 0 {
			db = db.Where("orders.sys_tenancy_id = ?", info.SysTenancyId)
		}
	}

	if info.Date != "" {
		db = filterDate(db, info.Date, "orders")
	}

	if info.OrderType != "" && info.OrderType != "0" {
		db = db.Where("orders.order_type = ?", info.OrderType)
	}

	if info.ActivityType != "" {
		db = db.Where("orders.activity_type = ?", info.ActivityType)
	}

	if info.IsTrader != "" {
		db = db.Where("sys_tenancies.is_trader = ?", info.IsTrader)
	}
	if info.OrderSn != "" {
		db = db.Where("orders.order_sn like ?", info.OrderSn+"%")
	}

	if info.Keywords != "" {
		db = db.Where(g.TENANCY_DB.Where("orders.order_sn like ?", info.Keywords+"%").Or("orders.real_name like ?", info.Keywords+"%").Or("orders.user_phone like ?", info.Keywords+"%"))
	}

	if info.Username != "" {
		db = db.Where(g.TENANCY_DB.Where("orders.order_sn like ?", info.Keywords+"%").Or("orders.real_name like ?", info.Keywords+"%").Or("orders.user_phone like ?", info.Keywords+"%"))
	}
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where("orders."+isDelField, g.StatusFalse)
	}
	return db, nil
}

func getStat(info request.OrderPageInfo, ctx *gin.Context, stat []map[string]interface{}) ([]map[string]interface{}, error) {
	// 已支付订单数量
	var all int64
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Where("orders.paid =?", 1).Count(&all).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[0]["count"] = all
	}

	//实际支付金额
	var payPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {

		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).First(&payPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[1]["count"] = payPrice.Count
	}

	//已退款金额
	var orderIds []uint
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {

		err = db.Select("orders.id").Find(&orderIds).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if len(orderIds) > 0 {
			stat[2]["count"], _ = GetRefundOrder(orderIds, ctx)
		}
	}

	//微信支付金额
	var wxPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type in ?", []int{model.PayTypeWx, model.PayTypeRoutine, model.PayTypeH5}).First(&wxPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[3]["count"] = wxPayPrice.Count
	}

	//余额支付金额
	var blanPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type = ?", model.PayTypeBalance).First(&blanPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[4]["count"] = blanPayPrice.Count
	}

	//支付宝支付金额
	var aliPayPrice request.Result
	if db, err := getOrderSearch(info, ctx, g.TENANCY_DB.Model(&model.Order{}).Joins("left join sys_tenancies on orders.sys_tenancy_id = sys_tenancies.id").
		Joins("left join group_orders on orders.group_order_id = group_orders.id")); err != nil {
		return nil, err
	} else {
		err = db.Select("sum(orders.pay_price) as count").Where("orders.paid =?", 1).Where("orders.pay_type = ?", model.PayTypeAlipay).First(&aliPayPrice).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		stat[5]["count"] = aliPayPrice.Count
	}

	return stat, nil
}

// UpdateOrder
func UpdateOrder(id uint, order request.OrderRemarkAndUpdate, ctx *gin.Context) error {
	db := g.TENANCY_DB.Model(&model.Order{}).Where("id = ?", id)
	isDelField := GetIsDelField(ctx)
	if isDelField != "" {
		db = db.Where(isDelField, g.StatusFalse)
	}
	return db.Updates(map[string]interface{}{
		"pay_price":     order.PayPrice,
		"total_price":   order.TotalPrice,
		"total_postage": order.TotalPostage,
	}).Error
}

// DeleteOrder
func DeleteOrder(id uint) error {
	return g.TENANCY_DB.Model(&model.Order{}).Where("id = ?", id).Update("is_system_del", g.StatusTrue).Error
}
