package api

import (
	"GoTenancy/config/application"
	"GoTenancy/config/db"
	"GoTenancy/models/orders"
	"GoTenancy/models/products"
	"GoTenancy/models/users"
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	"github.com/qor/qor"
)

// New new api app
func New(config *Config) *App {
	if config.Prefix == "" {
		config.Prefix = "/api"
	}
	return &App{Config: config}
}

// App api app
type App struct {
	Config *Config
}

// Config api config struct
type Config struct {
	Prefix string
}

// ConfigureApplication configure application
func (app App) ConfigureApplication(application *application.Application) {
	API := admin.New(&qor.Config{DB: db.DB})

	Product := API.AddResource(&products.Product{})

	ColorVariationMeta := Product.Meta(&admin.Meta{Name: "ColorVariations"})
	ColorVariation := ColorVariationMeta.Resource
	ColorVariation.IndexAttrs("ID", "Color", "Images", "SizeVariations")
	ColorVariation.ShowAttrs("Color", "Images", "SizeVariations")

	SizeVariationMeta := ColorVariation.Meta(&admin.Meta{Name: "SizeVariations"})
	SizeVariation := SizeVariationMeta.Resource
	SizeVariation.IndexAttrs("ID", "Size", "AvailableQuantity")
	SizeVariation.ShowAttrs("ID", "Size", "AvailableQuantity")

	API.AddResource(&orders.Order{})

	API.AddResource(&users.User{})
	// User := API.AddResource(&users.User{})
	// userOrders, _ := User.AddSubResource("Orders")
	// userOrders.AddSubResource("OrderItems", &admin.Config{Name: "Items"})

	API.AddResource(&products.Category{})

	application.IrisApp.Any(app.Config.Prefix, iris.FromStd(API.NewServeMux(app.Config.Prefix)))
}
