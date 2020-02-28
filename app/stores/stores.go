package stores

import (
	"strings"

	"github.com/qor/admin"
	"github.com/qor/qor"
	"github.com/qor/qor/resource"
	"github.com/qor/qor/utils"
	"github.com/qor/validations"
	"go-tenancy/config/application"
	"go-tenancy/models/stores"
)

// New 新建 stores app
func New(config *Config) *App {
	return &App{Config: config}
}

// App stores app
type App struct {
	Config *Config
}

// Config stores 配置结构
type Config struct {
}

// ConfigureApplication 配置 application
func (app App) ConfigureApplication(application *application.Application) {
	app.ConfigureAdmin(application.Admin)
}

// ConfigureAdmin 配置 admin 接口
func (App) ConfigureAdmin(Admin *admin.Admin) {
	// Add Store
	store := Admin.AddResource(&stores.Store{}, &admin.Config{Menu: []string{"商店管理"}})
	store.Meta(&admin.Meta{Name: "Owner", Type: "single_edit"})
	store.AddValidator(&resource.Validator{
		Handler: func(record interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if meta := metaValues.Get("Name"); meta != nil {
				if name := utils.ToString(meta.Value); strings.TrimSpace(name) == "" {
					return validations.NewError(record, "Name", "Name can't be blank")
				}
			}
			return nil
		},
	})
}
