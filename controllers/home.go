package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type HomeController struct {
	Ctx iris.Context
}

func (c *HomeController) Get() mvc.Result {
	c.Ctx.ViewLayout(iris.NoLayout)
	authUser := GetAuthUser()
	return mvc.View{
		Name: "index.html",
		Data: iris.Map{
			"User": authUser,
		},
	}
}
