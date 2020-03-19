package middleware

import (
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/go-tenancy/sysinit"

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
	userId := sysinit.Sess.Start(ctx).GetInt64Default(sysinit.UserIDKey, 0)
	if userId == 0 {
		ctx.Redirect("/auth/login")
		return
	}

	if !c.Check(ctx.Request(), strconv.FormatInt(userId, 10)) {
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

	ok, _ := c.enforcer.Enforce(userId, path, method)
	return ok
}
