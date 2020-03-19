package controllers

import (
	"fmt"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
)

var captchaId = captcha.New()

type AuthController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *AuthController) getCurrentUserID() uint {
	userID := c.Session.GetInt64Default(sysinit.UserIDKey, 0)
	return uint(userID)
}

func (c *AuthController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *AuthController) logout() {
	c.Session.Destroy()
}

// GetLogin handles GET: http://localhost:8080/auth/login.
func (c *AuthController) GetLogin() mvc.Result {
	c.Ctx.ViewLayout(iris.NoLayout)
	if c.isLoggedIn() {
		c.logout()
	}

	return mvc.View{
		Name: "auth/login.html",
		Data: iris.Map{"CaptchaId": captchaId},
	}

}

// PostLogin handles GET: http://localhost:8080/auth/login.
func (c *AuthController) PostLogin() interface{} {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
		capId    = c.Ctx.FormValue("captchaId")
	)

	if !captcha.VerifyString(captchaId, capId) {
		return common.Response{Msg: "请输入正确验证码"}
	}

	user, found := c.Service.GetByUsername(username)
	if !found && user.ID > 0 {
		return common.Response{Msg: "请输入正确用户名"}
	}

	validatePassword, err := models.ValidatePassword(password, user.Password)
	if !validatePassword {
		return common.Response{Msg: fmt.Sprintf("密码错误 %v", err)}
	}

	c.Session.Set(sysinit.UserIDKey, user.ID)

	return common.Response{Status: true, Msg: "登陆成功", Data: user}
}

// GetMe handles GET: http://localhost:8080/auth/me.
func (c *AuthController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		return mvc.Response{Path: "/user/login"}
	}

	u, found := c.Service.GetByID(c.getCurrentUserID())
	if !found {
		c.logout()
		return c.GetMe()
	}

	return mvc.View{
		Name: "user/me.html",
		Data: iris.Map{
			"Title": "Profile of " + u.Username,
			"User":  u,
		},
	}
}

// AnyLogout handles All/Any HTTP Methods for: http://localhost:8080/auth/logout.
func (c *AuthController) AnyLogout() interface{} {
	if c.isLoggedIn() {
		c.logout()
	}

	return common.Response{Status: true, Msg: "退出登录"}
}
