package stores

import (
	"strings"

	"GoTenancy/config/application"
	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/validations"
	"GoTenancy/models/stores"
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
	store := Admin.AddResource(&stores.Store{}, &admin.Config{Menu: []string{"Store Management"}})
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
