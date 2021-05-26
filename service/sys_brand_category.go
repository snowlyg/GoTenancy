package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// CreateBrandCategory
func CreateBrandCategory(m request.CreateSysBrandCategory) (model.SysBrandCategory, error) {
	var brandCategory model.SysBrandCategory
	err := g.TENANCY_DB.Where("cate_name = ?", m.CateName).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}
	brandCategory.CateName = m.CateName
	brandCategory.Sort = m.Sort
	brandCategory.Path = m.Path
	brandCategory.IsShow = m.IsShow
	brandCategory.Level = m.Level
	brandCategory.Pid = m.Pid
	err = g.TENANCY_DB.Create(&brandCategory).Error
	return brandCategory, err
}

// GetBrandCategoryByID
func GetBrandCategoryByID(id float64) (model.SysBrandCategory, error) {
	var brandCategory model.SysBrandCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&brandCategory).Error
	return brandCategory, err
}

// UpdateBrandCategory
func UpdateBrandCategory(m request.UpdateSysBrandCategory) (model.SysBrandCategory, error) {
	var brandCategory model.SysBrandCategory
	err := g.TENANCY_DB.Where("cate_name = ?", m.CateName).Where("id <> ?", m.Id).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}
	data := map[string]interface{}{
		"pid":       m.Pid,
		"cate_name": m.CateName,
		"sort":      m.Sort,
		"path":      m.Path,
		"is_show":   m.IsShow,
		"level":     m.Level,
	}
	brandCategory.ID = m.Id
	err = g.TENANCY_DB.Model(&brandCategory).Updates(data).Error
	return brandCategory, err
}

// DeleteBrandCategory
func DeleteBrandCategory(id float64) error {
	var brandCategory model.SysBrandCategory
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brandCategory).Error
}

// GetBrandCategoryInfoList
func GetBrandCategoryInfoList(info request.PageInfo) ([]model.SysBrandCategory, int64, error) {
	var brandCategoryList []model.SysBrandCategory
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysBrandCategory{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return brandCategoryList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&brandCategoryList).Error
	return brandCategoryList, total, err
}
