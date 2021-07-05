package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// StoreOrder 订单表
type StoreOrder struct {
	g.TENANCY_MODEL

	// 订单组 id
	OrderSn        string    `gorm:"column:order_sn;type:varchar(32);not null" json:"orderSn"`                                        // 订单号
	RealName       string    `gorm:"column:real_name;type:varchar(32);not null" json:"realName"`                                      // 用户姓名
	UserPhone      string    `gorm:"column:user_phone;type:varchar(18);not null" json:"userPhone"`                                    // 用户电话
	UserAddress    string    `gorm:"column:user_address;type:varchar(128);not null" json:"userAddress"`                               // 详细地址
	TotalNum       uint      `gorm:"column:total_num;type:int unsigned;not null;default:0" json:"totalNum"`                           // 订单商品总数
	TotalPrice     float64   `gorm:"column:total_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPrice"`           // 订单总价
	TotalPostage   float64   `gorm:"column:total_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"totalPostage"`       // 邮费
	PayPrice       float64   `gorm:"column:pay_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"payPrice"`               // 实际支付金额
	PayPostage     float64   `gorm:"column:pay_postage;type:decimal(8,2) unsigned;not null;default:0.00" json:"payPostage"`           // 支付邮费
	ExtensionOne   float64   `gorm:"column:extension_one;type:decimal(8,2) unsigned;not null;default:0.00" json:"extensionOne"`       // 一级佣金
	ExtensionTwo   float64   `gorm:"column:extension_two;type:decimal(8,2) unsigned;not null;default:0.00" json:"extensionTwo"`       // 二级佣金
	CommissionRate float64   `gorm:"column:commission_rate;type:decimal(6,4) unsigned;not null;default:0.0000" json:"commissionRate"` // 平台手续费
	CouponPrice    float64   `gorm:"column:coupon_price;type:decimal(8,2) unsigned;not null;default:0.00" json:"couponPrice"`         // 优惠券金额
	OrderType      uint8     `gorm:"column:order_type;type:tinyint unsigned;default:0" json:"orderType"`                              // 0普通1自提
	Paid           uint8     `gorm:"column:paid;type:tinyint unsigned;not null;default:0" json:"paid"`                                // 支付状态
	PayTime        time.Time `gorm:"column:pay_time;type:timestamp" json:"payTime"`                                                   // 支付时间
	PayType        bool      `gorm:"column:pay_type;type:tinyint(1);not null" json:"payType"`                                         // 支付方式  1=微信 2=小程序 4=余额 3=h5
	CreateTime     time.Time `gorm:"column:create_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"createTime"`          // 创建时间
	Status         bool      `gorm:"column:status;type:tinyint(1);not null;default:0" json:"status"`                                  // 订单状态（4：待发货；1：待收货；2：待评价；3：已完成；5：已退款）
	DeliveryType   string    `gorm:"column:delivery_type;type:varchar(32)" json:"deliveryType"`                                       // 发货类型(1:发货 2: 送货 3: 虚拟)
	DeliveryName   string    `gorm:"column:delivery_name;type:varchar(64)" json:"deliveryName"`                                       // 快递名称/送货人姓名
	DeliveryID     string    `gorm:"column:delivery_id;type:varchar(64)" json:"deliveryId"`                                           // 快递单号/手机号
	Mark           string    `gorm:"column:mark;type:varchar(512);not null" json:"mark"`                                              // 备注
	Remark         string    `gorm:"column:remark;type:varchar(512)" json:"remark"`                                                   // 管理员备注
	AdminMark      string    `gorm:"column:admin_mark;type:varchar(512)" json:"adminMark"`                                            // 总后台备注
	VerifyCode     string    `gorm:"index:verify_code;column:verify_code;type:char(16)" json:"verifyCode"`                            // 核销码
	VerifyTime     time.Time `gorm:"column:verify_time;type:timestamp" json:"verifyTime"`                                             // 核销时间
	ActivityType   uint8     `gorm:"column:activity_type;type:tinyint unsigned;not null;default:0" json:"activityType"`               // 1:秒杀 2:预售 3:助力
	Cost           float64   `gorm:"column:cost;type:decimal(8,2) unsigned;not null" json:"cost"`                                     // 成本价
	IsDel          uint8     `gorm:"column:is_del;type:tinyint unsigned;not null;default:0" json:"isDel"`                             // 是否删除

	SysUserID        uint   `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
	SysTenancyID     uint   `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	GroupOrderID     int    `gorm:"column:group_order_id;type:int" json:"groupOrderId"`
	ReconciliationID uint8  `gorm:"column:reconciliation_id;type:tinyint unsigned;not null;default:0" json:"reconciliationId"` // 对账id
	VerifyServiceID  uint   `gorm:"column:verify_service_id;type:int unsigned;default:0" json:"verifyServiceId"`               // 核销客服 id
	CartID           string `gorm:"column:cart_id;type:varchar(256);not null" json:"cartId"`                                   // 购物车id
	// CouponID       string    `gorm:"column:coupon_id;type:varchar(128);not null;default:''" json:"couponId"`                          // 优惠券id
}
