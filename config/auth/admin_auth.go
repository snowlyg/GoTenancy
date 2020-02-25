package auth

import (
	"fmt"
	"net/http"
	"time"

	"GoTenancy/models/users"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	claims2 "github.com/qor/auth/claims"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*users.User).Role == "Admin"
	})
}

func NewAdminAuth(config *PathConfig) *AdminAuth {
	if config.Admin == "" {
		config.Admin = "/admin"
	}
	if config.Login == "" {
		config.Login = "/admin/auth/login"
	}
	if config.Logout == "" {
		config.Logout = "/admin/auth/logout"
	}
	return &AdminAuth{Paths: config}
}

type PathConfig struct {
	Admin  string
	Login  string
	Logout string
}

type AdminAuth struct {
	Paths *PathConfig
}

func (a *AdminAuth) LoginURL(c *admin.Context) string {
	return a.Paths.Login
}

func (a *AdminAuth) LogoutURL(c *admin.Context) string {
	return a.Paths.Logout
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	currentUser := Auth.GetCurrentUser(c.Request)
	if currentUser != nil {
		qorCurrentUser, ok := currentUser.(qor.CurrentUser)
		if !ok {
			fmt.Printf("User %#v haven't implement qor.CurrentUser interface\n", currentUser)
		}
		return qorCurrentUser
	}
	return nil
}

// GetLogin simply returns the login page
func (a *AdminAuth) GetLogin(c iris.Context) {
	if _, err := Auth.SessionStorer.Get(c.Request()); err == nil {
		c.Redirect(a.Paths.Admin)
		return
	}

	if err := c.View("login.tmpl"); err != nil {
		color.Red(fmt.Sprintf(" GetLogin - c.View: %v\n", err))
	}
}

// PostLogin is the handler to check if the user can connect
func (a *AdminAuth) PostLogin(c iris.Context) {
	claims := &claims2.Claims{}
	now := time.Now()
	claims.LastLoginAt = &now

	err := Auth.SessionStorer.Update(c.ResponseWriter(), c.Request(), claims)

	if err != nil {
		color.Red(fmt.Sprintf(" PostLogin: %v\n", err))
		c.Redirect(a.Paths.Login)
		return
	}
	c.Redirect(a.Paths.Admin)
}

// GetLogout allows the user to disconnect
func (a *AdminAuth) GetLogout(c iris.Context) {
	if err := Auth.SessionStorer.Delete(c.ResponseWriter(), c.Request()); err != nil {
		color.Red(fmt.Sprintf(" GetLogout %v\n", err))
	}
	c.Redirect(a.Paths.Login)
}
