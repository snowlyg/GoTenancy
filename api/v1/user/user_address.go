package user

import (
	"github.com/gin-gonic/gin"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/multi"
	"go.uber.org/zap"
)

// CreateAddress
func CreateAddress(ctx *gin.Context) {
	var address request.CreateAddress
	if errs := ctx.ShouldBindJSON(&address); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if returnAddress, err := service.CreateAddress(address, user_id); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("添加失败"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getAddressMap(returnAddress), "创建成功", ctx)
	}
}

// UpdateAddress
func UpdateAddress(ctx *gin.Context) {
	var address request.UpdateAddress
	if errs := ctx.ShouldBindJSON(&address); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnAddress, err := service.UpdateAddress(address); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(getAddressMap(returnAddress), "更新成功", ctx)
	}
}

// getAddressMap
func getAddressMap(returnAddress model.UserAddress) gin.H {
	return gin.H{"id": returnAddress.ID, "name": returnAddress.Name, "phone": returnAddress.Phone, "sex": returnAddress.Sex, "country": returnAddress.Country, "province": returnAddress.Province, "city": returnAddress.City, "district": returnAddress.District, "isDefault": returnAddress.IsDefault, "detail": returnAddress.Detail, "postcode": returnAddress.Postcode, "age": returnAddress.Age, "hospitalName": returnAddress.HospitalName, "locName": returnAddress.LocName, "bedNum": returnAddress.BedNum, "hospitalNo": returnAddress.HospitalNO, "disease": returnAddress.Disease}
}

// GetAddressList
func GetAddressList(ctx *gin.Context) {
	var pageInfo request.PageInfo
	if errs := ctx.ShouldBindJSON(&pageInfo); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if list, total, err := service.GetAddressInfoList(pageInfo, user_id); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetAddressById
func GetAddressById(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	address, err := service.GetAddressByID(reqId.Id, user_id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败:"+err.Error(), ctx)
	} else {
		response.OkWithData(address, ctx)
	}
}

// DeleteAddress
func DeleteAddress(ctx *gin.Context) {
	var reqId request.GetById
	if errs := ctx.ShouldBindJSON(&reqId); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := multi.GetUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if err := service.DeleteAddress(reqId.Id, user_id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败:"+err.Error(), ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
