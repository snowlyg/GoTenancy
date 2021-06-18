package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

func LoginTenancy(id uint) (response.LoginTenancy, error) {
	var loginTenancy response.LoginTenancy
	var token string
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Where("sys_tenancies.id = ?", id).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,sys_tenancy_infos.email, sys_tenancy_infos.phone, sys_tenancy_infos.nick_name, sys_tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join sys_tenancy_infos on sys_tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on sys_tenancy_infos.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&loginTenancy.Admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return loginTenancy, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return loginTenancy, err
	}
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(loginTenancy.Admin.ID), 10),
		Username:      loginTenancy.Admin.Username,
		TenancyId:     loginTenancy.Admin.TenancyId,
		TenancyName:   loginTenancy.Admin.TenancyName,
		AuthorityId:   loginTenancy.Admin.AuthorityId,
		AuthorityType: loginTenancy.Admin.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	if loginTenancy.Admin.ID == 0 {
		return loginTenancy, errors.New("用户名或者密码错误")
	}

	token, exp, err := multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return loginTenancy, err
	}
	loginTenancy.Token = token
	loginTenancy.Exp = exp
	loginTenancy.Url = "http://localhost:9528/merchant"

	return loginTenancy, nil
}

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
func GetTenancyByID(id uint) (model.SysTenancy, error) {
	var tenancy model.SysTenancy
	err := g.TENANCY_DB.Where("id = ?", id).First(&tenancy).Error
	return tenancy, err
}

// SetTenancyRegionByID
func SetTenancyRegionByID(regionCode request.SetRegionCode) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", regionCode.Id).Update("sys_region_code", regionCode.SysRegionCode).Error
}

// ChangeTenancyStatus
func ChangeTenancyStatus(changeStatus request.ChangeStatus) error {
	return g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", changeStatus.Id).Update("status", changeStatus.Status).Error
}

// UpdateTenancy
func UpdateTenancy(tenancy model.SysTenancy, id uint) (model.SysTenancy, error) {
	err := g.TENANCY_DB.Where("name = ?", tenancy.Name).Not("id = ?", id).First(&tenancy).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return tenancy, errors.New("名称已被注冊")
	}

	err = g.TENANCY_DB.Where("id = ?", id).Omit("uuid").Updates(&tenancy).Error
	return tenancy, err
}

// UpdateClientTenancy
func UpdateClientTenancy(req request.UpdateClientTenancy, id uint) error {

	err := g.TENANCY_DB.Model(&model.SysTenancy{}).
		Where("id = ?", id).Omit("uuid").
		Updates(map[string]interface{}{"avatar": req.Avatar, "banner": req.Banner, "info": req.Info, "tele": req.Tele, "state": req.State}).Error
	return err
}

// DeleteTenancy
func DeleteTenancy(id uint) error {
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

// GetTenancyInfo
func GetTenancyInfo(ctx *gin.Context) (response.TenancyInfo, error) {
	var info response.TenancyInfo
	err := g.TENANCY_DB.Model(&model.SysTenancy{}).Where("id = ?", multi.GetTenancyId(ctx)).Find(&info).Error
	return info, err
}

// GetUpdateTenancyMap
func GetUpdateTenancyMap(ctx *gin.Context) (Form, error) {
	var form Form
	var formStr string
	id := multi.GetTenancyId(ctx)
	tenancy, err := GetTenancyByID(id)
	if err != nil {
		return form, err
	}
	formStr = fmt.Sprintf(`{"rule":[{"type":"input","field":"info","value":"%s","title":"店铺简介","props":{"type":"textarea","placeholder":"请输入店铺简介"},"validate":[{"message":"请输入店铺简介","required":true,"type":"string","trigger":"change"}]},{"type":"input","field":"tele","value":"%s","title":"服务电话","props":{"type":"text","placeholder":"请输入服务电话"},"validate":[{"message":"请输入服务电话","required":true,"type":"string","trigger":"change"}]},{"type":"frame","field":"banner","value":"%s","title":"店铺Banner(710*200px)","props":{"type":"image","maxLength":1,"title":"请选择店铺Banner(710*200px)","src":"\/merchant\/setting\/uploadPicture?field=banner&type=1","modal":{"modal":false},"width":"896px","height":"480px","footer":false}},{"type":"frame","field":"avatar","value":"%s","title":"店铺头像(120*120px)","props":{"type":"image","maxLength":1,"title":"请选择店铺头像(120*120px)","src":"\/merchant\/setting\/uploadPicture?field=avatar&type=1","modal":{"modal":false},"width":"896px","height":"480px","footer":false}},{"type":"switch","field":"state","value":%d,"title":"是否开启","col":{"span":12},"props":{"activeValue":1,"inactiveValue":2,"inactiveText":"关闭","activeText":"开启"}}],"action":"%s/%d","method":"POST","title":"编辑店铺信息","config":{}}`, tenancy.Info, tenancy.Tele, tenancy.Banner, tenancy.Avatar, tenancy.State, "/client/tenancy/updateTenancy", id)

	err = json.Unmarshal([]byte(formStr), &form)
	if err != nil {
		return form, err
	}
	return form, err
}
