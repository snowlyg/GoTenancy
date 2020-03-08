//go:generate go run config/db/seeds/main.go config/db/seeds/seeds.go
//go:generate go run main.go -compile-templates=true
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/kataras/iris/v12"
	"go-tenancy/config"
	"go-tenancy/config/bindatafs"
	"go-tenancy/irisapp"
)

func main() {

	// 命令参数处理
	cmdLine := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)
	compileTemplate := cmdLine.Bool("compile-templates", false, "Compile Templates")
	if err := cmdLine.Parse(os.Args[1:]); err != nil {
		color.Red(fmt.Sprintf(" cmdLine.Parse error :%v", err))
	}

	// 初始化 iris
	irisApp := irisapp.New()

	if *compileTemplate { //处理前端静态文件
		if err := bindatafs.AssetFS.Compile(); err != nil {
			color.Red(fmt.Sprintf("bindatafs error %v", err))
		}
	} else {

		if config.Config.HTTPS {
			// 启动服务
			if err := irisApp.Run(iris.TLS(":443", "./config/local_certs/server.crt", "./config/local_certs/server.key")); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		} else {
			// iris 配置设置
			irisConfig := iris.WithConfiguration(
				iris.Configuration{
					DisableStartupLog:                 true,
					FireMethodNotAllowed:              true,
					DisableBodyConsumptionOnUnmarshal: true,
					TimeFormat:                        "Mon, 01 Jan 2006 15:04:05 GMT",
					Charset:                           "UTF-8",
				})
			// 启动服务
			if err := irisApp.Run(iris.Addr(fmt.Sprintf(":%d", config.Config.Port)), iris.WithoutServerError(iris.ErrServerClosed), irisConfig); err != nil {
				color.Red(fmt.Sprintf("app.Listen %v", err))
				panic(err)
			}
		}
	}
}
