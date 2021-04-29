package core

import (
	stdContext "context"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/initialize"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	if g.TENANCY_CONFIG.System.UseMultipoint {
		// 初始化redis服务
		initialize.Redis()
	}
	Router := initialize.Routers()
	// Router.Static("/form-generator", "./resource/page")
	address := fmt.Sprintf(":%d", g.TENANCY_CONFIG.System.Addr)
	idleConnsClosed := make(chan struct{})
	iris.RegisterOnInterrupt(func() {
		timeout := 10 * time.Second
		ctx, cancel := stdContext.WithTimeout(stdContext.Background(), timeout)
		defer cancel()
		// close all hosts.
		Router.Shutdown(ctx)
		close(idleConnsClosed)
	})
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	g.TENANCY_LOG.Info("server run success on ", zap.String("address", address))

	fmt.Printf("默认监听地址:http://127.0.0.1%s\n", address)
	err := s.ListenAndServe().Error()
	g.TENANCY_LOG.Error(err)
	<-idleConnsClosed
}
