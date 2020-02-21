package static

import (
	"net/http"
	"strings"

	"GoTenancy/config/application"
	"github.com/kataras/iris/v12"
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
	Prefixs []string
	Handler http.Handler
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	for _, prefix := range app.Config.Prefixs {
		application.IrisApp.Get("/"+strings.TrimPrefix(prefix, "/"), iris.FromStd(app.Config.Handler))
	}
}
