package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var UserBill = new(userBill)

type userBill struct{}

var userBills = []model.UserBill{
	{SysUserID: 3, LinkID: 2, Pm: 0, Title: "购买商品", Category: "now_money", Type: "pay_product", Number: 89.00, Balance: 999910.99, Mark: "余额支付支付89元购买商品", Status: 1},
	{SysUserID: 3, LinkID: 3, Pm: 0, Title: "购买商品", Category: "now_money", Type: "pay_product", Number: 89.00, Balance: 999821.99, Mark: "余额支付支付89元购买商品", Status: 1},
	{SysUserID: 3, LinkID: 4, Pm: 0, Title: "购买商品", Category: "now_money", Type: "pay_product", Number: 356.00, Balance: 999465.99, Mark: "余额支付支付356元购买商品", Status: 1},
	{SysUserID: 3, LinkID: 1, Pm: 1, Title: "退款增加余额", Category: "now_money", Type: "refund", Number: 89.00, Balance: 999554.99, Mark: "退款增加89余额，退款订单号:wx1625619291308852720", Status: 1},
}

func (m *userBill) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.UserBill{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> user_bills 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&userBills).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> user_bills 表初始数据成功!")
		return nil
	})
}
