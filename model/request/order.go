package request

type OrderPageInfo struct {
	Page         int    `json:"page" form:"page" binding:"required"`
	PageSize     int    `json:"pageSize" form:"pageSize" binding:"required"`
	Date         string `json:"date" form:"date"`
	IsTrader     string `json:"isTrader" form:"isTrader"`
	Keywords     string `json:"keywords" form:"keywords"`
	SysTenancyId int    `json:"sysTenancyId" form:"sysTenancyId"`
	OrderSn      string `json:"orderSn" form:"orderSn"`
	Status       string `json:"status" form:"status"`
	Username     string `json:"username" form:"username"`
}

type CartInfo struct {
	Product     CartInfoProduct     `json:"product" form:"product"`
	ProductAttr CartInfoProductAttr `json:"productAttr" form:"productAttr"`
}

type CartInfoProduct struct {
	Image     string `json:"image" form:"image"`
	StoreName string `json:"storeName" form:"storeName"`
}
type CartInfoProductAttr struct {
	Price float64 `json:"price" form:"price"`
	Sku   string  `json:"sku" form:"sku"`
}
