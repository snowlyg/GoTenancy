package request

type AttrTemplate struct {
	TemplateName  string  `json:"templateName"`  // 规格名称
	TemplateValue []Value `json:"templateValue"` // 规格值
}

type Value struct {
	Detail []string `json:"detail"`
	Value  string   `json:"value"`
}
