//go:generate go run config/db/seeds/main.go config/db/seeds/seeds.go
//go:generate go run main.go -compile-templates=true
package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/qor/admin"
	"github.com/qor/qor/utils"
	"go-tenancy/app/backend/account"
	adminapp "go-tenancy/app/backend/admin"
	"go-tenancy/app/backend/static"
	tenantaccount "go-tenancy/app/tenant/account"
	tenantapp "go-tenancy/app/tenant/admin"
	"go-tenancy/config"
	"go-tenancy/config/application"
	"go-tenancy/config/auth"
	"go-tenancy/config/bindatafs"
	"go-tenancy/config/db"
	"go-tenancy/utils/funcmapmaker"
)

func main() {

	// 命令参数处理
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	if err := cmdLine.Parse(os.Args[1:]); err != nil {
		color.Red(fmt.Sprintf(" cmdLine.Parse error :%v", err))
	}

	var (
		adminAuth = auth.NewAdminAuth(&auth.PathConfig{})
		Admin     = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancyAdmin",
			Auth:     adminAuth,
			DB:       db.DB,
		})

		tenantAuth = auth.NewAdminAuth(&auth.PathConfig{})
		Tenant     = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancyTenant",
			Auth:     tenantAuth,
			DB:       db.DB,
		})

		Iris        = iris.Default()
		Application = application.New(&application.Config{
			IrisApplication: application.IrisApplication{
				Iris:        Iris,
				AdminParty:  Iris.Subdomain("admin"),
				TenantParty: Iris.Subdomain("tenant"),
			},
			Admin:  Admin,
			Tenant: Tenant,
			DB:     db.DB,
		})
	)

	f := NewLogFile()
	defer f.Close()
	Iris.Logger().SetOutput(f) //记录日志

	Iris.Logger().SetLevel("debug")
	Iris.Use(logger.New())
	Iris.Use(recover2.New())

	// 认证相关视图渲染
	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)
	funcmapmaker.AddFuncMapMaker(auth.TenantAuth.Config.Render)

	// 加载应用
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/system"},
		Handler: utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))),
	}))
	prefixs := []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"}
	Application.Use(static.New(&static.Config{
		Prefixs: prefixs,
		Handler: bindatafs.AssetFS.FileServer("public", prefixs...),
	}))

	Application.Use(tenantapp.New(&tenantapp.Config{}))
	Application.Use(tenantaccount.New(&tenantaccount.Config{}))

	if *compileTemplate { //处理前端静态文件
		if err := bindatafs.AssetFS.Compile(); err != nil {
			color.Red(fmt.Sprintf("bindatafs error %v", err))
		}
	} else {
		if config.Config.HTTPS {
			if err := Iris.Run(iris.TLS(":443", "./config/local_certs/server.crt", "./config/local_certs/server.key")); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		} else {
			irisConfig := iris.WithConfiguration(
				iris.Configuration{
					DisableStartupLog:                 true,
					FireMethodNotAllowed:              true,
					DisableBodyConsumptionOnUnmarshal: true,
					TimeFormat:                        "Mon, 01 Jan 2006 15:04:05 GMT",
					Charset:                           "UTF-8",
				})
			if err := Iris.Run(iris.Addr(fmt.Sprintf(":%d", config.Config.Port)), iris.WithoutServerError(iris.ErrServerClosed), irisConfig); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		}
	}
}
