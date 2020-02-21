package banner_editor

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/assetfs"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/serializable_meta"
	mobiledetect "github.com/Shaked/gomobiledetect"
	"github.com/jinzhu/gorm"
)

var (
	registeredElements           []*Element
	registeredExternalStylePaths []string
	viewPaths                    []string
	assetFileSystem              assetfs.Interface
)

const (
	// Laptop default laptop platform name
	Laptop = "Laptop"
	// Mobile default mobile platform name
	Mobile = "Mobile"
)

func init() {
	assetFileSystem = assetfs.AssetFS().NameSpace("banner_editor")
}

// BannerEditorConfig configure display elements and setting model
type BannerEditorConfig struct {
	MediaLibrary    *admin.Resource
	Elements        []string
	SettingResource *admin.Resource
	Platforms       []Platform
}

// QorBannerEditorSettingInterface interface to support customize setting model
type QorBannerEditorSettingInterface interface {
	GetID() uint
	serializable_meta.SerializableMetaInterface
}

// QorBannerEditorSetting default setting model
type QorBannerEditorSetting struct {
	gorm.Model
	serializable_meta.SerializableMeta
}

// Element represent a button/element in banner_editor toolbar
type Element struct {
	Icon     string
	Name     string
	Label    string
	Template string
	Resource *admin.Resource
	Context  func(context *admin.Context, setting interface{}) interface{}
}

// Size contains width ad height
type Size struct {
	Width  int
	Height int
}

// Platform used to defined how many platform a banner need to support
type Platform struct {
	Name     string
	SafeArea Size
}

type configurePlatform struct {
	Name  string
	Value string
}

func init() {
	admin.RegisterViewPath("GoTenancy/libs/banner_editor/views")
}

// RegisterElement register a element
func RegisterElement(e *Element) {
	registeredElements = append(registeredElements, e)
}

// RegisterExternalStylePath register a asset path
func RegisterExternalStylePath(path string) {
	registeredExternalStylePaths = append(registeredExternalStylePaths, path)
}

// ConfigureQorMeta configure route and funcmap for banner_editor meta
func (config *BannerEditorConfig) ConfigureQorMeta(metaor resource.Metaor) {
	if meta, ok := metaor.(*admin.Meta); ok {
		meta.Type = "banner_editor"
		Admin := meta.GetBaseResource().(*admin.Resource).GetAdmin()

		if config.SettingResource == nil {
			config.SettingResource = Admin.NewResource(&QorBannerEditorSetting{})
		}

		if config.MediaLibrary == nil {
			panic("BannerEditor: MediaLibrary can't be blank.")
		} else {
			if getMediaLibraryResourceURLMethod(config.MediaLibrary.NewStruct()).IsNil() {
				panic("BannerEditor: MediaLibrary's struct doesn't have any field implement URL method, please refer media_library.MediaLibrary{}.")
			}

			config.MediaLibrary.Meta(&admin.Meta{
				Name: "BannerEditorUrl",
				Type: "hidden",
				Valuer: func(v interface{}, c *qor.Context) interface{} {
					values := getMediaLibraryResourceURLMethod(v).Call([]reflect.Value{})
					if len(values) > 0 {
						return values[0]
					}
					return ""
				},
			})

			config.MediaLibrary.IndexAttrs(config.MediaLibrary.IndexAttrs(), "BannerEditorUrl")
		}

		if len(config.Platforms) == 0 {
			config.Platforms = []Platform{
				{Name: Laptop},
				{Name: Mobile},
			}
		}

		router := Admin.GetRouter()
		res := config.SettingResource
		router.Get(fmt.Sprintf("%v/new", res.ToParam()), New, &admin.RouteConfig{Resource: res})
		router.Post(fmt.Sprintf("%v", res.ToParam()), Create, &admin.RouteConfig{Resource: res})
		router.Put(fmt.Sprintf("%v/%v", res.ToParam(), res.ParamIDName()), Update, &admin.RouteConfig{Resource: res})
		Admin.RegisterResourceRouters(res, "read", "update")

		Admin.RegisterFuncMap("formatted_banner_edit_value", formattedValue)
		Admin.RegisterFuncMap("banner_editor_configure", func(config *BannerEditorConfig) string {
			type element struct {
				Label     string
				CreateURL string
				Icon      string
			}
			type platform struct {
				Name   string
				Width  int
				Height int
			}
			var (
				selectedElements = registeredElements
				elements         = []element{}
				newElementURL    = router.Prefix + fmt.Sprintf("/%v/new", res.ToParam())
			)
			if len(config.Elements) != 0 {
				selectedElements = []*Element{}
				for _, name := range config.Elements {
					if e := GetElement(name); e != nil {
						selectedElements = append(selectedElements, e)
					}
				}
			}
			for _, e := range selectedElements {
				element := element{Icon: e.Icon, Label: e.Label, CreateURL: fmt.Sprintf("%v?kind=%v", newElementURL, template.URLQueryEscaper(e.Name))}
				if element.Label == "" {
					element.Label = e.Name
				}
				elements = append(elements, element)
			}

			platforms := []platform{}
			for _, p := range config.Platforms {
				platforms = append(platforms, platform{Name: p.Name, Width: p.SafeArea.Width, Height: p.SafeArea.Height})
			}
			results, err := json.Marshal(struct {
				Elements          []element
				ExternalStylePath []string
				EditURL           string
				Platforms         []platform
			}{
				Elements:          elements,
				ExternalStylePath: registeredExternalStylePaths,
				EditURL:           fmt.Sprintf("%v/%v/:id/edit", router.Prefix, res.ToParam()),
				Platforms:         platforms,
			})
			if err != nil {
				return err.Error()
			}
			return string(results)
		})
	}
}

// GetElement returnn element struct by name
func GetElement(name string) *Element {
	for _, e := range registeredElements {
		if e.Name == name {
			return e
		}
	}
	return nil
}

// GetID return setting ID
func (setting QorBannerEditorSetting) GetID() uint {
	return setting.ID
}

// GetSerializableArgumentResource return setting's resource
func (setting QorBannerEditorSetting) GetSerializableArgumentResource() *admin.Resource {
	element := GetElement(setting.Kind)
	if element != nil {
		return element.Resource
	}
	return nil
}

// GetContent return HTML string by detector, detector could be a string (Platform, Mobile or other), http request and nil
func GetContent(value string, detector interface{}) string {
	if platform, ok := detector.(string); ok {
		return getContentByPlatform(value, platform)
	} else if req, ok := detector.(*http.Request); ok {
		detect := mobiledetect.NewMobileDetect(req, nil)
		if detect.IsMobile() && !detect.IsTablet() {
			return getContentByPlatform(value, Mobile)
		}
		return getContentByPlatform(value, Laptop)
	}
	return getContentByPlatform(value, Laptop)
}

func getContentByPlatform(value string, platform string) string {
	configurePlatforms := []configurePlatform{}
	if err := json.Unmarshal([]byte(value), &configurePlatforms); err == nil {
		if len(configurePlatforms) == 0 {
			return ""
		}
		for _, p := range configurePlatforms {
			if p.Name == platform && strings.TrimSpace(p.Value) != "" {
				return unescapeValue(p.Value)
			}
		}
		return unescapeValue(configurePlatforms[0].Value)
	}
	return unescapeValue(value)
}

func unescapeValue(value string) string {
	if val, err := url.QueryUnescape(value); err == nil {
		return val
	}
	return value
}

func formattedValue(value string) string {
	if value == "" {
		return "[]"
	}
	configurePlatforms := []configurePlatform{}
	if err := json.Unmarshal([]byte(value), &configurePlatforms); err == nil {
		return value
	}
	jsonValue, err := json.Marshal(&[]configurePlatform{{Name: Laptop, Value: value}})
	if err != nil {
		return fmt.Sprintf("BannerEditor: format value to json failure, got %v", err.Error())
	}
	return string(jsonValue)
}

func getMediaLibraryResourceURLMethod(i interface{}) reflect.Value {
	value := reflect.Indirect(reflect.ValueOf(i))
	for i := 0; i < value.NumField(); i++ {
		field := value.Field(i)
		if urlMethod := field.MethodByName("URL"); urlMethod.IsValid() {
			return urlMethod
		}
	}
	return reflect.Value{}
}
