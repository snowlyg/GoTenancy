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
	"go-tenancy/app/account"
	adminapp "go-tenancy/app/admin"
	"go-tenancy/app/home"
	"go-tenancy/app/static"
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

	// 初始化 iris
	var (
		adminAuth = auth.NewAdminAuth(&auth.PathConfig{})
		//定义 admin 对象
		Admin = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancy", // 站点名称
			Auth:     adminAuth,
			DB:       db.DB,
		})

		//定义应用
		Application = application.New(&application.Config{
			Admin: Admin,
			DB:    db.DB,
		})
	)

	// 认证相关视图渲染
	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	// 全局中间件
	//irisApp.Use(middleware.AddHeader)
	Application.IrisApp.Logger().SetLevel("debug")
	Application.IrisApp.Use(logger.New())
	Application.IrisApp.Use(recover2.New())

	// 加载应用
	//Application.Use(api.New(&api.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(account.New(&account.Config{}))

	// 系统上传文件
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/system"},
		Handler: utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))),
	}))
	// 静态打包文件加载
	prefixs := []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"}
	Application.Use(static.New(&static.Config{
		Prefixs: prefixs, // 设置静态文件相关目录
		Handler: bindatafs.AssetFS.FileServer("public", prefixs...),
	}))

	if *compileTemplate { //处理前端静态文件
		if err := bindatafs.AssetFS.Compile(); err != nil {
			color.Red(fmt.Sprintf("bindatafs error %v", err))
		}
	} else {
		if config.Config.HTTPS {
			// 启动服务
			if err := Application.IrisApp.Run(iris.TLS(":443", "./config/local_certs/server.crt", "./config/local_certs/server.key")); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		} else {
			// iris 配置设置
			irisConfig := iris.WithConfiguration(
				iris.Configuration{
					DisableStartupLog:                 true,
					FireMethodNotAllowed:              true,
					DisableBodyConsumptionOnUnmarshal: true,
					TimeFormat:                        "Mon, 01 Jan 2006 15:04:05 GMT",
					Charset:                           "UTF-8",
				})
			// 启动服务
			if err := Application.IrisApp.Run(iris.Addr(fmt.Sprintf(":%d", config.Config.Port)), iris.WithoutServerError(iris.ErrServerClosed), irisConfig); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		}
	}
}
