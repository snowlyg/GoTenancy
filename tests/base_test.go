package tests

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect"
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
	multi.AuthDriver.Close()
}

func baseTester(t *testing.T) *httpexpect.Expect {
	handler := initialize.App()
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL: "http://127.0.0.1:8089/",
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
	obj := e.POST("v1/public/admin/login").
		WithJSON(map[string]interface{}{"username": "admin", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(1)
	user.Value("userName").String().Equal("admin")
	user.Value("email").String().Equal("admin@admin.com")
	user.Value("nickName").String().Equal("超级管理员")
	user.Value("authorityName").String().Equal("超级管理员")
	user.Value("authorityType").Number().Equal(multi.AdminAuthority)
	user.Value("authorityId").String().Equal("999")
	user.Value("defaultRouter").String().Equal("dashboard")
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func tenancyWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := baseTester(t)
	obj := e.POST("v1/public/merchant/login").
		WithJSON(map[string]interface{}{"username": "a303176530", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(2)
	user.Value("userName").String().Equal("a303176530")
	user.Value("email").String().Equal("a303176530@admin.com")
	user.Value("nickName").String().Equal("商户管理员")
	user.Value("authorityName").String().Equal("商户管理员")
	user.Value("authorityType").Number().Equal(multi.TenancyAuthority)
	user.Value("authorityId").String().Equal("998")
	user.Value("defaultRouter").String().Equal("dashboard")
	user.Value("tenancyName").String().Equal("宝安中心人民医院")
	user.Value("tenancyId").Number().Equal(1)
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func generalWithLoginTester(t *testing.T) *httpexpect.Expect {
	e := baseTester(t)
	obj := e.POST("v1/public/login").
		WithJSON(map[string]interface{}{"username": "oZM5VwD_PCaPKQZ8zRGt-NUdU2uM", "password": "123456", "captcha": "", "captchaId": ""}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("登录成功")
	data := obj.Value("data").Object()
	user := data.Value("user").Object()
	user.Value("id").Number().Equal(3)
	user.Value("userName").String().Equal("oZM5VwD_PCaPKQZ8zRGt-NUdU2uM")
	user.Value("email").String().Equal("a303176530@admin.com")
	user.Value("nickName").String().Equal("C端用户")
	user.Value("authorityName").String().Equal("C端用户")
	user.Value("authorityType").Number().Equal(multi.GeneralAuthority)
	user.Value("authorityId").String().Equal("997")
	user.Value("avatarUrl").String().NotEmpty()
	user.Value("sex").Number().Equal(2)
	user.Value("subscribe").Boolean().True()
	user.Value("openId").String().Equal("own1t5TysymNUqcZm-8giuEvT68M")
	user.Value("unionId").String().Equal("oZM5VwCgvGUZvkrnrGrdJZI4e12k")
	user.Value("country").String().Equal("")
	user.Value("province").String().Equal("")
	user.Value("city").String().Equal("")
	user.Value("idCard").String().Equal("445281199411285861")
	user.Value("isAuth").Boolean().False()
	user.Value("realName").String().Equal("余思琳")
	user.Value("birthday").String().Equal("1994-11-28T08:00:00+08:00")
	data.Value("AccessToken").NotNull()

	token := data.Value("AccessToken").String().Raw()
	return e.Builder(func(req *httpexpect.Request) {
		req.WithHeader("Authorization", "Bearer "+token)
	})
}

func baseLogOut(auth *httpexpect.Expect) {
	obj := auth.GET("v1/auth/logout").
		Expect().Status(http.StatusOK).JSON().Object()
	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("退出登录")
}
