package request

import "github.com/snowlyg/go-tenancy/model"

type SysDictionaryDetailSearch struct {
	model.SysDictionaryDetail
	PageInfo
}
