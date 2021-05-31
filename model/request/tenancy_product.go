package request

type CreateTenancyProduct struct {
	StoreName     string  `validate:"required" json:"storeName"`   // 商品名称
	StoreInfo     string  `validate:"required" json:"storeInfo"`   // 商品简介
	Keyword       string  `json:"keyword"`                         // 关键字
	BarCode       string  `validate:"required" json:"barCode"`     // 产品条码（一维码）
	IsShow        bool    `json:"isShow"`                          // 商户 状态（0：未上架，1：上架）
	Status        bool    `json:"status"`                          // 管理员 状态（0：审核中，1：审核通过 -1: 未通过 -2: 下架）
	TenancyStatus bool    `json:"tenancyStatus"`                   // 商铺状态是否 1.正常 0. 非正常
	UnitName      string  `validate:"required" json:"unitName"`    // 单位名
	Sort          int16   `json:"sort"`                            // 排序
	Rank          int16   `json:"rank"`                            // 总后台排序
	Sales         uint32  `json:"sales"`                           // 销量
	Price         float64 `validate:"required" json:"price"`       // 最低价格
	Cost          float64 `validate:"required" json:"cost"`        // 成本价
	OtPrice       float64 `validate:"required" json:"otPrice"`     // 原价
	Stock         uint    `validate:"required" json:"stock"`       // 总库存
	IsHot         uint8   `json:"isHot"`                           // 是否热卖
	IsBenefit     uint8   `json:"isBenefit"`                       // 促销单品
	IsBest        uint8   `json:"isBest"`                          // 是否精品
	IsNew         uint8   `json:"isNew"`                           // 是否新品
	IsGood        bool    `json:"isGood"`                          // 是否优品推荐
	ProductType   uint8   `validate:"required" json:"productType"` // 0.普通商品 1.秒杀商品,2.预售商品，3.助力商品
	Ficti         int32   `json:"ficti"`                           // 虚拟销量
	Browse        int     `json:"browse"`                          // 浏览量
	CodePath      string  `json:"codePath"`                        // 产品二维码地址(用户小程序海报)
	VideoLink     string  `json:"videoLink"`                       // 主图视频链接
	SpecType      int     `validate:"required" json:"specType"`    // 规格 0单 1多
	ExtensionType int     `json:"extensionType"`                   // 佣金比例 0.系统，1.自定义
	Refusal       string  `json:"refusal"`                         // 审核拒绝理由
	Rate          float64 `json:"rate"`                            // 评价分数
	ReplyCount    uint    `json:"replyCount"`                      // 评论数
	GiveCouponIDs string  `json:"giveCouponIds"`                   // 赠送优惠券
	IsGiftBag     bool    `json:"isGiftBag"`                       // 是否为礼包
	CareCount     int     `json:"careCount"`                       // 收藏数

	Image       string `json:"image"`       // 商品图片
	SliderImage string `json:"sliderImage"` // 轮播图

	OldID             int `json:"oldId"`                                 // 原商品ID
	TempID            int `json:"tempId"`                                // 运费模板ID
	SysTenancyID      int `validate:"required" json:"sysTenancyId"`      // 商户 id
	SysBrandID        int `validate:"required" json:"sysBrandId"`        // 品牌 id
	TenancyCategoryID int `validate:"required" json:"tenancyCategoryId"` // 分类id
}

type UpdateTenancyProduct struct {
	Id            uint    `json:"id" form:"id" validate:"required,gt=0"`
	StoreName     string  `validate:"required" json:"storeName"`   // 商品名称
	StoreInfo     string  `validate:"required" json:"storeInfo"`   // 商品简介
	Keyword       string  `json:"keyword"`                         // 关键字
	BarCode       string  `validate:"required" json:"barCode"`     // 产品条码（一维码）
	IsShow        bool    `json:"isShow"`                          // 商户 状态（0：未上架，1：上架）
	Status        bool    `json:"status"`                          // 管理员 状态（0：审核中，1：审核通过 -1: 未通过 -2: 下架）
	TenancyStatus bool    `json:"tenancyStatus"`                   // 商铺状态是否 1.正常 0. 非正常
	UnitName      string  `validate:"required" json:"unitName"`    // 单位名
	Sort          int16   `json:"sort"`                            // 排序
	Rank          int16   `json:"rank"`                            // 总后台排序
	Sales         uint32  `json:"sales"`                           // 销量
	Price         float64 `validate:"required" json:"price"`       // 最低价格
	Cost          float64 `validate:"required" json:"cost"`        // 成本价
	OtPrice       float64 `validate:"required" json:"otPrice"`     // 原价
	Stock         uint    `validate:"required" json:"stock"`       // 总库存
	IsHot         uint8   `json:"isHot"`                           // 是否热卖
	IsBenefit     uint8   `json:"isBenefit"`                       // 促销单品
	IsBest        uint8   `json:"isBest"`                          // 是否精品
	IsNew         uint8   `json:"isNew"`                           // 是否新品
	IsGood        bool    `json:"isGood"`                          // 是否优品推荐
	ProductType   uint8   `validate:"required" json:"productType"` // 0.普通商品 1.秒杀商品,2.预售商品，3.助力商品
	Ficti         int32   `json:"ficti"`                           // 虚拟销量
	Browse        int     `json:"browse"`                          // 浏览量
	CodePath      string  `json:"codePath"`                        // 产品二维码地址(用户小程序海报)
	VideoLink     string  `json:"videoLink"`                       // 主图视频链接
	SpecType      int     `validate:"required" json:"specType"`    // 规格 0单 1多
	ExtensionType int     `json:"extensionType"`                   // 佣金比例 0.系统，1.自定义
	Refusal       string  `json:"refusal"`                         // 审核拒绝理由
	Rate          float64 `json:"rate"`                            // 评价分数
	ReplyCount    uint    `json:"replyCount"`                      // 评论数
	GiveCouponIDs string  `json:"giveCouponIds"`                   // 赠送优惠券
	IsGiftBag     bool    `json:"isGiftBag"`                       // 是否为礼包
	CareCount     int     `json:"careCount"`                       // 收藏数

	Image       string `json:"image"`       // 商品图片
	SliderImage string `json:"sliderImage"` // 轮播图

	OldID             int `json:"oldId"`                                 // 原商品ID
	TempID            int `json:"tempId"`                                // 运费模板ID
	SysTenancyID      int `validate:"required" json:"sysTenancyId"`      // 商户 id
	SysBrandID        int `validate:"required" json:"sysBrandId"`        // 品牌 id
	TenancyCategoryID int `validate:"required" json:"tenancyCategoryId"` // 分类id
}
