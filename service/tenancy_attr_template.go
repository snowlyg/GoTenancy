package service

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// CreateAttrTemplate
func CreateAttrTemplate(brandAttrTemplate model.AttrTemplate, ctx *gin.Context) (model.AttrTemplate, error) {
	brandAttrTemplate.SysTenancyID = multi.GetTenancyId(ctx)
	err := g.TENANCY_DB.Create(&brandAttrTemplate).Error
	return brandAttrTemplate, err
}

// GetAttrTemplateByID
func GetAttrTemplateByID(id uint) (model.AttrTemplate, error) {
	var brandAttrTemplate model.AttrTemplate
	err := g.TENANCY_DB.Where("id = ?", id).First(&brandAttrTemplate).Error
	return brandAttrTemplate, err
}

// UpdateAttrTemplate
func UpdateAttrTemplate(brandAttrTemplate model.AttrTemplate) (model.AttrTemplate, error) {
	err := g.TENANCY_DB.Model(&brandAttrTemplate).Updates(brandAttrTemplate).Error
	return brandAttrTemplate, err
}

// DeleteAttrTemplate
func DeleteAttrTemplate(id uint) error {
	var brandAttrTemplate model.AttrTemplate
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brandAttrTemplate).Error
}

// GetAttrTemplateInfoList
func GetAttrTemplateInfoList(info request.PageInfo) ([]response.AttrTemplate, int64, error) {
	var brandAttrTemplateList []response.AttrTemplate
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.AttrTemplate{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return brandAttrTemplateList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&brandAttrTemplateList).Error
	return brandAttrTemplateList, total, err
}
