package service

import (
	"errors"
	"strconv"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateCategory
func CreateCategory(m request.CreateTenancyCategory) (model.TenancyCategory, error) {
	var brandCategory model.TenancyCategory
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

// GetCategoryByID
func GetCategoryByID(id float64) (model.TenancyCategory, error) {
	var brandCategory model.TenancyCategory
	err := g.TENANCY_DB.Where("id = ?", id).First(&brandCategory).Error
	return brandCategory, err
}

// UpdateCategory
func UpdateCategory(m request.UpdateTenancyCategory) (model.TenancyCategory, error) {
	var brandCategory model.TenancyCategory
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

// DeleteCategory
func DeleteCategory(id float64) error {
	var brandCategory model.TenancyCategory
	return g.TENANCY_DB.Where("id = ?", id).Delete(&brandCategory).Error
}

// GetCategoryInfoList
func GetCategoryInfoList() ([]response.TenancyCategory, error) {
	var brandCategoryList []response.TenancyCategory
	treeMap, err := getCategoryMap()
	brandCategoryList = treeMap["0"]
	for i := 0; i < len(brandCategoryList); i++ {
		err = getCategoryBaseChildrenList(&brandCategoryList[i], treeMap)
	}
	return brandCategoryList, err
}

// getCategoryMap
func getCategoryMap() (map[string][]response.TenancyCategory, error) {
	var brandCategoryList []response.TenancyCategory
	treeMap := make(map[string][]response.TenancyCategory)
	err := g.TENANCY_DB.Model(&model.TenancyCategory{}).Order("sort").Find(&brandCategoryList).Error
	for _, v := range brandCategoryList {
		treeMap[v.Pid] = append(treeMap[v.Pid], v)
	}
	return treeMap, err
}

// getCategoryBaseChildrenList
func getCategoryBaseChildrenList(cate *response.TenancyCategory, treeMap map[string][]response.TenancyCategory) (err error) {
	cate.Children = treeMap[strconv.Itoa(int(cate.ID))]
	for i := 0; i < len(cate.Children); i++ {
		err = getCategoryBaseChildrenList(&cate.Children[i], treeMap)
	}
	return err
}
