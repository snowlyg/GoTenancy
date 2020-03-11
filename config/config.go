package config

import (
	"os"

	"github.com/jinzhu/configor"
	"github.com/qor/auth/providers/facebook"
	"github.com/qor/auth/providers/github"
	"github.com/qor/auth/providers/google"
	"github.com/qor/auth/providers/twitter"
	"github.com/qor/location"
	"github.com/qor/mailer"
	"github.com/qor/mailer/logger"
	"github.com/qor/redirect_back"
	"github.com/qor/session/manager"
	"github.com/unrolled/render"
)

type SMTPConfig struct {
	Host     string
	Port     string
	User     string
	Password string
}

var Config = struct {
	HTTPS bool `default:"false" env:"HTTPS"`
	Port  uint `default:"8080" env:"PORT"`
	DB    struct {
		Name     string `env:"DBName" default:"qor_example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser" default:"root"`
		Password string `env:"DBPassword"`
	}
	S3 struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
		Region          string `env:"AWS_Region"`
		S3Bucket        string `env:"AWS_Bucket"`
	}
	SMTP         SMTPConfig
	Github       github.Config
	Google       google.Config
	Facebook     facebook.Config
	Twitter      twitter.Config
	GoogleAPIKey string `env:"GoogleAPIKey"`
	BaiduAPIKey  string `env:"BaiduAPIKey"`
}{}

var (
	Root   = os.Getenv("GOPATH") + "/src/github.com/snowlyg/go-tenancy"
	Mailer *mailer.Mailer
	Render = render.New()

	RedirectBack = redirect_back.New(&redirect_back.Config{
		SessionManager:  manager.SessionManager,
		IgnoredPrefixes: []string{"/auth"},
	})
)

func init() {
	if err := configor.Load(&Config, "config/database.yml", "config/smtp.yml", "config/application.yml"); err != nil {
		panic(err)
	}

	location.GoogleAPIKey = Config.GoogleAPIKey
	location.BaiduAPIKey = Config.BaiduAPIKey

	Mailer = mailer.New(&mailer.Config{
		Sender: logger.New(&logger.Config{}),
	})
}
