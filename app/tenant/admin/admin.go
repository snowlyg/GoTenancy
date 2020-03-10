package admin

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/media/asset_manager"
	registerviews "github.com/snowlyg/qor-registerviews"
	"go-tenancy/config/application"
	"go-tenancy/config/auth"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

// AssetManager asset manager
var AssetManager *admin.Resource

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/"
	}

	if config.AuthPerfix == "" {
		config.AuthPerfix = "/auth"
	}
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix     string
	AuthPerfix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	Tenant := application.Tenant
	TenantParty := application.IrisApplication.TenantParty

	// 支持 go mod 模式
	pkgnames := map[string][]string{
		"sorting":           {""},
		"seo":               {""},
		"notification":      {""},
		"location":          {""},
		"help":              {""},
		"banner_editor":     {""},
		"admin":             {""},
		"action_bar":        {""},
		"activity":          {""},
		"serializable_meta": {""},
		"worker":            {""},
		"media":             {"/media_library", ""},
		"l10n":              {"/publish", ""},
		"i18n":              {"/exchange_actions", "/inline_edit", ""},
	}
	registerPaths(pkgnames, Tenant)

	// 静态文件加载
	AssetManager = Tenant.AddResource(&asset_manager.AssetManager{}, &admin.Config{
		Invisible: true,
	})

	// Add action bar
	ActionBar = action_bar.New(Tenant)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/"})

	SetupDashboard(Tenant)

	// 使用 `iris.FromStd`创建一个 qor 处理器并覆盖到 iris
	// 注册 admin 路由和静态文件到 iris
	handler := iris.FromStd(Tenant.NewServeMux(app.Config.Prefix))
	TenantParty.Any(app.Config.Prefix, handler)
	TenantParty.Macros().Get("string").RegisterFunc("notHas",
		func(validNames []string) func(string) bool {
			return func(paramValue string) bool {
				for _, validName := range validNames {
					if validName == paramValue {
						return false
					}
				}
				return true
			}
		})
	TenantParty.Macros().Get("string").RegisterFunc("has",
		func(validNames []string) func(string) bool {
			return func(paramValue string) bool {
				for _, validName := range validNames {
					if validName == paramValue {
						return true
					}
				}
				return false
			}
		})
	//子资源 ，例如User Management等等,但是不覆盖登录了相关路由
	TenantParty.Any(app.Config.Prefix+"/{name:string notHas([login,logout,password])}", handler)
	TenantParty.Any(app.Config.Prefix+"/{name:string}/{p:path}", handler)

	// 注册 auth 路由到 iris
	authHandler := iris.FromStd(auth.Auth.NewServeMux())
	TenantParty.Any(app.Config.AuthPerfix+"/{name:string has([login,logout])}", authHandler)
	TenantParty.Any(app.Config.AuthPerfix+"/{name:string}/{p:path}", authHandler) // 提交登陆表单,静态资源
}

// registerPaths 循环注册视图
func registerPaths(pkgnames map[string][]string, Tenant *admin.Admin) {
	for pkgname, subpaths := range pkgnames {
		for _, subpath := range subpaths {
			registerPath(Tenant, pkgname, subpath)
		}
	}
}

// registerPath 注册视图
func registerPath(Tenant *admin.Admin, pkgname, subpath string) {
	if err := Tenant.AssetFS.RegisterPath(registerviews.DetectViewsDir("github.com/qor", pkgname, subpath)); err != nil {
		color.Red(fmt.Sprintf("Admin.AssetFS.RegisterPath  %v/%v %v\n", pkgname, subpath, err))
	}
}
