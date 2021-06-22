package response

import "gorm.io/datatypes"

type AttrTemplate struct {
	TenancyResponse
	TemplateName  string         `json:"templateName"`  // 规格名称
	TemplateValue datatypes.JSON `json:"templateValue"` // 规格值

	SysTenancyID int `json:"sysTenancyId"` // 商户 id
}
