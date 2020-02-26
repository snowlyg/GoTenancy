package admin

import (
	"fmt"

	"GoTenancy/config/application"
	"GoTenancy/config/i18n"
	"GoTenancy/models/settings"
	"GoTenancy/utils/registerviews"
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

	// 静态文件加载
	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// 增加媒体库
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"站点管理"}})

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// 增加翻译
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"站点管理"}, Priority: -1})

	// 增加设置
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "店铺设置", Menu: []string{"站点管理"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	SetupSEO(Admin)
	SetupDashboard(Admin)

	// 使用 `iris.FromStd`创建一个 qor 处理器并覆盖到 iris
	// 注册 admin 路由和静态文件到 iris
	// 静态文件路由,可以使用 IrisApp.HandleDir() 替换
	handler := iris.FromStd(Admin.NewServeMux(app.Config.Prefix))
	application.IrisApp.Any(app.Config.Prefix, handler)

}
