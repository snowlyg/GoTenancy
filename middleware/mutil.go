package middleware

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/multi"
)

func JWTAuth() iris.Handler {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
