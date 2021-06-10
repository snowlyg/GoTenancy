package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"

	"gorm.io/gorm"
)

// CreateApi 新增基础api
func CreateApi(api model.SysApi) (model.SysApi, error) {
	err := g.TENANCY_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return model.SysApi{}, errors.New("存在相同api")
	}
	err = g.TENANCY_DB.Create(&api).Error
	return api, err
}

// DeleteApi 删除基础api
func DeleteApi(req request.DeleteApi) error {
	err := g.TENANCY_DB.Delete(&model.SysApi{}, req.Id).Error
	ClearCasbin(1, req.Path, req.Method)
	return err
}

// GetAPIInfoList 分页获取数据
func GetAPIInfoList(pageInfo request.SearchApiParams) ([]model.SysApi, int64, error) {
	limit := pageInfo.PageSize
	offset := pageInfo.PageSize * (pageInfo.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if pageInfo.Path != "" {
		db = db.Where("path LIKE ?", "%"+pageInfo.Path+"%")
	}

	if pageInfo.Description != "" {
		db = db.Where("description LIKE ?", "%"+pageInfo.Description+"%")
	}

	if pageInfo.Method != "" {
		db = db.Where("method = ?", pageInfo.Method)
	}

	if pageInfo.ApiGroup != "" {
		db = db.Where("api_group = ?", pageInfo.ApiGroup)
	}
	var total int64
	err := db.Count(&total).Error

	if err != nil {
		return apiList, total, err
	} else {
		db = db.Limit(limit).Offset(offset)
		if pageInfo.OrderKey != "" {
			var OrderStr string
			if pageInfo.Desc {
				OrderStr = pageInfo.OrderKey + " desc"
			} else {
				OrderStr = pageInfo.OrderKey
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return apiList, total, err
}

// GetAllApis 获取所有的api
func GetAllApis() ([]model.SysApi, error) {
	var apis []model.SysApi
	err := g.TENANCY_DB.Find(&apis).Error
	return apis, err
}

// GetApiById 根据id获取api
func GetApiById(id uint) (model.SysApi, error) {
	var api model.SysApi
	err := g.TENANCY_DB.Where("id = ?", id).First(&api).Error
	return api, err
}

// UpdateApi 根据id更新api
func UpdateApi(api model.SysApi) error {
	var oldA model.SysApi
	err := g.TENANCY_DB.Where("id = ?", api.ID).First(&oldA).Error
	if err != nil {
		return err
	}
	api.CreatedAt = oldA.CreatedAt
	if oldA.Path != api.Path || oldA.Method != api.Method {
		err := g.TENANCY_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}

	err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
	if err != nil {
		return err
	} else {
		err = g.TENANCY_DB.Save(&api).Error
	}

	return err
}

// DeleteApisByIds 删除选中API
func DeleteApisByIds(ids request.IdsReq) error {
	return g.TENANCY_DB.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
}
