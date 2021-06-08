package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var TenancyProduct = new(product)

type product struct{}

var products = []model.TenancyProduct{
	{
		StoreName:         "领立裁腰带短袖连衣裙",
		StoreInfo:         "短袖连衣裙",
		Keyword:           "短袖连衣裙",
		BarCode:           "",
		IsShow:            1,
		Status:            1,
		UnitName:          "件",
		Sort:              40,
		Rank:              0,
		Sales:             1,
		Price:             80.00,
		Cost:              50.00,
		OtPrice:           100.00,
		Stock:             399,
		IsHot:             0,
		IsBenefit:         0,
		IsBest:            0,
		IsNew:             0,
		IsGood:            1,
		ProductType:       model.GeneralSale,
		Ficti:             100,
		Browse:            0,
		CodePath:          "",
		VideoLink:         "",
		SpecType:          model.SingleSpec,
		ExtensionType:     model.SystemExtensionType,
		Refusal:           "",
		Rate:              5.0,
		ReplyCount:        0,
		GiveCouponIDs:     "",
		IsGiftBag:         2,
		CareCount:         0,
		Image:             "http://qmplusimg.henrongyi.top/head.png",
		SliderImage:       "http://qmplusimg.henrongyi.top/head.png",
		OldID:             0,
		TempID:            0,
		SysTenancyID:      1,
		SysBrandID:        1,
		TenancyCategoryID: 1,
	},
}

func (m *product) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.TenancyProduct{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_products 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&products).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_products 表初始数据成功!")
		return nil
	})
}
