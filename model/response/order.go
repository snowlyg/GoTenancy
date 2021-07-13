package response

import (
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/datatypes"
)

type OrderList struct {
	TenancyResponse

	model.BaseOrder

	GroupOrderSn     string         `json:"groupOrderSn" form:"groupOrderSn"`
	TenancyName      string         `json:"tenancyName" form:"tenancyName"`
	IsTrader         int            `json:"isTrader" form:"isTrader"`
	SysUserID        uint           `json:"sysUserId" form:"sysUserId"`
	SysTenancyID     uint           `json:"sysTenancyId"`
	GroupOrderID     int            `json:"groupOrderId"`
	ReconciliationID uint8          `json:"reconciliationId"`
	CartID           uint           `json:"cartId"`
	OrderProduct     []OrderProduct `gorm:"-" json:"orderProduct"`
}

type OrderProduct struct {
	ID       uint           `json:"id"`
	CartInfo datatypes.JSON `json:"cartInfo"`
	model.BaseOrderProduct
	OrderID   uint `json:"orderID"`
	ProductID uint `json:"productId"` // 商品ID
}

type OrderCondition struct {
	Type       int                    `json:"type"`
	Name       string                 `json:"name"`
	Conditions map[string]interface{} `json:"conditions"`
	IsDeleted  bool                   `json:"is_deleted"`
}

type OrderDetail struct {
	TenancyResponse

	model.BaseOrder

	SysUserID        uint   `json:"sysUserId" form:"sysUserId"`
	SysTenancyID     uint   `json:"sysTenancyId"`
	GroupOrderID     int    `json:"groupOrderId"`
	ReconciliationID uint8  `json:"reconciliationId"`
	CartID           uint   `json:"cartId"`
	UserNickName     string `json:"userNickName" form:"userNickName"`
}
