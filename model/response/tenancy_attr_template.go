package response

type TenancyAttrTemplate struct {
	TenancyResponse
	TemplateName  string `json:"templateName"`  // 规格名称
	TemplateValue string `json:"templateValue"` // 规格值

	SysTenancyID int `json:"sysTenancyId"` // 商户 id
}
