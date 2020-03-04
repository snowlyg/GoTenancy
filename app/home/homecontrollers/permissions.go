package homecontrollers

import (
	"time"

	"github.com/kataras/iris/v12"
	gf "github.com/snowlyg/gotransformer"
	"go-tenancy/app/home/hometransformer"
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
