package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
)

// CreateSysOperationRecord 创建记录
func CreateSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return g.TENANCY_DB.Create(&sysOperationRecord).Error
}

// DeleteSysOperationRecordByIds 批量删除记录
func DeleteSysOperationRecordByIds(ids request.IdsReq) error {
	return g.TENANCY_DB.Delete(&[]model.SysOperationRecord{}, "id in (?)", ids.Ids).Error
}

// DeleteSysOperationRecord 删除操作记录
func DeleteSysOperationRecord(sysOperationRecord model.SysOperationRecord) error {
	return g.TENANCY_DB.Delete(&sysOperationRecord).Error
}

// GetSysOperationRecord 根据id获取单条操作记录
func GetSysOperationRecord(id uint) (model.SysOperationRecord, error) {
	var sysOperationRecord model.SysOperationRecord
	err := g.TENANCY_DB.Where("id = ?", id).First(&sysOperationRecord).Error
	return sysOperationRecord, err
}

// GetSysOperationRecordInfoList 分页获取操作记录列表
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch) ([]model.SysOperationRecord, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := g.TENANCY_DB.Model(&model.SysOperationRecord{})
	var sysOperationRecords []model.SysOperationRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Method != "" {
		db = db.Where("method = ?", info.Method)
	}
	if info.Path != "" {
		db = db.Where("path LIKE ?", "%"+info.Path+"%")
	}
	if info.Status != 0 {
		db = db.Where("status = ?", info.Status)
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Order("id desc").Limit(limit).Offset(offset).Preload("User").Find(&sysOperationRecords).Error
	return sysOperationRecords, total, err
}
