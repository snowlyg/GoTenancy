package g

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/utils/timer"
)

var (
	TENANCY_DB     *gorm.DB
	TENANCY_REDIS  redis.UniversalClient
	TENANCY_CONFIG config.Server
	TENANCY_VP     *viper.Viper
	TENANCY_LOG    *zap.Logger
	TENANCY_Timer  timer.Timer = timer.NewTimerTask()
)
