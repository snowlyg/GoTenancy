package g

import (
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/snowlyg/go-tenancy/config"
)

var (
	TENANCY_DB     *gorm.DB
	TENANCY_REDIS  *redis.Client
	TENANCY_CONFIG config.Server
	TENANCY_VP     *viper.Viper
	TENANCY_LOG    *zap.Logger
	// TENANCY_Timer timer.Timer = timer.NewTimerTask()
)
