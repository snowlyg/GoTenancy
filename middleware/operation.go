package middleware

import (
	"bytes"
	"net/http"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
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
		claims := ctx.Values().Get("claims")
		// if  {
		waitUse := claims.(*request.CustomClaims)
		userId = int(waitUse.ID)
		// } else {
		// 	id, err := strconv.Atoi(ctx.Request.Header.Get("x-user-id"))
		// 	if err != nil {
		// 		userId = 0
		// 	}
		// 	userId = id
		// }

		record := model.SysOperationRecord{

			Ip:     ctx.RemoteAddr(),
			Method: ctx.Method(),
			Path:   ctx.Path(),
			Agent:  ctx.Request().UserAgent(),
			Body:   string(body),
			UserID: userId,
		}
		// 存在某些未知错误 TODO
		//values := c.Request.Header.Values("content-type")
		//if len(values) >0 && strings.Contains(values[0], "boundary") {
		//	record.Body = "file"
		//}
		// writer := responseBodyWriter{
		// 	ResponseWriter: ctx.Writer,
		// 	body:           &bytes.Buffer{},
		// }
		// ctx.Writer = writer
		now := time.Now()

		ctx.Next()

		latency := time.Now().Sub(now)
		record.ErrorMessage = ctx.GetErr().Error()
		record.Status = ctx.GetStatusCode()
		record.Latency = latency
		// record.Resp = writer.body.String()

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
