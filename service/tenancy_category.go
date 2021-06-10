package service

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// CreateCategory
func CreateCategory(m request.CreateTenancyCategory, ctx *gin.Context) (model.TenancyCategory, error) {
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
	brandCategory.Pic = m.Pic
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
		"pic":       m.Pic,
	}
	brandCategory.ID = m.Id
	err = g.TENANCY_DB.Model(&brandCategory).Updates(data).Error
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
	brandCategoryList = treeMap["0"]
	for i := 0; i < len(brandCategoryList); i++ {
		err = getCategoryBaseChildrenList(&brandCategoryList[i], treeMap)
	}
	return brandCategoryList, err
}

// getCategoryMap
func getCategoryMap(ctx *gin.Context) (map[string][]response.TenancyCategory, error) {
	var brandCategoryList []response.TenancyCategory
	treeMap := make(map[string][]response.TenancyCategory)
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
func getCategoryBaseChildrenList(cate *response.TenancyCategory, treeMap map[string][]response.TenancyCategory) (err error) {
	cate.Children = treeMap[strconv.Itoa(int(cate.ID))]
	for i := 0; i < len(cate.Children); i++ {
		err = getCategoryBaseChildrenList(&cate.Children[i], treeMap)
	}
	return err
}
