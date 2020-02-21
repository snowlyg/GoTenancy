package clean

import (
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"GoTenancy/libs/auth"
	"GoTenancy/libs/auth/auth_identity"
	"GoTenancy/libs/auth/claims"
	"GoTenancy/libs/auth/providers/password"
	"GoTenancy/libs/i18n"
	"GoTenancy/libs/i18n/backends/yaml"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/render"
)

// ErrPasswordConfirmationNotMatch password confirmation not match error
var ErrPasswordConfirmationNotMatch = errors.New("password confirmation doesn't match password")

// New initialize clean theme
func New(config *auth.Config) *auth.Auth {
	if config == nil {
		config = &auth.Config{}
	}
	config.ViewPaths = append(config.ViewPaths, "GoTenancy/libs/auth_themes/clean/views")

	if config.DB == nil {
		fmt.Print("Please configure *gorm.DB for Auth theme clean")
	}

	if config.Render == nil {
		yamlBackend := yaml.New()
		I18n := i18n.New(yamlBackend)
		for _, gopath := range append([]string{filepath.Join(utils.AppRoot, "vendor")}, utils.GOPATH()...) {
			filePath := filepath.Join(gopath, "src", "GoTenancy/libs/auth_themes/clean/locales/en-US.yml")
			if content, err := ioutil.ReadFile(filePath); err == nil {
				translations, _ := yamlBackend.LoadYAMLContent(content)
				for _, translation := range translations {
					I18n.AddTranslation(translation)
				}
				break
			}
		}

		config.Render = render.New(&render.Config{
			FuncMapMaker: func(render *render.Render, req *http.Request, w http.ResponseWriter) template.FuncMap {
				return template.FuncMap{
					"t": func(key string, args ...interface{}) template.HTML {
						return I18n.T(utils.GetLocale(&qor.Context{Request: req}), key, args...)
					},
				}
			},
		})
	}

	Auth := auth.New(config)

	Auth.RegisterProvider(password.New(&password.Config{
		Confirmable: true,
		RegisterHandler: func(context *auth.Context) (*claims.Claims, error) {
			context.Request.ParseForm()

			if context.Request.Form.Get("confirm_password") != context.Request.Form.Get("password") {
				return nil, ErrPasswordConfirmationNotMatch
			}

			return password.DefaultRegisterHandler(context)
		},
		// 修复登陆报错
		AuthorizeHandler: func(context *auth.Context) (*claims.Claims, error) {
			var (
				authInfo    auth_identity.AuthIdentity
				req         = context.Request
				tx          = context.Auth.GetDB(req)
				provider, _ = context.Provider.(*password.Provider)
			)

			req.ParseForm()
			authInfo.Provider = provider.GetName()
			authInfo.UID = strings.TrimSpace(req.Form.Get("login"))

			if tx.Model(context.Auth.AuthIdentityModel).Where(authInfo).Scan(&authInfo).RecordNotFound() {
				return nil, auth.ErrInvalidAccount
			}

			if provider.Config.Confirmable && authInfo.ConfirmedAt == nil {
				currentUser, _ := context.Auth.UserStorer.Get(authInfo.ToClaims(), context)
				provider.Config.ConfirmMailer(authInfo.UID, context, authInfo.ToClaims(), currentUser)

				return nil, password.ErrUnconfirmed
			}

			if err := provider.Encryptor.Compare(authInfo.EncryptedPassword, strings.TrimSpace(req.Form.Get("password"))); err == nil {
				return authInfo.ToClaims(), err
			}

			return nil, auth.ErrInvalidPassword
		},
	}))

	if Auth.Config.DB != nil {
		// Migrate Auth Identity model
		Auth.Config.DB.AutoMigrate(Auth.Config.AuthIdentityModel)
	}
	return Auth
}
