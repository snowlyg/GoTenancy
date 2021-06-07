package g

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/utils/timer"
)

const (
	StatusUnknown int = iota
	StatusTrue
	StatusFalse
)

var (
	TENANCY_DB     *gorm.DB
	TENANCY_CONFIG config.Server
	TENANCY_VP     *viper.Viper
	TENANCY_LOG    *zap.Logger
	TENANCY_Timer  timer.Timer = timer.NewTimerTask()
)
