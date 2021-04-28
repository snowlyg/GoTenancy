package service

import (
	"errors"
	"time"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

// JsonInBlacklist 拉黑jwt
func JsonInBlacklist(jwtList model.JwtBlacklist) (err error) {
	err = g.TENANCY_DB.Create(&jwtList).Error
	return
}

// IsBlacklist 判断JWT是否在黑名单内部
func IsBlacklist(jwt string) bool {
	isNotFound := errors.Is(g.TENANCY_DB.Where("jwt = ?", jwt).First(&model.JwtBlacklist{}).Error, gorm.ErrRecordNotFound)
	return !isNotFound
}

// GetRedisJWT 从redis取jwt
func GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = g.TENANCY_REDIS.Get(userName).Result()
	return err, redisJWT
}

// SetRedisJWT jwt存入redis并设置过期时间
func SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(g.TENANCY_CONFIG.JWT.ExpiresTime) * time.Second
	err = g.TENANCY_REDIS.Set(userName, jwt, timer).Err()
	return err
}
