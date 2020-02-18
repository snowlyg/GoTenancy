package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
	_ "github.com/qor/qor"
)

// Define a GORM-backend model
type User struct {
	gorm.Model
	Name string
}

// Define another GORM-backend model
type Product struct {
	gorm.Model
	Name        string
	Description string
}

func main() {
	DB, err := gorm.Open("sqlite3", "demo.db")
	if err != nil {
		color.Red(fmt.Sprintf("数据库连接出错 ： %v", err))
	}
	DB.AutoMigrate(&User{}, &Product{})

	// Initalize
	Admin := admin.New(&admin.AdminConfig{DB: DB})

	// Create resources from GORM-backend model
	Admin.AddResource(&User{})
	Admin.AddResource(&Product{})

	mux := http.NewServeMux()
	Admin.MountTo("/admin", mux)
	http.ListenAndServe(":8080", mux)

	//app := iris.Default()
	//ser := &http.Server{Addr: ":8080", Handler: mux}
	//_ = app.Run(iris.Raw(ser.ListenAndServe))
}
