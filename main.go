package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
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
		ctx.ViewData("Message", ctx.Values().
			GetStringDefault("message", "The page you're looking for doesn't exist"))
		if err := ctx.View("shared/404.html"); err != nil {
			panic(err)
		}
	})

	app.Get("/init", controllers.GetInitMenus)
	app.Get("/menus", controllers.GetMenus)

	home := mvc.New(app.Party("/"))
	home.Router.Use(middleware.Auth)
	home.Handle(new(controllers.HomeController))

	control := mvc.New(app.Party("/control"))
	control.Router.Use(middleware.Auth)
	control.Handle(new(controllers.ControlController))

	menu := mvc.New(app.Party("/menu"))
	menu.Register(
		sysinit.PermService,
	)
	menu.Router.Use(middleware.Auth)
	menu.Handle(new(controllers.MenuController))

	iris.RegisterOnInterrupt(func() {
		_ = sysinit.Db.Close()
	})

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
		iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT"),
	); err != nil {
		fmt.Println("App is closed")
	}
}
