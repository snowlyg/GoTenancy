package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// CreateReceipt
func CreateReceipt(m request.CreateReceipt, user_id int) (model.GeneralReceipt, error) {
	var receipt model.GeneralReceipt
	receipt.ReceiptType = m.ReceiptType
	receipt.ReceiptTitle = m.ReceiptTitle
	receipt.ReceiptTitleType = m.ReceiptTitleType
	receipt.DutyGaragraph = m.DutyGaragraph
	receipt.Email = m.Email
	receipt.BankName = m.BankName
	receipt.BankCode = m.BankCode
	receipt.Address = m.Address
	receipt.Tel = m.Tel
	receipt.IsDefault = m.IsDefault
	receipt.SysUserID = user_id
	err := g.TENANCY_DB.Create(&receipt).Error
	return receipt, err
}

// GetReceiptByID
func GetReceiptByID(id float64, user_id int) (model.GeneralReceipt, error) {
	var receipt model.GeneralReceipt
	err := g.TENANCY_DB.Where("id = ?", id).Where("sys_user_id = ?", user_id).First(&receipt).Error
	return receipt, err
}

// UpdateReceipt
func UpdateReceipt(m request.UpdateReceipt) (model.GeneralReceipt, error) {
	data := map[string]interface{}{
		"receipt_type":       m.ReceiptType,
		"receipt_title":      m.ReceiptTitle,
		"receipt_title_type": m.ReceiptTitleType,
		"duty_garagraph":     m.DutyGaragraph,
		"email":              m.Email,
		"bank_name":          m.BankName,
		"bank_code":          m.BankCode,
		"address":            m.Address,
		"tel":                m.Tel,
		"is_default":         m.IsDefault,
	}
	receipt := model.GeneralReceipt{TENANCY_MODEL: g.TENANCY_MODEL{ID: m.Id}}
	err := g.TENANCY_DB.Model(&receipt).Updates(data).Error
	return receipt, err
}

// DeleteReceipt
func DeleteReceipt(id float64, user_id int) error {
	var receipt model.GeneralReceipt
	return g.TENANCY_DB.Where("id = ?", id).Where("sys_user_id = ?", user_id).Delete(&receipt).Error
}

// GetReceiptInfoList
func GetReceiptInfoList(info request.PageInfo, user_id int) ([]model.GeneralReceipt, int64, error) {
	var receiptList []model.GeneralReceipt
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.GeneralReceipt{}).Where("sys_user_id = ?", user_id)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return receiptList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&receiptList).Error
	return receiptList, total, err
}
