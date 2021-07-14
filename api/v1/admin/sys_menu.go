package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"go.uber.org/zap"
)

// GetAddMenuMap 添加表单
func GetAddMenuMap(ctx *gin.Context) {
	if menus, err := service.GetMenuMap(0, ctx, false); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}

// GetAddTenancyMenuMap 添加商户菜单表单
func GetAddTenancyMenuMap(ctx *gin.Context) {
	if menus, err := service.GetMenuMap(0, ctx, true); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}

// GetMenu 编辑表单
func GetEditMenuMap(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if menus, err := service.GetMenuMap(req.Id, ctx, false); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}

// GetMenu 获取用户动态路由
func GetMenu(ctx *gin.Context) {
	if menus, err := service.GetMenuTree(ctx); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}

// GetBaseMenuTree 获取基础路由树
func GetBaseMenuTree(ctx *gin.Context) {
	if menus, err := service.GetBaseMenuTree(); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menus, "获取成功", ctx)
	}
}

// AddMenuAuthority 增加menu和角色关联关系
func AddMenuAuthority(ctx *gin.Context) {
	var authorityMenu request.AddMenuAuthorityInfo
	if errs := ctx.ShouldBindJSON(&authorityMenu); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.AddMenuAuthority(authorityMenu.Menus, authorityMenu.AuthorityId); err != nil {
		g.TENANCY_LOG.Error("添加失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败"+err.Error(), ctx)
	} else {
		response.OkWithMessage("添加成功", ctx)
	}
}

// GetMenuAuthority 获取指定角色menu
func GetMenuAuthority(ctx *gin.Context) {
	var param request.GetAuthorityId
	if errs := ctx.ShouldBindJSON(&param); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if menus, err := service.GetMenuAuthority(&param); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithDetailed(menus, "获取失败", ctx)
	} else {
		response.OkWithDetailed(gin.H{"menus": menus}, "获取成功", ctx)
	}
}

// AddBaseMenu 新增菜单
func AddBaseMenu(ctx *gin.Context) {
	var menu model.SysBaseMenu
	if errs := ctx.ShouldBindJSON(&menu); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	menu.IsTenancy = g.StatusFalse
	if menu, err := service.AddBaseMenu(menu); err != nil {
		g.TENANCY_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menu, "添加成功", ctx)
	}
}

// AddTenancyBaseMenu 新增菜单
func AddTenancyBaseMenu(ctx *gin.Context) {
	var menu model.SysBaseMenu
	if errs := ctx.ShouldBindJSON(&menu); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	menu.IsTenancy = g.StatusTrue
	if menu, err := service.AddBaseMenu(menu); err != nil {
		g.TENANCY_LOG.Error("添加失败!", zap.Any("err", err))

		response.FailWithMessage("添加失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menu, "添加成功", ctx)
	}
}

// DeleteBaseMenu 删除菜单
func DeleteBaseMenu(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.DeleteBaseMenu(req.Id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}

// UpdateBaseMenu 更新菜单
func UpdateBaseMenu(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	var menu model.SysBaseMenu
	if errs := ctx.ShouldBindJSON(&menu); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if err := service.UpdateBaseMenu(req.Id, menu); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("更新成功", ctx)
	}
}

// GetBaseMenuById 根据id获取菜单
func GetBaseMenuById(ctx *gin.Context) {
	var req request.GetById
	if errs := ctx.ShouldBindUri(&req); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if menu, err := service.GetBaseMenuById(req.Id); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(menu, "获取成功", ctx)
	}
}

// GetMenuList 分页获取基础menu列表
func GetMenuList(ctx *gin.Context) {
	if menuList, err := service.GetInfoList(1); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    0,
			Page:     1,
			PageSize: 10,
		}, "获取成功", ctx)
	}
}

// GetClientMenuList 分页获取基础menu列表
func GetClientMenuList(ctx *gin.Context) {
	if menuList, err := service.GetInfoList(2); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     menuList,
			Total:    0,
			Page:     1,
			PageSize: 10,
		}, "获取成功", ctx)
	}
}
