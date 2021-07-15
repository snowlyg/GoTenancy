package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
)

func SetNowMoneyMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	user, err := GetGeneralDetail(id)
	if err != nil {
		return Form{}, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"radio","field":"type","value":1,"title":"修改余额","props":{},"validate":[{"message":"请选择修改余额","required":true,"type":"string","trigger":"change"}],"options":[{"label":"增加","value":1},{"label":"减少","value":2}]},{"type":"inputNumber","field":"nowMoney","value":%f,"title":"金额","props":{"placeholder":"请输入金额","min":0},"validate":[{"message":"请输入金额","required":true,"type":"number","trigger":"change"}]}],"action":"","method":"POST","title":"修改用户余额","config":{}}`, user.NowMoney)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	form.SetAction(fmt.Sprintf("%s/%d", "/cuser/setNowMoney", id), ctx)
	return form, err
}

// SetNowMoney
func SetNowMoney(id uint, req request.SetNowMoney) error {
	user, err := GetGeneralDetail(id)
	if err != nil {
		return err
	}
	nowMoney := user.NowMoney
	// 增加
	if req.Type == 1 {
		nowMoney = user.NowMoney + req.NowMoney
	} else if req.Type == 2 {
		if user.NowMoney <= req.NowMoney {
			nowMoney = 0
		} else {
			nowMoney = user.NowMoney - req.NowMoney
		}
	}
	if err := g.TENANCY_DB.Model(&model.GeneralInfo{}).Where("sys_user_id = ?", id).Updates(map[string]interface{}{"now_money": nowMoney}).Error; err != nil {
		return err
	}
	return err
}

func GetGeneralDetail(id uint) (response.GeneralUserDetail, error) {
	var user response.GeneralUserDetail
	generalAuthorityIds, err := GetUserAuthorityIds()
	if err != nil {
		return user, err
	}

	err = g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id,general_infos.avatar_url,general_infos.nick_name,general_infos.now_money,general_infos.pay_count,general_infos.pay_price, sum(orders.pay_price) as total_pay_price, count(orders.id) as total_pay_count").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id").
		Joins("left join orders on orders.sys_user_id = sys_users.id").
		Where("sys_users.authority_id IN (?)", generalAuthorityIds).
		Where("sys_users.id = ?", id).
		First(&user).Error
	return user, err
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.UserPageInfo) ([]response.GeneralUser, int64, error) {
	var userList []response.GeneralUser
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	generalAuthorityIds, err := GetUserAuthorityIds()
	if err != nil {
		return userList, 0, err
	}

	db := g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id as uid,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, general_infos.*,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,user_groups.group_name").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join user_groups on general_infos.group_id = user_groups.id").
		Where("sys_users.authority_id IN (?)", generalAuthorityIds)
	if info.UserTimeType != "" && info.UserTime != "" {
		userTimes := strings.Split(info.UserTime, "-")
		start, err := time.Parse("2006/01/02", userTimes[0])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		end, err := time.Parse("2006/01/02", userTimes[1])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		if info.UserTimeType == "add_time" {
			db = db.Where("general_infos.created_at BETWEEN ? AND ?", start, end)
		} else if info.UserTimeType == "visit" {
			db = db.Where("general_infos.last_time BETWEEN ? AND ?", start, end)
		}
	}

	if info.PayCount != "" {
		if info.PayCount == "0" {
			db = db.Where("general_infos.pay_count = ?", info.PayCount)
		} else {
			db = db.Where("general_infos.pay_count >= ?", info.PayCount)
		}
	}
	if info.GroupId != "" {
		db = db.Where("general_infos.group_id = ?", info.GroupId)
	}
	if info.LabelId != "" {
		db = db.Where("general_infos.label_id = ?", info.LabelId)
	}
	if info.Sex != "" {
		db = db.Where("general_infos.sex = ?", info.Sex)
	}
	if info.NickName != "" {
		db = db.Where("general_infos.nick_name like ?", info.NickName+"%")
	}

	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userList).Error
	if err != nil {
		return userList, total, err
	}

	if len(userList) > 0 {
		userLabelIds := []uint{}
		for _, user := range userList {
			userLabelIds = append(userLabelIds, user.LabelID)
		}
		userLabels, err := GetUserLabelByIds(userLabelIds)
		if err != nil {
			return userList, total, err
		}
		for i := 0; i < len(userList); i++ {
			for _, userLabel := range userLabels {
				if userList[i].LabelID == userLabel.ID {
					userList[i].Label = append(userList[i].Label, userLabel)
				}
			}
		}
	}

	return userList, total, err
}
