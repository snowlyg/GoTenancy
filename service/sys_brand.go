package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// CreateBrand
func CreateBrand(m request.CreateSysBrand) (model.SysBrand, error) {
	var brand model.SysBrand
	err := g.TENANCY_DB.Where("brand_name = ?", m.BrandName).First(&brand).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brand, errors.New("名称已被注冊")
	}
	brand.BrandName = m.BrandName
	brand.Sort = m.Sort
	brand.Pic = m.Pic
	brand.IsShow = m.IsShow
	brand.BrandCategoryID = m.BrandCategoryID
	err = g.TENANCY_DB.Create(&brand).Error
	return brand, err
}

// GetBrandByID
func GetBrandByID(id uint) (model.SysBrand, error) {
	var brand model.SysBrand
	err := g.TENANCY_DB.Where("id = ?", id).First(&brand).Error
	return brand, err
}

// SetBrandCate
func SetBrandCate(setSysBrand request.SetSysBrand) error {
	return g.TENANCY_DB.Model(&model.SysBrand{}).Where("id = ?", setSysBrand.Id).Update("brand_category_id", setSysBrand.BrandCategoryID).Error
}

// UpdateBrand
func UpdateBrand(m request.UpdateSysBrand) (model.SysBrand, error) {
	var brand model.SysBrand
	err := g.TENANCY_DB.Where("brand_name = ?", m.BrandName).Where("id <> ?", m.Id).First(&brand).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brand, errors.New("名称已被注冊")
	}
	data := map[string]interface{}{
		"brand_name": m.BrandName,
		"sort":       m.Sort,
		"pic":        m.Pic,
		"is_show":    m.IsShow,
	}
	brand.ID = m.Id
	err = g.TENANCY_DB.Model(&brand).Updates(data).Error
	return brand, err
}

// DeleteBrand
func DeleteBrand(id uint) error {
	var brand model.SysBrand
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brand).Error
}

// GetBrandInfoList
func GetBrandInfoList(info request.PageInfo) ([]model.SysBrand, int64, error) {
	var brandList []model.SysBrand
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysBrand{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return brandList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&brandList).Error
	return brandList, total, err
}
