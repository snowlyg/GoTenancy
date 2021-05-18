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
	SUCCESS                    = 0
	BAD_REQUEST_ERROR          = 4000
	UNAUTHORIZED_REQUEST_ERROR = 4001
	FORBIDDEN_REQUEST_ERROR    = 4003
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
	Result(BAD_REQUEST_ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func UnauthorizedFailWithMessage(message string, ctx iris.Context) {
	Result(UNAUTHORIZED_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func ForbiddenFailWithMessage(message string, ctx iris.Context) {
	Result(FORBIDDEN_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithMessage(message string, ctx iris.Context) {
	Result(BAD_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx iris.Context) {
	Result(BAD_REQUEST_ERROR, data, message, ctx)
}
