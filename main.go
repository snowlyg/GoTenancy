package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"GoTenancy/app/account"
	adminapp "GoTenancy/app/admin"
	"GoTenancy/app/api"
	"GoTenancy/app/home"
	"GoTenancy/app/orders"
	"GoTenancy/app/pages"
	"GoTenancy/app/products"
	"GoTenancy/app/static"
	"GoTenancy/config"
	"GoTenancy/config/application"
	"GoTenancy/config/auth"
	"GoTenancy/config/bindatafs"
	"GoTenancy/config/db"
	"GoTenancy/libs/admin"
	"GoTenancy/libs/publish2"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/middleware"
	"GoTenancy/utils/funcmapmaker"
	"github.com/fatih/color"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
)

func main() {

	// 命令参数处理
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	if err := cmdLine.Parse(os.Args[1:]); err != nil {
		color.Red(fmt.Sprintf(" cmdLine.Parse error :%v", err))
	}

	var (
		//Router = chi.NewRouter() // 定义路由
		IrisApp = iris.New()
		//定义 admin 对象
		Admin = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancy", // 站点名称
			Auth: &auth.AdminAuth{Paths: auth.PathConfig{
				Login:  "/auth/login",
				Logout: "/auth/logout",
			}},
			DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
		})

		//定义应用
		Application = application.New(&application.Config{
			IrisApp: IrisApp,
			Admin:   Admin,
			DB:      db.DB,
		})
	)

	// 认证相关视图渲染
	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	// 全局中间件
	IrisApp.Use(middleware.AddHeader)
	IrisApp.Logger().SetLevel("debug")
	IrisApp.Use(logger.New())
	IrisApp.Use(recover2.New())

	// 本地化 && publish2.PreviewByDB
	IrisApp.Use(middleware.Locale)

	// 加载应用
	Application.Use(api.New(&api.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(products.New(&products.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(orders.New(&orders.Config{}))
	Application.Use(pages.New(&pages.Config{}))
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
			//if err := app.Listen(fmt.Sprintf(":%d", config.Config.Port)); err != nil {
			//	panic(err)
			//}
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
			if err := IrisApp.Run(iris.Addr(fmt.Sprintf(":%d", config.Config.Port)), iris.WithoutServerError(iris.ErrServerClosed), irisConfig); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		}
	}
}
