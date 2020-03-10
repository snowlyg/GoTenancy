package admin

import (
	"github.com/qor/admin"
	qor_seo "github.com/qor/seo"
	"go-tenancy/models/seo"
)

// SetupSEO add seo
func SetupSEO(Tenant *admin.Admin) {
	seo.SEOCollection = qor_seo.New("Common SEO")
	seo.SEOCollection.RegisterGlobalVaribles(&seo.SEOGlobalSetting{SiteName: "GoTenancy"})
	seo.SEOCollection.SettingResource = Tenant.AddResource(&seo.MySEOSetting{}, &admin.Config{Invisible: true})
	seo.SEOCollection.RegisterSEO(&qor_seo.SEO{
		Name: "Default Page",
	})
	Tenant.AddResource(seo.SEOCollection, &admin.Config{Name: "SEO Setting", Menu: []string{"Site Management"}, Singleton: true, Priority: 2})
}
