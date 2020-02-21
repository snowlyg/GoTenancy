package publish

import (
	"fmt"
	"net/http"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/l10n"
	"GoTenancy/libs/publish"
	"GoTenancy/libs/qor"
	"github.com/jinzhu/gorm"
)

type availableLocalesInterface interface {
	AvailableLocales() []string
}

type publishableLocalesInterface interface {
	PublishableLocales() []string
}

type editableLocalesInterface interface {
	EditableLocales() []string
}

func getPublishableLocales(req *http.Request, currentUser interface{}) []string {
	if user, ok := currentUser.(publishableLocalesInterface); ok {
		return user.PublishableLocales()
	}

	if user, ok := currentUser.(editableLocalesInterface); ok {
		return user.EditableLocales()
	}

	if user, ok := currentUser.(availableLocalesInterface); ok {
		return user.AvailableLocales()
	}
	return []string{l10n.Global}
}

// RegisterL10nForPublish register l10n language switcher for publish
func RegisterL10nForPublish(Publish *publish.Publish, Admin *admin.Admin) {
	searchHandler := Publish.SearchHandler
	Publish.SearchHandler = func(db *gorm.DB, context *qor.Context) *gorm.DB {
		if context != nil {
			if context.Request != nil && context.Request.URL.Query().Get("locale") == "" {
				publishableLocales := getPublishableLocales(context.Request, context.CurrentUser)
				return searchHandler(db, context).Set("l10n:mode", "unscoped").Scopes(func(db *gorm.DB) *gorm.DB {
					scope := db.NewScope(db.Value)
					if l10n.IsLocalizable(scope) {
						return db.Where(fmt.Sprintf("%v.language_code IN (?)", scope.QuotedTableName()), publishableLocales)
					}
					return db
				})
			}
			return searchHandler(db, context).Set("l10n:mode", "locale")
		}
		return searchHandler(db, context).Set("l10n:mode", "unscoped")
	}

	Admin.RegisterViewPath("GoTenancy/libs/l10n/publish/views")

	Admin.RegisterFuncMap("publishable_locales", func(context admin.Context) []string {
		return getPublishableLocales(context.Request, context.CurrentUser)
	})
}
