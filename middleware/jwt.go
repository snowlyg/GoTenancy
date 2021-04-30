package middleware

import (
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
)

const issuer = "GOTENANCY"

func JWTAuth() iris.Handler {
	verifier := jwt.NewVerifier(jwt.HS256, g.TENANCY_CONFIG.JWT.SigningKey, jwt.Expected{Issuer: issuer})
	// Enable server-side token block feature (even before its expiration time):
	verifier.WithDefaultBlocklist()
	// Enable payload decryption with:
	if g.TENANCY_CONFIG.JWT.EncKey != "" {
		verifier.WithDecryption([]byte(g.TENANCY_CONFIG.JWT.EncKey), nil)
	}
	verifier.Extractors = []jwt.TokenExtractor{jwt.FromHeader} // extract token only from Authorization: Bearer $token
	return verifier.Verify(func() interface{} {
		return new(request.CustomClaims)
	})
}

// var (
// 	TokenExpired     = errors.New("Token is expired")
// 	TokenNotValidYet = errors.New("Token not active yet")
// 	TokenMalformed   = errors.New("That's not even a token")
// 	TokenInvalid     = errors.New("Couldn't handle this token:")
// )

// 创建一个token
func CreateToken(claims request.CustomClaims) (string, int64, error) {
	signer := jwt.NewSigner(jwt.HS256, g.TENANCY_CONFIG.JWT.SigningKey, 10*time.Minute)
	token, err := signer.Sign(claims, jwt.Claims{
		ID:     claims.UUID.String(),
		Issuer: issuer,
	})
	if g.TENANCY_CONFIG.JWT.EncKey != "" {
		signer.WithEncryption([]byte(g.TENANCY_CONFIG.JWT.EncKey), nil)
	}

	return string(token), signer.MaxAge.Milliseconds(), err
}

// // 解析 token
// func (j *JWT) ParseToken(tokenString string) (*request.CustomClaims, error) {
// 	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		if ve, ok := err.(*jwt.ValidationError); ok {
// 			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
// 				return nil, TokenMalformed
// 			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
// 				// Token is expired
// 				return nil, TokenExpired
// 			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
// 				return nil, TokenNotValidYet
// 			} else {
// 				return nil, TokenInvalid
// 			}
// 		}
// 	}
// 	if token != nil {
// 		if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
// 			return claims, nil
// 		}
// 		return nil, TokenInvalid

// 	} else {
// 		return nil, TokenInvalid

// 	}

// }

// // RefreshToken 更新token
// func (j *JWT) RefreshToken(tokenString string) (string, error) {
// 	jwt.TimeFunc = func() time.Time {
// 		return time.Unix(0, 0)
// 	}
// 	token, err := jwt.ParseWithClaims(tokenString, &request.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return j.SigningKey, nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	if claims, ok := token.Claims.(*request.CustomClaims); ok && token.Valid {
// 		jwt.TimeFunc = time.Now
// 		claims.StandardClaims.ExpiresAt = time.Now().Unix() + 60*60*24*7
// 		return j.CreateToken(*claims)
// 	}
// 	return "", TokenInvalid
// }
