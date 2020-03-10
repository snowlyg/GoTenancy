package db

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/l10n"
	"github.com/qor/media"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"go-tenancy/config"
)

// DB 全局 DB 链接
var (
	DB *gorm.DB
	CE *casbin.Enforcer
)

func init() {
	var (
		err error
		c   *gormadapter.Adapter
	)

	if config.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", getConn())
		// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if config.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", getConn())
	} else if config.Config.DB.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", getConn())
	} else {
		panic(errors.New("not supported database adapter"))
	}

	c, err = gormadapter.NewAdapter(config.Config.DB.Adapter, getConn(), true) // Your driver and data source.
	if err != nil {
		color.Red(fmt.Sprintf("gormadapter.NewAdapter 错误: %v", err))
	}

	CE, err = casbin.NewEnforcer(filepath.Join("config", "casbin", "rbac_model.conf"), c)
	if err != nil {
		color.Red(fmt.Sprintf("NewEnforcer 错误: %v", err))
	}

	if err = CE.LoadPolicy(); err != nil {
		color.Red(fmt.Sprintf("LoadPolicy error %v\n", err))
	}

	if err == nil {
		if os.Getenv("DEBUG") != "" {
			DB.LogMode(true)
		}

		// 注册回调到 gorm DB
		l10n.RegisterCallbacks(DB)
		sorting.RegisterCallbacks(DB)
		validations.RegisterCallbacks(DB)
		media.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}

func getConn() string {
	dbConfig := config.Config.DB
	var conn string
	if config.Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	} else if config.Config.DB.Adapter == "postgres" {
		conn = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
	} else if config.Config.DB.Adapter == "sqlite" {
		conn = fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name)
	}

	return conn
}
