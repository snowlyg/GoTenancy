package i18n

import (
	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor/utils"
)

type i18nController struct {
	*I18n
}

func (controller *i18nController) Index(context *admin.Context) {
	context.Execute("index", controller.I18n)
}

func (controller *i18nController) Update(context *admin.Context) {
	form := context.Request.Form
	translation := Translation{Key: form.Get("Key"), Locale: form.Get("Locale"), Value: utils.HTMLSanitizer.Sanitize(form.Get("Value"))}

	if err := controller.I18n.SaveTranslation(&translation); err == nil {
		context.Writer.Write([]byte("OK"))
	} else {
		context.Writer.WriteHeader(422)
		context.Writer.Write([]byte(err.Error()))
	}
}
