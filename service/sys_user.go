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
func Register(u model.SysUser) (model.SysUser, error) {
	var user model.SysUser
	if !errors.Is(g.TENANCY_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return user, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Create(&u).Error
	return u, err
}

// Login 用户登录
func Login(u *model.SysUser) (*model.SysUser, error) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Where("username = ? AND password = ?", u.Username, u.Password).Preload("Authority").First(&user).Error
	return &user, err
}

// ChangePassword 修改用户密码
func ChangePassword(u *model.SysUser, newPassword string) (*model.SysUser, error) {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Where("username = ? AND password = ?", u.Username, u.Password).First(&user).Update("password", utils.MD5V([]byte(newPassword))).Error
	return u, err
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo) ([]model.SysUser, int64, error) {
	var userList []model.SysUser
	var adminAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.AdminAuthority).Select("authority_id").Find(&adminAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", adminAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("AdminInfo").Find(&userList).Error
	return userList, total, err
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo) ([]model.SysUser, int64, error) {
	var userList []model.SysUser
	var tenancyAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", tenancyAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("TenancyInfo").Find(&userList).Error
	return userList, total, err
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.PageInfo) ([]model.SysUser, int64, error) {
	var userList []model.SysUser
	var generalAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.GeneralAuthority).Select("authority_id").Find(&generalAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("authority_id IN (?)", generalAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).Preload("Authority").Preload("GeneralInfo").Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority  设置一个用户的权限
func SetUserAuthority(id float64, authorityId string) error {
	return g.TENANCY_DB.Where("id = ?", id).First(&model.SysUser{}).Update("authority_id", authorityId).Error
}

// DeleteUser 删除用户
func DeleteUser(id float64) (err error) {
	var user model.SysUser
	return g.TENANCY_DB.Where("id = ?", id).Delete(&user).Error
}

// SetUserAdminInfo 设置admin信息
func SetUserAdminInfo(reqUser model.SysAdminInfo, update bool) (model.SysAdminInfo, error) {
	if update {
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		err := g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserTenancyInfo 设置商户信息
func SetUserTenancyInfo(reqUser model.SysTenancyInfo, update bool) (model.SysTenancyInfo, error) {
	if update {
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		err := g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserGeneralInfo 设置普通用户信息
func SetUserGeneralInfo(reqUser model.SysGeneralInfo, update bool) (model.SysGeneralInfo, error) {
	if update {
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		err := g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// FindUserById 通过id获取用户信息
func FindUserById(id int) (*model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").Preload("AdminInfo").Preload("TenancyInfo").Preload("GeneralInfo").First(&u).Error
	return &u, err
}
