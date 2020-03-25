package sysinit

import (
	"errors"
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/services"
)

var (
	Db            *gorm.DB
	UserService   services.UserService
	PermService   services.PermService
	RoleService   services.RoleService
	TenantService services.TenantService
)

func init() {

	var err error
	var conn string
	if config.Config.DB.Adapter == "mysql" {
		conn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Port, config.Config.DB.Name)
	} else if config.Config.DB.Adapter == "postgres" {
		conn = fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", config.Config.DB.User, config.Config.DB.Password, config.Config.DB.Host, config.Config.DB.Name)
	} else if config.Config.DB.Adapter == "sqlite3" {
		conn = fmt.Sprintf("%v/%v", os.TempDir(), config.Config.DB.Name)
	} else {
		panic(errors.New("not supported database adapter"))
	}

	Db, err = gorm.Open(config.Config.DB.Adapter, conn)
	if err != nil {
		panic(err)
	}

	// 注册模型服务
	UserService = services.NewUserService(Db, Enforcer)
	PermService = services.NewPermService(Db)
	RoleService = services.NewRoleService(Db, Enforcer, PermService)
	TenantService = services.NewTenantService(Db, UserService, RoleService)

}
