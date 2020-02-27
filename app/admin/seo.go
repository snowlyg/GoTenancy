package admin

import (
	"GoTenancy/models/seo"
	"github.com/qor/admin"
	qor_seo "github.com/qor/seo"
)

// SetupSEO add seo
func SetupSEO(Admin *admin.Admin) {
	seo.SEOCollection = qor_seo.New("Common SEO")
	seo.SEOCollection.RegisterGlobalVaribles(&seo.SEOGlobalSetting{SiteName: "GoTenancy"})
	seo.SEOCollection.SettingResource = Admin.AddResource(&seo.MySEOSetting{}, &admin.Config{Invisible: true})
	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name: "Default Page",
	})
	Admin.AddResource(seo.SEOCollection, &admin.Config{Name: "SEO 设置", Menu: []string{"系统设置"}, Singleton: true, Priority: 2})
}
