package middleware

import (
	"bytes"
	"net/http"
	"strconv"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

func OperationRecord() iris.Handler {
	return func(ctx iris.Context) {
		var body []byte
		var userId int
		if ctx.Method() != http.MethodGet {
			var err error
			body, err = ctx.GetBody()
			if err != nil {
				g.TENANCY_LOG.Error("read body from request error:", zap.Any("err", err))
			} else {
				ctx.Recorder().SetBody(body)
			}
		}
		claims := jwt.Get(ctx).(*request.CustomClaims)
		if claims != nil {
			id, err := strconv.Atoi(claims.ID)
			if err != nil {
				userId = 0
			}
			userId = id
		} else {
			id, err := strconv.Atoi(ctx.GetHeader("X-USER-ID"))
			if err != nil {
				userId = 0
			}
			userId = id
		}

		record := model.SysOperationRecord{
			Ip:     ctx.RemoteAddr(),
			Method: ctx.Method(),
			Path:   ctx.Path(),
			Agent:  ctx.Request().UserAgent(),
			Body:   string(body),
			UserID: userId,
		}

		writer := responseBodyWriter{
			ResponseWriter: ctx.ResponseWriter().Clone(),
			body:           &bytes.Buffer{},
		}
		ctx.ResetResponseWriter(writer)
		now := time.Now()

		ctx.Next()

		latency := time.Since(now)
		errorMessage := ""
		if ctx.GetErr() != nil {
			errorMessage = ctx.GetErr().Error()
		}
		record.ErrorMessage = errorMessage
		record.Status = ctx.GetStatusCode()
		record.Latency = latency
		record.Resp = writer.body.String()

		if err := service.CreateSysOperationRecord(record); err != nil {
			g.TENANCY_LOG.Error("create operation record error:", zap.Any("err", err))
		}
	}
}

type responseBodyWriter struct {
	context.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}
