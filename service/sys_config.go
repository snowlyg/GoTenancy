package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateConfig
func CreateConfig(m model.SysConfig) (model.SysConfig, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", m.Name).Where("type = ?", m.Type).First(&model.SysConfig{}).Error, gorm.ErrRecordNotFound) {
		return m, errors.New("设置名称已经使用")
	}
	err := g.TENANCY_DB.Create(&m).Error
	return m, err
}

// GetConfigByName
func GetConfigByName(name, style string) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("name = ?", name).Where("type = ?", style).First(&config).Error
	return config, err
}

// GetConfigInfoList
func GetConfigInfoList(info request.PageInfo) ([]response.SysConfig, int64, error) {
	var configList []response.SysConfig
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysConfig{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return configList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&configList).Error
	return configList, total, err
}

// UpdateConfig
func UpdateConfig(m model.SysConfig) (model.SysConfig, error) {
	if !errors.Is(g.TENANCY_DB.Where("name = ?", m.Name).Where("id <> ?", m.ID).Where("type = ?", m.Type).First(&model.SysConfig{}).Error, gorm.ErrRecordNotFound) {
		return m, errors.New("设置名称已经使用")
	}
	err := g.TENANCY_DB.Updates(&m).Error
	return m, err
}

// DeleteConfig
func DeleteConfig(id float64) error {
	var config model.SysConfig
	return g.TENANCY_DB.Where("id = ?", id).Delete(&config).Error
}
