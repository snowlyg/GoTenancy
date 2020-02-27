package api

import (
	"GoTenancy/config/application"
	"GoTenancy/config/db"
	"GoTenancy/models/users"
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	"github.com/qor/qor"
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

	API.AddResource(&users.User{})
	// User := API.AddResource(&users.User{})
	// userOrders, _ := User.AddSubResource("Orders")
	// userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})

	application.IrisApp.Any(app.Config.Prefix, iris.FromStd(API.NewServeMux(app.Config.Prefix)))
}
