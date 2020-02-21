package seo

import (
	"reflect"
	"text/template"

	"GoTenancy/libs/admin"
)

func seoSections(context *admin.Context, collection *Collection) []interface{} {
	settings := []interface{}{}
	for _, seo := range collection.registeredSEO {
		s := collection.SettingResource.NewStruct()
		db := context.GetDB()
		db.Where("name = ?", seo.Name).First(s)
		if db.NewRecord(s) {
			s.(QorSEOSettingInterface).SetName(seo.Name)
			s.(QorSEOSettingInterface).SetSEOType(seo.Name)
			db.Save(s)
		}
		s.(QorSEOSettingInterface).SetCollection(collection)
		settings = append(settings, s)
	}
	return settings
}

func seoSettingMetas(collection *Collection) []*admin.Section {
	return collection.SettingResource.EditAttrs()
}

func seoGlobalSetting(context *admin.Context, collection *Collection) interface{} {
	s := collection.SettingResource.NewStruct()
	db := context.GetDB()
	db.Where("is_global_seo = ? AND name = ?", true, collection.Name).First(s)
	if db.NewRecord(s) {
		s.(QorSEOSettingInterface).SetName(collection.Name)
		s.(QorSEOSettingInterface).SetSEOType(collection.Name)
		s.(QorSEOSettingInterface).SetIsGlobalSEO(true)
		db.Save(s)
	}
	return s
}

func seoGlobalSettingValue(collection *Collection, setting QorSEOSettingInterface) interface{} {
	value := reflect.Indirect(reflect.ValueOf(collection.globalResource.NewStruct()))
	settingValue := setting.GetGlobalSetting()
	for i := 0; i < value.NumField(); i++ {
		fieldName := value.Type().Field(i).Name
		if settingValue[fieldName] != "" {
			value.Field(i).Set(reflect.ValueOf(settingValue[fieldName]))
		}
	}
	return value.Interface()
}

func seoGlobalSettingMetas(collection *Collection) []*admin.Section {
	return collection.globalResource.NewAttrs()
}

func seoTagsByType(seo *SEO) (tags []string) {
	if seo == nil {
		return []string{}
	}
	value := reflect.Indirect(reflect.ValueOf(seo.collection.globalSetting))
	for i := 0; i < value.NumField(); i++ {
		tags = append(tags, value.Type().Field(i).Name)
	}
	for _, s := range seo.Varibles {
		tags = append(tags, s)
	}
	return tags
}

func seoAppendDefaultValue(context *admin.Context, seo *SEO, resourceSeoValue interface{}) interface{} {
	db := context.GetDB()
	globalInteface := seo.collection.SettingResource.NewStruct()
	db.Where("name = ?", seo.Name).First(globalInteface)
	globalSetting := globalInteface.(QorSEOSettingInterface)
	setting := resourceSeoValue.(Setting)
	if !setting.EnabledCustomize && setting.Title == "" && setting.Description == "" && setting.Keywords == "" {
		setting.Title = globalSetting.GetTitle()
		setting.Description = globalSetting.GetDescription()
		setting.Keywords = globalSetting.GetKeywords()
	}

	if !setting.EnabledCustomize {
		if setting.OpenGraphTitle == "" {
			setting.OpenGraphTitle = globalSetting.GetOpenGraphTitle()
		}
		if setting.OpenGraphDescription == "" {
			setting.OpenGraphDescription = globalSetting.GetOpenGraphDescription()
		}
		if setting.OpenGraphURL == "" {
			setting.OpenGraphURL = globalSetting.GetOpenGraphURL()
		}
		if setting.OpenGraphType == "" {
			setting.OpenGraphType = globalSetting.GetOpenGraphType()
		}
		if setting.OpenGraphImageURL == "" {
			setting.OpenGraphImageURL = globalSetting.GetOpenGraphImageURL()
		}
		if setting.OpenGraphImageFromMediaLibrary.URL() == "" {
			setting.OpenGraphImageFromMediaLibrary = globalSetting.GetOpenGraphImageFromMediaLibrary()
		}
		if len(setting.OpenGraphMetadata) == 0 {
			setting.OpenGraphMetadata = globalSetting.GetOpenGraphMetadata()
		}
	}

	return setting
}

func seoURL(collection *Collection, name string) string {
	return collection.SEOSettingURL(name)
}

func registerFuncMap(a *admin.Admin) {
	funcMaps := template.FuncMap{
		"seo_sections":             seoSections,
		"seo_setting_metas":        seoSettingMetas,
		"seo_global_setting_value": seoGlobalSettingValue,
		"seo_global_setting_metas": seoGlobalSettingMetas,
		"seo_global_setting":       seoGlobalSetting,
		"seo_tags_by_type":         seoTagsByType,
		"seo_append_default_value": seoAppendDefaultValue,
		"seo_url_for":              seoURL,
	}

	for key, value := range funcMaps {
		a.RegisterFuncMap(key, value)
	}
}
