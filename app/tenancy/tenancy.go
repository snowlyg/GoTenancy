package tenancy

import (
	"github.com/qor/admin"
	"go-tenancy/config/application"
	"go-tenancy/models/tenancy"
)

// New new tenancy app
func New(config *Config) *App {
	return &App{Config: config}
}

// App tenancy app
type App struct {
	Config *Config
}

// Config tenancy config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	app.ConfigureAdmin(application.Admin)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	Admin.AddMenu(&admin.Menu{Name: "Tenancy Management", Priority: 1})
	_ = Admin.AddResource(&tenancy.RabcUser{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	_ = Admin.AddResource(&tenancy.RabcRole{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	_ = Admin.AddResource(&tenancy.RabcPermission{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	_ = Admin.AddResource(&tenancy.OauthToken{}, &admin.Config{Menu: []string{"Tenancy Management"}})

	// Add Tenancy
	tenant := Admin.AddResource(&tenancy.Tenant{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	tenant.Meta(&admin.Meta{Name: "RabcUsers", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})
}
