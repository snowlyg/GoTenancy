package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"go.uber.org/zap"
)

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery(stack bool) iris.Handler {
	return func(ctx iris.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(ctx.Request(), false)
				if brokenPipe {
					g.TENANCY_LOG.Error(ctx.Path(),
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					ctx.Writef(err.(error).Error()) // nolint: errcheck
					ctx.StatusCode(http.StatusInternalServerError)
					return
				}

				if stack {
					g.TENANCY_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					g.TENANCY_LOG.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				ctx.StatusCode(http.StatusInternalServerError)
			}
		}()
		ctx.Next()
	}
}
