package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/services"
)

type ControlController struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *ControlController) Get() mvc.Result {
	return mvc.View{
		Name: "control.html",
	}
}
