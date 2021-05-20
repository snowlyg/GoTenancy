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

// CreateMini
func CreateMini(m request.CreateSysMini) (model.SysMini, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("name = ?", m.Name).First(&mini).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mini, errors.New("商户名称已被注冊")
	}
	mini.UUID = uuid.NewV4()
	mini.Name = m.Name
	mini.AppID = m.AppID
	mini.AppSecret = m.AppSecret
	mini.Remark = m.Remark
	err = g.TENANCY_DB.Create(&mini).Error
	return mini, err
}

// GetMiniByID
func GetMiniByID(id float64) (model.SysMini, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("id = ?", id).First(&mini).Error
	return mini, err
}

// UpdateMini
func UpdateMini(m request.UpdateSysMini) (model.SysMini, error) {
	var mini model.SysMini
	err := g.TENANCY_DB.Where("name = ?", m.Name).Where("id <> ?", m.Id).First(&mini).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return mini, errors.New("商户名称已被注冊")
	}
	mini.ID = m.Id
	mini.Name = m.Name
	mini.AppID = m.AppID
	mini.AppSecret = m.AppSecret
	mini.Remark = m.Remark
	err = g.TENANCY_DB.Updates(&mini).Error
	return mini, err
}

// DeleteMini
func DeleteMini(id float64) error {
	var mini model.SysMini
	return g.TENANCY_DB.Where("id = ?", id).Delete(&mini).Error
}

// GetMiniInfoList
func GetMiniInfoList(info request.PageInfo) ([]response.SysMini, int64, error) {
	var miniList []response.SysMini
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
