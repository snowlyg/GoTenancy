package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// CreateSysDictionary 创建字典数据
func CreateSysDictionary(sysDictionary model.SysDictionary) error {
	if (!errors.Is(g.TENANCY_DB.First(&model.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	return g.TENANCY_DB.Create(&sysDictionary).Error
}

// DeleteSysDictionary 删除字典数据
func DeleteSysDictionary(sysDictionary model.SysDictionary) error {
	return g.TENANCY_DB.Delete(&sysDictionary).Delete(&sysDictionary.SysDictionaryDetails).Error
}

// UpdateSysDictionary 更新字典数据
func UpdateSysDictionary(sysDictionary *model.SysDictionary) error {
	var err error
	var dict model.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Type":   sysDictionary.Type,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := g.TENANCY_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Type == sysDictionary.Type {
		err = db.Updates(sysDictionaryMap).Error
	} else {
		if (!errors.Is(g.TENANCY_DB.First(&model.SysDictionary{}, "type = ?", sysDictionary.Type).Error, gorm.ErrRecordNotFound)) {
			return errors.New("存在相同的type，不允许创建")
		}
		err = db.Updates(sysDictionaryMap).Error

	}
	return err
}

// GetSysDictionary 根据id或者type获取字典单条数据
func GetSysDictionary(Type string, Id uint) (model.SysDictionary, error) {
	var sysDictionary model.SysDictionary
	err := g.TENANCY_DB.Where("type = ? OR id = ?", Type, Id).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	return sysDictionary, err
}

// GetSysDictionaryInfoList 分页获取字典列表
func GetSysDictionaryInfoList(info request.SysDictionarySearch) ([]model.SysDictionary, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := g.TENANCY_DB.Model(&model.SysDictionary{})
	var sysDictionarys []model.SysDictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Type != "" {
		db = db.Where("`type` LIKE ?", "%"+info.Type+"%")
	}
	if info.Status != nil {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return sysDictionarys, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return sysDictionarys, total, err
}
