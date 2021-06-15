package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// GetTenancyCategoryMap
func GetTenancyCategoryMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		cate, err := GetCategoryByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false},"filterable":true},"validate":[{"required":true,"type":"integer","trigger":"change"}]},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"\/admin\/setting\/uploadPicture?field=pic&type=1","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"%s/%d","method":"PUT","title":"编辑分类","config":{}}`, cate.Pid, cate.CateName, cate.Pic, cate.Status, cate.Sort, "/admin/category/updateCategory", id)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"pid","value":%d,"title":"上级分类","props":{"type":"other","options":[],"placeholder":"请选择上级分类","props":{"checkStrictly":true,"emitPath":false},"filterable":true},"validate":[{"required":true,"type":"integer","trigger":"change"}]},{"type":"input","field":"cateName","value":"%s","title":"分类名称","props":{"type":"text","placeholder":"请输入分类名称"},"validate":[{"message":"请输入分类名称","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"pic","value":"%s","title":"分类图片(110*110px)","props":{"type":"image","maxLength":1,"title":"请选择分类图片(110*110px)","src":"\/admin\/setting\/uploadPicture?field=pic&type=1","width":"896px","height":"480px","footer":false,"modal":{"modal":false,"custom-class":"suibian-modal"}}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"%s","method":"POST","title":"添加分类","config":{}}`, 0, "", "", 2, 0, "/admin/category/createCategory")
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	opts, err := GetTenacyCategoriesOptions(ctx)
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	return form, err
}

// CreateCategory
func CreateCategory(brandCategory model.TenancyCategory, ctx *gin.Context) (model.TenancyCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", brandCategory.CateName).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}

	brandCategory.SysTenancyID = multi.GetTenancyId(ctx)
	err = g.TENANCY_DB.Create(&brandCategory).Error
	return brandCategory, err
}

// GetCategoryByID
func GetCategoryByID(id uint) (model.TenancyCategory, error) {
	var brandCategory model.TenancyCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&brandCategory).Error
	return brandCategory, err
}

// ChangeCategoryStatus
func ChangeCategoryStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.TenancyCategory{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateCategory
func UpdateCategory(brandCategory model.TenancyCategory, id uint) (model.TenancyCategory, error) {
	err := g.TENANCY_DB.Where("cate_name = ?", brandCategory.CateName).Where("id <> ?", id).First(&brandCategory).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return brandCategory, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(brandCategory).Error
	return brandCategory, err
}

// DeleteCategory
func DeleteCategory(id uint) error {
	var brandCategory model.TenancyCategory
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brandCategory).Error
}

// GetCategoryInfoList
func GetCategoryInfoList(ctx *gin.Context) ([]response.TenancyCategory, error) {
	var brandCategoryList []response.TenancyCategory
	treeMap, err := getCategoryMap(ctx)
	brandCategoryList = treeMap[0]
	for i := 0; i < len(brandCategoryList); i++ {
		err = getCategoryBaseChildrenList(&brandCategoryList[i], treeMap)
	}
	return brandCategoryList, err
}

// getCategoryMap
func getCategoryMap(ctx *gin.Context) (map[int32][]response.TenancyCategory, error) {
	var brandCategoryList []response.TenancyCategory
	treeMap := make(map[int32][]response.TenancyCategory)
	db := g.TENANCY_DB.Model(&model.TenancyCategory{})
	if !multi.IsAdmin(ctx) {
		db = db.Where("sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	err := db.Order("sort").Find(&brandCategoryList).Error
	for _, v := range brandCategoryList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// getCategoryBaseChildrenList
func getCategoryBaseChildrenList(cate *response.TenancyCategory, treeMap map[int32][]response.TenancyCategory) (err error) {
	cate.Children = treeMap[int32(cate.ID)]
	for i := 0; i < len(cate.Children); i++ {
		err = getCategoryBaseChildrenList(&cate.Children[i], treeMap)
	}
	return err
}

// GetTenacyCategoriesOptions
func GetTenacyCategoriesOptions(ctx *gin.Context) ([]Option, error) {
	var options []Option
	options = append(options, Option{Label: "请选择", Value: 0})
	treeMap, err := getCategoryMap(ctx)

	for _, opt := range treeMap[0] {
		options = append(options, Option{Label: opt.CateName, Value: opt.ID})
	}
	for i := 0; i < len(options); i++ {
		getCategoriesOption(&options[i], treeMap)
	}

	return options, err
}

// getCategoriesOption
func getCategoriesOption(op *Option, treeMap map[int32][]response.TenancyCategory) {
	id, ok := op.Value.(uint)
	if ok {
		for _, opt := range treeMap[int32(id)] {
			op.Children = append(op.Children, Option{Label: opt.CateName, Value: opt.ID})
		}
		for i := 0; i < len(op.Children); i++ {
			getCategoriesOption(&op.Children[i], treeMap)
		}
	}
}
