package home

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"go-tenancy/config/application"
)

// New new home app
func New(config *Config) *App {
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
}

// ConfigureApplication configure application
func (App) ConfigureApplication(application *application.Application) {

	Register(application.IrisApp)

	application.IrisApp.HandleDir("/static", "app/home/views/static")
	application.IrisApp.RegisterView(iris.HTML("./app/home/views", ".html"))
	application.IrisApp.Get("/", func(ctx iris.Context) {
		if err := ctx.View("index.html"); err != nil {
			color.Red(fmt.Sprintf("Home Index View error: %v\n", err))
		}
	})
}
