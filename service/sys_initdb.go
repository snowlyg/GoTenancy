package service

import (
	"database/sql"
	"fmt"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/source"
	"github.com/snowlyg/go-tenancy/utils"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// WriteConfig 回写配置
func WriteConfig(viper *viper.Viper, mysql config.Mysql) error {
	g.TENANCY_CONFIG.Mysql = mysql
	cs := utils.StructToMap(g.TENANCY_CONFIG)
	for k, v := range cs {
		viper.Set(k, v)
	}
	return viper.WriteConfig()
}

// createTable 创建数据库(mysql)
func createTable(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}

func initDB(InitDBFunctions ...model.InitDBFunc) error {
	for _, v := range InitDBFunctions {
		err := v.Init()
		if err != nil {
			return err
		}
	}
	return nil
}

// InitDB 创建数据库并初始化
func InitDB(conf request.InitDB) error {
	BaseMysql := config.Mysql{
		Path:     "",
		Dbname:   "",
		Username: "",
		Password: "",
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}

	if conf.Port == "" {
		conf.Port = "3306"
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", conf.UserName, conf.Password, conf.Host, conf.Port)
	createSql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", conf.DBName)

	if err := createTable(dsn, "mysql", createSql); err != nil {
		return err
	}

	MysqlConfig := config.Mysql{
		Path:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Dbname:   conf.DBName,
		Username: conf.UserName,
		Password: conf.Password,
		Config:   "charset=utf8mb4&parseTime=True&loc=Local",
	}

	if err := WriteConfig(g.TENANCY_VP, MysqlConfig); err != nil {
		return err
	}
	m := g.TENANCY_CONFIG.Mysql
	if m.Dbname == "" {
		return nil
	}

	linkDns := m.Username + ":" + m.Password + "@tcp(" + m.Path + ")/" + m.Dbname + "?" + m.Config
	mysqlConfig := mysql.Config{
		DSN:                       linkDns, // DSN data source name
		DefaultStringSize:         191,     // string 类型字段的默认长度
		DisableDatetimePrecision:  true,    // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,    // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,    // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,   // 根据版本自动配置
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true}); err != nil {
		//g.TENANCY_LOG.Error("MySQL启动异常", zap.Any("err", err))
		//os.Exit(0)
		//return nil
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(m.MaxIdleConns)
		sqlDB.SetMaxOpenConns(m.MaxOpenConns)
		g.TENANCY_DB = db
	}

	err := g.TENANCY_DB.AutoMigrate(
		model.SysUser{},
		model.SysAdminInfo{},
		model.SysTenancyInfo{},
		model.SysGeneralInfo{},
		model.SysAuthority{},
		model.SysApi{},
		model.SysBaseMenu{},
		model.SysBaseMenuParameter{},
		model.SysDictionary{},
		model.SysDictionaryDetail{},
		model.SysOperationRecord{},
		model.SysTenancy{},
		model.SysRegion{},
	)
	if err != nil {
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
		return err
	}
	err = initDB(
		source.Admin,
		source.Api,
		source.AuthorityMenu,
		source.Authority,
		source.AuthoritiesMenus,
		source.Casbin,
		source.DataAuthorities,
		source.BaseMenu,
		source.Tenancy,
		source.Region,
	)
	if err != nil {
		_ = WriteConfig(g.TENANCY_VP, BaseMysql)
		return err
	}
	return nil
}
