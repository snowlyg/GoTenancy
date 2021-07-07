package request

type RefundOrderPageInfo struct {
	Page          int    `json:"page" form:"page" binding:"required"`
	PageSize      int    `json:"pageSize" form:"pageSize" binding:"required"`
	Date          string `json:"date" form:"date"`
	IsTrader      string `json:"isTrader" form:"isTrader"`
	OrderSn       string `json:"orderSn" form:"orderSn"`
	RefundOrderSn string `json:"refundOrderSn" form:"refundOrderSn"`
	Status        string `json:"status" form:"status"`
}
