package controllers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/services"
	"github.com/snowlyg/go-tenancy/sysinit"
	"github.com/snowlyg/go-tenancy/validatas"
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

// Get handles Post: http://localhost:8080/user.
// 使用 ReadJSON 获取数据前端数据需要格式化成json, JSON.stringify(data.field),
// 并且前端提交的数据格式需要是"application/x-www-form-urlencoded"，  dataType: "application/x-www-form-urlencoded",
func (c *UserController) Post() interface{} {

	var user models.User

	if err := c.Ctx.ReadJSON(&user); err != nil {
		return common.ActionResponse{Code: 0, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := validatas.Vaild(user); err != nil {
		return common.ActionResponse{Code: 0, Msg: fmt.Sprintf("数据验证错误：%v", err)}
	}

	if err := c.Service.Create(string(user.Password), &user); err != nil {
		return common.ActionResponse{Code: 0, Msg: fmt.Sprintf("用户创建错误：%v", err)}
	}

	return common.ActionResponse{Code: 1, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/user.
func (c *UserController) PostBy(id uint) interface{} {
	user, _ := c.Service.GetByID(id)
	return mvc.View{
		Name: "user/edit.html",
		Data: iris.Map{
			"User": user,
		},
	}
}
