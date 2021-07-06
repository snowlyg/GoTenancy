package g

import (
	"fmt"
	"time"

	"github.com/bwmarrin/snowflake"
	"gorm.io/gorm"
)

type TENANCY_MODEL struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// CreateOrderSn 生成订单号 payType+20060102150405+随机数
func CreateOrderSn(payType int) string {
	node, _ := snowflake.NewNode(1)
	id := node.Generate().Int64()
	now := time.Now().Format("20060102150405")
	return fmt.Sprintf("%d%s%d", payType, now, id)
}
