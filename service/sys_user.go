package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"gorm.io/gorm"
)

// Register 用户注册
func Register(u model.SysUser) (err error, userInter model.SysUser) {
	var user model.SysUser
	if !errors.Is(g.TENANCY_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return errors.New("用户名已注册"), userInter
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err = g.TENANCY_DB.Create(&u).Error
	return err, u
}

// Login 用户登录
func Login(u *model.SysUser) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = g.TENANCY_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return err, &user
}

// ChangePassword 修改用户密码
func ChangePassword(u *model.SysUser, newPassword string) (err error, userInter *model.SysUser) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err = g.TENANCY_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return err, u
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	var userList []model.SysUser
	var adminAuthorityIds []int
	err = g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.AdminAuthority).Select("authority_id").Find(&adminAuthorityIds).Error
	if err != nil {
		return err, userList, total
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", adminAuthorityIds)

	err = db.Count(&total).Error
	if err != nil {
		return err, userList, total
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("AdminInfo").Find(&userList).Error
	return err, userList, total
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	var userList []model.SysUser
	var tenancyAuthorityIds []int
	err = g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return err, userList, total
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", tenancyAuthorityIds)
	err = db.Count(&total).Error
	if err != nil {
		return err, userList, total
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("TenancyInfo").Find(&userList).Error
	return err, userList, total
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	var userList []model.SysUser
	var generalAuthorityIds []int
	err = g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.GeneralAuthority).Select("authority_id").Find(&generalAuthorityIds).Error
	if err != nil {
		return err, userList, total
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", generalAuthorityIds)
	err = db.Count(&total).Error
	if err != nil {
		return err, userList, total
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("GeneralInfo").Find(&userList).Error
	return err, userList, total
}

// SetUserAuthority  设置一个用户的权限
func SetUserAuthority(id float64, authorityId string) (err error) {
	err = g.TENANCY_DB.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityId).Error
	return err
}

// DeleteUser 删除用户
func DeleteUser(id float64) (err error) {
	var user model.SysUser
	err = g.TENANCY_DB.Where("id = ?", id).Delete(&user).Error
	return err
}

// SetUserAdminInfo 设置admin信息
func SetUserAdminInfo(reqUser model.SysAdminInfo, update bool) (err error, user model.SysAdminInfo) {
	if update {
		err = g.TENANCY_DB.Updates(&reqUser).Error
	} else {
		err = g.TENANCY_DB.Create(&reqUser).Error
	}
	return err, reqUser
}

// SetUserTenancyInfo 设置商户信息
func SetUserTenancyInfo(reqUser model.SysTenancyInfo, update bool) (err error, user model.SysTenancyInfo) {
	if update {
		err = g.TENANCY_DB.Updates(&reqUser).Error
	} else {
		err = g.TENANCY_DB.Create(&reqUser).Error
	}
	return err, reqUser
}

// SetUserGeneralInfo 设置普通用户信息
func SetUserGeneralInfo(reqUser model.SysGeneralInfo, update bool) (err error, user model.SysGeneralInfo) {
	if update {
		err = g.TENANCY_DB.Updates(&reqUser).Error
	} else {
		err = g.TENANCY_DB.Create(&reqUser).Error
	}
	return err, reqUser
}

// FindUserById 通过id获取用户信息
func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").Preload("AdminInfo").Preload("TenancyInfo").Preload("GeneralInfo").First(&u).Error
	return err, &u
}
