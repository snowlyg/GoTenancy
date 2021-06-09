package service

import (
	"errors"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateTenancy
func CreateTenancy(tenancy model.SysTenancy) (model.SysTenancy, error) {
	err := g.TENANCY_DB.Where("name = ?", tenancy.Name).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}
	tenancy.UUID = uuid.NewV4()
	err = g.TENANCY_DB.Create(&tenancy).Error
	return tenancy, err
}

// GetTenancyByID
func GetTenancyByID(id string) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return tenancy, err
}

// SetTenancyRegionByID
func SetTenancyRegionByID(regionCode request.SetRegionCode) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", regionCode.Id).Update("sys_region_code", regionCode.SysRegionCode).Error
}

// ChangeTenancyStatus
func ChangeTenancyStatus(changeStatus request.ChangeTenancyStatus) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateTenany
func UpdateTenany(tenancy model.SysTenancy, id string) (model.SysTenancy, error) {
	err := g.TENANCY_DB.Where("name = ?", tenancy.Name).Not("id = ?", id).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}

	err = g.TENANCY_DB.Where("id = ?", id).Omit("uuid").Updates(&tenancy).Error
	return tenancy, err
}

// DeleteTenancy
func DeleteTenancy(id string) error {
	return g.TENANCY_DB.Where("id = ?", id).Delete(&model.SysTenancy{}).Error
}

// GetTenanciesInfoList
func GetTenanciesInfoList(info request.TenancyPageInfo) ([]response.SysTenancy, int64, error) {
	var tenancyList []response.SysTenancy
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("status = ?", info.Status)
	if info.Keyword != "" {
		db = db.Where(g.TENANCY_DB.Where("name like ?", info.Keyword+"%").Or("tele like ?", info.Keyword+"%"))
	}
	if info.Date != "" {
		db = filterDate(db, info.Date)
	}

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return tenancyList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&tenancyList).Error
	return tenancyList, total, err
}

// GetTenanciesByRegion
func GetTenanciesByRegion(p_code string) ([]response.SysTenancy, error) {
	var tenancyList []response.SysTenancy
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("sys_region_code = ?", p_code).Find(&tenancyList).Error
	return tenancyList, err
}

type Result struct {
	ID   int
	Name string
	Age  int
}

// GetTenancyCount
func GetTenancyCount() (gin.H, error) {
	var counts response.Counts
	err := g.TENANCY_DB.Raw("SELECT sum(case when status = ? then 1 else 0 end) as 'valid',sum(case when status = ? then 1 else 0 end) as 'invalid' FROM sys_tenancies WHERE ISNULL(deleted_at)", g.StatusTrue, g.StatusFalse).Scan(&counts).Error
	return gin.H{
		"invalid": counts.Invalid,
		"valid":   counts.Valid,
	}, err
}
