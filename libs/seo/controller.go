package seo

import (
	"net/http"
	"net/url"
	"path"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/responder"
)

type seoController struct {
	Collection *Collection
}

func (sc seoController) Index(context *admin.Context) {
	context = context.NewResourceContext(sc.Collection.SettingResource)
	context.Execute("index", struct {
		Collection      *Collection
		SettingResource *admin.Resource
	}{
		Collection:      sc.Collection,
		SettingResource: sc.Collection.SettingResource,
	})
}

func (sc seoController) InlineEdit(context *admin.Context) {
	settingContext := context.NewResourceContext(sc.Collection.SettingResource)
	result := sc.Collection.SettingResource.NewStruct()
	name, err := url.QueryUnescape(context.Request.Form.Get("name"))
	if err != nil {
		settingContext.AddError(err)
	}
	context.DB.Where("name = ?", name).First(result)

	if seoSetting, ok := result.(QorSEOSettingInterface); ok {
		seoSetting.SetCollection(sc.Collection)
	}

	responder.With("html", func() {
		settingContext.Execute("edit", struct {
			Setting interface{}
			EditURL string
			Metas   []*admin.Section
		}{
			Setting: result,
			EditURL: sc.Collection.SEOSettingURL(name),
			Metas:   seoSettingMetas(sc.Collection),
		})
	}).With("json", func() {
		settingContext.JSON("edit", result)
	}).Respond(context.Request)
}

func (sc seoController) Update(context *admin.Context) {
	settingResource := sc.Collection.SettingResource
	settingContext := context.NewResourceContext(settingResource)
	result := settingResource.NewStruct()

	name, err := url.QueryUnescape(context.Request.Form.Get("name"))
	if err != nil {
		settingContext.AddError(err)
	}
	context.DB.Where("name = ?", name).First(result)
	if context.DB.NewRecord(result) {
		context.Request.Form["QorResource.Name"] = []string{name}
		context.Request.Form["QorResource.Setting.Type"] = []string{name}
	}

	seoSettingInterface := result.(QorSEOSettingInterface)
	if seoSettingInterface.GetIsGlobalSEO() {
		globalSetting := make(map[string]string)
		for fieldWithPrefix := range context.Request.Form {
			if strings.HasPrefix(fieldWithPrefix, "QorResource") {
				field := strings.Replace(fieldWithPrefix, "QorResource.", "", -1)
				globalSetting[field] = context.Request.Form.Get(fieldWithPrefix)
			}
		}
		seoSettingInterface.SetGlobalSetting(globalSetting)
	}

	res := settingContext.Resource
	if !settingContext.HasError() {
		if settingContext.AddError(res.Decode(settingContext.Context, result)); !settingContext.HasError() {
			settingContext.AddError(res.CallSave(result, settingContext.Context))
		}
	}

	responder.With("html", func() {
		http.Redirect(context.Writer, context.Request, path.Join(settingResource.GetAdmin().GetRouter().Prefix, context.Resource.ToParam()), http.StatusFound)
	}).With("json", func() {
		if settingContext.HasError() {
			context.Writer.WriteHeader(admin.HTTPUnprocessableEntity)
			settingContext.JSON("edit", map[string]interface{}{"errors": settingContext.GetErrors()})
		} else {
			settingContext.JSON("show", result)
		}
	}).Respond(context.Request)
}
