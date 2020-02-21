package migrations

import (
	"GoTenancy/app/admin"
	"GoTenancy/config/db"
	"GoTenancy/libs/activity"
	"GoTenancy/libs/auth/auth_identity"
	"GoTenancy/libs/banner_editor"
	"GoTenancy/libs/help"
	"GoTenancy/libs/media/asset_manager"
	"GoTenancy/libs/transition"
	"GoTenancy/models/blogs"
	"GoTenancy/models/orders"
	"GoTenancy/models/products"
	"GoTenancy/models/seo"
	"GoTenancy/models/settings"
	"GoTenancy/models/stores"
	"GoTenancy/models/users"
)

func init() {
	AutoMigrate(&asset_manager.AssetManager{})

	AutoMigrate(&products.Product{}, &products.ProductVariation{}, &products.ProductImage{}, &products.ColorVariation{}, &products.ColorVariationImage{}, &products.SizeVariation{})
	AutoMigrate(&products.Color{}, &products.Size{}, &products.Material{}, &products.Category{}, &products.Collection{})

	AutoMigrate(&users.User{}, &users.Address{})

	AutoMigrate(&orders.Order{}, &orders.OrderItem{})

	AutoMigrate(&orders.DeliveryMethod{})

	AutoMigrate(&stores.Store{})

	AutoMigrate(&settings.Setting{}, &settings.MediaLibrary{})

	AutoMigrate(&transition.StateChangeLog{})

	AutoMigrate(&activity.QorActivity{})

	AutoMigrate(&admin.QorWidgetSetting{})

	AutoMigrate(&blogs.Page{}, &blogs.Article{})

	AutoMigrate(&seo.MySEOSetting{})

	AutoMigrate(&help.QorHelpEntry{})

	AutoMigrate(&auth_identity.AuthIdentity{})

	AutoMigrate(&banner_editor.QorBannerEditorSetting{})
}

// AutoMigrate run auto migration
func AutoMigrate(values ...interface{}) {
	for _, value := range values {
		db.DB.AutoMigrate(value)
	}
}
