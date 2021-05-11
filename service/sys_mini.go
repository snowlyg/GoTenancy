package service

import (
	"errors"

	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// CreateMini
func CreateMini(m model.SysMini) (model.SysMini, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", m.Name).First(&model.SysMini{}).Error, gorm.ErrRecordNotFound) {
		return m, errors.New("商户名称已被注冊")
	}
	m.UUID = uuid.NewV4()
	err := g.TENANCY_DB.Create(&m).Error
	return m, err
}

// GetMiniByID
func GetMiniByID(id float64) (model.SysMini, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("id = ?", id).First(&mini).Error
	return mini, err
}

// UpdateMini
func UpdateMini(m model.SysMini) (model.SysMini, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", m.Name).Where("id <> ?", m.ID).First(&model.SysTenancy{}).Error, gorm.ErrRecordNotFound) {
		return m, errors.New("商户名称已被注冊")
	}
	err := g.TENANCY_DB.Updates(&m).Error
	return m, err
}

// DeleteMini
func DeleteMini(id float64) error {
	var mini model.SysMini
	return g.TENANCY_DB.Where("id = ?", id).Delete(&mini).Error
}

// GetMiniInfoList
func GetMiniInfoList(info request.PageInfo) ([]model.SysMini, int64, error) {
	var miniList []model.SysMini
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysMini{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return miniList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&miniList).Error
	return miniList, total, err
}
