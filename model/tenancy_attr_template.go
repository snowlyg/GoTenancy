package model

import "github.com/snowlyg/go-tenancy/g"

// TenancyAttrTemplate 商品规则值(规格)表
type TenancyAttrTemplate struct {
	g.TENANCY_MODEL
	TemplateName  string `gorm:"column:template_name;type:varchar(32);not null" json:"templateName"` // 规格名称
	TemplateValue string `gorm:"column:template_value;type:text;not null" json:"templateValue"`      // 规格值

	TenancyID int `gorm:"index:tenancy_id;column:tenancy_id;type:int;not null" json:"tenancyId"` // 商户 id
}
