package seo

import (
	"fmt"
	"html/template"
	"net/url"
	"reflect"
	"regexp"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/media"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
)

func init() {
	admin.RegisterViewPath("GoTenancy/libs/seo/views")
}

// New initialize a SeoCollection instance
func New(name string) *Collection {
	return &Collection{Name: name}
}

// Collection will hold registered seo configures and global setting definition and other configures
type Collection struct {
	Name            string
	SettingResource *admin.Resource

	registeredSEO  []*SEO
	resource       *admin.Resource
	globalResource *admin.Resource
	globalSetting  interface{}
}

// SEO represents a seo object for a page
type SEO struct {
	Name       string
	Varibles   []string
	OpenGraph  *OpenGraphConfig
	Context    func(...interface{}) map[string]string
	collection *Collection
}

// OpenGraphConfig open graph config
type OpenGraphConfig struct {
	ImageResource *admin.Resource
	Size          *media.Size
}

// RegisterGlobalVaribles register global setting struct and will represents as 'Site-wide Settings' part in admin
func (collection *Collection) RegisterGlobalVaribles(s interface{}) {
	collection.globalSetting = s
}

// RegisterSEO register a seo
func (collection *Collection) RegisterSEO(seo *SEO) {
	seo.collection = collection
	if seo.OpenGraph == nil {
		seo.OpenGraph = &OpenGraphConfig{}
	}
	collection.registeredSEO = append(collection.registeredSEO, seo)
}

// GetSEOSetting return SEO title, keywords and description and open graph settings
func (collection Collection) GetSEOSetting(context *qor.Context, name string, objects ...interface{}) Setting {
	var (
		seoSetting Setting
		db         = context.GetDB()
		seo        = collection.GetSEO(name)
	)

	// If passed objects has customzied SEO Setting field
	for _, obj := range objects {
		if value := reflect.Indirect(reflect.ValueOf(obj)); value.IsValid() && value.Kind() == reflect.Struct {
			for i := 0; i < value.NumField(); i++ {
				if value.Field(i).Type() == reflect.TypeOf(Setting{}) {
					seoSetting = value.Field(i).Interface().(Setting)
					break
				}
			}
		}
	}

	if !seoSetting.EnabledCustomize {
		globalSeoSetting := collection.SettingResource.NewStruct().(QorSEOSettingInterface)
		if !db.Where("name = ?", name).First(globalSeoSetting).RecordNotFound() {
			seoSetting = globalSeoSetting.GetSEOSetting()
		}
	}

	siteWideSetting := collection.SettingResource.NewStruct()
	db.Where("is_global_seo = ? AND name = ?", true, collection.Name).First(siteWideSetting)
	tagValues := siteWideSetting.(QorSEOSettingInterface).GetGlobalSetting()

	if tagValues == nil {
		tagValues = map[string]string{}
	}

	if seo.Context != nil {
		for key, value := range seo.Context(objects...) {
			tagValues[key] = value
		}
	}

	return replaceTags(seoSetting, seo.Varibles, tagValues)
}

// Render render SEO Setting
func (collection Collection) Render(context *qor.Context, name string, objects ...interface{}) template.HTML {
	seoSetting := collection.GetSEOSetting(context, name, objects...)
	return seoSetting.FormattedHTML(context)
}

// GetSEO get a Seo by name
func (collection *Collection) GetSEO(name string) *SEO {
	for _, s := range collection.registeredSEO {
		if s.Name == name {
			return s
		}
	}

	return &SEO{Name: name, collection: collection}
}

// SEOSettingURL get setting inline edit url by name
func (collection *Collection) SEOSettingURL(name string) string {
	qorAdmin := collection.resource.GetAdmin()
	return fmt.Sprintf("%v/%v/!seo_setting?name=%v", qorAdmin.GetRouter().Prefix, collection.resource.ToParam(), url.QueryEscape(name))
}

// ConfigureQorResource configure seoCollection for qor admin
func (collection *Collection) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		Admin := res.GetAdmin()
		collection.resource = res
		if collection.SettingResource == nil {
			collection.SettingResource = Admin.AddResource(&QorSEOSetting{}, &admin.Config{Invisible: true})
		}

		collection.SettingResource.UseTheme("seo")
		collection.SettingResource.EditAttrs("Name", "Setting")
		if nameMeta := collection.SettingResource.GetMeta("Name"); nameMeta != nil {
			nameMeta.Type = "hidden"
		}

		globalSettingRes := Admin.AddResource(collection.globalSetting, &admin.Config{Invisible: true})
		collection.globalResource = globalSettingRes

		res.Config.Singleton = true
		res.UseTheme("seo")

		router := Admin.GetRouter()
		controller := seoController{Collection: collection}
		router.Get(res.ToParam(), controller.Index)
		router.Put(fmt.Sprintf("%v/!seo_setting", res.ToParam()), controller.Update)
		router.Get(fmt.Sprintf("%v/!seo_setting", res.ToParam()), controller.InlineEdit)

		registerFuncMap(Admin)
	}
}

// Helpers
func replaceTags(seoSetting Setting, validTags []string, values map[string]string) Setting {
	replace := func(str string) string {
		re := regexp.MustCompile("{{([a-zA-Z0-9]*)}}")
		matches := re.FindAllStringSubmatch(str, -1)
		for _, match := range matches {
			str = strings.Replace(str, match[0], values[match[1]], 1)
		}
		return str
	}

	seoSetting.Title = replace(seoSetting.Title)
	seoSetting.Description = replace(seoSetting.Description)
	seoSetting.Keywords = replace(seoSetting.Keywords)
	seoSetting.Type = replace(seoSetting.Type)
	seoSetting.OpenGraphTitle = replace(seoSetting.OpenGraphTitle)
	seoSetting.OpenGraphDescription = replace(seoSetting.OpenGraphDescription)
	seoSetting.OpenGraphURL = replace(seoSetting.OpenGraphURL)
	seoSetting.OpenGraphImageURL = replace(seoSetting.OpenGraphImageURL)
	seoSetting.OpenGraphType = replace(seoSetting.OpenGraphType)
	for idx, metadata := range seoSetting.OpenGraphMetadata {
		seoSetting.OpenGraphMetadata[idx] = OpenGraphMetadata{
			Property: replace(metadata.Property),
			Content:  replace(metadata.Content),
		}
	}
	return seoSetting
}
