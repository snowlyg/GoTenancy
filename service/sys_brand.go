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

// GetBrandMap
func GetBrandMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		brand, err := GetBrandByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"brandCategoryId","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"emitPath":false}}},{"type":"input","field":"brandName","value":"%s","title":"品牌名称","props":{"type":"text","placeholder":"请输入品牌名称"},"validate":[{"message":"请输入品牌名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"\/admin\/setting\/uploadPicture?field=pic&type=1","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"PUT","title":"添加品牌","config":{}}`, brand.BrandCategoryID, brand.BrandName, brand.Pic, brand.Status, brand.Sort)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"brandCategoryId","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"emitPath":false}}},{"type":"input","field":"brandName","value":"%s","title":"品牌名称","props":{"type":"text","placeholder":"请输入品牌名称"},"validate":[{"message":"请输入品牌名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"\/admin\/setting\/uploadPicture?field=pic&type=1","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"","method":"POST","title":"添加品牌","config":{}}`, 0, "", "", 2, 0)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/brand/updateBrand/%d", id), ctx)
	} else {
		form.SetAction("/brand/createBrand", ctx)
	}
	opts, err := GetBrandCategoriesOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	return form, err
}

// CreateBrand
func CreateBrand(brand model.SysBrand) (model.SysBrand, error) {
	err := g.TENANCY_DB.Where("brand_name = ?", brand.BrandName).First(&brand).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brand, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Create(&brand).Error
	return brand, err
}

// GetBrandByID
func GetBrandByID(id uint) (model.SysBrand, error) {
	var brand model.SysBrand
	err := g.TENANCY_DB.Where("id = ?", id).First(&brand).Error
	return brand, err
}

// ChangeBrandStatus
func ChangeBrandStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysBrand{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateBrand
func UpdateBrand(brand model.SysBrand, id uint) (model.SysBrand, error) {
	err := g.TENANCY_DB.Where("brand_name = ?", brand.BrandName).Where("id <> ?", id).First(&brand).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brand, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(brand).Error
	return brand, err
}

// DeleteBrand
func DeleteBrand(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.SysBrand{}).Error
}

// GetBrandInfoList
func GetBrandInfoList(info request.BrandPageInfo) ([]model.SysBrand, int64, error) {
	var brandList []model.SysBrand
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysBrand{})
	if info.BrandCategoryId > 0 {
		db = db.Where("brand_category_id =?", info.BrandCategoryId)
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return brandList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&brandList).Error
	return brandList, total, err
}
