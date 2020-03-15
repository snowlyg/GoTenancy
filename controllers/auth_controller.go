package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/api_types"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type AuthController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *AuthController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(sysinit.UserIDKey, 0)
	return userID
}

func (c *AuthController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *AuthController) logout() {
	c.Session.Destroy()
}

var loginStaticView = mvc.View{
	Name: "login.html",
	Data: iris.Map{"Title": "登陆"},
}

// GetLogin handles GET: http://localhost:8080/auth/login.
func (c *AuthController) GetLogin() mvc.Result {
	c.Ctx.ViewLayout(iris.NoLayout)
	if c.isLoggedIn() {
		c.logout()
	}

	return loginStaticView
}

// PostLogin handles GET: http://localhost:8080/auth/login.
func (c *AuthController) PostLogin() interface{} {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	user, found := c.Service.GetByUsernameAndPassword(username, password)

	if !found {
		return api_types.Response{Msg: "用户名或者密码错误"}
	}

	c.Session.Set(sysinit.UserIDKey, user.ID)

	return api_types.Response{Status: true, Msg: "登陆成功", Data: user}
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

	return api_types.Response{Status: true, Msg: "退出登录"}
}
