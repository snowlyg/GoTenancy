package api

import (
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	"github.com/qor/qor"
	"go-tenancy/config/application"
	"go-tenancy/config/db"
	"go-tenancy/models/tenant"
)

// New new api app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/v1/admin"
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

	API.AddResource(&tenant.RabcUser{})
	//API.AddResource(&rabc.Role{})
	//API.AddResource(&rabc.Permission{})
	//API.AddResource(&rabc.OauthToken{})

	std := iris.FromStd(API.NewServeMux(app.Config.Prefix))
	application.IrisApp.Any(app.Config.Prefix, std)

}
