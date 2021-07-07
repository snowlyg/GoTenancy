package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var RefundOrder = new(refundOrder)

type refundOrder struct{}

var refundOrders = []model.RefundOrder{
	{RefundOrderSn: g.CreateRefundOrderSn(), OrderID: 1, SysUserID: 3, SysTenancyID: 1, RefundType: 1, RefundMessage: "收货地址填错了", RefundPrice: 89.00, RefundNum: 1, Status: model.RefundStatusAudit, StatusTime: time.Now(), IsDel: g.StatusFalse, IsSystemDel: g.StatusFalse},
}

var refundProducts = []model.RefundProduct{
	{RefundOrderID: 1, OrderProductID: 1, RefundNum: 1},
}

var refundStatus = []model.RefundStatus{
	{RefundOrderID: 1, ChangeType: "create", ChangeMessage: "创建退款单", ChangeTime: time.Now()},
}

//@description: refundOrders 表数据初始化
func (a *refundOrder) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.RefundOrder{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> refund_orders 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&refundOrders).Error; err != nil { // 遇到错误时回滚事务
			return err
		}

		if err := tx.Model(&model.RefundProduct{}).Create(&refundProducts).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		if err := tx.Model(&model.RefundStatus{}).Create(&refundStatus).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> refund_orders 表初始数据成功!")
		return nil
	})
}
