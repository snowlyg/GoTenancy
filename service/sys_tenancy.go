package service

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateTenancy
func CreateTenancy(t request.CreateSysTenancy) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("name = ?", t.Name).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}
	tenancy.UUID = uuid.NewV4()
	tenancy.Address = t.Address
	tenancy.Tele = t.Tele
	tenancy.Name = t.Name
	tenancy.BusinessTime = t.BusinessTime
	tenancy.SysRegionCode = t.SysRegionCode
	err = g.TENANCY_DB.Create(&tenancy).Error
	return tenancy, err
}

// GetTenancyByID
func GetTenancyByID(id float64) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return tenancy, err
}

// SetTenancyRegionByID
func SetTenancyRegionByID(id float64, sysRegionCode int) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", id).Update("sys_region_code", sysRegionCode).Error
}

// UpdateTenany
func UpdateTenany(t request.UpdateSysTenancy) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("name = ?", t.Name).Not("id = ?", t.Id).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}

	tenancy.ID = t.Id
	tenancy.Address = t.Address
	tenancy.Tele = t.Tele
	tenancy.Name = t.Name
	tenancy.BusinessTime = t.BusinessTime
	tenancy.SysRegionCode = t.SysRegionCode
	err = g.TENANCY_DB.Updates(&tenancy).Error
	return tenancy, err
}

// DeleteTenancy
func DeleteTenancy(id float64) error {
	var tenancy model.SysTenancy
	return g.TENANCY_DB.Where("id = ?", id).Delete(&tenancy).Error
}

// GetTenanciesInfoList
func GetTenanciesInfoList(info request.PageInfo) ([]response.SysTenancy, int64, error) {
	var tenancyList []response.SysTenancy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysTenancy{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}

// GetTenanciesByRegion
func GetTenanciesByRegion(p_code string) ([]response.SysTenancy, error) {
	var tenancyList []response.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("sys_region_code = ?", p_code).Find(&tenancyList).Error
	return tenancyList, err
}
