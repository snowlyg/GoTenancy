package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/lib"
	"github.com/snowlyg/go-tenancy/sysinit"
	"net/http"

	"github.com/casbin/casbin/v2"
)

/*
	Updated for the casbin 2.x released 3 days ago.
	2019-07-15
*/

// New returns the auth service which receives a casbin enforcer.
//
// Adapt with its `Wrapper` for the entire application
// or its `ServeHTTP` for specific routes or parties.
func New(e *casbin.Enforcer) *Casbin {
	return &Casbin{enforcer: e}
}

// ServeHTTP 修改此方法，同时验证用户身份和授权pa
func (c *Casbin) ServeHTTP(ctx context.Context) {
	authCookie := ctx.GetCookie(common.UserCookieName, iris.CookieDecode(sysinit.SC.Decode))
	if len(authCookie) == 0 {
		ctx.Redirect("/auth/login")
		return
	}

	common.GetAuthInfo(authCookie)

	if common.AuthUserId == 0 {
		ctx.Redirect("/auth/login")
		return
	}

	if !c.Check(ctx.Request(), lib.UintToString(common.AuthUserId)) {
		ctx.StatusCode(http.StatusForbidden) // Status Forbidden
		ctx.StopExecution()
		return
	}

	ctx.Next()
}

// Casbin is the auth services which contains the casbin enforcer.
type Casbin struct {
	enforcer *casbin.Enforcer
}

// Check checks the userId, request's method and path and
// returns true if permission grandted otherwise false.
func (c *Casbin) Check(r *http.Request, userId string) bool {
	method := r.Method
	path := r.URL.Path

	ok, err := c.enforcer.Enforce(userId, path, method)
	if err != nil {
		panic(err)
	}
	return ok
}
