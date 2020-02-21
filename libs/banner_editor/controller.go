package banner_editor

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/responder"
)

// New handle new setting page
func New(context *admin.Context) {
	setting := context.Resource.NewStruct().(QorBannerEditorSettingInterface)
	kind := context.Request.URL.Query().Get("kind")
	if GetElement(kind) != nil {
		setting.SetSerializableArgumentKind(kind)
	} else {
		context.AddError(fmt.Errorf("BannerEditor: It isn't any element match %v", kind))
	}
	context.Execute("new", setting)
}

// Create handle create setting
func Create(context *admin.Context) {
	var (
		res     = context.Resource
		result  = res.NewStruct()
		kind    = context.Request.Form.Get("QorResource.Kind")
		element = GetElement(kind)
	)
	if context.AddError(res.Decode(context.Context, result)); !context.HasError() {
		context.AddError(res.CallSave(result, context.Context))
	}

	c := element.Context(context, result)
	html, err := render(element.Template, c)
	if err != nil {
		context.AddError(fmt.Errorf("BannerEditor: can't parse %v's template, got %v", kind, err))
	}
	if context.HasError() {
		responder.With("html", func() {
			context.Writer.WriteHeader(admin.HTTPUnprocessableEntity)
			context.Execute("new", result)
		}).With([]string{"json"}, func() {
			context.Writer.WriteHeader(admin.HTTPUnprocessableEntity)
			context.Encode("index", map[string]interface{}{"errors": context.GetErrors()})
		}).Respond(context.Request)
	} else {
		responder.With("html", func() {
			context.Flash(string(res.GetAdmin().T(context.Context, "qor_admin.form.successfully_created", "{{.Name}} was successfully created", res)), "success")
			http.Redirect(context.Writer, context.Request, context.URLFor(result, res), http.StatusFound)
		}).With([]string{"json"}, func() {
			json, err := settingJSON(html, result)
			if err != nil {
				panic(err)
			}
			context.Writer.Write(json)
		}).Respond(context.Request)
	}
}

// Update handle update setting
func Update(context *admin.Context) {
	var (
		res     = context.Resource
		kind    = context.Request.Form.Get("QorResource.Kind")
		element = GetElement(kind)
		html    template.HTML
		err     error
	)
	result, err := context.FindOne()

	if err != nil {
		context.AddError(err)
	} else {
		if context.AddError(res.Decode(context.Context, result)); !context.HasError() {
			context.AddError(res.CallSave(result, context.Context))
		}

		c := element.Context(context, result)
		html, err = render(element.Template, c)
		if err != nil {
			context.AddError(fmt.Errorf("BannerEditor: can't parse %v's template, got %v", kind, err))
		}
	}

	if context.HasError() {
		context.Writer.WriteHeader(admin.HTTPUnprocessableEntity)
		responder.With("html", func() {
			context.Execute("edit", result)
		}).With([]string{"json", "xml"}, func() {
			context.Encode("edit", map[string]interface{}{"errors": context.GetErrors()})
		}).Respond(context.Request)
	} else {
		responder.With("html", func() {
			context.Flash(string(res.GetAdmin().T(context.Context, "qor_admin.form.successfully_updated", "{{.Name}} was successfully updated", res)), "success")
			context.Execute("show", result)
		}).With([]string{"json"}, func() {
			json, err := settingJSON(html, result)
			if err != nil {
				panic(err)
			}
			context.Writer.Write(json)
		}).Respond(context.Request)
	}
}

func settingJSON(html template.HTML, result interface{}) ([]byte, error) {
	jsonValue := &bytes.Buffer{}
	encoder := json.NewEncoder(jsonValue)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(struct {
		ID       uint
		Template template.HTML
	}{
		ID:       result.(QorBannerEditorSettingInterface).GetID(),
		Template: html,
	})
	if err != nil {
		return nil, err
	}
	return jsonValue.Bytes(), nil
}
