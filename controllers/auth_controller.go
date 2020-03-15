package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/services"
)

type AuthController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

const userIDKey = "UserID"

func (c *AuthController) getCurrentUserID() int64 {
	userID := c.Session.GetInt64Default(userIDKey, 0)
	return userID
}

func (c *AuthController) isLoggedIn() bool {
	return c.getCurrentUserID() > 0
}

func (c *AuthController) logout() {
	c.Session.Destroy()
}

var registerStaticView = mvc.View{
	Name: "user/register.html",
	Data: iris.Map{"Title": "User Registration"},
}

// GetRegister handles GET: http://localhost:8080/auth/register.
func (c *AuthController) GetRegister() mvc.Result {
	if c.isLoggedIn() {
		c.logout()
	}

	return registerStaticView
}

// PostRegister handles POST: http://localhost:8080/auth/register.
func (c *AuthController) PostRegister() mvc.Result {

	var (
		firstname = c.Ctx.FormValue("firstname")
		username  = c.Ctx.FormValue("username")
		password  = c.Ctx.FormValue("password")
	)

	u, err := c.Service.Create(password, models.User{
		Username:  username,
		Firstname: firstname,
	})

	c.Session.Set(userIDKey, u.ID)

	return mvc.Response{
		Err:  err,
		Path: "/user/me",
	}
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

// PostLogin handles POST: http://localhost:8080/auth/register.
func (c *AuthController) PostLogin() mvc.Result {
	var (
		username = c.Ctx.FormValue("username")
		password = c.Ctx.FormValue("password")
	)

	u, found := c.Service.GetByUsernameAndPassword(username, password)

	if !found {
		return mvc.Response{
			Path: "/user/register",
		}
	}

	c.Session.Set(userIDKey, u.ID)

	return mvc.Response{
		Path: "/user/me",
	}
}

// GetMe handles GET: http://localhost:8080/auth/me.
func (c *AuthController) GetMe() mvc.Result {
	if !c.isLoggedIn() {
		// if it's not logged in then redirect user to the login page.
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
func (c *AuthController) AnyLogout() {
	if c.isLoggedIn() {
		c.logout()
	}

	c.Ctx.Redirect("/auth/login")
}
