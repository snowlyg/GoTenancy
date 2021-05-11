package service

import (
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/response"
)

// GetRegion 获取行政区域
func GetRegion(pCode int) ([]response.SysRegion, error) {
	var regions []response.SysRegion
	err := g.TENANCY_DB.Where("p_code", pCode).Find(&regions).Error
	return regions, err
}
