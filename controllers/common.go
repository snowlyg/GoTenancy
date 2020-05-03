package controllers

import (
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func GetAuthUser() models.User {
	authUser, _ := sysinit.UserService.GetByID(common.AuthUserId)

	return authUser
}
