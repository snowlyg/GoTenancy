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
func CreateTenancy(t model.SysTenancy) (model.SysTenancy, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", t.Name).First(&model.SysTenancy{}).Error, gorm.ErrRecordNotFound) {
		return t, errors.New("商户名称已被注冊")
	}
	t.UUID = uuid.NewV4()
	err := g.TENANCY_DB.Create(&t).Error
	return t, err
}

// GetTenancyByID
func GetTenancyByID(id float64) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return tenancy, err
}

// UpdateTenany
func UpdateTenany(t model.SysTenancy) (model.SysTenancy, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", t.Name).Where("id <> ?", t.ID).First(&model.SysTenancy{}).Error, gorm.ErrRecordNotFound) {
		return t, errors.New("商户名称已被注冊")
	}
	err := g.TENANCY_DB.Updates(&t).Error
	return t, err
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
func GetTenanciesByRegion(p_code int) ([]response.SysTenancy, error) {
	var tenancyList []response.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("sys_region_code = ?", p_code).Find(&tenancyList).Error
	return tenancyList, err
}
