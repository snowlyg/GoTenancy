// +build ignore

package main

import (
	"fmt"
	"time"

	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func CreateAdminUsers() {
	admin := &models.User{
		Username:  "username",
		Firstname: "超级管理员",
		CreatedAt: time.Now(),
	}

	if err := sysinit.UserService.Create("password", admin); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}

}
