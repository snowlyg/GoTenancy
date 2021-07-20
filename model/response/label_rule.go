package response

import "github.com/snowlyg/go-tenancy/model"

type LabelRule struct {
	model.LabelRule
	LabelName string `json:"labelName" ` // 标签名称
}
