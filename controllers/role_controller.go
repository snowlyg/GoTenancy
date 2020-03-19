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

type RoleController struct {
	Ctx     iris.Context
	Service services.RoleService
}

// GetRoles handles GET: http://localhost:8080/role/table.
func (c *RoleController) GetTable() interface{} {

	var pagination common.Pagination
	if err := c.Ctx.ReadQuery(&pagination); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("分页参数获取错误：%v", err)}
	}

	args := map[string]interface{}{}
	count, roles := sysinit.RoleService.GetAll(args, &pagination, false)

	return common.Table{Code: 0, Msg: "", Count: count, Data: roles}
}

// Get handles GET: http://localhost:8080/role.
func (c *RoleController) Get() mvc.Result {
	return mvc.View{
		Name: "role/index.html",
	}
}

// Get handles GET: http://localhost:8080/role/create.
func (c *RoleController) GetCreate() mvc.Result {
	return mvc.View{
		Name: "role/add.html",
	}
}

// Get handles GET: http://localhost:8080/role/id.
func (c *RoleController) GetBy(id uint) mvc.Result {
	role, _ := c.Service.GetByID(id)
	return mvc.View{
		Name: "role/edit.html",
		Data: iris.Map{
			"Role": role,
		},
	}
}

// Get handles Post: http://localhost:8080/role.
// 使用 ReadJSON 获取数据前端数据需要格式化成json, JSON.stringify(data.field),
func (c *RoleController) Post() interface{} {

	var role models.Role

	if err := c.Ctx.ReadJSON(&role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := validatas.Vaild(role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据验证错误：%v", err)}
	}

	if err := c.Service.Create(&role, []uint{}); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色创建错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/role/id.
func (c *RoleController) PostBy(id uint) interface{} {

	var role models.Role

	if err := c.Ctx.ReadJSON(&role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := validatas.Vaild(role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据验证错误：%v", err)}
	}

	if err := c.Service.Update(id, &role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色更新错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/role/id.
func (c *RoleController) DeleteBy(id uint) interface{} {
	if err := c.Service.DeleteByID(id); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色删除错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/user/deletes.
func (c *RoleController) PostDeletes() interface{} {
	var ids []common.Id

	if err := c.Ctx.ReadJSON(&ids); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := c.Service.DeleteMnutil(ids); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色删除错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}
