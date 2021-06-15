package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// GetBrandCategoryMap
func GetBrandCategoryMap(id uint) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		cate, err := GetBrandCategoryByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"%s/%d","method":"PUT","title":"添加分类","config":{}}`, cate.Pid, cate.CateName, cate.Status, cate.Sort, "/admin/brandCategory/updateBrandCategory", id)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"%s","method":"POST","title":"添加分类","config":{}}`, 0, "", 2, 0, "/admin/brandCategory/createBrandCategory")
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	opts, err := GetBrandCategoriesOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	return form, err
}

// CreateBrandCategory
func CreateBrandCategory(brandCategory model.SysBrandCategory) (model.SysBrandCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", brandCategory.CateName).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Create(&brandCategory).Error
	return brandCategory, err
}

// GetBrandCategoryByID
func GetBrandCategoryByID(id uint) (model.SysBrandCategory, error) {
	var brandCategory model.SysBrandCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&brandCategory).Error
	return brandCategory, err
}

// ChangeBrandCategoryStatus
func ChangeBrandCategoryStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysBrandCategory{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateBrandCategory
func UpdateBrandCategory(brandCategory model.SysBrandCategory, id uint) (model.SysBrandCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", brandCategory.CateName).Where("id <> ?", id).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Model(&brandCategory).Where("id = ?", id).Updates(brandCategory).Error
	return brandCategory, err
}

// DeleteBrandCategory
func DeleteBrandCategory(id uint) error {
	var brandCategory model.SysBrandCategory
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brandCategory).Error
}

// GetBrandCategoryInfoList
func GetBrandCategoryInfoList() ([]response.SysBrandCategory, error) {
	var brandCategoryList []response.SysBrandCategory
	treeMap, err := getBrandCategoryMap()
	brandCategoryList = treeMap[0]
	for i := 0; i < len(brandCategoryList); i++ {
		err = getBrandCategoryBaseChildrenList(&brandCategoryList[i], treeMap)
	}
	return brandCategoryList, err
}

// getBrandCategoryMap
func getBrandCategoryMap() (map[int32][]response.SysBrandCategory, error) {
	var brandCategoryList []response.SysBrandCategory
	treeMap := make(map[int32][]response.SysBrandCategory)
	err := g.TENANCY_DB.Model(&model.SysBrandCategory{}).Order("sort").Find(&brandCategoryList).Error
	for _, v := range brandCategoryList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// getBrandCategoryBaseChildrenList
func getBrandCategoryBaseChildrenList(cate *response.SysBrandCategory, treeMap map[int32][]response.SysBrandCategory) (err error) {
	cate.Children = treeMap[int32(cate.ID)]
	for i := 0; i < len(cate.Children); i++ {
		err = getBrandCategoryBaseChildrenList(&cate.Children[i], treeMap)
	}
	return err
}

// GetBrandCategoriesOptions
func GetBrandCategoriesOptions() ([]Option, error) {
	var options []Option
	options = append(options, Option{Label: "请选择", Value: 0})
	treeMap, err := getBrandCategoryMap()

	for _, opt := range treeMap[0] {
		options = append(options, Option{Label: opt.CateName, Value: opt.ID})
	}
	for i := 0; i < len(options); i++ {
		getBrandCategoriesOption(&options[i], treeMap)
	}

	return options, err
}

// getBrandCategoriesOption
func getBrandCategoriesOption(op *Option, treeMap map[int32][]response.SysBrandCategory) {
	id, ok := op.Value.(uint)
	if ok {
		for _, opt := range treeMap[int32(id)] {
			op.Children = append(op.Children, Option{Label: opt.CateName, Value: opt.ID})
		}
		for i := 0; i < len(op.Children); i++ {
			getBrandCategoriesOption(&op.Children[i], treeMap)
		}
	}
}
