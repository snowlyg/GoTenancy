package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"github.com/snowlyg/go-tenancy/services"
)

type Controlcontroller struct {
	Ctx     iris.Context
	Service services.UserService
	Session *sessions.Session
}

func (c *Controlcontroller) Get() mvc.Result {
	c.Ctx.ViewLayout(iris.NoLayout)
	return mvc.View{
		Name: "control.html",
	}
}
