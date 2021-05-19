package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

// CreateAddress
func CreateAddress(m request.CreateAddress) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	err := g.TENANCY_DB.Create(&address).Error
	return address, err
}

// GetAddressByID
func GetAddressByID(id float64) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	err := g.TENANCY_DB.Where("id = ?", id).First(&address).Error
	return address, err
}

// UpdateAddress
func UpdateAddress(m request.UpdateAddress) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	err := g.TENANCY_DB.Updates(&address).Error
	return address, err
}

// DeleteAddress
func DeleteAddress(id float64) error {
	var address model.GeneralAddress
	return g.TENANCY_DB.Where("id = ?", id).Delete(&address).Error
}

// GetAddressInfoList
func GetAddressInfoList(info request.PageInfo) ([]response.GeneralAddress, int64, error) {
	var addressList []response.GeneralAddress
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.GeneralAddress{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return addressList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&addressList).Error
	return addressList, total, err
}
