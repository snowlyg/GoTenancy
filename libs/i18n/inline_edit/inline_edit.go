package inline_edit

import (
	"fmt"
	"html/template"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/i18n"
)

func init() {
	admin.RegisterViewPath("libs/i18n/inline_edit/views")
}

// FuncMap generate func map for inline edit
func FuncMap(I18n *i18n.I18n, locale string, enableInlineEdit bool) template.FuncMap {
	return template.FuncMap{
		"t": InlineEdit(I18n, locale, enableInlineEdit),
	}
}

// InlineEdit enable inline edit
func InlineEdit(I18n *i18n.I18n, locale string, isInline bool) func(string, ...interface{}) template.HTML {
	return func(key string, args ...interface{}) template.HTML {
		// Get Translation Value
		var value template.HTML
		var defaultValue string
		if len(args) > 0 {
			if args[0] == nil {
				defaultValue = key
			} else {
				defaultValue = fmt.Sprint(args[0])
			}
			value = I18n.Default(defaultValue).T(locale, key, args[1:]...)
		} else {
			value = I18n.T(locale, key)
		}

		// Append inline-edit script/tag
		if isInline {
			var editType string
			if len(value) > 25 {
				editType = "data-type=\"textarea\""
			}
			prefix := I18n.Resource.GetAdmin().GetRouter().Prefix
			assetsTag := fmt.Sprintf("<script data-prefix=\"%v\" src=\"%v/assets/javascripts/i18n-checker.js?theme=i18n\"></script>", prefix, prefix)
			return template.HTML(fmt.Sprintf("%s<span class=\"qor-i18n-inline\" %s data-locale=\"%s\" data-key=\"%s\">%s</span>", assetsTag, editType, locale, key, string(value)))
		}
		return value
	}
}
