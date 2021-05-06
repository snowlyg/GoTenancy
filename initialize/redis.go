package initialize

import (
	"context"
	"net"
	"os"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := g.TENANCY_CONFIG.Redis
	client := redis.NewUniversalClient(&redis.UniversalOptions{
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
	})
	ping, err := client.Ping(context.Background()).Result()
	if err != nil {
		g.TENANCY_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
		os.Exit(0)
	} else {
		g.TENANCY_LOG.Info("redis connect ping response:", zap.String("ping", ping))
		g.TENANCY_REDIS = client
	}
}
