package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// GetMenu 获取用户动态路由
func GetMenu(ctx iris.Context) {
	if menus, err := service.GetMenuTree(getUserAuthorityId(ctx)); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.SysMenusResponse{Menus: menus}, "获取成功", ctx)
	}
}

// GetBaseMenuTree 获取用户动态路由
func GetBaseMenuTree(ctx iris.Context) {
	if menus, err := service.GetBaseMenuTree(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.SysBaseMenusResponse{Menus: menus}, "获取成功", ctx)
	}
}

// AddMenuAuthority 增加menu和角色关联关系
func AddMenuAuthority(ctx iris.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	if errs := utils.Verify(ctx.ReadJSON(&authorityMenu)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		g.TENANCY_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败", ctx)
	} else {
		response.OkWithMessage("添加成功", ctx)
	}
}

// GetMenuAuthority 获取指定角色menu
func GetMenuAuthority(ctx iris.Context) {
	var param request.GetAuthorityId
	if errs := utils.Verify(ctx.ReadJSON(&param)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err, menus := service.GetMenuAuthority(&param); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithDetailed(response.SysMenusResponse{Menus: menus}, "获取失败", ctx)
	} else {
		response.OkWithDetailed(iris.Map{"menus": menus}, "获取成功", ctx)
	}
}

// AddBaseMenu 新增菜单
func AddBaseMenu(ctx iris.Context) {
	var menu model.SysBaseMenu
	if errs := utils.Verify(ctx.ReadJSON(&menu)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.AddBaseMenu(menu); err != nil {
		g.TENANCY_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败", ctx)
	} else {
		response.OkWithMessage("添加成功", ctx)
	}
}

// DeleteBaseMenu 删除菜单
func DeleteBaseMenu(ctx iris.Context) {
	var menu request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&menu)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteBaseMenu(menu.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateBaseMenu 更新菜单
func UpdateBaseMenu(ctx iris.Context) {
	var menu model.SysBaseMenu
	if errs := utils.Verify(ctx.ReadJSON(&menu)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateBaseMenu(menu); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetBaseMenuById 根据id获取菜单
func GetBaseMenuById(ctx iris.Context) {
	var idInfo request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&idInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if menu, err := service.GetBaseMenuById(idInfo.Id); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.SysBaseMenuResponse{Menu: menu}, "获取成功", ctx)
	}
}

// GetMenuList 分页获取基础menu列表
func GetMenuList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if menuList, err := service.GetInfoList(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    0,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}
