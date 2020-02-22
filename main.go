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
			Auth:     auth.AdminAuth{},
			DB:       db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
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
	IrisApp.Use(func(ctx iris.Context) {
		// 演示设置，请勿在生产环境使用
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	IrisApp.Use(func(ctx iris.Context) {
		ctx.Request().Header.Del("Authorization")
		ctx.Next()
	})

	IrisApp.Logger().SetLevel("debug")
	IrisApp.Use(logger.New())
	IrisApp.Use(recover2.New())

	// 本地化
	//IrisApp.Use(iris.FromStd(func(next http.Handler)  http.Handler {
	//	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
	//		var (
	//			tx         = db.DB
	//			qorContext = &qor.Context{Request: req, Writer: w}
	//		)
	//
	//		if locale := utils.GetLocale(qorContext); locale != "" {
	//			tx = tx.Set("l10n:locale", locale)
	//		}
	//
	//		ctx := context2.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
	//		next.ServeHTTP(w, req.WithContext(ctx))
	//	})
	//}))

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
			// 启动服务
			if err := IrisApp.Run(iris.Addr(fmt.Sprintf(":%d", config.Config.Port)), iris.WithoutServerError(iris.ErrServerClosed), iris.WithConfiguration(
				iris.Configuration{
					DisableStartupLog:                 true,
					FireMethodNotAllowed:              true,
					DisableBodyConsumptionOnUnmarshal: true,
					TimeFormat:                        "Mon, 01 Jan 2006 15:04:05 GMT",
					Charset:                           "UTF-8",
				}),
			); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		}

		// 使用 net/http 原生包
		//color.Yellow(fmt.Sprintf("Listening on: %v\n", config.Config.Port))
		//if config.Config.HTTPS {
		//	if err := http.ListenAndServeTLS(fmt.Sprintf(":%d", config.Config.Port), "config/local_certs/server.crt", "config/local_certs/server.key", Application.NewServeMux()); err != nil {
		//		panic(err)
		//	}
		//} else {
		//	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Port), Application.NewServeMux()); err != nil {
		//		panic(err)
		//	}
		//}
	}
}
