package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var ShippingTemplate = new(shippingTemplate)

type shippingTemplate struct{}

var shippingTemplates = []model.ShippingTemplate{
	// 	102	陕西包邮新疆不配送其它1元	0	1	1	65	0	9	2020-07-10 10:50:21
	// 103	邮费	1	0	1	66	0	0	2020-07-13 15:04:28
	// 104	邮费	1	0	1	66	0	0	2020-07-13 15:10:30
	// 105	包邮	0	1	1	65	0	0	2020-07-13 17:30:51
	{BaseShippingTemplate: model.BaseShippingTemplate{Name: "陕西包邮新疆不配送其它1元", Type: 2, Appoint: 1, Undelivery: 1, IsDefault: 2, Sort: 9}, SysTenancyID: 1},
	{BaseShippingTemplate: model.BaseShippingTemplate{Name: "邮费", Type: 1, Appoint: 2, Undelivery: 1, IsDefault: 2, Sort: 0}, SysTenancyID: 1},
	{BaseShippingTemplate: model.BaseShippingTemplate{Name: "邮费", Type: 1, Appoint: 2, Undelivery: 1, IsDefault: 2, Sort: 0}, SysTenancyID: 1},
	{BaseShippingTemplate: model.BaseShippingTemplate{Name: "包邮", Type: 2, Appoint: 1, Undelivery: 1, IsDefault: 2, Sort: 0}, SysTenancyID: 1},
}

func (m *shippingTemplate) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.ShippingTemplate{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> shipping_templates 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&shippingTemplates).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> shipping_templates 表初始数据成功!")
		return nil
	})
}

var ShippingTemplateFree = new(shippingTemplateFree)

type shippingTemplateFree struct{}

var shippingTemplateFrees = []model.ShippingTemplateFree{
	{Number: 1, Price: 1.00, Code: 1, ShippingTemplateID: 1},
	{Number: 1, Price: 1.00, Code: 1, ShippingTemplateID: 4},
	{Number: 1, Price: 1.00, Code: 1, ShippingTemplateID: 4},
}

func (m *shippingTemplateFree) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.ShippingTemplateFree{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> shipping_template_frees 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&shippingTemplateFrees).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> shipping_template_frees 表初始数据成功!")
		return nil
	})
}

var ShippingTemplateRegion = new(shippingTemplateRegion)

type shippingTemplateRegion struct{}

var shippingTemplateRegions = []model.ShippingTemplateRegion{
	{First: 1, FirstPrice: 1.00, Continue: 1.00, ContinuePrice: 1.00, Code: 1, ShippingTemplateID: 1},
	{First: 1, FirstPrice: 1.00, Continue: 10.00, ContinuePrice: 1.00, Code: 1, ShippingTemplateID: 2},
	{First: 1, FirstPrice: 1.00, Continue: 10.00, ContinuePrice: 1.00, Code: 1, ShippingTemplateID: 3},
	{First: 1, FirstPrice: 1.00, Continue: 1.00, ContinuePrice: 1.00, Code: 1, ShippingTemplateID: 4},
}

func (m *shippingTemplateRegion) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.ShippingTemplateRegion{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> shipping_template_regions 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&shippingTemplateRegions).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> shipping_template_regions 表初始数据成功!")
		return nil
	})
}

var ShippingTemplateUndelivery = new(shippingTemplateUndelivery)

type shippingTemplateUndelivery struct{}

var shippingTemplateUndeliveries = []model.ShippingTemplateUndelivery{
	{Code: 1, ShippingTemplateID: 1},
	{Code: 1, ShippingTemplateID: 2},
	{Code: 1, ShippingTemplateID: 3},
	{Code: 1, ShippingTemplateID: 4},
}

func (m *shippingTemplateUndelivery) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.ShippingTemplateUndelivery{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> shipping_template_undeliveries 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&shippingTemplateUndeliveries).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> shipping_template_undeliveries 表初始数据成功!")
		return nil
	})
}
