package v1

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"github.com/snowlyg/go-tenancy/service"
	"github.com/snowlyg/go-tenancy/utils"
	"go.uber.org/zap"
)

// CreateReceipt
func CreateReceipt(ctx iris.Context) {
	var receipt request.CreateReceipt
	if errs := utils.Verify(ctx.ReadJSON(&receipt)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := getUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if returnReceipt, err := service.CreateReceipt(receipt, user_id); err != nil {
		g.TENANCY_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", ctx)
	} else {
		response.OkWithDetailed(getReceiptMap(returnReceipt), "创建成功", ctx)
	}
}

// UpdateReceipt
func UpdateReceipt(ctx iris.Context) {
	var receipt request.UpdateReceipt
	if errs := utils.Verify(ctx.ReadJSON(&receipt)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}
	if returnReceipt, err := service.UpdateReceipt(receipt); err != nil {
		g.TENANCY_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", ctx)
	} else {
		response.OkWithDetailed(getReceiptMap(returnReceipt), "更新成功", ctx)
	}
}

// getReceiptMap
func getReceiptMap(returnReceipt model.GeneralReceipt) context.Map {
	return iris.Map{
		"id":               returnReceipt.ID,
		"receiptType":      returnReceipt.ReceiptType,
		"receiptTitle":     returnReceipt.ReceiptTitle,
		"receiptTitleType": returnReceipt.ReceiptTitleType,
		"dutyGaragraph":    returnReceipt.DutyGaragraph,
		"email":            returnReceipt.Email,
		"bankName":         returnReceipt.BankName,
		"bankCode":         returnReceipt.BankCode,
		"address":          returnReceipt.Address,
		"tel":              returnReceipt.Tel,
		"isDefault":        returnReceipt.IsDefault,
	}
}

// GetReceiptList
func GetReceiptList(ctx iris.Context) {
	var pageInfo request.PageInfo
	if errs := utils.Verify(ctx.ReadJSON(&pageInfo)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := getUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if list, total, err := service.GetReceiptInfoList(pageInfo, user_id); err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", ctx)
	}
}

// GetReceiptById
func GetReceiptById(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := getUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	receipt, err := service.GetReceiptByID(reqId.Id, user_id)
	if err != nil {
		g.TENANCY_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", ctx)
	} else {
		response.OkWithData(receipt, ctx)
	}
}

// DeleteReceipt
func DeleteReceipt(ctx iris.Context) {
	var reqId request.GetById
	if errs := utils.Verify(ctx.ReadJSON(&reqId)); errs != nil {
		response.FailWithMessage(errs.Error(), ctx)
		return
	}

	user_id := getUserId(ctx)
	if user_id == 0 {
		g.TENANCY_LOG.Error("更新失败! general_user is 0")
		response.FailWithMessage("请求失败", ctx)
		return
	}

	if err := service.DeleteReceipt(reqId.Id, user_id); err != nil {
		g.TENANCY_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", ctx)
	} else {
		response.OkWithMessage("删除成功", ctx)
	}
}
