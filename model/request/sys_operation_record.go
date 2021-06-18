package request

import "github.com/snowlyg/go-tenancy/model"

type SysOperationRecordSearch struct {
	model.BaseOperationRecord
	PageInfo
}
