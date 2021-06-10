package service

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Register 用户注册
func Register(u model.SysUser, authorityType int) (model.SysUser, error) {
	var user model.SysUser
	if !errors.Is(g.TENANCY_DB.
		Where("sys_users.username = ?", u.Username).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&user).Error, gorm.ErrRecordNotFound) { // 判断用户名是否注册
		return user, errors.New("用户名已注册")
	}
	// 否则 附加uuid 密码md5简单加密 注册
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Create(&u).Error
	return u, err
}

// Login 用户登录
func Login(u *model.SysUser, authorityType int) (response.LoginResponse, error) {
	switch {
	case authorityType == multi.AdminAuthority:
		return adminLogin(u)
	case authorityType == multi.TenancyAuthority:
		return tenancyLogin(u)
	case authorityType == multi.GeneralAuthority:
		return generalLogin(u)
	default:
		return response.LoginResponse{
			User:  nil,
			Token: "",
		}, errors.New("用户名或者密码错误")
	}
}

// adminLogin
func adminLogin(u *model.SysUser) (response.LoginResponse, error) {
	var admin response.SysAdminUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.AdminAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_admin_infos.email, sys_admin_infos.phone, sys_admin_infos.nick_name, sys_admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join sys_admin_infos on sys_admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&admin).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, err
	}
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(admin.ID), 10),
		Username:      admin.Username,
		AuthorityId:   admin.AuthorityId,
		AuthorityType: admin.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	if admin.ID == 0 {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return response.LoginResponse{
			User:  admin,
			Token: token,
		}, err
	}

	println(token)

	return response.LoginResponse{
		User:  admin,
		Token: token,
	}, nil
}

// tenancyLogin
func tenancyLogin(u *model.SysUser) (response.LoginResponse, error) {
	var tenancy response.SysTenancyUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.TenancyAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,sys_tenancy_infos.email, sys_tenancy_infos.phone, sys_tenancy_infos.nick_name, sys_tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join sys_tenancy_infos on sys_tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on sys_tenancy_infos.sys_tenancy_id = sys_tenancies.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&tenancy).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, err
	}
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(tenancy.ID), 10),
		Username:      tenancy.Username,
		TenancyId:     tenancy.TenancyId,
		TenancyName:   tenancy.TenancyName,
		AuthorityId:   tenancy.AuthorityId,
		AuthorityType: tenancy.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	if tenancy.ID == 0 {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return response.LoginResponse{
			User:  tenancy,
			Token: token,
		}, err
	}

	return response.LoginResponse{
		User:  tenancy,
		Token: token,
	}, nil
}

// generalLogin
func generalLogin(u *model.SysUser) (response.LoginResponse, error) {
	var general response.SysGeneralUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.GeneralAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_general_infos.email,sys_general_infos.phone,sys_general_infos.nick_name,sys_general_infos.avatar_url,sys_general_infos.sex,sys_general_infos.subscribe,sys_general_infos.open_id,sys_general_infos.union_id,sys_general_infos.country,sys_general_infos.province,sys_general_infos.city,sys_general_infos.id_card,sys_general_infos.is_auth,sys_general_infos.real_name,sys_general_infos.birthday,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join sys_general_infos on sys_general_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&general).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return response.LoginResponse{
			User:  general,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}
	if err != nil {
		return response.LoginResponse{
			User:  general,
			Token: token,
		}, err
	}
	claims := &multi.CustomClaims{
		ID:            strconv.FormatUint(uint64(general.ID), 10),
		Username:      general.Username,
		AuthorityId:   general.AuthorityId,
		AuthorityType: general.AuthorityType,
		LoginType:     multi.LoginTypeWeb,
		AuthType:      multi.AuthPwd,
		CreationDate:  time.Now().Local().Unix(),
		ExpiresIn:     multi.RedisSessionTimeoutWeb.Milliseconds(),
	}

	if general.ID == 0 {
		return response.LoginResponse{
			User:  general,
			Token: token,
		}, errors.New("用户名或者密码错误")
	}

	token, _, err = multi.AuthDriver.GenerateToken(claims)
	if err != nil {
		return response.LoginResponse{
			User:  general,
			Token: token,
		}, err
	}

	return response.LoginResponse{
		User:  general,
		Token: token,
	}, nil
}

// ChangePassword 修改用户密码
func ChangePassword(u *model.SysUser, newPassword string, authorityType int) error {
	var user model.SysUser
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", authorityType).
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("修改失败，原密码与当前账户不符")
	}
	if user.ID == 0 {
		return errors.New("修改失败，原密码与当前账户不符")
	}
	err = g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", user.ID).Update("password", utils.MD5V([]byte(newPassword))).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo) ([]response.SysAdminUser, int64, error) {
	var userList []response.SysAdminUser
	var adminAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.AdminAuthority).Select("authority_id").Find(&adminAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", adminAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_admin_infos.email, sys_admin_infos.phone, sys_admin_infos.nick_name, sys_admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join sys_admin_infos on sys_admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Find(&userList).Error
	return userList, total, err
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo) ([]response.SysTenancyUser, int64, error) {
	var userList []response.SysTenancyUser
	var tenancyAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", tenancyAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_tenancy_infos.email, sys_tenancy_infos.phone, sys_tenancy_infos.nick_name, sys_tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join sys_tenancy_infos on sys_tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join sys_tenancies on sys_tenancy_infos.sys_tenancy_id = sys_tenancies.id").
		Find(&userList).Error
	return userList, total, err
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.PageInfo) ([]response.SysGeneralUser, int64, error) {
	var userList []response.SysGeneralUser
	var generalAuthorityIds []int
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.GeneralAuthority).Select("authority_id").Find(&generalAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", generalAuthorityIds)

	var total int64
	err = db.Count(&total).Error
	if err != nil {
		return userList, total, err
	}
	err = db.Limit(limit).Offset(offset).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, sys_general_infos.email, sys_general_infos.phone, sys_general_infos.nick_name, sys_general_infos.avatar_url,sys_general_infos.sex, sys_general_infos.subscribe,sys_general_infos.open_id,sys_general_infos.union_id,sys_general_infos.country,sys_general_infos.province,sys_general_infos.city,sys_general_infos.id_card,sys_general_infos.is_auth,sys_general_infos.real_name,sys_general_infos.birthday,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join sys_general_infos on sys_general_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Find(&userList).Error
	return userList, total, err
}

// SetUserAuthority  设置一个用户的权限
func SetUserAuthority(id uint, authorityId string) error {
	return g.TENANCY_DB.Model(&model.SysUser{}).Where("id = ?", id).Update("authority_id", authorityId).Error
}

// DeleteUser 删除用户
func DeleteUser(id uint) (err error) {
	var user model.SysUser
	return g.TENANCY_DB.Where("id = ?", id).Delete(&user).Error
}

// SetUserAdminInfo 设置admin信息
func SetUserAdminInfo(reqUser model.SysAdminInfo, infoId uint, userId string) (model.SysAdminInfo, error) {
	if infoId > 0 {
		reqUser.ID = infoId
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			return reqUser, err
		}
		reqUser.SysUserID = id
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserTenancyInfo 设置商户信息
func SetUserTenancyInfo(reqUser model.SysTenancyInfo, infoId uint, userId string) (model.SysTenancyInfo, error) {
	if infoId > 0 {
		reqUser.ID = infoId
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			return reqUser, err
		}
		reqUser.SysUserID = id
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserGeneralInfo 设置普通用户信息
func SetUserGeneralInfo(reqUser model.SysGeneralInfo, infoId uint, userId string) (model.SysGeneralInfo, error) {
	if infoId > 0 {
		reqUser.ID = infoId
		err := g.TENANCY_DB.Updates(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	} else {
		id, err := strconv.Atoi(userId)
		if err != nil {
			return reqUser, err
		}
		reqUser.SysUserID = id
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// FindUserById 通过id获取用户信息
func FindUserById(id string) (*model.SysUser, error) {
	var u model.SysUser
	err := g.TENANCY_DB.Where("`id` = ?", id).Preload("Authority").Preload("AdminInfo").Preload("TenancyInfo").Preload("GeneralInfo").First(&u).Error
	return &u, err
}

// DelToken 删除token
func DelToken(token string) error {
	err := multi.AuthDriver.DelUserTokenCache(token)
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		return fmt.Errorf("del token %w", err)
	}
	return nil
}

// CleanToken 清空 token
func CleanToken(userId string) error {
	err := multi.AuthDriver.CleanUserTokenCache(userId)
	if err != nil {
		g.TENANCY_LOG.Error("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}
