package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/iris-contrib/httpexpect/v2"
	"github.com/kataras/iris/v12/httptest"
	"github.com/snowlyg/go-tenancy/core"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"
	"github.com/snowlyg/multi"
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
	// 初始化认证服务
	initialize.Auth()

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

func baseWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "authorityType": multi.AdminAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("登录成功")
	data := obj.Value("data").Object()
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func generalWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := baseTester(t)
	obj := e.POST("/v1/public/login").
		WithJSON(map[string]interface{}{"username": "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM", "password": "123456", "authorityType": multi.GeneralAuthority}).
		Expect().Status(httptest.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("登录成功")
	data := obj.Value("data").Object()
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func baseLogOut(auth *httpexpect.Expect) {
	obj := auth.GET("/v1/auth/logout").
		Expect().Status(httptest.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("code", "data", "msg")
	obj.Value("code").Number().Equal(0)
	obj.Value("msg").String().Equal("退出登录")
}
