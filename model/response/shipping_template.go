package response

import "github.com/snowlyg/go-tenancy/model"

type ShippingTemplateList struct {
	TenancyResponse
	model.BaseShippingTemplate
}

type ShippingTemplateDetail struct {
	TenancyResponse
	model.BaseShippingTemplate
	Free      []model.ShippingTemplateFree     `gorm:"-" json:"free"`
	Region    []model.ShippingTemplateRegion   `gorm:"-" json:"region"`
	Undelives model.ShippingTemplateUndelivery `gorm:"-" json:"undelives"`
}

type ShippingTemplateSelect struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
