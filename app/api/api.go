package api

import (
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"go-tenancy/config/application"
	"go-tenancy/config/db"
)

// New new api app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/api"
	}
	return &App{Config: config}
}

// App api app
type App struct {
	Config *Config
}

// Config api config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	API := admin.New(&qor.Config{DB: db.DB})

	//API.AddResource(&rabc.RabcUser{})
	//API.AddResource(&rabc.Role{})
	//API.AddResource(&rabc.Permission{})
	//API.AddResource(&rabc.OauthToken{})

	application.IrisApp.Any(app.Config.Prefix, iris.FromStd(API.NewServeMux(app.Config.Prefix)))
}
