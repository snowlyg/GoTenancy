package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

// SaveConfigValue
// TODO: configKey 没有使用，可用于过滤参数
func SaveConfigValue(values map[string]string, configKey string) error {
	for key, value := range values {
		configValue := model.SysConfigValue{}
		err := g.TENANCY_DB.Where("config_key = ?", key).First(&configValue).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			g.TENANCY_DB.Model(&model.SysConfigValue{}).Where("id = ?", configValue.ID).Update("value", value)
		} else {
			g.TENANCY_DB.Create(&model.SysConfigValue{ConfigKey: key, Value: value})
		}
	}
	return nil
}
