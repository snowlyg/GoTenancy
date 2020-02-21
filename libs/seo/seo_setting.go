package seo

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"html/template"
	"log"
	"net/url"
	"time"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/media/media_library"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
)

// QorSEOSettingInterface support customize Seo model
type QorSEOSettingInterface interface {
	GetName() string
	SetName(string)
	GetSEOSetting() Setting
	GetGlobalSetting() map[string]string
	SetGlobalSetting(map[string]string)
	GetSEOType() string
	SetSEOType(string)
	GetIsGlobalSEO() bool
	SetIsGlobalSEO(bool)
	GetTitle() string
	GetDescription() string
	GetKeywords() string
	SetCollection(*Collection)
	GetOpenGraphTitle() string
	GetOpenGraphDescription() string
	GetOpenGraphURL() string
	GetOpenGraphType() string
	GetOpenGraphImageURL() string
	GetOpenGraphImageFromMediaLibrary() media_library.MediaBox
	GetOpenGraphMetadata() []OpenGraphMetadata
}

// QorSEOSetting default seo model
type QorSEOSetting struct {
	Name        string `gorm:"primary_key"`
	Setting     Setting
	IsGlobalSEO bool

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`

	collection *Collection
}

// Setting defined meta's attributes
type Setting struct {
	Title                          string `gorm:"size:4294967295"`
	Description                    string
	Keywords                       string
	Type                           string
	OpenGraphTitle                 string
	OpenGraphDescription           string
	OpenGraphURL                   string
	OpenGraphType                  string
	OpenGraphImageURL              string
	OpenGraphImageFromMediaLibrary media_library.MediaBox
	OpenGraphMetadata              []OpenGraphMetadata
	EnabledCustomize               bool
	GlobalSetting                  map[string]string
}

// OpenGraphMetadata open graph meta data
type OpenGraphMetadata struct {
	Property string
	Content  string
}

// GetSEOSetting get seo setting
func (s QorSEOSetting) GetSEOSetting() Setting {
	return s.Setting
}

// GetName get QorSeoSetting's name
func (s QorSEOSetting) GetName() string {
	return s.Name
}

// SetName set QorSeoSetting's name
func (s *QorSEOSetting) SetName(name string) {
	s.Name = name
}

// GetSEOType get QorSeoSetting's type
func (s QorSEOSetting) GetSEOType() string {
	return s.Setting.Type
}

// SetSEOType set QorSeoSetting's type
func (s *QorSEOSetting) SetSEOType(t string) {
	s.Setting.Type = t
}

// GetIsGlobalSEO get QorSEOSetting's isGlobal
func (s QorSEOSetting) GetIsGlobalSEO() bool {
	return s.IsGlobalSEO
}

// SetIsGlobalSEO set QorSeoSetting's isGlobal
func (s *QorSEOSetting) SetIsGlobalSEO(isGlobal bool) {
	s.IsGlobalSEO = isGlobal
}

// GetGlobalSetting get QorSeoSetting's globalSetting
func (s QorSEOSetting) GetGlobalSetting() map[string]string {
	return s.Setting.GlobalSetting
}

// SetGlobalSetting set QorSeoSetting's globalSetting
func (s *QorSEOSetting) SetGlobalSetting(globalSetting map[string]string) {
	s.Setting.GlobalSetting = globalSetting
}

func (s QorSEOSetting) GetOpenGraphTitle() string {
	return s.Setting.OpenGraphTitle
}
func (s QorSEOSetting) GetOpenGraphDescription() string {
	return s.Setting.OpenGraphDescription
}
func (s QorSEOSetting) GetOpenGraphURL() string {
	return s.Setting.OpenGraphURL
}
func (s QorSEOSetting) GetOpenGraphType() string {
	return s.Setting.OpenGraphType
}
func (s QorSEOSetting) GetOpenGraphImageURL() string {
	return s.Setting.OpenGraphImageURL
}
func (s QorSEOSetting) GetOpenGraphImageFromMediaLibrary() media_library.MediaBox {
	return s.Setting.OpenGraphImageFromMediaLibrary
}
func (s QorSEOSetting) GetOpenGraphMetadata() []OpenGraphMetadata {
	return s.Setting.OpenGraphMetadata
}

// GetTitle get Setting's title
func (s QorSEOSetting) GetTitle() string {
	return s.Setting.Title
}

// GetDescription get Setting's description
func (s QorSEOSetting) GetDescription() string {
	return s.Setting.Description
}

// GetKeywords get Setting's keywords
func (s QorSEOSetting) GetKeywords() string {
	return s.Setting.Keywords
}

// SetCollection set Setting's collection
func (s *QorSEOSetting) SetCollection(collection *Collection) {
	s.collection = collection
}

// GetSEO get Setting's SEO configure
func (s QorSEOSetting) GetSEO() *SEO {
	return s.collection.GetSEO(s.Name)
}

// Scan scan value from database into struct
func (setting *Setting) Scan(value interface{}) error {
	if bytes, ok := value.([]byte); ok {
		json.Unmarshal(bytes, setting)
	} else if str, ok := value.(string); ok {
		json.Unmarshal([]byte(str), setting)
	} else if strs, ok := value.([]string); ok {
		for _, str := range strs {
			json.Unmarshal([]byte(str), setting)
		}
	}
	return nil
}

// Value get value from struct, and save into database
func (setting Setting) Value() (driver.Value, error) {
	result, err := json.Marshal(setting)
	return string(result), err
}

// FormattedHTML return formated seo setting as HTML
func (setting Setting) FormattedHTML(context *qor.Context) template.HTML {
	toAbsoluteURL := func(str string) string {
		if u, err := url.Parse(str); err == nil {
			if u.IsAbs() {
				return str
			}

			if u.Host == "" && context.Request != nil {
				u.Host = context.Request.Host
			}

			if u.Scheme == "" {
				if context.Request != nil && context.Request.URL.Scheme != "" {
					u.Scheme = context.Request.URL.Scheme
				} else {
					u.Scheme = "http"
				}
			}
			return u.String()
		}
		return ""
	}

	openGraphData := map[string]string{}

	if setting.OpenGraphURL != "" {
		openGraphData["og:url"] = toAbsoluteURL(setting.OpenGraphURL)
	}

	if setting.OpenGraphType != "" {
		openGraphData["og:type"] = setting.OpenGraphType
	}

	if len(setting.OpenGraphImageFromMediaLibrary.Files) > 0 {
		openGraphData["og:image"] = toAbsoluteURL(setting.OpenGraphImageFromMediaLibrary.URL())
	} else if setting.OpenGraphImageURL != "" {
		openGraphData["og:image"] = toAbsoluteURL(setting.OpenGraphImageURL)
	}

	for _, metavalue := range setting.OpenGraphMetadata {
		openGraphData[metavalue.Property] = metavalue.Content
	}

	if _, ok := openGraphData["og:title"]; !ok {
		title := setting.Title
		if setting.OpenGraphTitle != "" {
			title = setting.OpenGraphTitle
		}
		openGraphData["og:title"] = title
	}

	if _, ok := openGraphData["og:description"]; !ok {
		desc := setting.Description
		if setting.OpenGraphDescription != "" {
			desc = setting.OpenGraphDescription
		}

		openGraphData["og:description"] = desc
	}

	var buf bytes.Buffer
	err := seoTmpl.Execute(&buf, map[string]interface{}{
		"title":       setting.Title,
		"description": setting.Description,
		"keywords":    setting.Keywords,
		"ogs":         openGraphData,
	})
	if err != nil {
		var requestURL string
		if context != nil && context.Request != nil && context.Request.URL != nil {
			requestURL = context.Request.URL.String()
		}
		log.Printf("Error: exec seo tmpl has err (%v) in %s", err, requestURL)
		return ""
	}

	return template.HTML(buf.String())
}

var seoTmpl = template.Must(
	template.New("seo_tmpl").Parse(`<title>{{.title}}</title>
<meta name="description" content="{{.description}}">
<meta name="keywords" content="{{.keywords}}">
{{range $key, $val := .ogs -}}
{{if ne $val "" -}}
<meta property="{{$key}}" name="{{$key}}" content="{{$val}}">
{{end -}}
{{end -}}`),
)

// ConfigureQorMetaBeforeInitialize configure SEO setting for qor admin
func (setting Setting) ConfigureQorMetaBeforeInitialize(meta resource.Metaor) {
	if meta, ok := meta.(*admin.Meta); ok {
		meta.Type = "seo"
		if res, ok := meta.GetBaseResource().(*admin.Resource); ok {
			res.UseTheme("seo_meta")
		}
	}
}

// ConfigureQorResource configure resource for seo setting
func (setting Setting) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		res.Meta(&admin.Meta{Name: "Title", Label: "HTML Title"})
		res.Meta(&admin.Meta{Name: "Description", Label: "Meta Description"})
		res.Meta(&admin.Meta{Name: "Keywords", Label: "Meta Keywords"})
		res.Meta(&admin.Meta{Name: "Type", Type: "hidden"})
		res.Meta(&admin.Meta{Name: "EnabledCustomize", Type: "hidden"})
		res.Meta(&admin.Meta{Name: "OpenGraphImageFromMediaLibrary", Label: "Open Graph Image", Config: &media_library.MediaBoxConfig{
			Max:       1,
			AllowType: media_library.ALLOW_TYPE_IMAGE,
		}})

		metadataResource := res.Meta(&admin.Meta{Name: "OpenGraphMetadata"}).Resource
		metadataResource.NewAttrs(&admin.Section{Rows: [][]string{{"Property", "Content"}}})
		metadataResource.EditAttrs(&admin.Section{Rows: [][]string{{"Property", "Content"}}})

		res.EditAttrs(
			&admin.Section{
				Title: "Basic",
				Rows:  [][]string{{"Title"}, {"Description"}, {"Keywords"}},
			},
			&admin.Section{
				Title: "Open Graph Information",
				Rows: [][]string{
					{"OpenGraphTitle", "OpenGraphDescription"},
					{"OpenGraphURL", "OpenGraphType"},
					{"OpenGraphImageURL", "OpenGraphImageFromMediaLibrary"}, {"OpenGraphMetadata"},
				},
			},
			"Type", "EnabledCustomize",
		)
	}
}
