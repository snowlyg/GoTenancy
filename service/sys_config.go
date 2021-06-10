package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"gorm.io/gorm"
)

// GetConfigMap
func GetConfigMap(id string) (Form, error) {
	var form Form
	var formStr string
	if id != "" {
		config, err := GetConfigByID(id)
		if err != nil {
			return form, err
		}
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"sysConfigCategoryId","value":%d,"title":"配置分类","props":{"type":"other","options":[],"placeholder":"请选择分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"userType","value":%d,"title":"后台类型","props":{"multiple":false,"placeholder":"请选择后台类型"},"validate":[{"message":"请选择后台类型","required":true,"type":"number","trigger":"change"}],"options":[{"label":"总后台配置","value":2},{"label":"商户后台配置","value":1}]},{"type":"input","field":"configName","value":"%s","title":"配置名称","props":{"type":"text","placeholder":"请输入配置名称"},"validate":[{"message":"请输入配置名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"configKey","value":"%s","title":"配置key","props":{"type":"text","placeholder":"请输入配置key"},"validate":[{"message":"请输入配置key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"说明","props":{"type":"text","placeholder":"请输入说明"}},{"type":"select","field":"configType","value":"%s","title":"配置类型","props":{"multiple":false,"placeholder":"请选择配置类型"},"validate":[{"message":"请选择配置类型","required":true,"type":"string","trigger":"change"}],"options":[]},{"type":"input","field":"configRule","value":"","title":"规则","props":{"type":"textarea","placeholder":"请输入规则"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"required","value":%d,"title":"必填","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":0,"inactiveText":"关闭","activeText":"开启"}}],"action":"%s/%s","method":"PUT","title":"添加配置","config":{}}`, config.SysConfigCategoryID, config.UserType, config.ConfigKey, config.ConfigName, config.Info, config.ConfigType, config.Sort, config.Required, config.Status, "/admin/config/updateConfig", id)
	} else {
		formStr = fmt.Sprintf(`{"rule":[{"type":"cascader","field":"sysConfigCategoryId","value":%d,"title":"配置分类","props":{"type":"other","options":[],"placeholder":"请选择分类","props":{"checkStrictly":true,"emitPath":false}}},{"type":"select","field":"userType","value":%d,"title":"后台类型","props":{"multiple":false,"placeholder":"请选择后台类型"},"validate":[{"message":"请选择后台类型","required":true,"type":"number","trigger":"change"}],"options":[{"label":"总后台配置","value":2},{"label":"商户后台配置","value":1}]},{"type":"input","field":"configName","value":"%s","title":"配置名称","props":{"type":"text","placeholder":"请输入配置名称"},"validate":[{"message":"请输入配置名称","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"configKey","value":"%s","title":"配置key","props":{"type":"text","placeholder":"请输入配置key"},"validate":[{"message":"请输入配置key","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"info","value":"%s","title":"说明","props":{"type":"text","placeholder":"请输入说明"}},{"type":"select","field":"configType","value":"%s","title":"配置类型","props":{"multiple":false,"placeholder":"请选择配置类型"},"validate":[{"message":"请选择配置类型","required":true,"type":"string","trigger":"change"}],"options":[]},{"type":"input","field":"configRule","value":"","title":"规则","props":{"type":"textarea","placeholder":"请输入规则"}},{"type":"inputNumber","field":"sort","value":%d,"title":"排序","props":{"placeholder":"请输入排序"}},{"type":"switch","field":"required","value":%d,"title":"必填","props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}},{"type":"switch","field":"status","value":%d,"title":"是否显示","props":{"activeValue":1,"inactiveValue":0,"inactiveText":"关闭","activeText":"开启"}}],"action":"%s","method":"POST","title":"添加配置","config":{}}`, 0, 2, "", "", "", "", 0, 2, 1, "/admin/config/createConfig")
	}
	err := json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	opts, err := GetConfigCategoriesOptions()
	if err != nil {
		return form, err
	}
	form.Rule[0].Props["options"] = opts
	form.Rule[5].Options = ConfigTypes
	return form, err
}

// CreateConfig
func CreateConfig(m model.SysConfig) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("config_key = ?", m.ConfigKey).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, fmt.Errorf("设置key:%s已经使用", m.ConfigKey)
	}
	err = g.TENANCY_DB.Create(&m).Error
	return m, err
}

// GetConfigByKey
func GetConfigByKey(config_key string) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("config_key = ?", config_key).First(&config).Error
	return config, err
}

// GetConfigByID
func GetConfigByID(id string) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("id = ?", id).First(&config).Error
	return config, err
}

// GetConfigInfoList
func GetConfigInfoList(info request.PageInfo) ([]model.SysConfig, int64, error) {
	var configList []model.SysConfig
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysConfig{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return configList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&configList).Error
	if err != nil {
		return configList, total, err
	}
	// 获取类型名称
	for i := 0; i < len(configList); i++ {
		configList[i].TypeName = GetConfigTypeName(configList[i].ConfigType)
	}
	return configList, total, err
}

// ChangeConfigStatus
func ChangeConfigStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysConfig{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateConfig
func UpdateConfig(m model.SysConfig, id string) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("config_key = ?", m.ConfigKey).Where("id <> ?", id).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, fmt.Errorf("设置key:%s已经使用", m.ConfigKey)
	}
	err = g.TENANCY_DB.Where("id= ?", id).Updates(&m).Error
	return m, err
}

// DeleteConfig
func DeleteConfig(id string) error {
	var config model.SysConfig
	return g.TENANCY_DB.Unscoped().Where("id = ?", id).Delete(&config).Error
}
