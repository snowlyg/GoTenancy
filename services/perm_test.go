package services

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
	"github.com/snowlyg/go-tenancy/models"
)

var (
	DB *gorm.DB
)

func init() {
	var err error

	if DB, err = OpenTestConnection(); err != nil {
		panic(fmt.Sprintf("No error should happen when connecting to test database, but got err=%+v", err))
	}

	runMigration()
}

func runMigration() {
	values := []interface{}{&models.Perm{}, &models.Tenant{}, &models.Role{}, &models.User{}}
	for _, value := range values {
		DB.DropTable(value)
	}
	if err := DB.AutoMigrate(values...).Error; err != nil {
		panic(fmt.Sprintf("No error should happen when create table, but got %+v", err))
	}
}

func OpenTestConnection() (db *gorm.DB, err error) {
	dbDSN := os.Getenv("GORM_DSN")

	if dbDSN == "" {
		dbDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True"
	}
	db, err = gorm.Open("mysql", dbDSN)

	// db.SetLogger(Logger{log.New(os.Stdout, "\r\n", 0)})
	// db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	if debug := os.Getenv("DEBUG"); debug == "true" {
		db.LogMode(true)
	} else if debug == "false" {
		db.LogMode(false)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	return
}

func TestNewPermService(t *testing.T) {
	type args struct {
		gdb *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want PermService
	}{
		{
			name: "success",
			args: args{gdb: DB},
			want: NewPermService(DB),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPermService(tt.args.gdb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPermService() = %v, want %v", got, tt.want)
			}
		})
	}
}
