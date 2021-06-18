package middleware

import (
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// ErrorToEmail
func ErrorToEmail() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var username string
		waitUse := multi.Get(ctx)
		if waitUse != nil {
			username = waitUse.Username
		} else {
			user, err := service.FindUserById(ctx.GetHeader("X-USER-ID"))
			if err != nil {
				username = "Unknown"
			}
			username = user.Username
		}
		body, _ := ioutil.ReadAll(ctx.Request.Body)
		record := model.SysOperationRecord{
			BaseOperationRecord: model.BaseOperationRecord{
				Ip:     ctx.ClientIP(),
				Method: ctx.Request.Method,
				Path:   ctx.Request.URL.Path,
				Agent:  ctx.Request.UserAgent(),
				Body:   string(body),
			},
		}
		now := time.Now()

		ctx.Next()

		latency := time.Now().Sub(now)
		status := ctx.Writer.Status()
		record.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				g.TENANCY_LOG.Error("ErrorToEmail Failed, err:", zap.Any("err", err))
			}
		}
	}
}
