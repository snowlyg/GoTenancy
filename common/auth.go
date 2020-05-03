package common

import (
	"github.com/snowlyg/go-authcode-1"
	"github.com/snowlyg/go-tenancy/lib"
	"strings"
)

var (
	AuthUserId       uint
	AuthUserTenantId uint
	AuthUserRoleIds  []uint
	UserCookieName   = "go_tenancy_user"
	AdminAuthKey     = "XXD45@##@afdsa"
)

// 从 cookie 中获取认证用户信息
func GetAuthInfo(authStr string) {
	authCode := authcode.AuthCode(authStr, "DECODE", AdminAuthKey, 0)
	authStrs := strings.Split(authCode, "||")

	userId := authStrs[0]
	userTenantId := authStrs[2]
	userRoleIds := authStrs[3]

	AuthUserId = lib.StringToUint(userId)
	AuthUserTenantId = lib.StringToUint(userTenantId)

	for _, item := range strings.Split(userRoleIds, userRoleIds) {
		id := lib.StringToUint(item)
		AuthUserRoleIds = append(AuthUserRoleIds, uint(id))
	}
}
