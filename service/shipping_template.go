package service

import (
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateShippingTemplate
func CreateShippingTemplate(shippingTem model.ShippingTemplate, tenancyId uint) (model.ShippingTemplate, error) {
	shippingTem.SysTenancyID = tenancyId
	err := g.TENANCY_DB.Create(&shippingTem).Error
	return shippingTem, err
}

// GetShippingTemplateByID
func GetShippingTemplateByID(id uint) (response.ShippingTemplateDetail, error) {
	var shippingTem response.ShippingTemplateDetail
	err := g.TENANCY_DB.Model(&model.ShippingTemplate{}).
		Where("id = ?", id).
		First(&shippingTem).Error
	return shippingTem, err
}

// UpdateShippingTemplate
func UpdateShippingTemplate(req request.UpdateShippingTemplate, id uint) error {
	shippingTem := model.ShippingTemplate{BaseShippingTemplate: model.BaseShippingTemplate{}}
	err := g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Updates(&shippingTem).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteShippingTemplate
func DeleteShippingTemplate(id uint) error {
	var product model.ShippingTemplate
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// GetShippingTemplateInfoList
func GetShippingTemplateInfoList(info request.ShippingTemplatePageInfo) ([]response.ShippingTemplateList, int64, error) {
	var shippingTemList []response.ShippingTemplateList
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.ShippingTemplate{})

	if info.Name != "" {
		db = db.Where("name like ?", fmt.Sprintf("%s%%", info.Name))
	}

	err := db.Count(&total).Error
	if err != nil {
		return shippingTemList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&shippingTemList).Error

	return shippingTemList, total, err
}

// GetShippingTemplateInfoSelect
func GetShippingTemplateInfoSelect() ([]response.ShippingTemplateSelect, error) {
	var shippingTemList []response.ShippingTemplateSelect
	err := g.TENANCY_DB.Model(&model.ShippingTemplate{}).Select("id,name").Find(&shippingTemList).Error
	if err != nil {
		return nil, err
	}
	return shippingTemList, err
}
