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

// CreateAutoLabel
func CreateAutoLabel(labelRule request.LabelRule) (request.LabelRule, error) {
	var label model.UserLabel
	err := g.TENANCY_DB.Model(&model.UserLabel{}).Where("label_name = ?", labelRule.LabelName).
		Where("sys_tenancy_id = ?", labelRule.SysTenancyID).
		Where("type = ?", model.UserLabelTypeZD).
		First(&label).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return labelRule, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		label := model.UserLabel{LabelName: labelRule.LabelName, Type: model.UserLabelTypeZD, SysTenancyID: labelRule.SysTenancyID}
		if err := tx.Model(&model.UserLabel{}).Create(&label).Error; err != nil {
			return err
		}
		ru := model.LabelRule{Max: labelRule.Max, Min: labelRule.Min, Type: labelRule.Type, UserLabelID: label.ID, SysTenancyID: labelRule.SysTenancyID}
		if err := tx.Model(&model.LabelRule{}).Create(&ru).Error; err != nil {
			return err
		}
		labelRule.LabelRule = ru
		return nil
	})
	if err != nil {
		return labelRule, err
	}
	return labelRule, nil
}

// GetLabelRuleById
func GetLabelRuleById(id uint, ctx *gin.Context) (response.LabelRule, error) {
	var labelRule response.LabelRule
	err := g.TENANCY_DB.Model(&model.UserLabel{}).
		Select("label_rules.*,user_labels.label_name").
		Joins("left join label_rules on label_rules.user_label_id = user_labels.id").
		Where("label_rules.id = ?", id).
		Where("user_labels.sys_tenancy_id", multi.GetTenancyId(ctx)).
		Find(&labelRule).Error
	if err != nil {
		return labelRule, fmt.Errorf("get label rule %w", err)
	}
	return labelRule, nil
}

// UpdateAutoLabel
func UpdateAutoLabel(labelRule request.LabelRule, id uint, ctx *gin.Context) error {
	rule, err := GetLabelRuleById(id, ctx)
	if err != nil {
		return err
	}
	var label model.UserLabel
	err = g.TENANCY_DB.Model(&model.UserLabel{}).Where("label_name = ?", labelRule.LabelName).
		Where("sys_tenancy_id = ?", labelRule.SysTenancyID).
		Where("type = ?", model.UserLabelTypeZD).
		Where("id <> ?", rule.UserLabelID).
		First(&label).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("名称已被注冊")
	}
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		ru := map[string]interface{}{"max": labelRule.Max, "min": labelRule.Min, "type": labelRule.Type}
		if err := tx.Model(&model.LabelRule{}).Where("id = ?", id).Updates(ru).Error; err != nil {
			return err
		}
		if err := tx.Model(&model.UserLabel{}).Where("id = ?", rule.UserLabelID).Updates(map[string]interface{}{"label_name": labelRule.LabelName}).Error; err != nil {
			return err
		}
		return nil
	})
}

// DeleteAutoLabel
func DeleteAutoLabel(id uint, ctx *gin.Context) error {
	rule, err := GetLabelRuleById(id, ctx)
	if err != nil {
		return err
	}
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.LabelRule{}, id).Error; err != nil {
			return err
		}
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Delete(&model.UserLabel{}, rule.UserLabelID).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}
		// 返回 nil 提交事务
		return nil
	})
}

func GetAutoUserLabelInfoList(info request.UserLabelPageInfo, ctx *gin.Context) ([]response.LabelRule, int64, error) {
	var userLabelList []response.LabelRule
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.UserLabel{}).
		Select("label_rules.*,user_labels.label_name").
		Joins("left join label_rules on label_rules.user_label_id = user_labels.id and  label_rules.sys_tenancy_id =user_labels.sys_tenancy_id")
	var total int64
	if info.LabelType > 0 {
		db = db.Where("user_labels.type = ?", info.LabelType)
	}
	db = db.Where("user_labels.sys_tenancy_id", multi.GetTenancyId(ctx))
	err := db.Count(&total).Error
	if err != nil {
		return userLabelList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userLabelList).Error
	return userLabelList, total, err
}
