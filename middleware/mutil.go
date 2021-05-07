package middleware

import (
	"errors"
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

const issuer = "GOTENANCY"

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

func JWTAuth() iris.Handler {
	verifier := multi.NewVerifier()
	verifier.Extractors = []multi.TokenExtractor{multi.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify()
}

// CreateToken 创建token
func CreateToken(claims *multi.CustomClaims) (string, int64, error) {
	if g.TENANCY_AUTH.IsUserTokenOver(claims.ID) {
		return "", 0, errors.New("以达到同时登录设备上限")
	}
	token, err := multi.GetToken()
	if err != nil {
		return "", 0, err
	}
	err = g.TENANCY_AUTH.ToCache(token, claims)
	if err != nil {
		return "", 0, err
	}
	if err = g.TENANCY_AUTH.SyncUserTokenCache(token); err != nil {
		return "", 0, err
	}

	return token, int64(claims.ExpiresIn), err
}

// DelToken 删除token
func DelToken(token string) error {
	err := g.TENANCY_AUTH.DelUserTokenCache(token)
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		return fmt.Errorf("del token %w", err)
	}
	return nil
}

// CleanToken 清空 token
func CleanToken(userId string) error {
	err := g.TENANCY_AUTH.CleanUserTokenCache(userId)
	if err != nil {
		g.TENANCY_LOG.Error("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}
