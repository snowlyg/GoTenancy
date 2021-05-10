package service

import (
	"github.com/snowlyg/go-tenancy/config"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// GetSystemConfig 读取配置文件
func GetSystemConfig() config.Server {
	return g.TENANCY_CONFIG
}

// SetSystemConfig 设置配置文件
func SetSystemConfig(system model.System) error {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		g.TENANCY_VP.Set(k, v)
	}
	return g.TENANCY_VP.WriteConfig()
}

// GetServerInfo 获取服务器信息
func GetServerInfo() (*utils.Server, error) {
	var s utils.Server
	var err error
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		g.TENANCY_LOG.Error("func utils.InitCPU() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Rrm, err = utils.InitRAM(); err != nil {
		g.TENANCY_LOG.Error("func utils.InitRAM() Failed!", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		g.TENANCY_LOG.Error("func utils.InitDisk() Failed!", zap.String("err", err.Error()))
		return &s, err
	}

	return &s, nil
}
