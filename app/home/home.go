package home

import (
	"github.com/kataras/iris/v12"
	"github.com/qor/render"
	"go-tenancy/config/application"
	"go-tenancy/utils/funcmapmaker"
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
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("home")}, "app/home/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	application.IrisApp.HandleDir("/static", "app/home/views/static")
	application.IrisApp.RegisterView(iris.HTML("app/home/views", ".html"))
	application.IrisApp.Get("/", controller.Index)
	//application.IrisApp.Get("/switch_locale", controller.SwitchLocale)
}
