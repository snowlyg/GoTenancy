package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"

	"gorm.io/gorm"
)

// CreateApi 新增基础api
func CreateApi(api model.SysApi) (err error) {
	if !errors.Is(g.TENANCY_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同api")
	}
	return g.TENANCY_DB.Create(&api).Error
}

// DeleteApi 删除基础api
func DeleteApi(api model.SysApi) (err error) {
	err = g.TENANCY_DB.Delete(&api).Error
	ClearCasbin(1, api.Path, api.Method)
	return err
}

// GetAPIInfoList 分页获取数据
func GetAPIInfoList(api model.SysApi, info request.PageInfo, order string, desc bool) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysApi{})
	var apiList []model.SysApi

	if api.Path != "" {
		db = db.Where("path LIKE ?", "%"+api.Path+"%")
	}

	if api.Description != "" {
		db = db.Where("description LIKE ?", "%"+api.Description+"%")
	}

	if api.Method != "" {
		db = db.Where("method = ?", api.Method)
	}

	if api.ApiGroup != "" {
		db = db.Where("api_group = ?", api.ApiGroup)
	}

	err = db.Count(&total).Error

	if err != nil {
		return err, apiList, total
	} else {
		db = db.Limit(limit).Offset(offset)
		if order != "" {
			var OrderStr string
			if desc {
				OrderStr = order + " desc"
			} else {
				OrderStr = order
			}
			err = db.Order(OrderStr).Find(&apiList).Error
		} else {
			err = db.Order("api_group").Find(&apiList).Error
		}
	}
	return err, apiList, total
}

// GetAllApis 获取所有的api
func GetAllApis() (err error, apis []model.SysApi) {
	err = g.TENANCY_DB.Find(&apis).Error
	return
}

// GetApiById 根据id获取api
func GetApiById(id float64) (err error, api model.SysApi) {
	err = g.TENANCY_DB.Where("id = ?", id).First(&api).Error
	return
}

// UpdateApi 根据id更新api
func UpdateApi(api model.SysApi) (err error) {
	var oldA model.SysApi
	err = g.TENANCY_DB.Where("id = ?", api.ID).First(&oldA).Error
	if oldA.Path != api.Path || oldA.Method != api.Method {
		if !errors.Is(g.TENANCY_DB.Where("path = ? AND method = ?", api.Path, api.Method).First(&model.SysApi{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同api路径")
		}
	}
	if err != nil {
		return err
	} else {
		err = UpdateCasbinApi(oldA.Path, api.Path, oldA.Method, api.Method)
		if err != nil {
			return err
		} else {
			err = g.TENANCY_DB.Save(&api).Error
		}
	}
	return err
}

// DeleteApisByIds 删除选中API
func DeleteApisByIds(ids request.IdsReq) (err error) {
	err = g.TENANCY_DB.Delete(&[]model.SysApi{}, "id in ?", ids.Ids).Error
	return err
}
