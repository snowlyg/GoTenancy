package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// GetUserGroupMap
func GetUserGroupMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		userGroup, err := GetUserGroupByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"groupName","value":"%s","title":"用户分组名称","props":{"type":"text","placeholder":"请输入用户分组名称"},"validate":[{"message":"请输入用户分组名称","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"PUT","title":"添加用户分组","config":{}}`, userGroup.GroupName)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"groupName","value":"%s","title":"用户分组名称","props":{"type":"text","placeholder":"请输入用户分组名称"},"validate":[{"message":"请输入用户分组名称","required":true,"type":"string","trigger":"change"}]}],"action":"","method":"POST","title":"添加用户分组","config":{}}`, "")
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/userGroup/updateUserGroup/%d", id), ctx)
	} else {
		form.SetAction("/userGroup/createUserGroup", ctx)
	}
	return form, err
}

// GetUserGroupOptions
func GetUserGroupOptions() ([]Option, error) {
	var options []Option
	var opts []Opt
	err := g.TENANCY_DB.Model(&model.UserGroup{}).Select("id as value,group_name as label").Find(&opts).Error
	if err != nil {
		return options, err
	}
	options = append(options, Option{Label: "请选择", Value: ""})

	for _, opt := range opts {
		options = append(options, Option{Label: opt.Label, Value: opt.Value})
	}

	return options, err
}

// CreateUserGroup
func CreateUserGroup(userGroup model.UserGroup) (model.UserGroup, error) {
	err := g.TENANCY_DB.Where("group_name = ?", userGroup.GroupName).First(&userGroup).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return userGroup, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Create(&userGroup).Error
	return userGroup, err
}

// GetUserGroupByID
func GetUserGroupByID(id uint) (model.UserGroup, error) {
	var userGroup model.UserGroup
	err := g.TENANCY_DB.Where("id = ?", id).First(&userGroup).Error
	return userGroup, err
}

// UpdateUserGroup
func UpdateUserGroup(userGroup model.UserGroup, id uint) (model.UserGroup, error) {
	err := g.TENANCY_DB.Where("group_name = ?", userGroup.GroupName).Where("id <> ?", id).First(&userGroup).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return userGroup, errors.New("名称已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(userGroup).Error
	return userGroup, err
}

// DeleteUserGroup
func DeleteUserGroup(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.UserGroup{}).Error
}

// GetUserGroupInfoList
func GetUserGroupInfoList(info request.PageInfo) ([]model.UserGroup, int64, error) {
	var userGroupList []model.UserGroup
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.UserGroup{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return userGroupList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&userGroupList).Error
	return userGroupList, total, err
}
