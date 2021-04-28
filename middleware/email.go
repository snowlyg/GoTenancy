package middleware

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// ErrorToEmail
func ErrorToEmail() iris.Handler {
	return func(ctx iris.Context) {
		var username string
		claims := ctx.Values().Get("claims")
		// if ok {
		waitUse := claims.(*request.CustomClaims)
		username = waitUse.Username
		// } else {
		// 	id, _ := strconv.Atoi(ctx.GetHeader("x-user-id"))
		// 	err, user := service.FindUserById(id)
		// 	if err != nil {
		// 		username = "Unknown"
		// 	}
		// 	username = user.Username
		// }
		body, _ := ctx.GetBody()
		record := model.SysOperationRecord{
			Ip:     ctx.RemoteAddr(),
			Method: ctx.Method(),
			Path:   ctx.Path(),
			Agent:  ctx.Request().UserAgent(),
			Body:   string(body),
		}
		now := time.Now()

		ctx.Next()

		latency := time.Now().Sub(now)
		status := ctx.GetStatusCode()
		// record.ErrorMessage = ctx.Errors.ByType(iris.ErrorTypePrivate).String()
		record.ErrorMessage = ctx.GetErr().Error()
		str := "接收到的请求为" + record.Body + "\n" + "请求方式为" + record.Method + "\n" + "报错信息如下" + record.ErrorMessage + "\n" + "耗时" + latency.String() + "\n"
		if status != 200 {
			subject := username + "" + record.Ip + "调用了" + record.Path + "报错了"
			if err := utils.ErrorToEmail(subject, str); err != nil {
				g.TENANCY_LOG.Error("ErrorToEmail Failed, err:", zap.Any("err", err))
			}
		}
	}
}
