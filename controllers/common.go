package controllers

import (
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func GetAuthUser(sess *sessions.Session) models.User {
	userID := sess.GetInt64Default(sysinit.UserIDKey, 0)
	authUser, _ := sysinit.UserService.GetByID(uint(userID))

	return authUser
}
