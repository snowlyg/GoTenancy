package main

import (
	"fmt"
	"time"

	"github.com/dchest/captcha"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/controllers"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func main() {

	app := iris.New()
	app.Logger().SetLevel("debug")

	tmpl := iris.HTML("./views", ".html").
		Layout("shared/layout.html").
		Reload(true)
	app.RegisterView(tmpl)
	app.HandleDir("/public", "./public")

	app.OnAnyErrorCode(func(ctx iris.Context) {
		ctx.ViewData("Status", common.ErrorStatus[ctx.GetStatusCode()])
		if err := ctx.View("shared/error.html"); err != nil {
			panic(err)
		}
	})

	iris.RegisterOnInterrupt(func() {
		_ = sysinit.Db.Close()
	})

	//验证码
	app.Get("/captcha/{id:string}", iris.FromStd(captcha.Server(captcha.StdWidth, captcha.StdHeight)))

	//表格接口
	init := mvc.New(app.Party("/init"))
	init.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	init.Handle(new(controllers.InitController))

	home := mvc.New(app.Party("/"))
	home.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	home.Handle(new(controllers.HomeController))

	control := mvc.New(app.Party("/control"))
	control.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	control.Handle(new(controllers.ControlController))

	menu := mvc.New(app.Party("/menu"))
	menu.Register(sysinit.PermService)
	menu.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	menu.Handle(new(controllers.MenuController))

	role := mvc.New(app.Party("/role"))
	role.Register(sysinit.RoleService)
	role.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	role.Handle(new(controllers.RoleController))

	user := mvc.New(app.Party("/user"))
	user.Register(
		sysinit.UserService,
		sysinit.RoleService,
	)
	user.Router.Use(middleware.New(sysinit.Enforcer).ServeHTTP)
	user.Handle(new(controllers.UserController))

	auth := mvc.New(app.Party("/auth"))
	auth.Register(
		sysinit.UserService,
		sysinit.Sess.Start,
	)
	auth.Handle(new(controllers.AuthController))

	if err := app.Run(
		iris.Addr(fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port)),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat(time.RFC3339),
	); err != nil {
		fmt.Println("App is closed")
	}
}
