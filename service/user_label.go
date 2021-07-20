package service

import (
	"encoding/json"
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

// GetUserLabelMap
func GetUserLabelMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		userLabel, err := GetUserLabelByID(id, multi.GetTenancyId(ctx))
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"labelName","value":"%s","title":"用户标签名称","props":{"type":"text","placeholder":"请输入用户标签名称"},"validate":[{"message":"请输入用户标签名称","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"PUT","title":"添加用户标签","config":{}}`, userLabel.LabelName)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"labelName","value":"%s","title":"用户标签名称","props":{"type":"text","placeholder":"请输入用户标签名称"},"validate":[{"message":"请输入用户标签名称","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"添加用户标签","config":{}}`, "")
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/userLabel/updateUserLabel/%d", id), ctx)
	} else {
		form.SetAction("/userLabel/createUserLabel", ctx)
	}
	return form, err
}

// GetUserLabelOptions
func GetUserLabelOptions(tenancyId uint) ([]Option, error) {
	var options []Option
	var opts []Opt
	err := g.TENANCY_DB.Model(&model.UserLabel{}).Select("id as value,label_name as label").Where("sys_tenancy_id", tenancyId).Find(&opts).Error
	if err != nil {
		return options, err
	}
	options = append(options, Option{Label: "请选择", Value: ""})

	for _, opt := range opts {
		options = append(options, Option{Label: opt.Label, Value: opt.Value})
	}

	return options, err
}

// CreateUserLabel
func CreateUserLabel(userLabel model.UserLabel) (model.UserLabel, error) {
	err := g.TENANCY_DB.Where("label_name = ?", userLabel.LabelName).Where("sys_tenancy_id = ?", userLabel.SysTenancyID).First(&userLabel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return userLabel, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Create(&userLabel).Error
	return userLabel, err
}

// GetUserLabelByID
func GetUserLabelByID(id uint, tenancyId uint) (model.UserLabel, error) {
	var userLabel model.UserLabel
	err := g.TENANCY_DB.Where("id = ?", id).Where("sys_tenancy_id", tenancyId).First(&userLabel).Error
	return userLabel, err
}

// GetUserLabelByUserIds
func GetUserLabelByUserIds(ids []uint, tenancyId uint) ([]response.UserLabelWithUserId, error) {
	var userLabels []response.UserLabelWithUserId
	err := g.TENANCY_DB.Model(&model.UserLabel{}).
		Select("user_labels.*,user_user_labels.sys_user_id").
		Joins("left join user_user_labels on user_user_labels.user_label_id = user_labels.id and sys_tenancy_id =?", tenancyId).
		Where("user_user_labels.sys_user_id in ?", ids).
		Where("user_labels.sys_tenancy_id", tenancyId).
		Find(&userLabels).Error
	return userLabels, err
}

// GetUserLabelByIds
func GetUserLabelByIds(ids []string, tenancyId uint) ([]response.UserLabelWithUserId, error) {
	var userLabels []response.UserLabelWithUserId
	err := g.TENANCY_DB.Model(&model.UserLabel{}).
		Select("user_labels.*,user_user_labels.sys_user_id").
		Joins("left join user_user_labels on user_user_labels.user_label_id = user_labels.id").
		Where("user_user_labels.user_label_id in ?", ids).
		Where("user_labels.sys_tenancy_id", tenancyId).
		Find(&userLabels).Error
	return userLabels, err
}

// UpdateUserLabel
func UpdateUserLabel(userLabel model.UserLabel, id uint) (model.UserLabel, error) {
	err := g.TENANCY_DB.Where("label_name = ?", userLabel.LabelName).Where("sys_tenancy_id = ?", userLabel.SysTenancyID).Where("id <> ?", id).First(&userLabel).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return userLabel, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(userLabel).Error
	return userLabel, err
}

// DeleteUserLabel
func DeleteUserLabel(id uint, tenancyId uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Where("sys_tenancy_id", tenancyId).Delete(&model.UserLabel{}).Error
}

// GetUserLabelInfoList
func GetUserLabelInfoList(info request.UserLabelPageInfo, tenancyId uint) ([]model.UserLabel, int64, error) {
	var userLabelList []model.UserLabel
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.UserLabel{})
	var total int64
	if info.LabelType > 0 {
		db = db.Where("type = ?", info.LabelType)
	}
	db = db.Where("sys_tenancy_id", tenancyId)
	err := db.Count(&total).Error
	if err != nil {
		return userLabelList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userLabelList).Error
	return userLabelList, total, err
}
