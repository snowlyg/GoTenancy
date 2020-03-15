package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func Auth(ctx iris.Context) {
	if userID := sysinit.Sess.Start(ctx).GetInt64Default(sysinit.UserIDKey, 0); userID == 0 {
		ctx.Redirect("/auth/login")
		return
	}

	ctx.Next()
}
