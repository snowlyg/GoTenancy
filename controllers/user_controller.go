package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
)

type UserController struct {
	Ctx     iris.Context
	Service services.UserService
}

// GetUsers handles GET: http://localhost:8080/user/table.
func (c *UserController) GetTable() interface{} {
	args := map[string]interface{}{}
	users := sysinit.UserService.GetAll(args, false)

	return common.Table{Code: 0, Msg: "", Count: len(users), Data: users}
}

// Get handles GET: http://localhost:8080/user.
func (c *UserController) Get() mvc.Result {
	return mvc.View{
		Name: "user/index.html",
	}
}

// Get handles GET: http://localhost:8080/user/create.
func (c *UserController) GetCreate() mvc.Result {
	return mvc.View{
		Name: "user/add.html",
	}
}

// Get handles GET: http://localhost:8080/user/id.
func (c *UserController) GetBy(id uint) mvc.Result {
	user, _ := c.Service.GetByID(id)
	return mvc.View{
		Name: "user/edit.html",
		Data: iris.Map{
			"User": user,
		},
	}
}
