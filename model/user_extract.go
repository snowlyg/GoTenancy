package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// UserExtract 用户提现表
type UserExtract struct {
	g.TENANCY_MODEL

	RealName     string    `gorm:"column:real_name;type:varchar(64)" json:"realName"`                                      // 姓名
	ExtractType  int       `gorm:"column:extract_type;type:tinyint(1);default:0" json:"extractType"`                       // 0 银行卡 1 支付宝 2微信
	BankCode     string    `gorm:"column:bank_code;type:varchar(32);default:0" json:"bankCode"`                            // 银行卡
	BankAddress  string    `gorm:"column:bank_address;type:varchar(256);default:''" json:"bankAddress"`                    // 开户地址
	AlipayCode   string    `gorm:"column:alipay_code;type:varchar(64);default:''" json:"alipayCode"`                       // 支付宝账号
	Wechat       string    `gorm:"column:wechat;type:varchar(15)" json:"wechat"`                                           // 微信号
	ExtractPic   string    `gorm:"column:extract_pic;type:varchar(128)" json:"extractPic"`                                 // 收款码
	ExtractPrice float64   `gorm:"column:extract_price;type:decimal(8,2) unsigned;default:0.00" json:"extractPrice"`       // 提现金额
	Balance      float64   `gorm:"column:balance;type:decimal(8,2) unsigned;default:0.00" json:"balance"`                  // 余额
	Mark         string    `gorm:"column:mark;type:varchar(512)" json:"mark"`                                              // 管理员备注
	AdminID      int       `gorm:"column:admin_id;type:int;default:0" json:"adminId"`                                      // 审核管理员
	FailMsg      string    `gorm:"column:fail_msg;type:varchar(128)" json:"failMsg"`                                       // 无效原因
	StatusTime   time.Time `gorm:"column:status_time;type:timestamp;not null;default:CURRENT_TIMESTAMP" json:"statusTime"` // 无效时间
	Status       int       `gorm:"column:status;type:tinyint;default:2" json:"status"`                                     // 1 审核中 2 已提现 3 未通过

	SysUserID uint `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
