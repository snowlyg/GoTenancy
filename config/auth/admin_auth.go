package auth

import (
	"fmt"
	"net/http"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
	"go-tenancy/libs"
	"go-tenancy/models/users"
)

func init() {
	// 注册超级管理员
	roles.Register("super_admin", func(req *http.Request, currentUser interface{}) bool {
		return libs.InStringSlices(req.RemoteAddr, []string{"127.0.0.1", "admin.gotenant.com"}) && currentUser != nil && currentUser.(*users.User).Role == "SuperAdmin"
	})
}

func NewAdminAuth(config *PathConfig) *AdminAuth {
	if config.Admin == "" {
		config.Admin = "/"
	}
	if config.Login == "" {
		config.Login = "/auth/login"
	}
	if config.Logout == "" {
		config.Logout = "/auth/logout"
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
