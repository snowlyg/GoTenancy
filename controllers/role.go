package controllers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
	"github.com/snowlyg/go-tenancy/transformer"
	"github.com/snowlyg/go-tenancy/validatas"
)

type RoleController struct {
	Ctx iris.Context
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

// GetRoletBy handles GET: http://localhost:8080/role/perm/:id.
func (c *RoleController) GetPermBy(id uint) interface{} {

	role, _ := sysinit.RoleService.GetByID(id)

	args := map[string]interface{}{}
	_, perms := sysinit.PermService.GetAll(args, false)

	return common.ActionResponse{Status: true, Msg: "", Data: c.transformerSelectPerms(perms, role.ID)}
}

// Get handles GET: http://localhost:8080/role/id.
func (c *RoleController) GetBy(id uint) mvc.Result {
	role, _ := sysinit.RoleService.GetByID(id)
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

	if err := sysinit.RoleService.Create(&role); err != nil {
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

	if err := sysinit.RoleService.UpdateRole(id, &role); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色更新错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/role/id.
func (c *RoleController) DeleteBy(id uint) interface{} {
	if err := sysinit.RoleService.DeleteByID(id); err != nil {
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

	if err := sysinit.RoleService.DeleteMnutil(ids); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("角色删除错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// transformerSelectPerms 角色权限下拉接口数据转换
func (c *RoleController) transformerSelectPerms(perms []*models.Perm, roleId uint) *transformer.PermSelect {

	tablePerm := &transformer.PermSelect{}
	var checkedId []uint
	var list []transformer.List

	for _, perm := range perms {
		var ll transformer.List
		ll.Id = perm.ID
		ll.Name = perm.Title
		if len(perm.Href) > 0 {
			ll.Name += " : " + perm.Href
		}
		ll.Pid = perm.ParentId.Int64

		rolePerms, err := sysinit.RoleService.GetPermsByID(roleId)
		if err == nil {
			for _, rolePerm := range rolePerms {
				if rolePerm.ID == perm.ID {
					checkedId = append(checkedId, perm.ID)
				}
			}
		}

		list = append(list, ll)
	}

	tablePerm.CheckedId = checkedId
	tablePerm.List = list

	return tablePerm
}
