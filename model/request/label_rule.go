package request

import "github.com/snowlyg/go-tenancy/model"

type LabelRule struct {
	model.LabelRule
	LabelName string `json:"labelName"`
}
