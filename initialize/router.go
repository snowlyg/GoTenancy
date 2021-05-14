package initialize

import (
	stdContext "context"
	"os"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/middleware"
	"github.com/snowlyg/go-tenancy/router"
	"go.uber.org/zap"
)

// 初始化总路由

var IdleConnsClosed = make(chan struct{})

func Routers() *iris.Application {
	Router := iris.New()
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts
		Router.Shutdown(ctx)
		close(IdleConnsClosed)
	})
	// Set default log level.
	Router.Logger().SetLevel(g.TENANCY_CONFIG.System.Level)

	// Register the accesslog middleware.
	logFile, err := os.OpenFile("./access.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err == nil {
		// Close the file on shutdown.
		Router.ConfigureHost(func(su *iris.Supervisor) {
			su.RegisterOnShutdown(func() {
				logFile.Close()
			})
		})

		ac := accesslog.New(logFile)
		ac.AddOutput(Router.Logger().Printer)
		Router.UseRouter(ac.Handler)
		Router.Logger().Debugf("Using <%s> to log requests", logFile.Name())
	}

	// Register the requestid middleware
	// before recover so current Context.GetID() contains the info on panic logs.
	// Router.UseRouter(requestid.New())
	// Router.Logger().Debugf("Using <UUID4> to identify requests")

	// Register the recovery, after accesslog and recover,
	// before end-developer's middleware.
	Router.UseRouter(recover.New())

	// Router.StaticFS(g.TENANCY_CONFIG.Local.Path, http.Dir(g.TENANCY_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	g.TENANCY_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors()) // 如需跨域可以打开
	g.TENANCY_LOG.Info("use middleware cors")
	// 方便统一添加路由组前缀 多服务器上线使用
	PublicGroup := Router.Party("/v1")
	{
		router.InitPublicRouter(PublicGroup) // 注册基础功能路由 不做鉴权
		router.InitInitRouter(PublicGroup)   // 自动初始化相关
	}
	AdminGroup := Router.Party("/v1/admin", middleware.JWTAuth(), middleware.CasbinHandler(), middleware.OperationRecord(), middleware.IsAdmin())
	{
		router.InitApiRouter(AdminGroup)                // 注册功能api路由
		router.InitUserRouter(AdminGroup)               // 注册用户路由
		router.InitTenancyRouter(AdminGroup)            // 注册商户路由
		router.InitMiniRouter(AdminGroup)               // 注册小程序路由
		router.InitMenuRouter(AdminGroup)               // 注册menu路由
		router.InitEmailRouter(AdminGroup)              // 邮件相关路由
		router.InitSystemRouter(AdminGroup)             // system相关路由
		router.InitCasbinRouter(AdminGroup)             // 权限相关路由
		router.InitAuthorityRouter(AdminGroup)          // 注册角色路由
		router.InitSysOperationRecordRouter(AdminGroup) // 操作记录
	}

	// TenancyGroup := Router.Party("/admin", middleware.JWTAuth(), middleware.CasbinHandler(), middleware.OperationRecord(), middleware.IsTenancy())
	// {

	// }

	// GeneralGroup := Router.Party("/admin", middleware.JWTAuth(), middleware.CasbinHandler(), middleware.OperationRecord(), middleware.IsGeneral())
	// {

	// }

	err = Router.Build()
	if err != nil {
		g.TENANCY_LOG.Error("router build", zap.Any("err", err))
	}
	g.TENANCY_LOG.Info("router register success")
	return Router
}
