package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

const (
	UnknownSale int32 = iota
	GeneralSale       // 普通商品
	FlashSale         //秒杀商品
	PreSale           // 预售商品
	AssistSale        // 助力商品
)

const (
	UnknownSpec int = iota
	SingleSpec      // 单规格
	DoubleSpec      // 多规格
)

const (
	UnknownExtensionType int = iota
	SystemExtensionType      //佣金比例 系统
	CustomExtensionType      //佣金比例 自定义
)

type Product struct {
	g.TENANCY_MODEL
	BaseProduct
	GiveCouponIDs string `gorm:"column:give_coupon_ids;type:varchar(500)" json:"giveCouponIds"`      // 赠送优惠券
	SliderImage   string `gorm:"column:slider_image;type:varchar(2000);not null" json:"sliderImage"` // 轮播图
}

type BaseProduct struct {
	StoreName     string  `gorm:"column:store_name;type:varchar(128);not null" json:"storeName" binding:"required"` // 商品名称
	StoreInfo     string  `gorm:"column:store_info;type:varchar(256);not null" json:"storeInfo" `                   // 商品简介
	Keyword       string  `gorm:"column:keyword;type:varchar(128);not null" json:"keyword"`                         // 关键字
	BarCode       string  `gorm:"column:bar_code;type:varchar(15);not null;default:''" json:"barCode"`              // 产品条码（一维码）
	IsShow        int     `gorm:"column:is_show;type:tinyint;not null;default:2" json:"isShow"`                     // 商户 状态（2：未上架，1：上架）
	Status        int     `gorm:"column:status;type:tinyint;not null;default:2" json:"status"`                      // 管理员 状态（2：审核中，1：审核通过 3: 未通过）
	UnitName      string  `gorm:"column:unit_name;type:varchar(16);not null" json:"unitName" `                      // 单位名
	Sort          int16   `gorm:"index;column:sort;type:smallint;not null;default:0" json:"sort"`                   // 排序
	Rank          int16   `gorm:"column:rank;type:smallint;not null;default:0" json:"rank"`                         // 总后台排序
	Sales         uint32  `gorm:"index:sales;column:sales;type:mediumint unsigned;not null;default:0" json:"sales"` // 销量
	Price         float64 `gorm:"column:price;type:decimal(10,2) unsigned;default:0.00" json:"price" `              // 最低价格
	Cost          float64 `gorm:"column:cost;type:decimal(10,2);default:0.00" json:"cost" `                         // 成本价
	OtPrice       float64 `gorm:"column:ot_price;type:decimal(10,2);default:0.00" json:"otPrice" `                  // 原价
	Stock         uint    `gorm:"column:stock;type:int unsigned;default:0" json:"stock" `                           // 总库存
	IsHot         int     `gorm:"column:is_hot;type:tinyint unsigned;not null;default:0" json:"isHot"`              // 是否热卖
	IsBenefit     int     `gorm:"column:is_benefit;type:tinyint unsigned;not null;default:0" json:"isBenefit"`      // 促销单品
	IsBest        int     `gorm:"column:is_best;type:tinyint unsigned;not null;default:0" json:"isBest"`            // 是否精品
	IsNew         int     `gorm:"column:is_new;type:tinyint unsigned;not null;default:0" json:"isNew"`              // 是否新品
	IsGood        int     `gorm:"column:is_good;type:tinyint;not null;default:2" json:"isGood"`                     // 是否优品推荐
	ProductType   int32   `gorm:"column:product_type;type:tinyint unsigned;not null;default:0" json:"productType" ` // 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
	Ficti         int32   `gorm:"column:ficti;type:mediumint;default:0" json:"ficti"`                               // 虚拟销量
	Browse        int     `gorm:"column:browse;type:int;default:0" json:"browse"`                                   // 浏览量
	CodePath      string  `gorm:"column:code_path;type:varchar(64);not null;default:''" json:"codePath"`            // 产品二维码地址(用户小程序海报)
	VideoLink     string  `gorm:"column:video_link;type:varchar(200);not null;default:''" json:"videoLink"`         // 主图视频链接
	SpecType      int     `gorm:"column:spec_type;type:tinyint;not null" json:"specType" `                          // 规格 1单 2多
	ExtensionType int     `gorm:"column:extension_type;type:tinyint" json:"extensionType"`                          // 佣金比例 1.系统，2.自定义
	Refusal       string  `gorm:"column:refusal;type:varchar(255)" json:"refusal"`                                  // 审核拒绝理由
	Rate          float64 `gorm:"column:rate;type:decimal(2,1);default:5.0" json:"rate"`                            // 评价分数
	ReplyCount    uint    `gorm:"column:reply_count;type:int unsigned;default:0" json:"replyCount"`                 // 评论数

	IsGiftBag int `gorm:"column:is_gift_bag;type:tinyint;default:2" json:"isGiftBag"`     // 是否为礼包
	CareCount int `gorm:"column:care_count;type:int;not null;default:0" json:"careCount"` // 收藏数
	// 原商品ID
	Image string `gorm:"column:image;type:varchar(256);not null" json:"image"` // 商品图片

	OldID        uint `gorm:"column:old_id;type:int;default:0" json:"oldId"`
	TempID       uint `gorm:"column:temp_id;type:int;not null;default:0" json:"tempId"`                         // 运费模板ID
	SysTenancyID uint `gorm:"index:sys_tenancy_id;column:sys_tenancy_id;type:int;not null" json:"sysTenancyId"` // 商户 id
	SysBrandID   uint `gorm:"column:sys_brand_id;type:int" json:"sysBrandId"`                                   // 品牌 id

	ProductCategoryID uint `gorm:"index:product_category_id;column:product_category_id;type:int;not null" json:"productCategoryId"` // 平台分类

}
