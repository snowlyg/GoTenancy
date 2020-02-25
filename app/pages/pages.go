package pages

import (
	"GoTenancy/config/application"
	"GoTenancy/config/db"
	"GoTenancy/models/blogs"
	"GoTenancy/utils/funcmapmaker"
	"github.com/qor/admin"
	"github.com/qor/page_builder"
	"github.com/qor/render"
	"github.com/qor/widget"
)

// New new page app
func New(config *Config) *App {
	return &App{Config: config}
}

// App page app
type App struct {
	Config *Config
}

// Config page config struct
type Config struct {
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("blog")}, "app/pages/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)
	application.IrisApp.Get("/blog", controller.Index)
}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	pageMenuName := "Pages Management"
	Admin.AddMenu(&admin.Menu{Name: pageMenuName, Priority: 4})

	// Blog Management
	article := Admin.AddResource(&blogs.Article{}, &admin.Config{Menu: []string{pageMenuName}, IconName: "Publish"})
	article.IndexAttrs("ID", "VersionName", "ScheduledStartAt", "ScheduledEndAt", "Author", "Title")

	// Setup pages
	PageBuilderWidgets := widget.New(&widget.Config{DB: db.DB})
	//PageBuilderWidgets.WidgetSettingResource = Admin.NewResource(&adminapp.QorWidgetSetting{}, &admin.Config{Name: "PageBuilderWidgets"})
	PageBuilderWidgets.WidgetSettingResource.NewAttrs(
		&admin.Section{
			Rows: [][]string{{"Kind"}, {"SerializableMeta"}},
		},
	)
	//PageBuilderWidgets.WidgetSettingResource.AddProcessor(&resource.Processor{
	//	Handler: func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
	//		if widgetSetting, ok := value.(*adminapp.QorWidgetSetting); ok {
	//			if widgetSetting.Name == "" {
	//				var count int
	//				context.GetDB().Set(admin.DisableCompositePrimaryKeyMode, "off").Model(&adminapp.QorWidgetSetting{}).Count(&count)
	//				widgetSetting.Name = fmt.Sprintf("%v %v", utils.ToString(metaValues.Get("Kind").Value), count)
	//			}
	//		}
	//		return nil
	//	},
	//})
	Admin.AddResource(PageBuilderWidgets, &admin.Config{Menu: []string{pageMenuName}, IconName: "Publish"})

	page := page_builder.New(&page_builder.Config{
		Admin:       Admin,
		PageModel:   &blogs.Page{},
		Containers:  PageBuilderWidgets,
		AdminConfig: &admin.Config{Name: "Pages", Menu: []string{pageMenuName}, IconName: "Publish", Priority: 1},
	})
	page.IndexAttrs("ID", "Title", "PublishLiveNow")
}
