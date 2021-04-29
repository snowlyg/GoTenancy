package service

import (
	"errors"
	"strconv"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// CreateAuthority 创建一个角色
func CreateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(g.TENANCY_DB.Where("authority_id = ?", auth.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), auth
	}
	err = g.TENANCY_DB.Create(&auth).Error
	return err, auth
}

// CopyAuthority 复制一个角色
func CopyAuthority(copyInfo response.SysAuthorityCopyResponse) (err error, authority model.SysAuthority) {
	var authorityBox model.SysAuthority
	if !errors.Is(g.TENANCY_DB.Where("authority_id = ?", copyInfo.Authority.AuthorityId).First(&authorityBox).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同角色id"), authority
	}
	copyInfo.Authority.Children = []model.SysAuthority{}
	err, menus := GetMenuAuthority(&request.GetAuthorityId{AuthorityId: copyInfo.OldAuthorityId})
	var baseMenu []model.SysBaseMenu
	for _, v := range menus {
		intNum, _ := strconv.Atoi(v.MenuId)
		v.SysBaseMenu.ID = uint(intNum)
		baseMenu = append(baseMenu, v.SysBaseMenu)
	}
	copyInfo.Authority.SysBaseMenus = baseMenu
	err = g.TENANCY_DB.Create(&copyInfo.Authority).Error

	paths := GetPolicyPathByAuthorityId(copyInfo.OldAuthorityId)
	err = UpdateCasbin(copyInfo.Authority.AuthorityId, paths)
	if err != nil {
		_ = DeleteAuthority(&copyInfo.Authority)
	}
	return err, copyInfo.Authority
}

// UpdateAuthority 更改一个角色
func UpdateAuthority(auth model.SysAuthority) (err error, authority model.SysAuthority) {
	err = g.TENANCY_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Updates(&auth).Error
	return err, auth
}

// DeleteAuthority 删除角色
func DeleteAuthority(auth *model.SysAuthority) (err error) {
	if !errors.Is(g.TENANCY_DB.Where("authority_id = ?", auth.AuthorityId).First(&model.SysUser{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色有用户正在使用禁止删除")
	}
	if !errors.Is(g.TENANCY_DB.Where("parent_id = ?", auth.AuthorityId).First(&model.SysAuthority{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("此角色存在子角色不允许删除")
	}
	db := g.TENANCY_DB.Preload("SysBaseMenus").Where("authority_id = ?", auth.AuthorityId).First(auth)
	err = db.Unscoped().Delete(auth).Error
	if len(auth.SysBaseMenus) > 0 {
		err = g.TENANCY_DB.Model(auth).Association("SysBaseMenus").Delete(auth.SysBaseMenus)
		//err = db.Association("SysBaseMenus").Delete(&auth)
	} else {
		err = db.Error
	}
	ClearCasbin(0, auth.AuthorityId)
	return err
}

// GetAuthorityInfoList 分页获取数据
func GetAuthorityInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB
	var authority []model.SysAuthority
	err = db.Limit(limit).Offset(offset).Preload("DataAuthorityId").Where("parent_id = 0").Find(&authority).Error
	if len(authority) > 0 {
		for k := range authority {
			err = findChildrenAuthority(&authority[k])
		}
	}
	return err, authority, total
}

// GetAuthorityInfo 获取所有角色信息
func GetAuthorityInfo(auth model.SysAuthority) (err error, sa model.SysAuthority) {
	err = g.TENANCY_DB.Preload("DataAuthorityId").Where("authority_id = ?", auth.AuthorityId).First(&sa).Error
	return err, sa
}

// SetDataAuthority 设置角色资源权限
func SetDataAuthority(auth model.SysAuthority) error {
	var s model.SysAuthority
	g.TENANCY_DB.Preload("DataAuthorityId").First(&s, "authority_id = ?", auth.AuthorityId)
	err := g.TENANCY_DB.Model(&s).Association("DataAuthorityId").Replace(&auth.DataAuthorityId)
	return err
}

// SetMenuAuthority 菜单与角色绑定
func SetMenuAuthority(auth *model.SysAuthority) error {
	var s model.SysAuthority
	g.TENANCY_DB.Preload("SysBaseMenus").First(&s, "authority_id = ?", auth.AuthorityId)
	err := g.TENANCY_DB.Model(&s).Association("SysBaseMenus").Replace(&auth.SysBaseMenus)
	return err
}

// findChildrenAuthority 查询子角色
func findChildrenAuthority(authority *model.SysAuthority) (err error) {
	err = g.TENANCY_DB.Preload("DataAuthorityId").Where("parent_id = ?", authority.AuthorityId).Find(&authority.Children).Error
	if len(authority.Children) > 0 {
		for k := range authority.Children {
			err = findChildrenAuthority(&authority.Children[k])
		}
	}
	return err
}