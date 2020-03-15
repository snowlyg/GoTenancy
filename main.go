package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/controllers"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/services"
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

	app.Get("/init", controllers.GetMenus)

	home := mvc.New(app.Party("/"))
	home.Router.Use(middleware.Auth)
	home.Handle(new(controllers.Homecontroller))

	control := mvc.New(app.Party("/control"))
	control.Router.Use(middleware.Auth)
	control.Handle(new(controllers.Controlcontroller))

	userService := services.NewUserService(sysinit.Db)
	auth := mvc.New(app.Party("/auth"))
	auth.Register(
		userService,
		sysinit.Sess.Start,
	)
	auth.Handle(new(controllers.AuthController))

	if err := app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
		iris.WithTimeFormat("Mon, 01 Jan 2006 15:04:05 GMT"),
	); err != nil {
		fmt.Println("App is closed")
	}
}
