package admin

import (
	"GoTenancy/config/application"
	"GoTenancy/config/i18n"
	"GoTenancy/libs/action_bar"
	"GoTenancy/libs/admin"
	"GoTenancy/libs/help"
	"GoTenancy/libs/media/asset_manager"
	"GoTenancy/libs/media/media_library"
	"GoTenancy/models/settings"
	"github.com/kataras/iris/v12"
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

	// 静态文件加载
	AssetManager = Admin.AddResource(&asset_manager.AssetManager{}, &admin.Config{Invisible: true})

	// Add Media Library
	Admin.AddResource(&media_library.MediaLibrary{}, &admin.Config{Menu: []string{"Site Management"}})

	// Add Help
	Help := Admin.NewResource(&help.QorHelpEntry{})
	Help.Meta(&admin.Meta{Name: "Body", Config: &admin.RichEditorConfig{AssetManager: AssetManager}})

	// Add action bar
	ActionBar = action_bar.New(Admin)
	ActionBar.RegisterAction(&action_bar.Action{Name: "Admin Dashboard", Link: "/admin"})

	// Add Translations
	Admin.AddResource(i18n.I18n, &admin.Config{Menu: []string{"Site Management"}, Priority: -1})

	// Add Setting
	Admin.AddResource(&settings.Setting{}, &admin.Config{Name: "Shop Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 1})

	SetupNotification(Admin)
	SetupWorker(Admin)
	SetupSEO(Admin)
	SetupWidget(Admin)
	SetupDashboard(Admin)

	// 使用 `iris.FromStd`创建一个 qor 处理器并覆盖到 iris
	// 注册 admin 路由和静态文件到 iris
	// 静态文件路由,可以使用 IrisApp.HandleDir() 替换
	handler := iris.FromStd(Admin.NewServeMux(app.Config.Prefix))
	application.IrisApp.Any(app.Config.Prefix, handler)
	application.IrisApp.Any(app.Config.Prefix+"/{p:path}", handler)
	//application.IrisApp.HandleDir("/static", "./assets", iris.DirOptions {
	//	Asset: Asset,
	//	AssetInfo: AssetInfo,
	//	AssetNames: AssetNames,
	//	Gzip: false,
	//})
}
