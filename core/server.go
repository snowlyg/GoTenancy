package core

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"

	"go.uber.org/zap"
)

func RunWindowsServer() {

	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", g.TENANCY_CONFIG.System.Addr)
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	g.TENANCY_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf("默认监听地址:http://127.0.0.1%s\n", address)
	err := Router.Listen(address)
	if err != nil {
		g.TENANCY_LOG.Error(err.Error())
	}
}
