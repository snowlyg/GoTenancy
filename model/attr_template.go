package model

import "github.com/snowlyg/go-tenancy/g"

// AttrTemplate 商品规则值(规格)表
type AttrTemplate struct {
	g.TENANCY_MODEL
	TemplateName  string `gorm:"column:template_name;type:varchar(32);not null" json:"templateName" binding:"required"` // 规格名称
	TemplateValue string `gorm:"column:template_value;type:text;not null" json:"templateValue" binding:"required"`      // 规格值

	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
}
