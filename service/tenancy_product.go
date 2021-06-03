package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

// CreateProduct
func CreateProduct(m request.CreateTenancyProduct, ctx *gin.Context) (model.TenancyProduct, error) {
	var product = model.TenancyProduct{
		StoreName: m.StoreName,
		StoreInfo: m.StoreInfo,
		Keyword:   m.Keyword,
		BarCode:   m.BarCode,
		IsShow:    m.IsShow,
		Status:    m.Status,

		UnitName:          m.UnitName,
		Sort:              m.Sort,
		Rank:              m.Rank,
		Sales:             m.Sales,
		Price:             m.Price,
		Cost:              m.Cost,
		OtPrice:           m.OtPrice,
		Stock:             m.Stock,
		IsHot:             m.IsHot,
		IsBenefit:         m.IsBenefit,
		IsBest:            m.IsBest,
		IsNew:             m.IsNew,
		IsGood:            m.IsGood,
		ProductType:       m.ProductType,
		Ficti:             m.Ficti,
		Browse:            m.Browse,
		CodePath:          m.CodePath,
		VideoLink:         m.VideoLink,
		SpecType:          m.SpecType,
		ExtensionType:     m.ExtensionType,
		Refusal:           m.Refusal,
		Rate:              m.Rate,
		ReplyCount:        m.ReplyCount,
		GiveCouponIDs:     m.GiveCouponIDs,
		IsGiftBag:         m.IsGiftBag,
		CareCount:         m.CareCount,
		Image:             m.Image,
		SliderImage:       m.SliderImage,
		OldID:             m.OldID,
		TempID:            m.TempID,
		SysTenancyID:      multi.GetTenancyId(ctx),
		SysBrandID:        m.SysBrandID,
		TenancyCategoryID: m.TenancyCategoryID,
	}

	err := g.TENANCY_DB.Create(&product).Error
	return product, err
}

// GetProductByID
func GetProductByID(id float64) (model.TenancyProduct, error) {
	var product model.TenancyProduct
	err := g.TENANCY_DB.Where("id = ?", id).First(&product).Error
	return product, err
}

// UpdateProduct
func UpdateProduct(m request.UpdateTenancyProduct) (model.TenancyProduct, error) {
	var product model.TenancyProduct

	data := map[string]interface{}{
		"store_name":        m.StoreName,
		"store_info":        m.StoreInfo,
		"keyword":           m.Keyword,
		"bar_code":          m.BarCode,
		"is_show":           m.IsShow,
		"status":            m.Status,
		"unit_name":         m.UnitName,
		"sort":              m.Sort,
		"rank":              m.Rank,
		"sales":             m.Sales,
		"Price":             m.Price,
		"Cost":              m.Cost,
		"OtPrice":           m.OtPrice,
		"Stock":             m.Stock,
		"IsHot":             m.IsHot,
		"IsBenefit":         m.IsBenefit,
		"IsBest":            m.IsBest,
		"IsNew":             m.IsNew,
		"IsGood":            m.IsGood,
		"ProductType":       m.ProductType,
		"Ficti":             m.Ficti,
		"Browse":            m.Browse,
		"CodePath":          m.CodePath,
		"VideoLink":         m.VideoLink,
		"SpecType":          m.SpecType,
		"ExtensionType":     m.ExtensionType,
		"Refusal":           m.Refusal,
		"Rate":              m.Rate,
		"ReplyCount":        m.ReplyCount,
		"GiveCouponIDs":     m.GiveCouponIDs,
		"IsGiftBag":         m.IsGiftBag,
		"CareCount":         m.CareCount,
		"Image":             m.Image,
		"SliderImage":       m.SliderImage,
		"OldID":             m.OldID,
		"TempID":            m.TempID,
		"SysTenancyID":      m.SysTenancyID,
		"SysBrandID":        m.SysBrandID,
		"TenancyCategoryID": m.TenancyCategoryID,
	}
	product.ID = m.Id
	err := g.TENANCY_DB.Model(&product).Updates(data).Error
	return product, err
}

// DeleteProduct
func DeleteProduct(id float64) error {
	var product model.TenancyProduct
	return g.TENANCY_DB.Where("id = ?", id).Delete(&product).Error
}

// GetProductInfoList
func GetProductInfoList(info request.PageInfo, ctx *gin.Context) ([]response.TenancyProduct, int64, error) {
	var tenancyList []response.TenancyProduct
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.TenancyProduct{})
	if !multi.IsAdmin(ctx) {
		fmt.Printf("\n\n %d \n\n", multi.GetTenancyId(ctx))
		db = db.Where("sys_tenancy_id = ?", multi.GetTenancyId(ctx))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	err = db.Select("tenancy_products.*,sys_tenancies.name as sys_tenancy_name").
		Joins("left join sys_tenancies on tenancy_products.sys_tenancy_id = sys_tenancies.id").
		Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}
