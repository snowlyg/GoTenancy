package service

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
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
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
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
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at,sys_tenancies.id  as tenancy_id,sys_tenancies.name as tenancy_name,tenancy_infos.email, tenancy_infos.phone, tenancy_infos.nick_name, tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join tenancy_infos on tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_tenancies on tenancy_infos.sys_tenancy_id = sys_tenancies.id").
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
	var general response.GeneralUser
	var token string
	u.Password = utils.MD5V([]byte(u.Password))
	err := g.TENANCY_DB.Model(&model.SysUser{}).
		Where("sys_users.username = ? AND sys_users.password = ?", u.Username, u.Password).
		Where("sys_authorities.authority_type = ?", multi.GeneralAuthority).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, general_infos.email,general_infos.phone,general_infos.nick_name,general_infos.avatar_url,general_infos.sex,general_infos.subscribe,general_infos.open_id,general_infos.union_id,general_infos.country,general_infos.province,general_infos.city,general_infos.id_card,general_infos.is_auth,general_infos.real_name,general_infos.birthday,sys_authorities.authority_name,sys_authorities.authority_type,sys_authorities.default_router,sys_users.authority_id").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id").
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

// ChangeProfile 修改用户信息
func ChangeProfile(user request.ChangeProfile, sysUserId uint) error {
	err := g.TENANCY_DB.Model(&model.AdminInfo{}).Where("sys_user_id = ?", sysUserId).
		Updates(map[string]interface{}{"nick_name": user.NickName, "phone": user.Phone}).Error
	if err != nil {
		return err
	}
	return nil
}

// GetAdminInfoList 分页获取数据
func GetAdminInfoList(info request.PageInfo) ([]response.SysAdminUser, int64, error) {
	var userList []response.SysAdminUser
	var adminAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.AdminAuthority).Select("authority_id").Find(&adminAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", adminAuthorityIds)
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, admin_infos.email, admin_infos.phone, admin_infos.nick_name, admin_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id").
		Joins("left join admin_infos on admin_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Find(&userList).Error
	return userList, total, err
}

// GetTenancyInfoList 分页获取数据
func GetTenancyInfoList(info request.PageInfo) ([]response.SysTenancyUser, int64, error) {
	var userList []response.SysTenancyUser
	var tenancyAuthorityIds []int
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	err := g.TENANCY_DB.Model(&model.SysAuthority{}).Where("authority_type", multi.TenancyAuthority).Select("authority_id").Find(&tenancyAuthorityIds).Error
	if err != nil {
		return userList, 0, err
	}
	db := g.TENANCY_DB.Model(&model.SysUser{}).Where("sys_users.authority_id IN (?)", tenancyAuthorityIds)
	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}
	err = db.
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, tenancy_infos.email, tenancy_infos.phone, tenancy_infos.nick_name, tenancy_infos.header_img,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,sys_tenancies.name as tenancy_name").
		Joins("left join tenancy_infos on tenancy_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join sys_tenancies on tenancy_infos.sys_tenancy_id = sys_tenancies.id").
		Find(&userList).Error
	return userList, total, err
}

// GetGeneralInfoList 分页获取数据
func GetGeneralInfoList(info request.UserPageInfo) ([]response.GeneralUser, int64, error) {
	var userList []response.GeneralUser
	var total int64
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	generalAuthorityIds, err := GetUserAuthorityIds()
	if err != nil {
		return userList, 0, err
	}

	db := g.TENANCY_DB.Model(&model.SysUser{}).
		Select("sys_users.id,sys_users.username,sys_users.authority_id,sys_users.created_at,sys_users.updated_at, general_infos.*,sys_authorities.authority_name,sys_authorities.authority_type,sys_users.authority_id,user_groups.group_name").
		Joins("left join general_infos on general_infos.sys_user_id = sys_users.id").
		Joins("left join sys_authorities on sys_authorities.authority_id = sys_users.authority_id").
		Joins("left join user_groups on general_infos.group_id = user_groups.id").
		Where("sys_users.authority_id IN (?)", generalAuthorityIds)
	if info.UserTimeType != "" && info.UserTime != "" {
		userTimes := strings.Split(info.UserTime, "-")
		start, err := time.Parse("2006/01/02", userTimes[0])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		end, err := time.Parse("2006/01/02", userTimes[1])
		if err != nil {
			return userList, total, fmt.Errorf("parse time %w", err)
		}
		if info.UserTimeType == "add_time" {
			db = db.Where("general_infos.created_at BETWEEN ? AND ?", start, end)
		} else if info.UserTimeType == "visit" {
			db = db.Where("general_infos.last_time BETWEEN ? AND ?", start, end)
		}
	}

	if info.PayCount != "" {
		if info.PayCount == "0" {
			db = db.Where("general_infos.pay_count = ?", info.PayCount)
		} else {
			db = db.Where("general_infos.pay_count >= ?", info.PayCount)
		}
	}
	if info.GroupId != "" {
		db = db.Where("general_infos.group_id = ?", info.GroupId)
	}
	if info.LabelId != "" {
		db = db.Where("general_infos.label_id = ?", info.LabelId)
	}
	if info.Sex != "" {
		db = db.Where("general_infos.sex = ?", info.Sex)
	}
	if info.NickName != "" {
		db = db.Where("general_infos.nick_name like ?", info.NickName+"%")
	}

	if limit > 0 {
		err = db.Count(&total).Error
		if err != nil {
			return userList, total, err
		}
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&userList).Error
	if err != nil {
		return userList, total, err
	}

	if len(userList) > 0 {
		userLabelIds := []uint{}
		for _, user := range userList {
			userLabelIds = append(userLabelIds, user.LabelID)
		}
		userLabels, err := GetUserLabelByIds(userLabelIds)
		if err != nil {
			return userList, total, err
		}
		for i := 0; i < len(userList); i++ {
			for _, userLabel := range userLabels {
				if userList[i].LabelID == userLabel.ID {
					userList[i].Label = append(userList[i].Label, userLabel)
				}
			}
		}
	}

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
func SetUserAdminInfo(reqUser model.AdminInfo, infoId uint, userId string) (model.AdminInfo, error) {
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
		reqUser.SysUserID = uint(id)
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserTenancyInfo 设置商户信息
func SetUserTenancyInfo(reqUser model.TenancyInfo, infoId uint, userId string) (model.TenancyInfo, error) {
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
		reqUser.SysUserID = uint(id)
		err = g.TENANCY_DB.Create(&reqUser).Error
		if err != nil {
			return reqUser, err
		}
	}
	return reqUser, nil
}

// SetUserGeneralInfo 设置普通用户信息
func SetUserGeneralInfo(reqUser model.GeneralInfo, infoId uint, userId string) (model.GeneralInfo, error) {
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
		reqUser.SysUserID = uint(id)
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
