package admin

import (
	"fmt"

	"GoTenancy/config/application"
	"GoTenancy/config/auth"
	registerviews "github.com/snowlyg/qor-registerviews"

	"GoTenancy/models/settings"
	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/qor/action_bar"
	"github.com/qor/admin"
	"github.com/qor/help"
	"github.com/qor/media/asset_manager"
	"github.com/qor/media/media_library"
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

	if err := Admin.AssetFS.RegisterPath(registerviews.DetectViewsDir("github.com/qor", "admin")); err != nil {
		color.Red(fmt.Sprintf("Admin.AssetFS.RegisterPath %v\n", err))
	}
	if err := Admin.AssetFS.RegisterPath(registerviews.DetectViewsDir("github.com/qor", "publish2")); err != nil {
		color.Red(fmt.Sprintf("Admin.AssetFS.RegisterPath %v\n", err))
	}

	// 静态文件加载
	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// 增加媒体库
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Name: "媒体库", Menu: []string{"系统设置"}})
	// 覆盖-系统设置-菜单的 IconName
	Admin.GetMenu("系统设置").IconName = "Site"

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// 增加翻译
	//Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"系统设置"}, IconName: "Site", Priority: -1})

	// 增加设置
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "店铺设置", Menu: []string{"系统设置"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	//SetupSEO(Admin)
	SetupDashboard(Admin)

	// 使用 `iris.FromStd`创建一个 qor 处理器并覆盖到 iris
	// 注册 admin 路由和静态文件到 iris
	handler := iris.FromStd(Admin.NewServeMux(app.Config.Prefix))
	application.IrisApp.Any(app.Config.Prefix, handler)
	application.IrisApp.Macros().Get("string").RegisterFunc("notAuth",
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
	//子资源 ，例如用户管理等等,但是不覆盖登录了相关路由
	application.IrisApp.Any(app.Config.Prefix+"/{name:string notAuth([login,logout,password])}", handler)
	application.IrisApp.Any(app.Config.Prefix+"/{name:string notAuth([login,logout,password])}/{p:path}", handler)

	// 注册 auth 路由和静态文件到 iris
	application.IrisApp.HandleDir("/admin/auth/assets", "public/auth_resource/assets")
	application.IrisApp.HandleDir("/admin/password/auth/assets", "public/auth_resource/assets")

	authHandler := iris.FromStd(auth.Auth.NewServeMux())
	application.IrisApp.Macros().Get("string").RegisterFunc("isAuth",
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
	application.IrisApp.Any(app.Config.Prefix+"/{name:string isAuth([login,logout])}", authHandler)
	application.IrisApp.Any(app.Config.Prefix+"/password/login", authHandler) // 提交登陆表单
}
