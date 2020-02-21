package location

import (
	"fmt"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
)

func init() {
	admin.RegisterViewPath("libs/location/views")
}

var (
	_ admin.MetaConfigInterface = &LocationConfig{}
	// GoogleAPIKey Key of Google Map API
	GoogleAPIKey string
	// BaiduAPIKey Key of Baidu Map API
	BaiduAPIKey string
)

// LocationConfig Location Meta's Config
type LocationConfig struct {
	Backend      string
	BaiduAPIKey  string
	GoogleAPIKey string
}

// ConfigureQorMeta configure Qor Meta to implement MetaConfig interface
func (locationConfig *LocationConfig) ConfigureQorMeta(meta resource.Metaor) {
	if GoogleAPIKey != "" && locationConfig.GoogleAPIKey == "" {
		locationConfig.GoogleAPIKey = GoogleAPIKey
	} else if BaiduAPIKey != "" && locationConfig.BaiduAPIKey == "" {
		locationConfig.BaiduAPIKey = BaiduAPIKey
	}

	if locationConfig.Backend == "" {
		if locationConfig.GoogleAPIKey != "" {
			locationConfig.Backend = "google"
		} else if locationConfig.BaiduAPIKey != "" {
			locationConfig.Backend = "baidu"
		}
	}
}

// Location is a struct, you could embedded it into your model to get the Location feature for your model
type Location struct {
	Address   string
	City      string
	Region    string
	Country   string
	Zip       string
	Latitude  float64
	Longitude float64
}

type locationInterface interface {
	GetLocation() *Location
}

// GetLocation get location from your model
func (location *Location) GetLocation() *Location {
	return location
}

// ConfigureQorResource configure qor locale for Qor Admin
func (*Location) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		Admin := res.GetAdmin()
		res.UseTheme("location")

		Admin.RegisterFuncMap("replace_suffix", func(str, suffix, newSuffix string) string {
			return fmt.Sprint(strings.TrimSuffix(str, suffix), newSuffix)
		})

		scope := Admin.DB.NewScope(res.Value)
		if field, ok := scope.GetModelStruct().ModelType.FieldByName("Location"); ok {
			labelName := field.Name
			if customName, ok := utils.ParseTagOption(field.Tag.Get("location"))["NAME"]; ok {
				labelName = customName
			}

			res.Meta(&admin.Meta{Name: field.Name, Label: labelName, Type: "location", Config: &LocationConfig{GoogleAPIKey: GoogleAPIKey, BaiduAPIKey: BaiduAPIKey}, Valuer: func(resource interface{}, ctx *qor.Context) interface{} {
				return resource.(locationInterface).GetLocation()
			}})

			res.IndexAttrs(res.IndexAttrs(), "-"+field.Name, "-Latitude", "-Longitude")

			res.OverrideNewAttrs(func() {
				res.NewAttrs(res.NewAttrs(), "-Address", "-City", "-Region", "-Country", "-Zip", "-Latitude", "-Longitude")
			})

			res.OverrideEditAttrs(func() {
				res.EditAttrs(res.EditAttrs(), "-Address", "-City", "-Region", "-Country", "-Zip", "-Latitude", "-Longitude")
			})

			res.OverrideShowAttrs(func() {
				res.ShowAttrs(res.ShowAttrs(), "-"+field.Name)
			})
		}
	}
}
