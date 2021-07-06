package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// GroupOrder 用户订单表
type GroupOrder struct {
	g.TENANCY_MODEL

	GroupOrderSn string  `gorm:"unique;column:group_order_sn;type:varchar(32);not null" json:"groupOrderSn"`                // 订单号
	TotalPostage float64 `gorm:"column:total_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPostage"` // 邮费
	TotalPrice   float64 `gorm:"column:total_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPrice"`     // 订单总额
	TotalNum     uint    `gorm:"column:total_num;type:int unsigned;not null;default:0" json:"totalNum"`                     // 商品数
	CouponPrice  float64 `gorm:"column:coupon_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"couponPrice"`   // 优惠金额
	RealName     string  `gorm:"column:real_name;type:varchar(32);not null" json:"realName"`                                // 联系人
	UserPhone    string  `gorm:"column:user_phone;type:varchar(18);not null" json:"userPhone"`                              // 联系电话
	UserAddress  string  `gorm:"column:user_address;type:varchar(128);not null" json:"userAddress"`                         // 收货地址
	PayPrice     float64 `gorm:"column:pay_price;type:decimal(8,2) unsigned;not null" json:"payPrice"`                      // 支付金额
	PayPostage   float64 `gorm:"column:pay_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"payPostage"`     // 支付邮费
	Cost         float64 `gorm:"column:cost;type:decimal(8,2) unsigned;not null" json:"cost"`                               // 成本价

	Paid     uint8     `gorm:"index:paid;column:paid;type:tinyint unsigned;not null;default:0" json:"paid"` // 是否支付
	PayTime  time.Time `gorm:"column:pay_time;type:timestamp" json:"payTime"`                               // 支付时间
	PayType  int       `gorm:"column:pay_type;type:tinyint(1);not null" json:"payType"`                     // 支付方式  1=微信 2=小程序 3=h5 4=余额  5=支付宝
	IsRemind uint8     `gorm:"column:is_remind;type:tinyint unsigned;not null;default:2" json:"isRemind"`
	// 是否提醒

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
