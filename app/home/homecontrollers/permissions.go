package homecontrollers

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"go-tenancy/app/home/hometransformer"
	"go-tenancy/app/home/homevalidates"
	"go-tenancy/libs"
	"go-tenancy/models/tenant"
)

/**
* @api {get} /admin/permissions/:id 根据id获取权限信息
* @apiName 根据id获取权限信息
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 根据id获取权限信息
* @apiSampleRequest /admin/permissions/:id
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission
 */
func GetPermission(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	perm := tenant.NewRabcPermission(id, "", "")
	perm.GetRabcPermissionById()

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, permTransform(perm), "操作成功"))
}

/**
* @api {post} /admin/permissions/ 新建权限
* @apiName 新建权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 新建权限
* @apiSampleRequest /admin/permissions/
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func CreatePermission(ctx iris.Context) {
	aul := new(homevalidates.RabcPermissionRequest)
	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}
	err := homevalidates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(homevalidates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(false, nil, e))
				return
			}
		}
	}

	perm := tenant.NewRabcPermissionByStruct(aul)
	perm.CreateRabcPermission()

	ctx.StatusCode(iris.StatusOK)
	if perm.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, perm, "操作失败"))
	} else {
		_, _ = ctx.JSON(ApiResource(true, perm, "操作成功"))
	}

}

/**
* @api {post} /admin/permissions/:id/update 更新权限
* @apiName 更新权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 更新权限
* @apiSampleRequest /admin/permissions/:id/update
* @apiParam {string} name 权限名
* @apiParam {string} display_name
* @apiParam {string} description
* @apiParam {string} level
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func UpdatePermission(ctx iris.Context) {
	aul := new(homevalidates.RabcPermissionRequest)

	if err := ctx.ReadJSON(aul); err != nil {
		ctx.StatusCode(iris.StatusOK)
		_, _ = ctx.JSON(ApiResource(false, nil, err.Error()))
		return
	}
	err := homevalidates.Validate.Struct(*aul)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs.Translate(homevalidates.ValidateTrans) {
			if len(e) > 0 {
				ctx.StatusCode(iris.StatusOK)
				_, _ = ctx.JSON(ApiResource(false, nil, e))
				return
			}
		}
	}

	id, _ := ctx.Params().GetUint("id")
	perm := tenant.NewRabcPermission(id, "", "")
	perm.UpdateRabcPermission(aul)

	ctx.StatusCode(iris.StatusOK)
	if perm.ID == 0 {
		_, _ = ctx.JSON(ApiResource(false, perm, "操作失败"))
	} else {
		_, _ = ctx.JSON(ApiResource(true, perm, "操作成功"))
	}

}

/**
* @api {delete} /admin/permissions/:id/delete 删除权限
* @apiName 删除权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 删除权限
* @apiSampleRequest /admin/permissions/:id/delete
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func DeletePermission(ctx iris.Context) {
	id, _ := ctx.Params().GetUint("id")
	perm := tenant.NewRabcPermission(id, "", "")
	perm.DeleteRabcPermissionById()
	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, nil, "删除成功"))
}

/**
* @api {get} /permissions 获取所有的权限
* @apiName 获取所有的权限
* @apiGroup Permissions
* @apiVersion 1.0.0
* @apiDescription 获取所有的权限
* @apiSampleRequest /permissions
* @apiSuccess {String} msg 消息
* @apiSuccess {bool} state 状态
* @apiSuccess {String} data 返回数据
* @apiPermission null
 */
func GetAllPermissions(ctx iris.Context) {
	offset := libs.ParseInt(ctx.URLParam("offset"), 1)
	limit := libs.ParseInt(ctx.URLParam("limit"), 20)
	name := ctx.FormValue("name")
	orderBy := ctx.FormValue("orderBy")

	permissions := tenant.GetAllRabcPermissions(name, orderBy, offset, limit)

	ctx.StatusCode(iris.StatusOK)
	_, _ = ctx.JSON(ApiResource(true, permsTransform(permissions), "操作成功"))
}

func permsTransform(perms []*tenant.RabcPermission) []*hometransformer.Permission {
	var rs []*hometransformer.Permission
	for _, perm := range perms {
		r := permTransform(perm)
		rs = append(rs, r)
	}
	return rs
}

func permTransform(perm *tenant.RabcPermission) *hometransformer.Permission {
	r := &hometransformer.Permission{}
	g := gf.NewTransform(r, perm, time.RFC3339)
	_ = g.Transformer()
	return r
}
