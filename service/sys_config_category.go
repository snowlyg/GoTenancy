package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// GetConfigCategoryMap
func GetConfigCategoryMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		cate, err := GetConfigCategoryByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"配置分类名称","props":{"type":"text","placeholder":"请输入配置分类名称"},"validate":[{"message":"请输入配置分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"key","value":"%s","title":"配置分类key","props":{"type":"text","placeholder":"请输入配置分类key"},"validate":[{"message":"请输入配置分类key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"配置分类说明","props":{"type":"text","placeholder":"请输入配置分类说明"}},{"type":"frame","field":"icon","value":"%s","title":"配置分类图标","props":{"type":"input","maxLength":1,"title":"请选择配置分类图标","src":"\/admin\/setting\/icons?field=icon","icon":"el-icon-circle-plus-outline","height":"338px","width":"700px","modal":{"modal":false}}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"PUT","title":"添加配置分类","config":{}}`, cate.Name, cate.Key, cate.Key, cate.Icon, cate.Sort, cate.Status)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"配置分类名称","props":{"type":"text","placeholder":"请输入配置分类名称"},"validate":[{"message":"请输入配置分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"key","value":"%s","title":"配置分类key","props":{"type":"text","placeholder":"请输入配置分类key"},"validate":[{"message":"请输入配置分类key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"配置分类说明","props":{"type":"text","placeholder":"请输入配置分类说明"}},{"type":"frame","field":"icon","value":"%s","title":"配置分类图标","props":{"type":"input","maxLength":1,"title":"请选择配置分类图标","src":"\/admin\/setting\/icons?field=icon","icon":"el-icon-circle-plus-outline","height":"338px","width":"700px","modal":{"modal":false}}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"","method":"POST","title":"添加配置分类","config":{}}`, "", "", "", "", 0, 1)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/configCategory/updateConfigCategory/%d", id), ctx)
	} else {
		form.SetAction("/configCategory/createConfigCategory", ctx)
	}
	return form, err
}

// CreateConfigCategory
func CreateConfigCategory(cate model.SysConfigCategory) (model.SysConfigCategory, error) {
	err := g.TENANCY_DB.Where("`key` = ?", cate.Key).First(&cate).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return cate, fmt.Errorf("KEY %s 已被使用", cate.Key)
	}
	err = g.TENANCY_DB.Create(&cate).Error
	return cate, err
}

// GetConfigCategoryByID
func GetConfigCategoryByID(id uint) (model.SysConfigCategory, error) {
	var cate model.SysConfigCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&cate).Error
	return cate, err
}

// GetConfigCategoryByKey
func GetConfigCategoryByKey(key string) (model.SysConfigCategory, error) {
	var cate model.SysConfigCategory
	err := g.TENANCY_DB.Where("key = ?", key).First(&cate).Error
	return cate, err
}

// SetConfigCategoryRegionByID
func SetConfigCategoryRegionByID(regionCode request.SetRegionCode) error {
	return g.TENANCY_DB.Model(&model.SysConfigCategory{}).Where("id = ?", regionCode.Id).Update("sys_region_code", regionCode.SysRegionCode).Error
}

// ChangeConfigCategoryStatus
func ChangeConfigCategoryStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysConfigCategory{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateConfigCategory
func UpdateConfigCategory(cate model.SysConfigCategory, id string) (model.SysConfigCategory, error) {
	err := g.TENANCY_DB.Where("`key` = ?", cate.Key).Not("id = ?", id).First(&cate).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return cate, fmt.Errorf("KEY %s 已被使用", cate.Key)
	}

	err = g.TENANCY_DB.Where("id = ?", id).Omit("uuid").Updates(&cate).Error
	return cate, err
}

// DeleteConfigCategory
func DeleteConfigCategory(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.SysConfigCategory{}).Error
}

// GetConfigCategoriesInfoList
func GetConfigCategoriesInfoList() ([]model.SysConfigCategory, error) {
	var cateList []model.SysConfigCategory
	err := g.TENANCY_DB.Find(&cateList).Error
	return cateList, err
}

// GetConfigCategoriesOptions
func GetConfigCategoriesOptions() ([]Option, error) {
	var options []Option
	var opts []Opt
	err := g.TENANCY_DB.Model(&model.SysConfigCategory{}).Select("id as value,name as label").Where("status = ?", g.StatusTrue).Find(&opts).Error
	if err != nil {
		return options, err
	}
	options = append(options, Option{Label: "请选择", Value: 0})

	for _, opt := range opts {
		options = append(options, Option{Label: opt.Label, Value: opt.Value})
	}

	return options, err
}
