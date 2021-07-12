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

// GetExpressMap
func GetExpressMap(id uint, ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	if id > 0 {
		express, err := GetExpressByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"快递公司名称","props":{"type":"text","placeholder":"请输入快递公司名称"},"validate":[{"message":"请输入快递公司名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"code","value":"%s","title":"快递公司编码","props":{"type":"text","placeholder":"请输入快递公司编码"},"validate":[{"message":"请输入快递公司编码","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"\/sys\/store\/express\/create.html","method":"PUT","title":"添加快递公司","config":{}}`, express.Name, express.Code, express.Status, express.Sort)

	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"name","value":"%s","title":"快递公司名称","props":{"type":"text","placeholder":"请输入快递公司名称"},"validate":[{"message":"请输入快递公司名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"code","value":"%s","title":"快递公司编码","props":{"type":"text","placeholder":"请输入快递公司编码"},"validate":[{"message":"请输入快递公司编码","required":true,"type":"string","trigger":"change"}]},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}}],"action":"\/sys\/store\/express\/create.html","method":"POST","title":"添加快递公司","config":{}}`, "", "", 1, 0)
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	if id > 0 {
		form.SetAction(fmt.Sprintf("/express/updateExpress/%d", id), ctx)
	} else {
		form.SetAction("/express/createExpress", ctx)
	}
	return form, err
}

// GetExpressOptions
func GetExpressOptions() ([]Option, error) {
	var options []Option
	var opts []Opt
	err := g.TENANCY_DB.Model(&model.Express{}).Select("id as value,name as label").Where("status = ?", g.StatusTrue).Find(&opts).Error
	if err != nil {
		return options, err
	}
	options = append(options, Option{Label: "请选择", Value: 0})

	for _, opt := range opts {
		options = append(options, Option{Label: opt.Label, Value: opt.Value})
	}

	return options, err
}

// CreateExpress
func CreateExpress(express model.Express) (model.Express, error) {
	err := g.TENANCY_DB.Where("code = ?", express.Code).First(&express).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return express, errors.New("物流代码已被注冊")
	}
	err = g.TENANCY_DB.Create(&express).Error
	return express, err
}

// GetExpressByID
func GetExpressByID(id uint) (model.Express, error) {
	var express model.Express
	err := g.TENANCY_DB.Where("id = ?", id).First(&express).Error
	return express, err
}

// ChangeExpressStatus
func ChangeExpressStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.Express{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateExpress
func UpdateExpress(express model.Express, id uint) (model.Express, error) {
	err := g.TENANCY_DB.Where("code = ?", express.Code).Where("id <> ?", id).First(&express).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return express, errors.New("物流代码已被注冊")
	}
	err = g.TENANCY_DB.Where("id = ?", id).Updates(express).Error
	return express, err
}

// DeleteExpress
func DeleteExpress(id uint) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.Express{}).Error
}

// GetExpressInfoList
func GetExpressInfoList(info request.ExpressPageInfo) ([]model.Express, int64, error) {
	var expressList []model.Express
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.Express{})
	if info.Name != "" {
		db = db.Where(g.TENANCY_DB.Where("name like ?", info.Name+"%").Or("code like ?", info.Name+"%"))
	}
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return expressList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&expressList).Error
	return expressList, total, err
}
