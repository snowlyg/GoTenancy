package auth

import (
	"fmt"
	"net/http"

	"GoTenancy/models/users"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/roles"
)

func init() {
	roles.Register("admin", func(req *http.Request, currentUser interface{}) bool {
		return currentUser != nil && currentUser.(*users.User).Role == "Admin"
	})
}

type PathConfig struct {
	Login  string
	Logout string
}

type AdminAuth struct {
	Paths PathConfig
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
