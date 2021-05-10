package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
)

// GetSubRegion 获取行政区域
func GetSubRegion(p_code int) (error, []model.SysRegion) {
	var regions []model.SysRegion
	err := g.TENANCY_DB.Where("p_code", p_code).Preload("SubRegions.SubRegions").Find(&regions).Error
	return err, regions
}
