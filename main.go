package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/SysInit"
	"github.com/snowlyg/go-tenancy/controllers"
	"github.com/snowlyg/go-tenancy/database"
	"github.com/snowlyg/go-tenancy/repositories"
	"github.com/snowlyg/go-tenancy/services"
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
		if err := ctx.View("shared/error.html"); err != nil {
			panic(err)
		}
	})

	db, err := database.LoadUsers(database.Memory)
	if err != nil {
		app.Logger().Fatalf("error while loading the users: %v", err)
		return
	}
	repo := repositories.NewUserRepository(db)
	userService := services.NewUserService(repo)

	auth := mvc.New(app.Party("/auth"))
	auth.Register(
		userService,
		SysInit.Sess.Start,
	)
	auth.Handle(new(controllers.AuthController))

	if err := app.Run(
		iris.Addr("localhost:8080"),
		iris.WithConfiguration(iris.YAML("./config/conf.yml")),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,
	); err != nil {
		fmt.Println("App is closed")
	}
}
