package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var AttrTemplate = new(attrTemplate)

type attrTemplate struct{}

var attrTemplates = []model.TenancyAttrTemplate{
	{TemplateName: "鞋类", TemplateValue: "[{\"value\":\"\\u989c\\u8272\",\"detail\":[\"\\u9ec4\\u8272\",\"\\u7ea2\\u8272\"]},{\"value\":\"\\u5927\\u5c0f\",\"detail\":[\"35\",\"36\",\"38\"]}]", SysTenancyID: 1},
	{TemplateName: "化妆品", TemplateValue: "[{\"value\":\"\\u96c5\\u8bd7\\u5170\\u9edb\",\"detail\":[\"15\",\"20\"]},{\"value\":\"\\u5170\\u853b\",\"detail\":[\"15\",\"20\"]}]", SysTenancyID: 1},
	{TemplateName: "手机", TemplateValue: "[{\"value\":\"\\u989c\\u8272\",\"detail\":[\"\\u9ed1\",\"\\u94f6\",\"\\u91d1\",\"\\u767d\",\"\\u65e0\\u8272\"]}]", SysTenancyID: 1},
}

func (m *attrTemplate) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2, 3}).Find(&[]model.TenancyAttrTemplate{}).RowsAffected == 3 {
			color.Danger.Println("\n[Mysql] --> attr_templates 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&attrTemplates).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> attr_templates 表初始数据成功!")
		return nil
	})
}
