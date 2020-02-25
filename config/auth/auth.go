package auth

import (
	"fmt"
	"time"

	"GoTenancy/config"
	"GoTenancy/config/bindatafs"
	"GoTenancy/config/db"
	"GoTenancy/models/users"
	"GoTenancy/utils/registerviews"
	"github.com/fatih/color"
	"github.com/qor/auth"
	"github.com/qor/auth/authority"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/auth_themes/clean"
	"github.com/qor/render"
)

var (
	// Auth 初始化用于认证的 Auth
	Auth = clean.New(&auth.Config{
		DB:         db.DB,
		Mailer:     config.Mailer,
		Render:     render.New(&render.Config{AssetFileSystem: bindatafs.AssetFS.NameSpace("auth")}),
		UserModel:  users.User{},
		Redirector: auth.Redirector{RedirectBack: config.RedirectBack},
	})

	// Authority 初始化用于认证的 Authority
	Authority = authority.New(&authority.Config{
		Auth: Auth,
	})
)

func init() {

	if err := Auth.Render.AssetFileSystem.RegisterPath(registerviews.DetectViewsDir("github.com/qor", "auth")); err != nil {
		color.Red(fmt.Sprintf(" Auth.Render.AssetFileSystem.RegisterPath %v\n", err))
	}

	if err := Auth.Render.AssetFileSystem.RegisterPath(registerviews.DetectViewsDir("github.com/qor", "auth_themes")); err != nil {
		color.Red(fmt.Sprintf(" Auth.Render.AssetFileSystem.RegisterPath %v\n", err))
	}

	Auth.RegisterProvider(github.New(&config.Config.Github))
	Auth.RegisterProvider(google.New(&config.Config.Google))
	Auth.RegisterProvider(facebook.New(&config.Config.Facebook))
	Auth.RegisterProvider(twitter.New(&config.Config.Twitter))

	Authority.Register("logged_in_half_hour", authority.Rule{TimeoutSinceLastLogin: time.Minute * 30})

}
