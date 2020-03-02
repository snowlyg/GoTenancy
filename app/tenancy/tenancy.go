package tenancy

import (
	"fmt"

	"github.com/qor/admin"
	"github.com/qor/render"
	"go-tenancy/config/application"
	"go-tenancy/models/rabc"
	"go-tenancy/models/tenancy"
	"go-tenancy/utils/funcmapmaker"
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

}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	Admin.AddMenu(&admin.Menu{Name: "Tenancy Management", Priority: 1})
	rabcUser := Admin.AddResource(&tenancy.RabcUser{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	rabcUser.Meta(&admin.Meta{Name: "RabcUsers", Type: "select_many"})
	rabcRole := Admin.AddResource(&tenancy.RabcRole{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	rabcPermission := Admin.AddResource(&tenancy.RabcPermission{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	rabcUser := Admin.AddResource(&tenancy.OauthToken{}, &admin.Config{Menu: []string{"Tenancy Management"}})

	// Add Tenancy
	tenant := Admin.AddResource(&tenancy.Tenant{}, &admin.Config{Menu: []string{"Tenancy Management"}})

	tenant.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if tenant, ok := record.(*tenancy.Tenant); ok {
				return fmt.Sprintf("/tenancy/%v", tenant.ID)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})
}
