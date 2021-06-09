package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/router"
	"github.com/snowlyg/go-tenancy/utils"
)

// 初始化总路由

func App() *gin.Engine {
	App := gin.Default()
	// 注册已定义验证方法
	utils.RegisterValidation()
	// 注册路由
	Routers(App)
	g.TENANCY_LOG.Info("router register success")
	return App
}

// Routers
func Routers(app *gin.Engine) {
	app.StaticFS(g.TENANCY_CONFIG.Local.Path, http.Dir(g.TENANCY_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	g.TENANCY_LOG.Info("use middleware logger")

	// 跨域
	app.Use(middleware.Cors()) // 如需跨域可以打开
	g.TENANCY_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := app.Group("/v1")
	{
		router.InitPublicRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup)   // 自动初始化相关
	}

	V1Group := app.Group("/v1", middleware.Auth(), middleware.CasbinHandler(), middleware.OperationRecord())
	{
		Auth := V1Group.Group("/auth")
		{
			router.InitAuthRouter(Auth) // 注册用户路由
		}

		// 商户和员工
		AdminGroup := V1Group.Group("/admin", middleware.IsAdmin())
		{
			router.InitApiRouter(AdminGroup)                // 注册功能api路由
			router.InitUserRouter(AdminGroup)               // 注册用户路由
			router.InitTenancyRouter(AdminGroup)            // 注册商户路由
			router.InitMiniRouter(AdminGroup)               // 注册小程序路由
			router.InitBrandRouter(AdminGroup)              // 注册品牌路由
			router.InitBrandCategoryRouter(AdminGroup)      // 注册品牌分类路由
			router.InitConfigCategoryRouter(AdminGroup)     // 注册系统配置分类路由
			router.InitConfigRouter(AdminGroup)             // 注册系统配置路由
			router.InitMenuRouter(AdminGroup)               // 注册menu路由
			router.InitEmailRouter(AdminGroup)              // 邮件相关路由
			router.InitSystemRouter(AdminGroup)             // system相关路由
			router.InitCasbinRouter(AdminGroup)             // 权限相关路由
			router.InitAuthorityRouter(AdminGroup)          // 注册角色路由
			router.InitSysOperationRecordRouter(AdminGroup) // 操作记录

			// 商户
			router.InitMediaRouter(AdminGroup)        // 媒体库路由
			router.InitCategoryRouter(AdminGroup)     // 商品分类路由
			router.InitAttrTemplateRouter(AdminGroup) // 规格模板路由
			router.InitProductRouter(AdminGroup)      // 商品路由

		}

		GeneralGroup := V1Group.Group("/general", middleware.IsGeneral())
		{
			router.InitAddressRouter(GeneralGroup) //我的地址管理
			router.InitReceiptRouter(GeneralGroup) //我的发票管理
		}
	}
}
