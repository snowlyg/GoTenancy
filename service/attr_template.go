package service

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// CreateAttrTemplate
func CreateAttrTemplate(req request.AttrTemplate, ctx *gin.Context) (model.AttrTemplate, error) {
	attrTemplate := model.AttrTemplate{
		TemplateName: req.TemplateName,
		SysTenancyID: multi.GetTenancyId(ctx),
	}
	value, err := json.Marshal(&req.TemplateValue)
	if err != nil {
		return attrTemplate, err
	}
	attrTemplate.TemplateValue = value
	err = g.TENANCY_DB.Create(&attrTemplate).Error
	return attrTemplate, err
}

// GetAttrTemplateByID
func GetAttrTemplateByID(id uint) (model.AttrTemplate, error) {
	var attrTemplate model.AttrTemplate
	err := g.TENANCY_DB.Where("id = ?", id).First(&attrTemplate).Error
	return attrTemplate, err
}

// UpdateAttrTemplate
func UpdateAttrTemplate(req request.AttrTemplate, id uint) error {
	attrTemplate := model.AttrTemplate{
		TemplateName: req.TemplateName,
	}
	value, err := json.Marshal(&req.TemplateValue)
	if err != nil {
		return err
	}
	attrTemplate.TemplateValue = value
	err = g.TENANCY_DB.Model(&model.AttrTemplate{}).Where("id = ?", id).Updates(attrTemplate).Error
	return err
}

// DeleteAttrTemplate
func DeleteAttrTemplate(id uint) error {
	var attrTemplate model.AttrTemplate
	return g.TENANCY_DB.Where("id = ?", id).Delete(&attrTemplate).Error
}

// GetAttrTemplateInfoList
func GetAttrTemplateInfoList(info request.PageInfo) ([]response.AttrTemplate, int64, error) {
	var attrTemplateList []response.AttrTemplate
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.AttrTemplate{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return attrTemplateList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&attrTemplateList).Error
	if err != nil {
		return attrTemplateList, total, err
	}
	return attrTemplateList, total, err
}
