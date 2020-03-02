package rabc

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
	controller := &Controller{View: render.New(&render.Config{AssetFileSystem: application.AssetFS.NameSpace("tenancy")}, "app/tenancy/views")}

	funcmapmaker.AddFuncMapMaker(controller.View)
	app.ConfigureAdmin(application.Admin)

	//application.IrisApp.PartyFunc("/admin/account", func(account iris.Party) {
	//	account.Use(middleware.Authorize)
	//	account.Post("/add_user_credit", middleware.AuthorizeloggedInHalfHour, controller.AddCredit) // role: logged_in_half_hour
	//	account.Get("/profile", controller.Profile)
	//	account.Post("/profile", controller.Update)
	//})
	//application.IrisApp.Get("/tenancys", controller.Index)
	//application.IrisApp.Get("/tenancy/{name:string}", controller.Show)

}

// ConfigureAdmin configure admin interface
func (App) ConfigureAdmin(Admin *admin.Admin) {
	Admin.AddMenu(&admin.Menu{Name: "RABC Management", Priority: 1})
	//color := Admin.AddResource(&tenancy.Color{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -5})
	//Admin.AddResource(&tenancy.Size{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -4})
	//Admin.AddResource(&tenancy.Material{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -4})

	//category := Admin.AddResource(&tenancy.Category{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -3})
	//category.Meta(&admin.Meta{Name: "Categories", Type: "select_many"})
	//
	//collection := Admin.AddResource(&tenancy.Collection{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -2})

	// Add TenancyImage as Media Libraray
	//TenancyImagesResource := Admin.AddResource(&tenancy.TenantImage{}, &admin.Config{Menu: []string{"Tenancy Management"}, Priority: -1})

	//TenancyImagesResource.Filter(&admin.Filter{
	//	Name:       "SelectedType",
	//	Label:      "Media Type",
	//	Operations: []string{"contains"},
	//	Config:     &admin.SelectOneConfig{Collection: [][]string{{"video", "Video"}, {"image", "Image"}, {"file", "File"}, {"video_link", "Video Link"}}},
	//})
	//TenancyImagesResource.Filter(&admin.Filter{
	//	Name:   "Color",
	//	Config: &admin.SelectOneConfig{RemoteDataResource: color},
	//})
	//TenancyImagesResource.Filter(&admin.Filter{
	//	Name:   "Category",
	//	Config: &admin.SelectOneConfig{RemoteDataResource: category},
	//})
	//TenancyImagesResource.IndexAttrs("File", "Title")

	// Add Tenancy
	rabcUser := Admin.AddResource(&rabc.RabcUser{}, &admin.Config{Menu: []string{"RABC Management"}})
	rabcRole := Admin.AddResource(&rabc.RabcRole{}, &admin.Config{Menu: []string{"RABC Management"}})
	rabcPermission := Admin.AddResource(&rabc.RabcPermission{}, &admin.Config{Menu: []string{"RABC Management"}})
	_ = Admin.AddResource(&rabc.OauthToken{}, &admin.Config{Menu: []string{"RABC Management"}})
	//tenant.Meta(&admin.Meta{Name: "Gender", Config: &admin.SelectOneConfig{Collection: Genders, AllowBlank: true}})

	//tenantPropertiesRes := tenant.Meta(&admin.Meta{Name: "TenancyProperties"}).Resource
	//tenantPropertiesRes.NewAttrs(&admin.Section{
	//	Rows: [][]string{{"Name", "Value"}},
	//})
	//tenantPropertiesRes.EditAttrs(&admin.Section{
	//	Rows: [][]string{{"Name", "Value"}},
	//})

	//tenant.Meta(&admin.Meta{Name: "Description", Config: &admin.RichEditorConfig{Plugins: []admin.RedactorPlugin{
	//	{Name: "medialibrary", Source: "/admin/assets/javascripts/qor_redactor_medialibrary.js"},
	//	{Name: "table", Source: "/vendors/redactor_table.js"},
	//},
	//	Settings: map[string]interface{}{
	//		"medialibraryUrl": "/admin/tenant_images",
	//	},
	//}})
	//tenant.Meta(&admin.Meta{Name: "Category", Config: &admin.SelectOneConfig{AllowBlank: true}})
	//tenant.Meta(&admin.Meta{Name: "Collections", Config: &admin.SelectManyConfig{SelectMode: "bottom_sheet"}})

	//tenant.Meta(&admin.Meta{Name: "MainImage", Config: &media_library.MediaBoxConfig{
	//	RemoteDataResource: TenancyImagesResource,
	//	Max:                1,
	//	Sizes: map[string]*media.Size{
	//		"main": {Width: 560, Height: 700},
	//	},
	//}})
	//tenant.Meta(&admin.Meta{Name: "MainImageURL", Valuer: func(record interface{}, context *qor.Context) interface{} {
	//	if p, ok := record.(*tenancy.Tenant); ok {
	//		result := bytes.NewBufferString("")
	//		tmpl, _ := template.New("").Parse("<img src='{{.image}}'></img>")
	//		tmpl.Execute(result, map[string]string{"image": p.MainImageURL()})
	//		return template.HTML(result.String())
	//	}
	//	return ""
	//}})

	//tenant.Filter(&admin.Filter{
	//	Name:   "Collections",
	//	Config: &admin.SelectOneConfig{RemoteDataResource: collection},
	//})

	//tenant.Filter(&admin.Filter{
	//	Name: "Featured",
	//})
	//
	//tenant.Filter(&admin.Filter{
	//	Name: "Name",
	//	Type: "string",
	//})
	//
	//tenant.Filter(&admin.Filter{
	//	Name: "Code",
	//})
	//
	//tenant.Filter(&admin.Filter{
	//	Name: "Price",
	//	Type: "number",
	//})
	//
	//tenant.Filter(&admin.Filter{
	//	Name: "CreatedAt",
	//})
	//
	//tenant.Action(&admin.Action{
	//	Name:        "Import Tenancy",
	//	URLOpenType: "slideout",
	//	URL: func(record interface{}, context *admin.Context) string {
	//		return "/admin/workers/new?job=Import Tenancys"
	//	},
	//	Modes: []string{"collection"},
	//})

	//type updateInfo struct {
	//	CategoryID  uint
	//	Category    *tenancy.Category
	//	MadeCountry string
	//	Gender      string
	//}
	//
	//updateInfoRes := Admin.NewResource(&updateInfo{})
	//tenant.Action(&admin.Action{
	//	Name:     "Update Info",
	//	Resource: updateInfoRes,
	//	Handler: func(argument *admin.ActionArgument) error {
	//		newTenancyInfo := argument.Argument.(*updateInfo)
	//		for _, record := range argument.FindSelectedRecords() {
	//			fmt.Printf("%#v\n", record)
	//			if tenant, ok := record.(*tenancy.Tenant); ok {
	//				if newTenancyInfo.Category != nil {
	//					tenant.Category = *newTenancyInfo.Category
	//				}
	//				if newTenancyInfo.MadeCountry != "" {
	//					tenant.MadeCountry = newTenancyInfo.MadeCountry
	//				}
	//				if newTenancyInfo.Gender != "" {
	//					tenant.Gender = newTenancyInfo.Gender
	//				}
	//				argument.Context.GetDB().Save(tenant)
	//			}
	//		}
	//		return nil
	//	},
	//	Modes: []string{"batch"},
	//})

	//tenant.UseTheme("grid")

	// variationsResource := tenant.Meta(&admin.Meta{Name: "Variations", Config: &variations.VariationsConfig{}}).Resource
	// if imagesMeta := variationsResource.GetMeta("Images"); imagesMeta != nil {
	// 	imagesMeta.Config = &media_library.MediaBoxConfig{
	// 		RemoteDataResource: TenancyImagesResource,
	// 		Sizes: map[string]*media.Size{
	// 			"icon":    {Width: 50, Height: 50},
	// 			"thumb":   {Width: 100, Height: 100},
	// 			"display": {Width: 300, Height: 300},
	// 		},
	// 	}
	// }

	// variationsResource.EditAttrs("-ID", "-Tenancy")
	// oldSearchHandler := tenant.SearchHandler
	// tenant.SearchHandler = func(keyword string, context *qor.Context) *gorm.DB {
	// 	context.SetDB(context.GetDB().Preload("Variations.Color").Preload("Variations.Size").Preload("Variations.Material"))
	// 	return oldSearchHandler(keyword, context)
	// }
	//colorVariationMeta := tenant.Meta(&admin.Meta{Name: "ColorVariations"})
	//colorVariation := colorVariationMeta.Resource
	//colorVariation.Meta(&admin.Meta{Name: "Images", Config: &media_library.MediaBoxConfig{
	//	RemoteDataResource: TenancyImagesResource,
	//	Sizes: map[string]*media.Size{
	//		"icon":    {Width: 50, Height: 50},
	//		"preview": {Width: 300, Height: 300},
	//		"listing": {Width: 640, Height: 640},
	//	},
	//}})

	//colorVariation.NewAttrs("-Tenancy", "-ColorCode")
	//colorVariation.EditAttrs("-Tenancy", "-ColorCode")
	//
	//sizeVariationMeta := colorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	//sizeVariation := sizeVariationMeta.Resource
	//sizeVariation.EditAttrs(
	//	&admin.Section{
	//		Rows: [][]string{
	//			{"Size", "AvailableQuantity"},
	//			{"ShareableVersion"},
	//		},
	//	},
	//)
	//sizeVariation.NewAttrs(sizeVariation.EditAttrs())

	//tenant.SearchAttrs("Name", "Code", "Category.Name", "Brand.Name")
	//tenant.IndexAttrs("MainImageURL", "Name", "Featured", "Price", "VersionName", "PublishLiveNow")
	//tenant.EditAttrs(
	//	&admin.Section{
	//		Title: "Seo Meta",
	//		Rows: [][]string{
	//			{"Seo"},
	//		}},
	//	&admin.Section{
	//		Title: "Basic Information",
	//		Rows: [][]string{
	//			{"Name", "Featured"},
	//			{"Code", "Price"},
	//			{"MainImage"},
	//		}},
	//	&admin.Section{
	//		Title: "Organization",
	//		Rows: [][]string{
	//			{"Category", "Gender"},
	//			{"Collections"},
	//		}},
	//	"TenancyProperties",
	//	"Description",
	//	"ColorVariations",
	//	"PublishReady",
	//)
	// tenant.ShowAttrs(tenant.EditAttrs())
	//tenant.NewAttrs(tenant.EditAttrs())
	//
	//for _, gender := range Genders {
	//	var gender = gender
	//	tenant.Scope(&admin.Scope{Name: gender, Group: "Gender", Handler: func(db *gorm.DB, ctx *qor.Context) *gorm.DB {
	//		return db.Where("gender = ?", gender)
	//	}})
	//}

	rabcUser.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if tenant, ok := record.(*tenancy.Tenant); ok {
				return fmt.Sprintf("/tenancy/%v", tenant.ID)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})
	rabcRole.Action(&admin.Action{
		Name: "View On Site",
		URL: func(record interface{}, context *admin.Context) string {
			if tenant, ok := record.(*tenancy.Tenant); ok {
				return fmt.Sprintf("/tenancy/%v", tenant.ID)
			}
			return "#"
		},
		Modes: []string{"menu_item", "edit"},
	})
	rabcPermission.Action(&admin.Action{
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
