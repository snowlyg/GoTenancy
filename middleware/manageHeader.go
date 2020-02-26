package middleware

import "github.com/kataras/iris/v12"

func AddHeader(ctx iris.Context) {
	// 演示设置，请勿在生产环境使用
	//ctx.Header("Access-Control-Allow-Origin", "*")
	//ctx.Request().Header.Del("Authorization")
	ctx.Next()
}
