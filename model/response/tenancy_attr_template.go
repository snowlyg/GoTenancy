package response

type AttrTemplate struct {
	TenancyResponse
	TemplateName  string  `json:"templateName"`  // 规格名称
	TemplateValue []Value `json:"templateValue"` // 规格值

	SysTenancyID int `json:"sysTenancyId"` // 商户 id
}

type Value struct {
	Detail []string `json:"detail"`
	Value  string   `json:"value"`
}
