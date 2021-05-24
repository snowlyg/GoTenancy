package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

// GeneralReceipt 发票
type GeneralReceipt struct {
	g.TENANCY_MODEL
	ReceiptType      int    `json:"receiptType" form:"receiptType" gorm:"column:receipt_type;default:0;comment:发票类型：1.普通，2.增值"`
	ReceiptTitle     string `json:"receiptTitle" form:"receiptTitle" gorm:"column:receipt_title;comment:发票抬头"`
	ReceiptTitleType int    `json:"receiptTitleType" form:"receiptTitleType" gorm:"column:receipt_title_type;comment:发票抬头类型：1.个人，2.企业"`
	DutyGaragraph    string `json:"dutyGaragraph" form:"dutyGaragraph" gorm:"column:duty_garagraph;comment:税号"`
	Email            string `json:"email" form:"email" gorm:"column:email;comment:邮箱"`
	BankName         string `json:"backName" form:"backName" gorm:"column:back_name;comment:开户行"`
	BankCode         string `json:"backCode" form:"backCode" gorm:"column:back_code;comment:银行账号"`
	Address          string `json:"address" form:"address" gorm:"column:address;comment:企业地址"`
	Tel              string `json:"tel" form:"tel" gorm:"column:tel;comment:企业电话"`
	IsDefault        int    `json:"isDefault" form:"isDefault" gorm:"column:is_default;comment:是否默认"`

	SysUserID int `json:"sysUserId" form:"sysUserId" gorm:"column:sys_user_id;comment:关联标记"`
}
