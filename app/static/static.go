package static

import (
	"net/http"
	"strings"

	"github.com/kataras/iris/v12"
	"go-tenancy/config/application"
)

// New 新建 static app
func New(config *Config) *App {
	return &App{Config: config}
}

// App static app
type App struct {
	Config *Config
}

// Config static 配置结构
type Config struct {
	Prefixs []string
	Handler http.Handler
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	for _, prefix := range app.Config.Prefixs {
		if app.Config.Handler != nil {
			application.IrisApp.Any("/"+strings.TrimPrefix(prefix, "/")+"/{p:path}", iris.FromStd(app.Config.Handler))
		}
		application.IrisApp.HandleDir("/"+strings.TrimPrefix(prefix, "/"), "./public/"+strings.TrimPrefix(prefix, "/"))
	}
}
