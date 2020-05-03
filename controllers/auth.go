package controllers

import (
	"fmt"
	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-authcode-1"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/lib"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

var captchaId = captcha.New()

type AuthController struct {
	Ctx iris.Context
}

func (c *AuthController) isLoggedIn() bool {
	return common.AuthUserId > 0
}

func (c *AuthController) logout() {
	c.Ctx.RemoveCookie(common.UserCookieName)
	common.AuthUserId = 0
	common.AuthUserTenantId = 0
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
		return common.ActionResponse{Status: false, Msg: "请输入正确验证码"}
	}

	user, found := sysinit.UserService.GetByUsername(username)
	if !found && user.ID > 0 {
		return common.ActionResponse{Status: false, Msg: "请输入正确用户名"}
	}

	validatePassword, err := models.ValidatePassword(password, user.Password)
	if !validatePassword {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("密码错误 %v", err)}
	}

	roleIds := lib.UnitJoin(user.RoleIds, ",")
	str := fmt.Sprintf("%d||%s||%d||%s", user.ID, lib.Sha1(user.Password), user.TenantId, roleIds)
	encode := authcode.AuthCode(str, "ENCODE", common.AdminAuthKey, 0)

	c.Ctx.SetCookieKV(common.UserCookieName, encode, iris.CookieEncode(sysinit.SC.Encode))

	return common.ActionResponse{Status: true, Msg: "登陆成功", Data: user}
}

// GetMe handles GET: http://localhost:8080/auth/me.
func (c *AuthController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		return mvc.Response{Path: "/user/login"}
	}

	u, found := sysinit.UserService.GetByID(common.AuthUserId)
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

	return common.ActionResponse{Status: true, Msg: "退出登录"}
}
