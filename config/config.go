package config

import (
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

var Config = struct {
	HTTPS bool   `default:"false" env:"HTTPS"`
	Port  uint   `default:"7000" env:"PORT"`
	Host  string `default:"localhost" env:"Host"`
	DB    struct {
		Name     string `env:"DBName" default:"qor_example"`
		Adapter  string `env:"DBAdapter" default:"mysql"`
		Host     string `env:"DBHost" default:"localhost"`
		Port     string `env:"DBPort" default:"3306"`
		User     string `env:"DBUser" default:"root"`
		Password string `env:"DBPassword"`
	}
}{}

var Root = os.Getenv("GOPATH") + "/src/github.com/snowlyg/go-tenancy"

func init() {
	if err := configor.Load(&Config, filepath.Join(Root, "config/application.yml")); err != nil {
		panic(err)
	}
}