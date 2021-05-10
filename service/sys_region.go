package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
)

// GetRegion 获取行政区域
func GetRegion(p_code int) ([]model.SysRegion, error) {
	var regions []model.SysRegion
	err := g.TENANCY_DB.Where("p_code", p_code).Preload("SubRegions.SubRegions").Find(&regions).Error
	return regions, err
}
