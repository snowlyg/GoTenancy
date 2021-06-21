package service

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
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
func GetSysOperationRecordInfoList(info request.SysOperationRecordSearch, ctx *gin.Context) ([]response.SysOperationRecord, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	var sysOperationRecords []response.SysOperationRecord
	var adminUsers []response.SysAdminUser
	var tenancyUsers []response.SysTenancyUser
	var err error
	db := g.TENANCY_DB.Model(&model.SysOperationRecord{})
	if multi.IsTenancy(ctx) {
		var userIds []int64
		tenancyId := multi.GetTenancyId(ctx)
		err = g.TENANCY_DB.Model(&model.SysTenancyInfo{}).Select("sys_user_id").Where("sys_tenancy_id = ?", tenancyId).Find(&userIds).Error
		if err == nil {
			db = db.Where("user_id in ?", userIds)
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Printf("err %+v\n\n\n\n", err)
			if err != nil {
				return nil, 0, err
			}
		}
	} else if multi.IsAdmin(ctx) {
		adminUsers, _, err = GetAdminInfoList(request.PageInfo{})
		if err != nil {
			return nil, 0, err
		}
	}
	tenancyUsers, _, err = GetTenancyInfoList(request.PageInfo{})
	if err != nil {
		return nil, 0, err
	}
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
	err = db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}
	err = db.Order("id desc").Limit(limit).Omit("tenancy_name,user_name,nick_name").Offset(offset).Find(&sysOperationRecords).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, total, err
	}

	if len(tenancyUsers) > 0 {
		for i := 0; i < len(sysOperationRecords); i++ {
			for _, tenancyUser := range tenancyUsers {
				if tenancyUser.ID == sysOperationRecords[i].UserID {
					sysOperationRecords[i].NickName = tenancyUser.NickName
					sysOperationRecords[i].TenancyName = tenancyUser.TenancyName
					sysOperationRecords[i].UserName = tenancyUser.Username
					continue
				}
			}
		}
	}

	if len(adminUsers) > 0 {
		for i := 0; i < len(sysOperationRecords); i++ {
			for _, adminUser := range adminUsers {
				if adminUser.ID == sysOperationRecords[i].UserID {
					sysOperationRecords[i].NickName = adminUser.NickName
					sysOperationRecords[i].UserName = adminUser.Username
					continue
				}
			}
		}
	}

	return sysOperationRecords, total, err
}
