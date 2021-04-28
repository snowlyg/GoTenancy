package utils

import (
	"os"

	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

// IsExists
func IsExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// CreateDir
func CreateDir(dirs ...string) (err error) {
	for _, v := range dirs {
		exist, err := IsExists(v)
		if err != nil {
			return err
		}
		if !exist {
			g.TENANCY_LOG.Debug("create directory" + v)
			err = os.MkdirAll(v, os.ModePerm)
			if err != nil {
				g.TENANCY_LOG.Error("create directory"+v, zap.Any(" error:", err))
			}
		}
	}
	return err
}
