package tenant

import (
	"fmt"
	"strconv"
	"time"

	"github.com/fatih/color"
	"github.com/iris-contrib/middleware/jwt"
	"github.com/jameskeane/bcrypt"
	"github.com/jinzhu/gorm"
	"go-tenancy/app/home/homevalidates"
	"go-tenancy/config/db"
	"go-tenancy/libs"
	"go-tenancy/utils"
)

type RabcUser struct {
	gorm.Model

	Name     string `gorm:"not null VARCHAR(191)"`
	Username string `gorm:"unique;VARCHAR(191)"`
	Password string `gorm:"not null VARCHAR(191)"`
}

func NewRabcUser(id uint, username string) *RabcUser {
	return &RabcUser{
		Model: gorm.Model{
			ID:        id,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: username,
	}
}

func NewRabcUserByStruct(ru *homevalidates.CreateUpdateRabcUserRequest) *RabcUser {
	return &RabcUser{
		Model: gorm.Model{
			ID:        0,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Username: ru.Username,
		Name:     ru.Name,
		Password: libs.HashPassword(ru.Password),
	}
}

func (u *RabcUser) GetRabcUserByUsername() {
	_ = db.DB.Where("username = ?", u.Username).First(u).Error
}

func (u *RabcUser) GetRabcUserById() {
	_ = db.DB.Where("id = ?", u.ID).First(u).Error
}

/**
* 通过 id 删除用户
* @method DeleteRabcUserById
 */
func (u *RabcUser) DeleteRabcUser() {
	if err := db.DB.Delete(u).Error; err != nil {
		color.Red(fmt.Sprintf("DeleteRabcUserByIdErr:%s \n ", err))
	}
}

/**
* 获取所有的账号
* @method GetAllRabcUser
* @param  {[type]} name string [description]
* @param  {[type]} username string [description]
* @param  {[type]} orderBy string [description]
* @param  {[type]} offset int    [description]
* @param  {[type]} limit int    [description]
 */
func GetAllRabcUsers(name, orderBy string, offset, limit int) []*RabcUser {
	var users []*RabcUser
	q := utils.GetAll(name, orderBy, offset, limit)
	if err := q.Find(&users).Error; err != nil {
		color.Red(fmt.Sprintf("GetAllRabcUserErr:%s \n ", err))
		return nil
	}
	return users
}

/**
* 创建
* @method CreateRabcUser
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (u *RabcUser) CreateRabcUser(roleIds []uint) error {
	if err := db.DB.Create(u).Error; err != nil {
		color.Red(fmt.Sprintf("CreateRabcUserErr:%s \n ", err))
		return err
	}

	addRoles(roleIds, u)

	return nil
}

/**
* 更新
* @method UpdateRabcUser
* @param  {[type]} kw string [description]
* @param  {[type]} cp int    [description]
* @param  {[type]} mp int    [description]
 */
func (u *RabcUser) UpdateRabcUser(uj *homevalidates.CreateUpdateRabcUserRequest) {
	uj.Password = libs.HashPassword(uj.Password)
	if err := utils.Update(u, uj); err != nil {
		color.Red(fmt.Sprintf("UpdateRabcUserErr:%s \n ", err))
	}

	addRoles(uj.RoleIds, u)
}

func addRoles(roleIds []uint, user *RabcUser) {
	if len(roleIds) > 0 {
		userId := strconv.FormatUint(uint64(user.ID), 10)
		if _, err := db.GetCasbinEnforcer().DeleteRolesForUser(userId); err != nil {
			color.Red(fmt.Sprintf("CreateRabcUserErr:%s \n ", err))
		}

		for _, roleId := range roleIds {
			roleId := strconv.FormatUint(uint64(roleId), 10)
			if _, err := db.GetCasbinEnforcer().AddRoleForUser(userId, roleId); err != nil {
				color.Red(fmt.Sprintf("CreateRabcUserErr:%s \n ", err))
			}
		}
	}
}

/**
* 判断用户是否登录
* @method CheckLogin
* @param  {[type]}  id       int    [description]
* @param  {[type]}  password string [description]
 */
func (u *RabcUser) CheckLogin(password string) (*Token, bool, string) {
	if u.ID == 0 {
		return nil, false, "用户不存在"
	} else {
		if ok := bcrypt.Match(password, u.Password); ok {
			token := jwt.NewTokenWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
				"iat": time.Now().Unix(),
			})
			tokenString, _ := token.SignedString([]byte("HS2JDFKhu7Y1av7b"))

			oauthToken := new(OauthToken)
			oauthToken.Token = tokenString
			oauthToken.UserId = u.ID
			oauthToken.Secret = "secret"
			oauthToken.Revoked = false
			oauthToken.ExpressIn = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			oauthToken.CreatedAt = time.Now()

			response := oauthToken.OauthTokenCreate()

			return response, true, "登陆成功"
		} else {
			return nil, false, "用户名或密码错误"
		}
	}
}

/**
* 用户退出登陆
* @method RabcUserAdminLogout
* @param  {[type]} ids string [description]
 */
func RabcUserAdminLogout(userId uint) bool {
	ot := OauthToken{}
	ot.UpdateOauthTokenByRabcUserId(userId)
	return ot.Revoked
}
