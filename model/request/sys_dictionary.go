package request

import "github.com/snowlyg/go-tenancy/model"

type SysDictionarySearch struct {
	model.SysDictionary
	PageInfo
}
