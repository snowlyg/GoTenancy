package sys_auth

import (
	"context"
	"errors"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
)

const (
	GtSessionTokenPrefix        = "GST:"           // token 缓存前缀
	GtSessionBindUserPrefix     = "GSBU:"          // token 绑定用户前缀
	GtSessionUserPrefix         = "GSU:"           // 用户前缀
	GtSessionUserMaxTokenPrefix = "GTUserMaxToken" // 用户最大 token 数前缀
)

var (
	ctx                                = context.Background()
	ErrTokenInvalid                    = errors.New("token is invalid")
	GtSessionUserMaxTokenDefault int64 = 10
)

const (
	NoneScope uint64 = iota
	AdminScope
)

const (
	NoAuth int = iota
	AuthPwd
	AuthCode
	AuthThirdParty
)

const (
	LoginTypeWeb int = iota
	LoginTypeApp
	LoginTypeWx
	LoginTypeAlipay
	LoginApplet
)

var (
	RedisSessionTimeoutWeb    = 30 * time.Minute
	RedisSessionTimeoutApp    = 24 * time.Hour
	RedisSessionTimeoutWx     = 5 * 52 * 168 * time.Hour
	RedisSessionTimeoutApplet = 7 * 24 * time.Hour
)

var authDriver Authentication

// NewAuthDriver 认证驱动
// redis 需要设置redis
// local 使用本地内存
func NewAuthDriver() Authentication {
	switch g.TENANCY_CONFIG.System.CacheType {
	case "redis":
		return NewRedisAuth()
	case "local":
		return NewLocalAuth()
	default:
		return NewLocalAuth()
	}
}

// Authentication  认证
type Authentication interface {
	ToCache(token string, rcc *request.CustomClaims) error
	SyncUserTokenCache(token string) error
	DelUserTokenCache(token string) error
	UserTokenExpired(token string) error
	UpdateUserTokenCacheExpire(token string) error
	GetCustomClaims(token string) (*request.CustomClaims, error)
	GetAuthId(token string) (uint, error)
	IsUserTokenOver(userId string) bool
	CleanUserTokenCache(userId string) error
	Close()
}
