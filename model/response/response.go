package response

import (
	"github.com/kataras/iris/v12"
)

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, ctx iris.Context) {
	// 开始时间
	ctx.JSON(Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx iris.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}

func OkWithMessage(message string, ctx iris.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx iris.Context) {
	Result(SUCCESS, data, "操作成功", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx iris.Context) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx iris.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func FailWithMessage(message string, ctx iris.Context) {
	Result(ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx iris.Context) {
	Result(ERROR, data, message, ctx)
}
