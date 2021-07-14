package initialize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/router/admin"
	"github.com/snowlyg/go-tenancy/router/client"
	"github.com/snowlyg/go-tenancy/router/general"
	"github.com/snowlyg/go-tenancy/router/public"
	"github.com/snowlyg/go-tenancy/utils"
)

// 初始化总路由

func App() *gin.Engine {
	gin.SetMode(g.TENANCY_CONFIG.System.Level)
	App := gin.Default()

	// 注册已定义验证方法
	utils.RegisterValidation()
	// 注册路由
	Routers(App)
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
		public.InitPublicRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		public.InitInitRouter(PublicGroup)   // 自动初始化相关
	}

	V1Group := app.Group("/v1", middleware.Auth(), middleware.CasbinHandler(), middleware.OperationRecord())
	{
		Auth := V1Group.Group("/auth")
		{
			public.InitAuthRouter(Auth) // 注册用户路由
		}

		// 商户和员工
		AdminGroup := V1Group.Group(g.TENANCY_CONFIG.System.AdminPreix, middleware.IsAdmin())
		{
			admin.InitApiRouter(AdminGroup)                // 注册功能api路由
			admin.InitUserRouter(AdminGroup)               // 注册用户路由
			admin.InitTenancyRouter(AdminGroup)            // 注册商户路由
			admin.InitMiniRouter(AdminGroup)               // 注册小程序路由
			admin.InitBrandRouter(AdminGroup)              // 注册品牌路由
			admin.InitBrandCategoryRouter(AdminGroup)      // 注册品牌分类路由
			admin.InitConfigCategoryRouter(AdminGroup)     // 注册系统配置分类路由
			admin.InitConfigRouter(AdminGroup)             // 注册系统配置路由
			admin.InitConfigValueRouter(AdminGroup)        // 注册系统配置值路由
			admin.InitMenuRouter(AdminGroup)               // 注册menu路由
			admin.InitEmailRouter(AdminGroup)              // 邮件相关路由
			admin.InitSystemRouter(AdminGroup)             // system相关路由
			admin.InitCasbinRouter(AdminGroup)             // 权限相关路由
			admin.InitAuthorityRouter(AdminGroup)          // 注册角色路由
			admin.InitMediaRouter(AdminGroup)              // 媒体库路由
			admin.InitCategoryRouter(AdminGroup)           // 商品分类路由
			admin.InitProductRouter(AdminGroup)            // 商品路由
			admin.InitOrderRouter(AdminGroup)              // 订单路由
			admin.InitRefundOrderRouter(AdminGroup)        // 退款订单路由
			admin.InitExpressRouter(AdminGroup)            // 物流公司路由
			admin.InitSysOperationRecordRouter(AdminGroup) // 操作记录
		}

		// 商户和员工
		ClientGroup := V1Group.Group(g.TENANCY_CONFIG.System.ClientPreix, middleware.IsTenancy())
		{

			client.InitTenancyRouter(ClientGroup)            // 注册商户路由
			client.InitBrandRouter(ClientGroup)              // 注册品牌路由
			client.InitConfigRouter(ClientGroup)             // 注册系统配置路由
			client.InitConfigValueRouter(ClientGroup)        // 注册系统配置值路由
			client.InitMenuRouter(ClientGroup)               // 注册menu路由
			client.InitMediaRouter(ClientGroup)              // 媒体库路由
			client.InitCategoryRouter(ClientGroup)           // 商品分类路由
			client.InitAttrTemplateRouter(ClientGroup)       // 规格模板路由
			client.InitProductRouter(ClientGroup)            // 商品路由
			client.InitShippingTemplateRouter(ClientGroup)   // 运费模板路由
			client.InitOrderRouter(ClientGroup)              // 订单路由
			client.InitRefundOrderRouter(ClientGroup)        // 退款订单路由
			client.InitExpressRouter(ClientGroup)            // 物流公司路由
			client.InitSysOperationRecordRouter(ClientGroup) // 操作记录
		}

		GeneralGroup := V1Group.Group("/general", middleware.IsGeneral())
		{
			general.InitAddressRouter(GeneralGroup) //我的地址管理
			general.InitReceiptRouter(GeneralGroup) //我的发票管理
		}
	}
}
