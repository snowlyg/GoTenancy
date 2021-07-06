package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/snowlyg/go-tenancy/g"
)

type SysTenancy struct {
	g.TENANCY_MODEL
	BaseTenancy
	Region SysRegion `json:"region" gorm:"foreignKey:SysRegionCode;references:code;comment:所属区域"`
}

type BaseTenancy struct {
	UUID         uuid.UUID `json:"uuid" gorm:"comment:UUID"`
	Name         string    `json:"name" form:"name" gorm:"column:name;comment:商户名称" binding:"required"`
	Tele         string    `json:"tele" form:"tele" gorm:"column:tele;comment:商户联系电话"`
	Address      string    `json:"address" form:"address" gorm:"column:address;comment:商户详细地址"`
	BusinessTime string    `json:"businessTime" form:"businessTime" gorm:"column:business_time;comment:商户营业时间"`
	Status       int       `gorm:"column:status;type:tinyint(1);not null;default:2" json:"status" binding:"required"` // 商户是否 禁用2锁定,1正常

	Keyword        string  `gorm:"column:keyword;type:varchar(64);not null" json:"Keyword"`                    // 商户关键字
	Avatar         string  `gorm:"column:avatar;type:varchar(128)" json:"Avatar"`                              // 商户头像
	Banner         string  `gorm:"column:banner;type:varchar(128)" json:"Banner"`                              // 商户banner图片
	Sales          uint    `gorm:"column:sales;type:int unsigned;default:0" json:"sales"`                      // 销量
	ProductScore   float64 `gorm:"column:product_score;type:decimal(11,1);default:5.0" json:"productScore"`    // 商品描述评分
	ServiceScore   float64 `gorm:"column:service_score;type:decimal(11,1);default:5.0" json:"serviceScore"`    // 服务评分
	PostageScore   float64 `gorm:"column:postage_score;type:decimal(11,1);default:5.0" json:"postageScore"`    // 物流评分
	Mark           string  `gorm:"column:mark;type:varchar(256);not null" json:"mark"`                         // 商户备注
	RegAdminID     uint    `gorm:"column:reg_admin_id;type:int unsigned;not null;default:0" json:"regAdminId"` // 总后台管理员ID
	Sort           uint    `gorm:"column:sort;type:int unsigned;not null;default:0" json:"sort"`
	CommissionRate float64 `gorm:"column:commission_rate;type:decimal(11,4)" json:"commissionRate"`                // 提成比例
	IsAudit        uint8   `gorm:"column:is_audit;type:tinyint unsigned;not null;default:2" json:"isAudit"`        // 添加的产品是否审核 0不审核1审核
	IsBroRoom      uint8   `gorm:"column:is_bro_room;type:tinyint unsigned;not null;default:1" json:"isBroRoom"`   // 是否审核直播间 0不审核1审核
	IsBroGoods     uint8   `gorm:"column:is_bro_goods;type:tinyint unsigned;not null;default:1" json:"isBroGoods"` // 是否审核直播商品 0不审核1审核
	IsBest         uint8   `gorm:"column:is_best;type:tinyint unsigned;not null;default:2" json:"isBest"`          // 是否推荐
	IsTrader       uint8   `gorm:"column:is_trader;type:tinyint unsigned;not null;default:2" json:"isTrader"`      // 是否自营
	State          int     `gorm:"column:state;type:tinyint unsigned;not null;default:0" json:"State"`             // 商户是否 1开启0关闭
	Info           string  `gorm:"column:info;type:varchar(256);not null;default:''" json:"Info"`                  // 店铺简介
	ServicePhone   string  `gorm:"column:service_phone;type:varchar(13);not null;default:''" json:"servicePhone"`  // 店铺电话
	CareCount      uint    `gorm:"column:care_count;type:int unsigned;default:0" json:"careCount"`                 // 关注总数
	CopyProductNum int     `gorm:"column:copy_product_num;type:int;default:0" json:"copyProductNum"`               // 剩余复制商品次数

	SysRegionCode uint `json:"sysRegionCode" form:"sysRegionCode" gorm:"column:sys_region_code;comment:商户所属区域code" binding:"required"`
}
