package main

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/qor/admin"
	"github.com/qor/auth"
	"github.com/qor/auth/auth_identity"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/password"
	"github.com/qor/auth/providers/twitter"
	_ "github.com/qor/qor"
	"github.com/qor/session/manager"
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

func init() {
	// Migrate AuthIdentity model, AuthIdentity will be used to save auth info, like username/password, oauth token, you could change that.
	DB.AutoMigrate(&auth_identity.AuthIdentity{})

	// Register Auth providers
	// Allow use username/password
	Auth.RegisterProvider(password.New(&password.Config{}))

	// Allow use Github
	Auth.RegisterProvider(github.New(&github.Config{
		ClientID:     "github client id",
		ClientSecret: "github client secret",
	}))

	// Allow use Google
	Auth.RegisterProvider(google.New(&google.Config{
		ClientID:     "google client id",
		ClientSecret: "google client secret",
	}))

	// Allow use Facebook
	Auth.RegisterProvider(facebook.New(&facebook.Config{
		ClientID:     "facebook client id",
		ClientSecret: "facebook client secret",
	}))

	// Allow use Twitter
	Auth.RegisterProvider(twitter.New(&twitter.Config{
		ClientID:     "twitter client id",
		ClientSecret: "twitter client secret",
	}))
}

func main() {
	DB, err := gorm.Open("sqlite3", "demo.db")
	if err != nil {
		color.Red(fmt.Sprintf("数据库连接出错 ： %v", err))
	}
	DB.AutoMigrate(&User{}, &Product{})

	// Initialize Auth with configuration
	Auth = auth.New(&auth.Config{
		DB: gormDB,
	})

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
