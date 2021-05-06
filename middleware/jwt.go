package middleware

import (
	"errors"
	"fmt"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/service/sys_auth"
	"go.uber.org/zap"
)

const issuer = "GOTENANCY"

var (
	TokenExpired     = errors.New("Token is expired")
	TokenNotValidYet = errors.New("Token not active yet")
	TokenMalformed   = errors.New("That's not even a token")
	TokenInvalid     = errors.New("Couldn't handle this token:")
)

type tokenValidator struct {
}

func (v tokenValidator) ValidateToken(token []byte, claims jwt.Claims, err error) error {
	if err != nil {
		return err
	}
	g.TENANCY_LOG.Info("jwt token", zap.Any("", string(token)))
	authDriver := sys_auth.NewAuthDriver()
	// defer authDriver.Close()
	rcc, err := authDriver.GetCustomClaims(string(token))
	if err != nil {
		authDriver.DelUserTokenCache(string(token))
		return err
	}
	if rcc == nil {
		return TokenInvalid
	}
	return nil
}

func JWTAuth() iris.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, g.TENANCY_CONFIG.JWT.SigningKey, jwt.Expected{Issuer: issuer})
	// Enable payload decryption with:
	if g.TENANCY_CONFIG.JWT.EncKey != "" {
		verifier.WithDecryption([]byte(g.TENANCY_CONFIG.JWT.EncKey), nil)
	}
	verifier.Extractors = []jwt.TokenExtractor{jwt.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify(func() interface{} {
		return new(request.CustomClaims)
	}, tokenValidator{})
}

// CreateToken 创建token
func CreateToken(claims request.CustomClaims) (string, int64, error) {
	authDriver := sys_auth.NewAuthDriver()
	// defer authDriver.Close()

	if authDriver.IsUserTokenOver(claims.ID) {
		return "", 0, errors.New("以达到同时登录设备上限")
	}

	signer := jwt.NewSigner(jwt.HS256, g.TENANCY_CONFIG.JWT.SigningKey, 10*time.Minute)
	token, err := signer.Sign(claims, jwt.Claims{
		ID:     claims.UUID.String(),
		Issuer: issuer,
	})
	if err != nil {
		return "", 0, err
	}

	if g.TENANCY_CONFIG.JWT.EncKey != "" {
		signer.WithEncryption([]byte(g.TENANCY_CONFIG.JWT.EncKey), nil)
	}

	err = authDriver.ToCache(string(token), &claims)
	if err != nil {
		return "", 0, err
	}
	if err = authDriver.SyncUserTokenCache(string(token)); err != nil {
		return "", 0, err
	}

	return string(token), signer.MaxAge.Milliseconds(), err
}

// DelToken 删除token
func DelToken(token string) error {
	authDriver := sys_auth.NewAuthDriver()
	// defer authDriver.Close()
	err := authDriver.DelUserTokenCache(token)
	if err != nil {
		g.TENANCY_LOG.Error("del token", zap.Any("err", err))
		return fmt.Errorf("del token %w", err)
	}
	return nil
}

// CleanToken 清空 token
func CleanToken(userId string) error {
	authDriver := sys_auth.NewAuthDriver()
	// defer authDriver.Close()
	err := authDriver.CleanUserTokenCache(userId)
	if err != nil {
		g.TENANCY_LOG.Error("clean token", zap.Any("err", err))
		return fmt.Errorf("clean token %w", err)
	}
	return nil
}
