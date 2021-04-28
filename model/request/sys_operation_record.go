package request

import "github.com/snowlyg/go-tenancy/model"

type SysOperationRecordSearch struct {
	model.SysOperationRecord
	PageInfo
}
