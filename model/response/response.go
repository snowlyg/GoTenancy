package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"status"`
	Data interface{} `json:"data"`
	Msg  string      `json:"message"`
}

const (
	SUCCESS                    = 200
	BAD_REQUEST_ERROR          = 4000
	UNAUTHORIZED_REQUEST_ERROR = 4001
	FORBIDDEN_REQUEST_ERROR    = 4003
)

func Result(code int, data interface{}, msg string, ctx *gin.Context) {
	// 开始时间
	ctx.JSON(http.StatusOK, Response{
		code,
		data,
		msg,
	})
}

func Ok(ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", ctx)
}

func OkWithMessage(message string, ctx *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, ctx)
}

func OkWithData(data interface{}, ctx *gin.Context) {
	Result(SUCCESS, data, "操作成功", ctx)
}

func OkWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(SUCCESS, data, message, ctx)
}

func Fail(ctx *gin.Context) {
	Result(BAD_REQUEST_ERROR, map[string]interface{}{}, "操作失败", ctx)
}

func UnauthorizedFailWithMessage(message string, ctx *gin.Context) {
	Result(UNAUTHORIZED_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func ForbiddenFailWithMessage(message string, ctx *gin.Context) {
	Result(FORBIDDEN_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithMessage(message string, ctx *gin.Context) {
	Result(BAD_REQUEST_ERROR, map[string]interface{}{}, message, ctx)
}

func FailWithDetailed(data interface{}, message string, ctx *gin.Context) {
	Result(BAD_REQUEST_ERROR, data, message, ctx)
}
