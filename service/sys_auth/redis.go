package sys_auth

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model/request"
	"go.uber.org/zap"
)

// RedisAuth
type RedisAuth struct {
	Client redis.UniversalClient
}

// NewRedisAuth
func NewRedisAuth() *RedisAuth {
	return &RedisAuth{
		Client: g.TENANCY_REDIS,
	}
}

// ToCache 缓存 token
func (ra *RedisAuth) ToCache(token string, rcc *request.CustomClaims) error {
	sKey := GtSessionTokenPrefix + token
	if _, err := ra.Client.HMSet(ctx, sKey,
		"user_id", rcc.GetID(),
		"login_type", rcc.LoginType,
		"auth_type", rcc.AuthType,
		"username", rcc.Username,
		"nickname", rcc.Nickname,
		"authority_id", rcc.AuthorityId,
		"creation_data", rcc.CreationDate,
		"expires_in", rcc.ExpiresIn,
		// "scope", rsv2.Scope,
	).Result(); err != nil {
		g.TENANCY_LOG.Error("to cache token", zap.Any("err", err))
		return fmt.Errorf("to cache token %w", err)
	}

	return nil
}

// GetAuthId
func (ra *RedisAuth) GetAuthId(token string) (uint, error) {
	sess, err := ra.GetCustomClaims(token)
	if err != nil {
		return 0, fmt.Errorf("get auth id %w", err)
	}
	id, err := strconv.ParseInt(sess.GetID(), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("get auth id parse int %w", err)
	}
	return uint(id), nil
}

//  GetCustomClaims session
func (ra *RedisAuth) GetCustomClaims(token string) (*request.CustomClaims, error) {
	sKey := GtSessionTokenPrefix + token

	_, err := ra.Client.Exists(context.Background(), sKey).Result()
	if err != nil {
		g.TENANCY_LOG.Error("get custom claims", zap.Error(err))
		return nil, ErrTokenInvalid
	}
	pp := new(request.CustomClaims)
	if err := ra.Client.HGetAll(ctx, sKey).Scan(pp); err != nil {
		return nil, fmt.Errorf("get custom claims redis hgetall %w", err)
	}
	return pp, nil
}

// IsUserTokenOver 超过登录设备限制
func (ra *RedisAuth) IsUserTokenOver(userId string) bool {
	g.TENANCY_LOG.Debug("user token count ", zap.Any("", ra.getUserTokenCount(userId)), zap.Any(" user max count ", ra.getUserTokenMaxCount()))
	return ra.getUserTokenCount(userId) >= ra.getUserTokenMaxCount()
}

// getUserTokenCount 获取登录数量
func (ra *RedisAuth) getUserTokenCount(userId string) int64 {
	count, err := ra.Client.SCard(ctx, GtSessionUserPrefix+userId).Result()
	if err != nil {
		g.TENANCY_LOG.Error("get user token count", zap.Any("err", err))
		return 0
	}
	return count
}

// getUserTokenMaxCount 最大登录限制
func (ra *RedisAuth) getUserTokenMaxCount() int64 {
	count, err := ra.Client.Get(ctx, GtSessionUserMaxTokenPrefix).Int64()
	if err != nil {
		return GtSessionUserMaxTokenDefault
	}
	return count
}

// UserTokenExpired 过期 token
func (ra *RedisAuth) UserTokenExpired(token string) error {
	uKey := GtSessionBindUserPrefix + token
	sKeys, err := ra.Client.SMembers(ctx, uKey).Result()
	if err != nil {
		g.TENANCY_LOG.Error("user token expired get members", zap.Any("err", err))
		return fmt.Errorf("user token expired %w", err)
	}
	for _, v := range sKeys {
		if !strings.Contains(v, GtSessionUserPrefix) {
			continue
		}
		_, err = ra.Client.SRem(ctx, v, token).Result()
		if err != nil {
			g.TENANCY_LOG.Error("user token expired do srem", zap.Any("err", err))
			continue
		}
	}
	if _, err = ra.Client.Del(ctx, uKey).Result(); err != nil {
		g.TENANCY_LOG.Error("user token expired del", zap.Any("err", err))
		return err
	}
	return nil
}

// GetUserScope 角色
func GetUserScope(userType string) uint64 {
	switch userType {
	case "admin":
		return AdminScope
	}
	return NoneScope
}

// SyncUserTokenCache 同步 token 到用户缓存
func (ra *RedisAuth) SyncUserTokenCache(token string) error {
	rcc, err := ra.GetCustomClaims(token)
	if err != nil {
		return fmt.Errorf("sysnc user token cache %w", err)
	}
	sKey := GtSessionUserPrefix + rcc.GetID()
	if _, err := ra.Client.SAdd(ctx, sKey, token).Result(); err != nil {
		g.TENANCY_LOG.Error("sync user token cache sadd", zap.Any("err", err))
		return fmt.Errorf("sync user token cache redis sadd %w", err)
	}
	sKey2 := GtSessionBindUserPrefix + token
	_, err = ra.Client.SAdd(ctx, sKey2, sKey).Result()
	if err != nil {
		g.TENANCY_LOG.Error("sync user token cache sadd", zap.Any("err", err))
		return fmt.Errorf("sync user token cache %w", err)
	}
	return nil
}

//UpdateUserTokenCacheExpire 更新过期时间
func (ra *RedisAuth) UpdateUserTokenCacheExpire(token string) error {
	rcc, err := ra.GetCustomClaims(token)
	if err != nil {
		return fmt.Errorf("update user token cache expire %w", err)
	}
	if rcc == nil {
		return errors.New("token cache is nil")
	}
	if _, err = ra.Client.Expire(ctx, GtSessionTokenPrefix+token, ra.getTokenExpire(rcc)).Result(); err != nil {
		g.TENANCY_LOG.Error("update user token cache expire", zap.Any("err", err))
		return fmt.Errorf("update user token cache expire redis expire %w", err)
	}
	return nil
}

// getTokenExpire 过期时间
func (ra *RedisAuth) getTokenExpire(rsv2 *request.CustomClaims) time.Duration {
	timeout := RedisSessionTimeoutApp
	if rsv2.LoginType == LoginTypeWeb {
		timeout = RedisSessionTimeoutWeb
	} else if rsv2.LoginType == LoginTypeWx {
		timeout = RedisSessionTimeoutWx
	} else if rsv2.LoginType == LoginTypeAlipay {
		timeout = RedisSessionTimeoutWx
	}
	return timeout
}

// DelUserTokenCache 删除token缓存
func (ra *RedisAuth) DelUserTokenCache(token string) error {
	rcc, err := ra.GetCustomClaims(token)
	if err != nil {
		return err
	}
	if rcc == nil {
		return errors.New("del user token, reids cache is nil")
	}
	sKey := GtSessionUserPrefix + rcc.GetID()
	_, err = ra.Client.SRem(ctx, sKey, token).Result()
	if err != nil {
		g.TENANCY_LOG.Error("del user token cache do srem", zap.Any("err", err))
		return fmt.Errorf("del user token cache redis srem %w", err)
	}
	err = ra.DelTokenCache(token)
	if err != nil {
		return err
	}

	return nil
}

// DelTokenCache 删除token缓存
func (ra *RedisAuth) DelTokenCache(token string) error {
	sKey2 := GtSessionBindUserPrefix + token
	_, err := ra.Client.Del(ctx, sKey2).Result()
	if err != nil {
		g.TENANCY_LOG.Error("del token cache del key2", zap.Any("err", err))
		return fmt.Errorf("del user token cache redis del2  %w", err)
	}

	sKey3 := GtSessionTokenPrefix + token
	_, err = ra.Client.Del(ctx, sKey3).Result()
	if err != nil {
		g.TENANCY_LOG.Error("del token cache del key3", zap.Any("err", err))
		return fmt.Errorf("del user token cache redis del3  %w", err)
	}

	return nil
}

// CleanUserTokenCache 清空token缓存
func (ra *RedisAuth) CleanUserTokenCache(userId string) error {
	sKey := GtSessionUserPrefix + userId
	var allTokens []string
	allTokens, err := ra.Client.SMembers(ctx, sKey).Result()
	if err != nil {
		g.TENANCY_LOG.Error("clean user token cache member", zap.Any("err", err))
		return fmt.Errorf("clean user token cache redis smembers  %w", err)
	}
	_, err = ra.Client.Del(ctx, sKey).Result()
	if err != nil {
		g.TENANCY_LOG.Error("clean user token cache del", zap.Any("err", err))
		return fmt.Errorf("clean user token cache redis del  %w", err)
	}

	for _, token := range allTokens {
		err = ra.DelTokenCache(token)
		if err != nil {
			return err
		}
	}
	return nil
}

// Close
func (ra *RedisAuth) Close() {
	ra.Client.Close()
}
