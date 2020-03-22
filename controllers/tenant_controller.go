package controllers

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/snowlyg/go-tenancy/common"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/services"

	"github.com/snowlyg/go-tenancy/validatas"
)

type TenantController struct {
	Ctx         iris.Context
	Service     services.TenantService
	UserService services.UserService
}

// GetTenants handles GET: http://localhost:8080/tenant/table.
func (c *TenantController) GetTable() interface{} {

	var pagination common.Pagination
	if err := c.Ctx.ReadQuery(&pagination); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("分页参数获取错误：%v", err)}
	}

	args := map[string]interface{}{}
	count, tenants := c.Service.GetAll(args, &pagination, false)

	return common.Table{Code: 0, Msg: "", Count: count, Data: tenants}
}

// Get handles GET: http://localhost:8080/tenant.
func (c *TenantController) Get() mvc.Result {
	return mvc.View{
		Name: "tenant/index.html",
	}
}

// Get handles GET: http://localhost:8080/tenant/create.
func (c *TenantController) GetCreate() mvc.Result {
	return mvc.View{
		Name: "tenant/add.html",
	}
}

// Get handles GET: http://localhost:8080/tenant/id.
func (c *TenantController) GetBy(id uint) mvc.Result {
	tenant, _ := c.Service.GetByID(id)

	return mvc.View{
		Name: "tenant/edit.html",
		Data: iris.Map{
			"Tenant": tenant,
		},
	}
}

// Get handles Post: http://localhost:8080/tenant.
// 使用 ReadJSON 获取数据前端数据需要格式化成json, JSON.stringify(data.field),
func (c *TenantController) Post() interface{} {

	var tenant models.Tenant

	if err := c.Ctx.ReadJSON(&tenant); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := validatas.Vaild(tenant); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据验证错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/tenant/id.
func (c *TenantController) PostBy(id uint) interface{} {

	var tenant models.Tenant

	if err := c.Ctx.ReadJSON(&tenant); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := validatas.Vaild(tenant); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据验证错误：%v", err)}
	}

	if err := c.Service.Update(id, &tenant); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("用户更新错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/tenant/id.
func (c *TenantController) DeleteBy(id uint) interface{} {
	if err := c.Service.DeleteByID(id); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("用户删除错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}

// Get handles Post: http://localhost:8080/tenant/deletes.
func (c *TenantController) PostDeletes() interface{} {
	var tenantIds []common.Id

	if err := c.Ctx.ReadJSON(&tenantIds); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("数据获取错误：%v", err)}
	}

	if err := c.Service.DeleteMnutil(tenantIds); err != nil {
		return common.ActionResponse{Status: false, Msg: fmt.Sprintf("用户删除错误：%v", err)}
	}

	return common.ActionResponse{Status: true, Msg: "操作成功"}
}
