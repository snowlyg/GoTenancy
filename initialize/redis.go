package initialize

import (
	"github.com/go-redis/redis"
	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

func Redis() {
	redisCfg := g.TENANCY_CONFIG.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping().Result()
	if err != nil {
		g.TENANCY_LOG.Error("redis connect ping failed, err:", zap.Any("err", err))
	} else {
		g.TENANCY_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		g.TENANCY_REDIS = client
	}
}
