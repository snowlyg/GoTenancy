package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"GoTenancy/app/account"
	adminapp "GoTenancy/app/admin"
	"GoTenancy/app/api"
	"GoTenancy/app/enterprise"
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
	"GoTenancy/utils/funcmapmaker"
	"github.com/fatih/color"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/kataras/iris/v12"
	"github.com/qor/admin"
	"github.com/qor/publish2"
	"github.com/qor/qor"
	"github.com/qor/qor/utils"
)

func main() {

	// 命令参数处理
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	if err := cmdLine.Parse(os.Args[1:]); err != nil {
		color.Red(fmt.Sprintf(" cmdLine.Parse error :%v", err))
	}

	var (
		Router = chi.NewRouter() // 定义路由
		//定义 admin 对象
		Admin = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancy", // 站点名称
			Auth:     auth.AdminAuth{},
			DB:       db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff),
		})
		//定义应用
		Application = application.New(&application.Config{
			Router: Router,
			Admin:  Admin,
			DB:     db.DB,
		})
	)

	// 认证相关视图渲染
	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	// 全局中间件
	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			// 演示设置，请勿在生产环境使用
			w.Header().Add("Access-Control-Allow-Origin", "*")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(func(handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			req.Header.Del("Authorization")
			handler.ServeHTTP(w, req)
		})
	})

	Router.Use(middleware.RealIP)
	Router.Use(middleware.Logger)
	Router.Use(middleware.Recoverer)
	// 本地化
	Router.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			var (
				tx         = db.DB
				qorContext = &qor.Context{Request: req, Writer: w}
			)

			if locale := utils.GetLocale(qorContext); locale != "" {
				tx = tx.Set("l10n:locale", locale)
			}

			ctx := context.WithValue(req.Context(), utils.ContextDBName, publish2.PreviewByDB(tx, qorContext))
			next.ServeHTTP(w, req.WithContext(ctx))
		})
	})

	// 加载应用
	Application.Use(api.New(&api.Config{}))
	Application.Use(adminapp.New(&adminapp.Config{}))
	Application.Use(home.New(&home.Config{}))
	Application.Use(products.New(&products.Config{}))
	Application.Use(account.New(&account.Config{}))
	Application.Use(orders.New(&orders.Config{}))
	Application.Use(pages.New(&pages.Config{}))
	Application.Use(enterprise.New(&enterprise.Config{}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"/system"},
		Handler: utils.FileServer(http.Dir(filepath.Join(config.Root, "public"))),
	}))
	Application.Use(static.New(&static.Config{
		Prefixs: []string{"javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"},
		Handler: bindatafs.AssetFS.FileServer(http.Dir("public"), "javascripts", "stylesheets", "images", "dist", "fonts", "vendors", "favicon.ico"),
	}))

	if *compileTemplate { //处理前端静态文件
		if err := bindatafs.AssetFS.Compile(); err != nil {
			color.Red(fmt.Sprintf("bindatafs error %v", err))
		}
	} else {
		color.Yellow(fmt.Sprintf("Listening on: %v\n", config.Config.Port))
		app := iris.Default()
		if config.Config.HTTPS {
			ser := &http.Server{Addr: fmt.Sprintf(":%d", config.Config.Port), Handler: Application.NewServeMux(), TLSConfig: &tls.Config{}}
			if err := app.Run(iris.Raw(ser.ListenAndServe)); err != nil {
				panic(err)
			}
		} else {
			ser := &http.Server{Addr: fmt.Sprintf(":%d", config.Config.Port), Handler: Application.NewServeMux()}
			if err := app.Run(iris.Raw(ser.ListenAndServe)); err != nil {
				panic(err)
			}
		}

		// 使用 net/http 原生包
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
