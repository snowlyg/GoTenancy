package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/multi"
)

func Auth() gin.HandlerFunc {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}
