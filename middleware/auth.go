package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func Auth(ctx iris.Context) {
	if auth, _ := sysinit.Sess.Start(ctx).GetBoolean("authenticated"); !auth {
		ctx.Redirect("/auth/login")
		ctx.WriteString("don't login")
		return
	}
	ctx.WriteString("The cake is a lie!")
	ctx.Next()
}
