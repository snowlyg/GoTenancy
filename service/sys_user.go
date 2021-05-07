package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
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

// GetUserInfoList 分页获取数据
func GetUserInfoList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{})
	var userList []model.SysUser
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Authority").Find(&userList).Error
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

// SetUserInfo 设置用户信息
func SetUserInfo(reqUser model.SysUser) (err error, user model.SysUser) {
	err = g.TENANCY_DB.Updates(&reqUser).Error
	return err, reqUser
}

// FindUserById 通过id获取用户信息
func FindUserById(id int) (err error, user *model.SysUser) {
	var u model.SysUser
	err = g.TENANCY_DB.Where("`id` = ?", id).First(&u).Error
	return err, &u
}

// FindUserByUuid 通过uuid获取用户信息
func FindUserByUuid(uuid string) (err error, user *model.SysUser) {
	var u model.SysUser
	if err = g.TENANCY_DB.Where("`uuid` = ?", uuid).First(&u).Error; err != nil {
		return errors.New("用户不存在"), &u
	}
	return nil, &u
}
