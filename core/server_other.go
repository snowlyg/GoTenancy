// +build !windows

package core

import (
	"time"

	"github.com/fvbock/endless"
)

func initServer(address string, router *iris.Engine) server {
	s := endless.NewServer(address, router)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20
	return s
}
