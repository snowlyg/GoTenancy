package router

import (
	"github.com/kataras/iris/v12"
	v1 "github.com/snowlyg/go-tenancy/api/v1"
)

func InitReceiptRouter(Router iris.Party) {
	ReceiptRouter := Router.Party("/receipt")
	{
		ReceiptRouter.Post("/createReceipt", v1.CreateReceipt)
		ReceiptRouter.Post("/getReceiptList", v1.GetReceiptList)
		ReceiptRouter.Post("/getReceiptById", v1.GetReceiptById)
		ReceiptRouter.Put("/updateReceipt", v1.UpdateReceipt)
		ReceiptRouter.Delete("/deleteReceipt", v1.DeleteReceipt)
	}
}
