package irisapp

import (
	"net/http"
	"path/filepath"

	"GoTenancy/app/account"
	adminapp "GoTenancy/app/admin"
	"GoTenancy/app/home"
	"GoTenancy/app/static"
	"GoTenancy/config"
	"GoTenancy/config/application"
	"GoTenancy/config/auth"
	"GoTenancy/config/bindatafs"
	"GoTenancy/config/db"
	"GoTenancy/utils/funcmapmaker"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	recover2 "github.com/kataras/iris/v12/middleware/recover"
	"github.com/qor/admin"
	"github.com/qor/qor/utils"
)

func New() *iris.Application {
	var (
		adminAuth = auth.NewAdminAuth(&auth.PathConfig{})
		irisApp   = iris.New()
		//定义 admin 对象
		Admin = admin.New(&admin.AdminConfig{
			SiteName: "GoTenancy", // 站点名称
			Auth:     adminAuth,
			DB:       db.DB,
		})

		//定义应用
		Application = application.New(&application.Config{
			IrisApp: irisApp,
			Admin:   Admin,
			DB:      db.DB,
		})
	)

	// 认证相关视图渲染
	funcmapmaker.AddFuncMapMaker(auth.Auth.Config.Render)

	// 全局中间件
	//irisApp.Use(middleware.AddHeader)
	irisApp.Logger().SetLevel("debug")
	irisApp.Use(logger.New())
	irisApp.Use(recover2.New())

	// 加载静态资源 for /admin
	irisApp.HandleDir("/assets", "public/architectui-html-free/assets")
	irisApp.HandleDir("/", "public/architectui-html-free/style")

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

	return irisApp
}
