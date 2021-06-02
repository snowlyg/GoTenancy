package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/multi"
)

func Auth() gin.HandlerFunc {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader} // extract token only from Authorization: Bearer $token
	verifier.ErrorHandler = func(ctx *gin.Context, err error) {
		response.UnauthorizedFailWithMessage(err.Error(), ctx)
		ctx.Abort()
	} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
