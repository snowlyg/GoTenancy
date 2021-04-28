// +build windows

package core

import (
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
)

func initServer(address string, router *iris.Application) server {
	return &http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
}
