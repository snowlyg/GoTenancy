package model

import (
	"time"

	"github.com/snowlyg/go-tenancy/g"
)

// ProductReply 商品评论表
type ProductReply struct {
	g.TENANCY_MODEL

	ProductScore         int       `gorm:"column:product_score;type:tinyint(1);not null" json:"productScore"`           // 商品分数
	ServiceScore         int       `gorm:"column:service_score;type:tinyint(1);not null" json:"serviceScore"`           // 服务分数
	PostageScore         int       `gorm:"column:postage_score;type:tinyint(1);not null" json:"postageScore"`           // 物流分数
	Rate                 float64   `gorm:"column:rate;type:float(2,1);default:5.0" json:"rate"`                         // 平均值
	Comment              string    `gorm:"column:comment;type:varchar(512);not null" json:"comment"`                    // 评论内容
	Pics                 string    `gorm:"column:pics;type:text;not null" json:"pics"`                                  // 评论图片
	MerchantReplyContent string    `gorm:"column:merchant_reply_content;type:varchar(300)" json:"merchantReplyContent"` // 管理员回复内容
	MerchantReplyTime    time.Time `gorm:"column:merchant_reply_time;type:timestamp" json:"merchantReplyTime"`          // 管理员回复时间
	IsDel                uint8     `gorm:"column:is_del;type:tinyint unsigned;not null;default:1" json:"isDel"`         // 2未删除1已删除
	IsReply              int       `gorm:"column:is_reply;type:tinyint(1);not null;default:1" json:"isReply"`           // 2未回复1已回复
	IsVirtual            int       `gorm:"column:is_virtual;type:tinyint(1);not null;default:1" json:"isVirtual"`       // 2不是虚拟评价1是虚拟评价
	Nickname             string    `gorm:"column:nickname;type:varchar(64);not null" json:"nickname"`                   // 用户名称
	Avatar               string    `gorm:"column:avatar;type:varchar(255);not null" json:"avatar"`                      // 用户头像

	SysUserID      int    `gorm:"index:sys_user_id;column:sys_user_id;type:int;not null" json:"sysUserId"`          // 用户ID
	SysTenancyID   int    `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	ProductID      uint   `gorm:"index:product_id;column:product_id;type:int;not null" json:"productId"`            // 商品id  // 商品id
	OrderProductID int    `gorm:"index:order_id;column:order_product_id;type:int;not null" json:"orderProductId"`   // 订单商品ID
	Unique         string `gorm:"uniqueIndex:order_id;column:unique;type:char(12)" json:"unique"`                   // 商品 sku
	ProductType    int8   `gorm:"column:product_type;type:tinyint;not null;default:1" json:"productType"`           // 1=普通商品
}
