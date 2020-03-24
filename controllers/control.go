package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ControlController struct {
	Ctx iris.Context
}

func (c *ControlController) Get() mvc.Result {
	return mvc.View{
		Name: "control.html",
	}
}
