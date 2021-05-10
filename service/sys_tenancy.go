package service

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// CreateTenancy
func CreateTenancy(t model.SysTenancy) (error, model.SysTenancy) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", t.Name).First(&model.SysTenancy{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("商户名称已被注冊"), t
	}
	t.UUID = uuid.NewV4()
	err := g.TENANCY_DB.Create(&t).Error
	return err, t
}

// GetTenancyByID
func GetTenancyByID(id float64) (error, model.SysTenancy) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return err, tenancy
}

// UpdateTenany
func UpdateTenany(t model.SysTenancy) (error, model.SysTenancy) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", t.Name).Where("id != ?", t.ID).First(&model.SysTenancy{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("商户名称已被注冊"), t
	}
	err := g.TENANCY_DB.Updates(&t).Error
	return err, t
}

// DeleteTenancy
func DeleteTenancy(id float64) error {
	var tenancy model.SysTenancy
	return g.TENANCY_DB.Where("id = ?", id).Delete(&tenancy).Error
}

// GetTenanciesInfoList
func GetTenanciesInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	var tenancyList []model.SysTenancy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysTenancy{})
	err = db.Count(&total).Error
	if err != nil {
		return err, tenancyList, total
	}
	err = db.Limit(limit).Offset(offset).Find(&tenancyList).Error
	return err, tenancyList, total
}

// GetTenanciesByRegion
func GetTenanciesByRegion(p_code int) (error, []model.SysTenancy) {
	var tenancyList []model.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("sys_region_code = ?", p_code).Find(&tenancyList).Error
	return err, tenancyList
}
