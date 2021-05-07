package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
)

func TestMain(m *testing.M) {
	g.TENANCY_VP = core.Viper()      // 初始化Viper
	g.TENANCY_LOG = core.Zap()       // 初始化zap日志库
	g.TENANCY_DB = initialize.Gorm() // gorm连接数据库
	initialize.Timer()
	if g.TENANCY_DB != nil {
		initialize.MysqlTables(g.TENANCY_DB) // 初始化表
		// 程序结束前关闭数据库链接
		g.TENANCY_DB.DB()
	}
	if g.TENANCY_CONFIG.System.CacheType == "redis" {
		// 初始化redis服务
		initialize.Auth()
	}

	// call flag.Parse() here if TestMain uses flags
	// 如果 TestMain 使用了 flags，这里应该加上 flag.Parse()
	os.Exit(m.Run())

	db, _ := g.TENANCY_DB.DB()
	db.Close()
	g.TENANCY_AUTH.Close()
}

func baseTester(t *testing.T) *httpexpect.Expect {
	handler := initialize.Routers()
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://127.0.0.1:8089",
		Client: &http.Client{
			Transport: httpexpect.NewBinder(handler),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})
}
