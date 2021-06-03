package response

type TenancyProduct struct {
	TenancyResponse
	StoreName     string  `json:"storeName"`     // 商品名称
	StoreInfo     string  `json:"storeInfo"`     // 商品简介
	Keyword       string  `json:"keyword"`       // 关键字
	BarCode       string  `json:"barCode"`       // 产品条码（一维码）
	IsShow        int     `json:"isShow"`        // 商户 状态（1：未上架，2：上架）
	Status        int     `json:"status"`        // 管理员 状态（1：审核中，2：审核通过 3: 未通过 4: 下架）
	UnitName      string  `json:"unitName"`      // 单位名
	Sort          int16   `json:"sort"`          // 排序
	Rank          int16   `json:"rank"`          // 总后台排序
	Sales         uint32  `json:"sales"`         // 销量
	Price         float64 `json:"price"`         // 最低价格
	Cost          float64 `json:"cost"`          // 成本价
	OtPrice       float64 `json:"otPrice"`       // 原价
	Stock         uint    `json:"stock"`         // 总库存
	IsHot         uint8   `json:"isHot"`         // 是否热卖
	IsBenefit     uint8   `json:"isBenefit"`     // 促销单品
	IsBest        uint8   `json:"isBest"`        // 是否精品
	IsNew         uint8   `json:"isNew"`         // 是否新品
	IsGood        int     `json:"isGood"`        // 是否优品推荐
	ProductType   uint8   `json:"productType"`   // 1.普通商品 2.秒杀商品,3.预售商品，4.助力商品
	Ficti         int32   `json:"ficti"`         // 虚拟销量
	Browse        int     `json:"browse"`        // 浏览量
	CodePath      string  `json:"codePath"`      // 产品二维码地址(用户小程序海报)
	VideoLink     string  `json:"videoLink"`     // 主图视频链接
	SpecType      int     `json:"specType"`      // 规格 1单 2多
	ExtensionType int     `json:"extensionType"` // 佣金比例 1.系统，2.自定义
	Refusal       string  `json:"refusal"`       // 审核拒绝理由
	Rate          float64 `json:"rate"`          // 评价分数
	ReplyCount    uint    `json:"replyCount"`    // 评论数
	GiveCouponIDs string  `json:"giveCouponIds"` // 赠送优惠券
	IsGiftBag     int     `json:"isGiftBag"`     // 是否为礼包
	CareCount     int     `json:"careCount"`     // 收藏数

	Image       string `json:"image"`       // 商品图片
	SliderImage string `json:"sliderImage"` // 轮播图

	OldID             int    `json:"oldId"`             // 原商品ID
	TempID            int    `json:"tempId"`            // 运费模板ID
	SysTenancyID      int    `json:"sysTenancyId"`      // 商户 id
	SysTenancyName    string `json:"sysTenancyName"`    // 商户名称
	SysBrandID        int    `json:"sysBrandId"`        // 品牌 id
	TenancyCategoryID int    `json:"tenancyCategoryId"` // 分类id
}
