package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/services"
)

type UserController struct {
	Ctx     iris.Context
	Service services.PermService
}

// Get handles GET: http://localhost:8080/user.
func (c *UserController) Get() mvc.Result {
	return mvc.View{
		Name: "user/index.html",
	}
}
