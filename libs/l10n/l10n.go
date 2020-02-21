package l10n

import (
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/roles"
)

// Global global language
var Global = "en-US"

type l10nInterface interface {
	IsGlobal() bool
	SetLocale(locale string)
}

// Locale embed this struct into GROM-backend models to enable localization feature for your model
type Locale struct {
	LanguageCode string `sql:"size:20" gorm:"primary_key"`
}

// IsGlobal return if current locale is global
func (l Locale) IsGlobal() bool {
	return l.LanguageCode == Global
}

// SetLocale set model's locale
func (l *Locale) SetLocale(locale string) {
	l.LanguageCode = locale
}

// LocaleCreatable if you embed it into your model, it will make the resource be creatable from locales, by default, you can only create it from global
type LocaleCreatable struct {
	Locale
}

// CreatableFromLocale a method to allow your mod=el be creatable from locales
func (LocaleCreatable) CreatableFromLocale() {}

type availableLocalesInterface interface {
	AvailableLocales() []string
}

type viewableLocalesInterface interface {
	ViewableLocales() []string
}

type editableLocalesInterface interface {
	EditableLocales() []string
}

func getAvailableLocales(req *http.Request, currentUser interface{}) []string {
	if user, ok := currentUser.(viewableLocalesInterface); ok {
		return user.ViewableLocales()
	}

	if user, ok := currentUser.(availableLocalesInterface); ok {
		return user.AvailableLocales()
	}
	return []string{Global}
}

func getEditableLocales(req *http.Request, currentUser interface{}) []string {
	if user, ok := currentUser.(editableLocalesInterface); ok {
		return user.EditableLocales()
	}

	if user, ok := currentUser.(availableLocalesInterface); ok {
		return user.AvailableLocales()
	}
	return []string{Global}
}

func getLocaleFromContext(context *qor.Context) string {
	if locale := utils.GetLocale(context); locale != "" {
		return locale
	}

	return Global
}

// LocalizeActionArgument localize action's argument
type LocalizeActionArgument struct {
	From string
	To   []string
}

// ConfigureQorResource configure qor locale for Qor Admin
func (l *Locale) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		Admin := res.GetAdmin()
		res.UseTheme("l10n")

		if res.Permission == nil {
			res.Permission = roles.NewPermission()
		}
		res.Permission.Allow(roles.CRUD, "locale_admin").Allow(roles.Read, "locale_reader")

		res.Meta(&admin.Meta{Name: "Localization", Type: "localization", Valuer: func(value interface{}, ctx *qor.Context) interface{} {
			var languageCodes []string
			var db = ctx.GetDB()
			var scope = db.NewScope(value)
			db.New().Set("l10n:mode", "unscoped").Model(res.Value).Where(fmt.Sprintf("%v = ?", scope.PrimaryKey()), scope.PrimaryKeyValue()).Pluck("language_code", &languageCodes)
			return utils.SliceUniq(languageCodes)
		}})

		res.OverrideIndexAttrs(func() {
			var attrs = res.ConvertSectionToStrings(res.IndexAttrs())
			var hasLocalization bool
			for _, attr := range attrs {
				if attr == "Localization" || attr == "-Localization" {
					hasLocalization = true
					break
				}
			}

			if hasLocalization {
				res.IndexAttrs(res.IndexAttrs(), "-LanguageCode")
			} else {
				res.IndexAttrs(res.IndexAttrs(), "-LanguageCode", "Localization")
			}
		})
		res.OverrideShowAttrs(func() {
			res.ShowAttrs(res.ShowAttrs(), "-LanguageCode", "-Localization")
		})
		res.NewAttrs(res.NewAttrs(), "-LanguageCode", "-Localization")
		res.EditAttrs(res.EditAttrs(), "-LanguageCode", "-Localization")

		// Set meta permissions
		for _, field := range Admin.DB.NewScope(res.Value).Fields() {
			if isSyncField(field.StructField) {
				if meta := res.GetMeta(field.Name); meta != nil {
					permission := meta.Meta.Permission
					if permission == nil {
						permission = roles.Allow(roles.CRUD, "global_admin").Allow(roles.Read, "locale_reader")
					} else {
						permission = permission.Allow(roles.CRUD, "global_admin").Allow(roles.Read, "locale_reader")
					}

					meta.SetPermission(permission)
				}
			}
		}

		// Roles
		role := res.Permission.Role
		if _, ok := role.Get("global_admin"); !ok {
			role.Register("global_admin", func(req *http.Request, currentUser interface{}) bool {
				if getLocaleFromContext(&qor.Context{Request: req}) == Global {
					for _, locale := range getEditableLocales(req, currentUser) {
						if locale == Global {
							return true
						}
					}
				}
				return false
			})
		}

		if _, ok := role.Get("locale_admin"); !ok {
			role.Register("locale_admin", func(req *http.Request, currentUser interface{}) bool {
				currentLocale := getLocaleFromContext(&qor.Context{Request: req})
				for _, locale := range getEditableLocales(req, currentUser) {
					if locale == currentLocale {
						return true
					}
				}
				return false
			})
		}

		if _, ok := role.Get("locale_reader"); !ok {
			role.Register("locale_reader", func(req *http.Request, currentUser interface{}) bool {
				currentLocale := getLocaleFromContext(&qor.Context{Request: req})
				for _, locale := range getAvailableLocales(req, currentUser) {
					if locale == currentLocale {
						return true
					}
				}
				return false
			})
		}

		// Inject for l10n
		admin.RegisterViewPath("libs/l10n/views")

		// Middleware
		Admin.GetRouter().Use(&admin.Middleware{
			Name: "l10n_set_locale",
			Handler: func(context *admin.Context, middleware *admin.Middleware) {
				db := context.GetDB().Set("l10n:locale", getLocaleFromContext(context.Context))
				if mode := context.Request.URL.Query().Get("locale_mode"); mode != "" {
					db = db.Set("l10n:mode", mode)
				}

				usingLanguageCodeAsPrimaryKey := false
				if res := context.Resource; res != nil {
					for idx, primaryField := range res.PrimaryFields {
						if primaryField.Name == "LanguageCode" {
							_, params := res.ToPrimaryQueryParams(res.GetPrimaryValue(context.Request), context.Context)
							if len(params) > idx {
								usingLanguageCodeAsPrimaryKey = true
								db = db.Set("l10n:locale", params[idx])

								// PUT usually used for localize
								if context.Request.Method == "PUT" {
									if _, ok := db.Get("l10n:localize_to"); !ok {
										db = db.Set("l10n:localize_to", getLocaleFromContext(context.Context))
									}
								}
								break
							}
						}
					}
				}

				if !usingLanguageCodeAsPrimaryKey {
					for key, values := range context.Request.URL.Query() {
						if regexp.MustCompile(`primary_key\[.+_language_code\]`).MatchString(key) {
							if len(values) > 0 {
								db = db.Set("l10n:locale", values[0])

								// PUT usually used for localize
								if context.Request.Method == "PUT" || context.Request.Method == "POST" {
									db = db.Set(key, "")
									if _, ok := db.Get("l10n:localize_to"); !ok {
										db = db.Set("l10n:localize_to", getLocaleFromContext(context.Context))
									}
								}
							}
						}
					}
				}

				if context.Request.URL.Query().Get("sorting") != "" {
					db = db.Set("l10n:mode", "locale")
				}
				context.SetDB(db)

				middleware.Next(context)
			},
		})

		// FunMap
		Admin.RegisterFuncMap("current_locale", func(context admin.Context) string {
			return getLocaleFromContext(context.Context)
		})

		Admin.RegisterFuncMap("global_locale", func() string {
			return Global
		})

		Admin.RegisterFuncMap("viewable_locales", func(context admin.Context) []string {
			return getAvailableLocales(context.Request, context.CurrentUser)
		})

		Admin.RegisterFuncMap("editable_locales", func(context admin.Context) []string {
			return getEditableLocales(context.Request, context.CurrentUser)
		})

		Admin.RegisterFuncMap("createable_locales", func(context admin.Context) []string {
			editableLocales := getEditableLocales(context.Request, context.CurrentUser)
			if _, ok := context.Resource.Value.(localeCreatableInterface); ok {
				return editableLocales
			}

			for _, locale := range editableLocales {
				if locale == Global {
					return []string{Global}
				}
			}
			return []string{}
		})

		if res.GetAction("Localize") == nil {
			argumentResource := Admin.NewResource(&LocalizeActionArgument{})
			argumentResource.Meta(&admin.Meta{
				Name: "From",
				Type: "select_one",
				Valuer: func(_ interface{}, context *qor.Context) interface{} {
					return Global
				},
				Collection: func(value interface{}, context *qor.Context) (results [][]string) {
					for _, locale := range getAvailableLocales(context.Request, context.CurrentUser) {
						results = append(results, []string{locale, locale})
					}
					return
				},
			})
			argumentResource.Meta(&admin.Meta{
				Name: "To",
				Type: "select_many",
				Valuer: func(_ interface{}, context *qor.Context) interface{} {
					return []string{getLocaleFromContext(context)}
				},
				Collection: func(value interface{}, context *qor.Context) (results [][]string) {
					for _, locale := range getEditableLocales(context.Request, context.CurrentUser) {
						results = append(results, []string{locale, locale})
					}
					return
				},
			})

			res.Action(&admin.Action{
				Name: "Localize",
				Handler: func(argument *admin.ActionArgument) error {
					var (
						db        = argument.Context.GetDB()
						arg       = argument.Argument.(*LocalizeActionArgument)
						results   = res.NewSlice()
						sqls      []string
						sqlParams []interface{}
					)

					for _, primaryValue := range argument.PrimaryValues {
						primaryQuerySQL, primaryParams := res.ToPrimaryQueryParams(primaryValue, argument.Context.Context)
						sqls = append(sqls, primaryQuerySQL)
						sqlParams = append(sqlParams, primaryParams...)
					}

					db.Set("l10n:locale", arg.From).Where(strings.Join(sqls, " OR "), sqlParams...).Find(results)

					reflectResults := reflect.Indirect(reflect.ValueOf(results))
					for i := 0; i < reflectResults.Len(); i++ {
						for _, to := range arg.To {
							if err := db.Set("l10n:localize_to", to).Unscoped().Save(reflectResults.Index(i).Interface()).Error; err != nil {
								return err
							}
						}
					}
					return nil
				},
				Modes:      []string{"index", "menu_item"},
				Permission: roles.Allow(roles.CRUD, roles.Anyone),
				Resource:   argumentResource,
			})
		}
	}
}
