package qiniu_test

import (
	"testing"

	"GoTenancy/libs/oss/qiniu"
	"GoTenancy/libs/oss/tests"
	"github.com/jinzhu/configor"
)

type Config struct {
	AccessID  string
	AccessKey string
	Region    string
	Bucket    string
	Endpoint  string
}

type AppConfig struct {
	Private Config
	Public  Config
}

var client *qiniu.Client
var privateClient *qiniu.Client

func init() {
	config := AppConfig{}
	configor.New(&configor.Config{ENVPrefix: "QINIU"}).Load(&config)
	if len(config.Private.AccessID) == 0 {
		return
	}

	client = qiniu.New(&qiniu.Config{
		AccessID:  config.Public.AccessID,
		AccessKey: config.Public.AccessKey,
		Region:    config.Public.Region,
		Bucket:    config.Public.Bucket,
		Endpoint:  config.Public.Endpoint,
	})
	privateClient = qiniu.New(&qiniu.Config{
		AccessID:   config.Private.AccessID,
		AccessKey:  config.Private.AccessKey,
		Region:     config.Private.Region,
		Bucket:     config.Private.Bucket,
		Endpoint:   config.Private.Endpoint,
		PrivateURL: true,
	})
}

func TestAll(t *testing.T) {
	if client == nil {
		t.Skip(`skip because of no config:


			`)
	}
	clis := []*qiniu.Client{client, privateClient}
	for _, cli := range clis {
		tests.TestAll(cli, t)
	}
}
