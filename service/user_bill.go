package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// GetUserBillInfoList
func GetUserBillInfoList(info request.PageInfo, sys_user_id uint) ([]model.UserBill, int64, error) {
	var userBillList []model.UserBill
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.UserBill{}).Where("sys_user_id = ?", sys_user_id)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return userBillList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userBillList).Error
	return userBillList, total, err
}
