package auth

import (
	"time"

	"github.com/qor/render"
	themes "github.com/snowlyg/qor-auth-theme"
	"go-tenancy/config"
	"go-tenancy/config/bindatafs"

	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"go-tenancy/config/db"
	"go-tenancy/models/users"
)

var (
	// Auth 初始化用于认证的 Auth
	TenantAuth = themes.New(&auth.Config{
		DB:     db.DB,
		Mailer: config.Mailer,
		Render: render.New(&render.Config{
			AssetFileSystem: bindatafs.AssetFS.NameSpace("auth"),
			ViewPaths:       []string{"app/tenant/views"},
		}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority 初始化用于认证的 Authority
	TenantAuthority = authority.New(&authority.Config{
		Auth: TenantAuth,
	})
)

func init() {
	TenantAuth.RegisterProvider(github.New(&config.Config.Github))
	TenantAuth.RegisterProvider(google.New(&config.Config.Google))
	TenantAuth.RegisterProvider(facebook.New(&config.Config.Facebook))
	TenantAuth.RegisterProvider(twitter.New(&config.Config.Twitter))

	TenantAuthority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})
}
