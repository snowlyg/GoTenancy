package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/l10n"
	"github.com/qor/media"
	iris_base_rabc "github.com/snowlyg/iris-base-rabc"

	//"github.com/qor/publish2"
	"github.com/qor/sorting"
	"github.com/qor/validations"
	"go-tenancy/config"
)

// DB 全局 DB 链接
var DB *gorm.DB

func init() {
	var err error

	conn := GetConn()
	if config.Config.DB.Adapter == "mysql" {
		DB, err = gorm.Open("mysql", conn)
		// DB = DB.Set("gorm:table_options", "CHARSET=utf8")
	} else if config.Config.DB.Adapter == "postgres" {
		DB, err = gorm.Open("postgres", conn)
	} else if config.Config.DB.Adapter == "sqlite" {
		DB, err = gorm.Open("sqlite3", conn)
	} else {
		panic(errors.New("not supported database adapter"))
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
		//publish2.RegisterCallbacks(DB)
	} else {
		panic(err)
	}
}

func GetConn() string {
	dbConfig := config.Config.DB
	var conn string
	if config.Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Name)
	} else if config.Config.DB.Adapter == "postgres" {
		conn = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Name)
	} else if config.Config.DB.Adapter == "sqlite" {
		conn = fmt.Sprintf("%v/%v", os.TempDir(), dbConfig.Name)
	}
	if err := iris_base_rabc.SetCasbinEnforcer(config.Config.DB.Adapter, conn); err != nil {
		color.Red(fmt.Sprintf("iris_base_rabc.SetCasbinEnforcer error: %v\n", err))
	}

	return conn
}
