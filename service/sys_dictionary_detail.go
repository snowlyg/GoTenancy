package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// CreateSysDictionaryDetail 创建字典详情数据
func CreateSysDictionaryDetail(sysDictionaryDetail model.SysDictionaryDetail) (err error) {
	err = g.TENANCY_DB.Create(&sysDictionaryDetail).Error
	return err
}

// DeleteSysDictionaryDetail 删除字典详情数据
func DeleteSysDictionaryDetail(sysDictionaryDetail model.SysDictionaryDetail) (err error) {
	err = g.TENANCY_DB.Delete(&sysDictionaryDetail).Error
	return err
}

//UpdateSysDictionaryDetail 更新字典详情数据
func UpdateSysDictionaryDetail(sysDictionaryDetail *model.SysDictionaryDetail) (err error) {
	err = g.TENANCY_DB.Save(sysDictionaryDetail).Error
	return err
}

// GetSysDictionaryDetail 根据id获取字典详情单条数据
func GetSysDictionaryDetail(id uint) (err error, sysDictionaryDetail model.SysDictionaryDetail) {
	err = g.TENANCY_DB.Where("id = ?", id).First(&sysDictionaryDetail).Error
	return
}

//GetSysDictionaryDetailInfoList 分页获取字典详情列表
func GetSysDictionaryDetailInfoList(info request.SysDictionaryDetailSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := g.TENANCY_DB.Model(&model.SysDictionaryDetail{})
	var sysDictionaryDetails []model.SysDictionaryDetail
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Label != "" {
		db = db.Where("label LIKE ?", "%"+info.Label+"%")
	}
	if info.Value != 0 {
		db = db.Where("value = ?", info.Value)
	}
	if info.Status != nil {
		db = db.Where("status = ?", info.Status)
	}
	if info.SysDictionaryID != 0 {
		db = db.Where("sys_dictionary_id = ?", info.SysDictionaryID)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&sysDictionaryDetails).Error
	return err, sysDictionaryDetails, total
}
