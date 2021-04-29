package initialize

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/router"
)

// 初始化总路由

func Routers() *iris.Application {
	var Router = iris.Default()
	// Router.StaticFS(g.TENANCY_CONFIG.Local.Path, http.Dir(g.TENANCY_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	g.TENANCY_LOG.Info("use middleware logger")
	// 跨域
	//Router.Use(middleware.Cors()) // 如需跨域可以打开
	g.TENANCY_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Party("")
	{
		router.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup) // 自动初始化相关
	}
	PrivateGroup := Router.Party("", middleware.JWTAuth(), middleware.CasbinHandler())
	fmt.Println(PrivateGroup)
	{
		router.InitApiRouter(PrivateGroup)                 // 注册功能api路由
		router.InitJwtRouter(PrivateGroup)                 // jwt相关路由
		router.InitUserRouter(PrivateGroup)                // 注册用户路由
		router.InitMenuRouter(PrivateGroup)                // 注册menu路由
		router.InitEmailRouter(PrivateGroup)               // 邮件相关路由
		router.InitSystemRouter(PrivateGroup)              // system相关路由
		router.InitCasbinRouter(PrivateGroup)              // 权限相关路由
		router.InitAuthorityRouter(PrivateGroup)           // 注册角色路由
		router.InitSysDictionaryRouter(PrivateGroup)       // 字典管理
		router.InitSysOperationRecordRouter(PrivateGroup)  // 操作记录
		router.InitSysDictionaryDetailRouter(PrivateGroup) // 字典详情管理
	}
	g.TENANCY_LOG.Info("router register success")
	return Router
}
