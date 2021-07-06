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
	{SysUserID: 3, GroupOrderSn: g.CreateOrderSn(1), TotalPostage: 20.00, TotalPrice: 50.00, TotalNum: 1, CouponPrice: 0, RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", PayPrice: 50.00, PayPostage: 30.00, Cost: 5.00, Paid: 0, PayTime: time.Now(), PayType: model.PayTypeWx, IsRemind: 1},
}
var orders = []model.Order{
	{SysUserID: 3, SysTenancyID: 1, GroupOrderID: 1, ReconciliationID: 0, CartID: 1, BaseOrder: model.BaseOrder{OrderSn: g.CreateOrderSn(1), RealName: "real_name", UserPhone: "user_phone", UserAddress: "user_address", TotalNum: 10, TotalPrice: 20.00, TotalPostage: 30.00, PayPrice: 50.00, PayPostage: 30.00, CommissionRate: 15.00, OrderType: model.OrderTypeGeneral, Paid: 0, PayTime: time.Now(), PayType: model.PayTypeWx, Status: model.OrderStatusNoDeliver, DeliveryType: model.DeliverTypeFH, DeliveryName: "delivery_name", DeliveryID: "delivery_id", Mark: "mark", Remark: "remark", AdminMark: "admin_mark", VerifyCode: "", VerifyTime: time.Now(), ActivityType: model.GeneralSale, Cost: 5.00, IsDel: g.StatusFalse}},
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
	}
	return orderProducts
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
		color.Info.Println("\n[Mysql] --> orders 表初始数据成功!")
		return nil
	})
}
