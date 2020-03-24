package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
)

type HomeController struct {
	Ctx     iris.Context
	Session *sessions.Session
}

func (c *HomeController) Get() mvc.Result {
	c.Ctx.ViewLayout(iris.NoLayout)
	authUser := GetAuthUser(c.Session)
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"User": authUser,
		},
	}
}
