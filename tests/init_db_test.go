package tests

import (
	"net/http"
	"testing"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/service"
)

func TestInitDB(t *testing.T) {
	e := baseTester(t)

	// 删除表和视图
	var sqls []string
	if err := g.TENANCY_DB.Raw("select CASE table_type WHEN 'VIEW' THEN concat('drop view ', table_name, ';') ELSE concat('drop table ', table_name, ';') END  from information_schema.tables where table_schema='tenancy';").Scan(&sqls).Error; err != nil {
		t.Fatalf("get drop table sql err %v\n", err)
	}

	for _, sql := range sqls {
		if err := g.TENANCY_DB.Exec(sql).Error; err != nil {
			continue
		}
	}

	MysqlConfig := config.Mysql{
		Path:     "",
		Dbname:   "",
		Username: "",
		Password: "",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if err := service.WriteConfig(g.TENANCY_VP, MysqlConfig); err != nil {
		t.Fatalf("write config err %v\n", err)
	}
	g.TENANCY_DB = nil
	obj := e.GET("/v1/init/checkdb").
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("前往初始化数据库")
	obj.Value("data").Object().Value("needInit").Boolean().True()

	obj = e.POST("/v1/init/initdb").
		WithJSON(map[string]interface{}{"host": "127.0.0.1", "port": "3306", "userName": "root", "password": "Chindeo", "dbName": "tenancy"}).
		Expect().Status(http.StatusOK).JSON().Object()

	obj.Keys().ContainsOnly("status", "data", "message")
	obj.Value("status").Number().Equal(200)
	obj.Value("message").String().Equal("操作成功")
	obj.Value("data").String().Equal("自动创建数据库成功")
}
