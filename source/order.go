package source

import (
	"encoding/json"
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

var Order = new(order)

type order struct{}

var groupOrders = []model.GroupOrder{
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(model.PayTypeWx), TotalPostage: 20.00, TotalPrice: 50.00, TotalNum: 1, RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", PayPrice: 50.00, PayPostage: 30.00, Cost: 5.00, Paid: 0, PayTime: time.Now(), PayType: model.PayTypeWx, IsRemind: 1},
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(model.PayTypeWx), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: 0, PayType: model.PayTypeWx, IsRemind: 1, TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now()}}},
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(model.PayTypeWx), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: 1, PayType: model.PayTypeWx, IsRemind: 1},
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(model.PayTypeWx), TotalPostage: 1.00, TotalPrice: 88.00, TotalNum: 1, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 89.00, PayPostage: 1.00, Cost: 100.00, Paid: 1, PayType: model.PayTypeWx, IsRemind: 1},
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(model.PayTypeWx), TotalPostage: 4.00, TotalPrice: 352.00, TotalNum: 4, RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", PayPrice: 356.00, PayPostage: 4.00, Cost: 400.00, Paid: 1, PayType: model.PayTypeWx, IsRemind: 1},
}
var orders = []model.Order{
	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 1, ReconciliationID: 0, CartID: 1, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.PayTypeWx), RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", TotalNum: 10, TotalPrice: 20.00, TotalPostage: 30.00, PayPrice: 50.00, PayPostage: 30.00, CommissionRate: 15.00, OrderType: model.OrderTypeGeneral, Paid: g.StatusFalse, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusRefund, DeliveryType: model.DeliverTypeFH, DeliveryName: "delivery_name", DeliveryID: "delivery_id", Mark: "mark", Remark: "remark", AdminMark: "admin_mark", VerifyCode: "", VerifyTime: time.Now(), ActivityType: model.GeneralSale, Cost: 5.00, IsDel: g.StatusFalse}},

	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 2, CartID: 2, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.PayTypeWx), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusFalse, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusNoReceive, Cost: 100.00, IsDel: g.StatusTrue}},

	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 3, CartID: 3, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.PayTypeWx), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusNoComment, Cost: 100.00, IsDel: g.StatusFalse}},

	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 4, CartID: 4, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.PayTypeWx), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 1, TotalPrice: 88.00, TotalPostage: 1.00, PayPrice: 89.00, PayPostage: 1.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusFinish, Cost: 100.00, IsDel: g.StatusFalse}},

	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 5, CartID: 5, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(model.PayTypeWx), RealName: "发斯蒂芬斯蒂芬", UserPhone: "13672286043", UserAddress: "北京市北京市东城区 的是非得失", TotalNum: 4, TotalPrice: 352.00, TotalPostage: 4.00, PayPrice: 356.00, PayPostage: 4.00, CommissionRate: 0.2000, OrderType: model.OrderTypeGeneral, Paid: g.StatusTrue, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusNoDeliver, Cost: 400.00, IsDel: g.StatusFalse}},
}

func getOrderProducts() []model.OrderProduct {
	cartInfo := request.CartInfo{
		Product: request.CartInfoProduct{
			Image:     "http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg",
			StoreName: "领立裁腰带短袖连衣裙",
		},
		ProductAttr: request.CartInfoProductAttr{
			Price: 50.00,
			Sku:   "L",
		},
	}
	ci, _ := json.Marshal(&cartInfo)

	orderProducts := []model.OrderProduct{
		{OrderID: 1, SysUserID: 3, CartID: 1, ProductID: 1, CartInfo: string(ci), BaseOrderProduct: model.BaseOrderProduct{ProductSku: "L", IsRefund: 0, ProductNum: 12, ProductType: model.GeneralSale, RefundNum: 5, IsReply: g.StatusFalse, ProductPrice: 50.00}},

		{OrderID: 1, SysUserID: 3, CartID: 1, ProductID: 2, CartInfo: string(ci), BaseOrderProduct: model.BaseOrderProduct{ProductSku: "L", IsRefund: 0, ProductNum: 12, ProductType: model.GeneralSale, RefundNum: 5, IsReply: g.StatusFalse, ProductPrice: 50.00}},

		{OrderID: 2, SysUserID: 3, CartID: 2, ProductID: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"isGiftBag\":0,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":99,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 1, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderID: 3, SysUserID: 3, CartID: 3, ProductID: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"isGiftBag\":0,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":98,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 1, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 0, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderID: 4, SysUserID: 3, CartID: 4, ProductID: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"isGiftBag\":0,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":97,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 1, ProductType: model.GeneralSale, RefundNum: 1, IsReply: g.StatusFalse, ProductPrice: 88.00}},

		{OrderID: 5, SysUserID: 3, CartID: 5, ProductID: 1, CartInfo: "{\"product\":{\"productId\":7,\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"storeName\":\"\u6885\u6e7e\u8857\u590d\u53e4\u96ea\u7eba\u7ffb\u9886\u4e0a\u8863\",\"isShow\":1,\"status\":1,\"isDel\":0,\"unitName\":\"\u4ef6\",\"price\":\"88.00\",\"tempId\":96,\"isGiftBag\":0,\"productType\":0,\"temp\":{\"id\":96,\"name\":\"\u8fd0\u8d39\u8bbe\u7f6e\",\"type\":0,\"appoint\":0,\"undelivery\":0,\"sysTenancyId\":64,\"isDefault\":0,\"sort\":0,\"createdAt\":\"2020-07-02 17:48:53\"}},\"productAttr\":{\"image\":\"http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg\",\"productId\":7,\"stock\":96,\"price\":\"88.00\",\"unique\":\"167a5a36ded0\",\"sku\":\"\",\"volume\":\"1.00\",\"weight\":\"1.00\",\"otPrice\":\"200.00\",\"cost\":\"100.00\"},\"productType\":0}", BaseOrderProduct: model.BaseOrderProduct{ProductSku: "167a5a36ded0", IsRefund: 0, ProductNum: 4, ProductType: model.GeneralSale, RefundNum: 4, IsReply: g.StatusFalse, ProductPrice: 352.00}},
	}
	return orderProducts
}

var orderStatus = []model.OrderStatus{
	{OrderID: 2, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderID: 2, ChangeType: "cancel", ChangeMessage: "取消订单[自动]", ChangeTime: time.Now()},
	{OrderID: 3, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderID: 3, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
	{OrderID: 4, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderID: 4, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
	{OrderID: 5, ChangeType: "create", ChangeMessage: "订单生成", ChangeTime: time.Now()},
	{OrderID: 5, ChangeType: "pay_success", ChangeMessage: "订单支付成功", ChangeTime: time.Now()},
}

//@description: orders 表数据初始化
func (a *order) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Order{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> orders 表的初始数据已存在!")
			return nil
		}
		if err := tx.Model(&model.GroupOrder{}).Create(&groupOrders).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Create(&orders).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		orderProducts := getOrderProducts()
		if err := tx.Model(&model.OrderProduct{}).Create(&orderProducts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.OrderStatus{}).Create(&orderStatus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> orders 表初始数据成功!")
		return nil
	})
}
