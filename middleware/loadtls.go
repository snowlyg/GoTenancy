package middleware

import (
	"github.com/kataras/iris/v12"
)

// 用https把这个中间件在router里面use一下就好

func LoadTls() iris.Handler {
	return func(ctx iris.Context) {
		// 	middleware := secure.New(secure.Options{
		// 		SSLRedirect: true,
		// 		SSLHost:     "localhost:443",
		// 	})
		// 	err := middleware.Process(ctx.Writer, ctx.Request)
		// 	if err != nil {
		// 		// 如果出现错误，请不要继续
		// 		fmt.Println(err)
		// 		return
		// 	}
		// 	// 继续往下处理
		ctx.Next()
	}
}
