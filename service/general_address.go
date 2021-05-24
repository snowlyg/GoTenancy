package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

// CreateAddress
func CreateAddress(m request.CreateAddress, user_id int) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	address.Name = m.Name
	address.Phone = m.Phone
	address.Sex = m.Sex
	address.Country = m.Country
	address.Province = m.Province
	address.City = m.City
	address.District = m.District
	address.IsDefault = m.IsDefault
	address.Detail = m.Detail
	address.Postcode = m.Postcode
	address.Age = m.Age
	address.HospitalName = m.HospitalName
	address.LocName = m.LocName
	address.BedNum = m.BedNum
	address.HospitalNO = m.HospitalNO
	address.Disease = m.Disease
	address.SysUserID = user_id
	err := g.TENANCY_DB.Create(&address).Error
	return address, err
}

// GetAddressByID
func GetAddressByID(id float64, user_id int) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	err := g.TENANCY_DB.Where("id = ?", id).Where("sys_user_id = ?", user_id).First(&address).Error
	return address, err
}

// UpdateAddress
func UpdateAddress(m request.UpdateAddress) (model.GeneralAddress, error) {
	var address model.GeneralAddress
	address.ID = m.Id
	address.Name = m.Name
	address.Phone = m.Phone
	address.Sex = m.Sex
	address.Country = m.Country
	address.Province = m.Province
	address.City = m.City
	address.District = m.District
	address.IsDefault = m.IsDefault
	address.Detail = m.Detail
	address.Postcode = m.Postcode
	address.Age = m.Age
	address.HospitalName = m.HospitalName
	address.LocName = m.LocName
	address.BedNum = m.BedNum
	address.HospitalNO = m.HospitalNO
	address.Disease = m.Disease
	err := g.TENANCY_DB.Updates(&address).Error
	return address, err
}

// DeleteAddress
func DeleteAddress(id float64, user_id int) error {
	var address model.GeneralAddress
	return g.TENANCY_DB.Where("id = ?", id).Where("sys_user_id = ?", user_id).Delete(&address).Error
}

// GetAddressInfoList
func GetAddressInfoList(info request.PageInfo, user_id int) ([]response.GeneralAddress, int64, error) {
	var addressList []response.GeneralAddress
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.GeneralAddress{}).Where("sys_user_id = ?", user_id)
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return addressList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&addressList).Error
	return addressList, total, err
}
