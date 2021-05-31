package initialize

import (
	"context"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

func Auth() {
	redisCfg := g.TENANCY_CONFIG.Redis
	err := multi.InitDriver(&multi.Config{
		DriverType: g.TENANCY_CONFIG.System.CacheType,
		UniversalOptions: &redis.UniversalOptions{
			Addrs:       strings.Split(redisCfg.Addr, ","),
			Password:    redisCfg.Password,
			PoolSize:    10,
			IdleTimeout: 300 * time.Second,
			Dialer: func(ctx context.Context, network, addr string) (net.Conn, error) {
				conn, err := net.Dial(network, addr)
				if err == nil {
					go func() {
						time.Sleep(5 * time.Second)
						conn.Close()
					}()
				}
				return conn, err
			},
		}})
	if err != nil {
		g.TENANCY_LOG.Error("new auth diver err:", zap.Any("err", err))
	}

	if multi.AuthDriver == nil {
		g.TENANCY_LOG.Error("new auth diver failed")
		os.Exit(0)
	}
}
