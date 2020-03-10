package admin

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/help"
	"github.com/qor/media/asset_manager"
	"github.com/qor/media/media_library"
	registerviews "github.com/snowlyg/qor-registerviews"
	"github.com/snowlyg/qortenant/backend/database"
	"go-tenancy/config/application"
	"go-tenancy/config/auth"
	"go-tenancy/config/i18n"
	"go-tenancy/models/settings"
)

// ActionBar admin action bar
var ActionBar *action_bar.ActionBar

// AssetManager asset manager
var AssetManager *admin.Resource

// New new home app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/admin"
	}
	return &App{Config: config}
}

// App home app
type App struct {
	Config *Config
}

// Config home config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	Admin := application.Admin

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
	registerPaths(pkgnames, Admin)

	// 静态文件加载
	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// 租户
	tenant := Admin.AddResource(database.Tenant{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: 2})
	tenant.Meta(&admin.Meta{Name: "TUsers", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})
	Admin.AddResource(database.TUser{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	Admin.AddResource(database.TRole{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	Admin.AddResource(database.TPermission{}, &admin.Config{Menu: []string{"Tenancy Management"}})
	Admin.AddResource(database.TOauthToken{}, &admin.Config{Menu: []string{"Tenancy Management"}})

	// 媒体库
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})
	// 覆盖-Site Management-菜单的 IconName
	Admin.GetMenu("Site Management").IconName = "Site"

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// 翻译
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: -1})

	// 设置
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Shop Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	SetupSEO(Admin)
	SetupDashboard(Admin)

	// 使用 `iris.FromStd`创建一个 qor 处理器并覆盖到 iris
	// 注册 admin 路由和静态文件到 iris
	handler := iris.FromStd(Admin.NewServeMux(app.Config.Prefix))
	application.IrisApp.Any(app.Config.Prefix, handler)
	application.IrisApp.Macros().Get("string").RegisterFunc("notHas",
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
	application.IrisApp.Macros().Get("string").RegisterFunc("has",
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
	application.IrisApp.Any(app.Config.Prefix+"/{name:string notHas([login,logout,password])}", handler)
	application.IrisApp.Any(app.Config.Prefix+"/{name:string}/{p:path}", handler)

	// 注册 auth 路由到 iris
	authHandler := iris.FromStd(auth.Auth.NewServeMux())
	application.IrisApp.Any(app.Config.Prefix+"/{name:string has([login,logout])}", authHandler)
	application.IrisApp.Any(app.Config.Prefix+"/password/login", authHandler) // 提交登陆表单
}

// registerPaths 循环注册视图
func registerPaths(pkgnames map[string][]string, Admin *admin.Admin) {
	for pkgname, subpaths := range pkgnames {
		for _, subpath := range subpaths {
			registerPath(Admin, pkgname, subpath)
		}
	}
}

// registerPath 注册视图
func registerPath(Admin *admin.Admin, pkgname, subpath string) {
	if err := Admin.AssetFS.RegisterPath(registerviews.DetectViewsDir("github.com/qor", pkgname, subpath)); err != nil {
		color.Red(fmt.Sprintf("Admin.AssetFS.RegisterPath  %v/%v %v\n", pkgname, subpath, err))
	}
}
