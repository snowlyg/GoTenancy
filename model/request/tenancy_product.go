package request

type CreateTenancyProduct struct {
	StoreName     string  `binding:"required" json:"storeName"`   // 商品名称
	StoreInfo     string  `binding:"required" json:"storeInfo"`   // 商品简介
	Keyword       string  `json:"keyword"`                        // 关键字
	BarCode       string  `json:"barCode"`                        // 产品条码（一维码）
	IsShow        int     `json:"isShow"`                         // 商户 状态（0：未上架，1：上架）
	Status        int     `json:"status"`                         // 管理员 状态（0：审核中，1：审核通过 -1: 未通过 -2: 下架）
	TenancyStatus int     `json:"tenancyStatus"`                  // 商铺状态是否 1.正常 0. 非正常
	UnitName      string  `binding:"required" json:"unitName"`    // 单位名
	Sort          int16   `json:"sort"`                           // 排序
	Rank          int16   `json:"rank"`                           // 总后台排序
	Sales         uint32  `json:"sales"`                          // 销量
	Price         float64 `binding:"required" json:"price"`       // 最低价格
	Cost          float64 `binding:"required" json:"cost"`        // 成本价
	OtPrice       float64 `binding:"required" json:"otPrice"`     // 原价
	Stock         uint    `binding:"required" json:"stock"`       // 总库存
	IsHot         uint8   `json:"isHot"`                          // 是否热卖
	IsBenefit     uint8   `json:"isBenefit"`                      // 促销单品
	IsBest        uint8   `json:"isBest"`                         // 是否精品
	IsNew         uint8   `json:"isNew"`                          // 是否新品
	IsGood        int     `json:"isGood"`                         // 是否优品推荐
	ProductType   int32   `binding:"required" json:"productType"` // 0.普通商品 1.秒杀商品,2.预售商品，3.助力商品
	Ficti         int32   `json:"ficti"`                          // 虚拟销量
	Browse        int     `json:"browse"`                         // 浏览量
	CodePath      string  `json:"codePath"`                       // 产品二维码地址(用户小程序海报)
	VideoLink     string  `json:"videoLink"`                      // 主图视频链接
	SpecType      int     `binding:"required" json:"specType"`    // 规格 0单 1多
	ExtensionType int     `json:"extensionType"`                  // 佣金比例 0.系统，1.自定义
	Refusal       string  `json:"refusal"`                        // 审核拒绝理由
	Rate          float64 `json:"rate"`                           // 评价分数
	ReplyCount    uint    `json:"replyCount"`                     // 评论数
	GiveCouponIDs string  `json:"giveCouponIds"`                  // 赠送优惠券
	IsGiftBag     int     `json:"isGiftBag"`                      // 是否为礼包
	CareCount     int     `json:"careCount"`                      // 收藏数

	Image       string `json:"image"`       // 商品图片
	SliderImage string `json:"sliderImage"` // 轮播图

	OldID             int `json:"oldId"`                                // 原商品ID
	TempID            int `json:"tempId"`                               // 运费模板ID
	SysTenancyID      int `json:"sysTenancyId"`                         // 商户 id
	SysBrandID        int `binding:"required" json:"sysBrandId"`        // 品牌 id
	TenancyCategoryID int `binding:"required" json:"tenancyCategoryId"` // 分类id
}

type UpdateTenancyProduct struct {
	Id            uint    `json:"id" form:"id" binding:"required,gt=0"`
	StoreName     string  `binding:"required" json:"storeName"`   // 商品名称
	StoreInfo     string  `binding:"required" json:"storeInfo"`   // 商品简介
	Keyword       string  `json:"keyword"`                        // 关键字
	BarCode       string  `json:"barCode"`                        // 产品条码（一维码）
	IsShow        int     `json:"isShow"`                         // 商户 状态（0：未上架，1：上架）
	Status        int     `json:"status"`                         // 管理员 状态（0：审核中，1：审核通过 -1: 未通过 -2: 下架）
	TenancyStatus int     `json:"tenancyStatus"`                  // 商铺状态是否 1.正常 0. 非正常
	UnitName      string  `binding:"required" json:"unitName"`    // 单位名
	Sort          int16   `json:"sort"`                           // 排序
	Rank          int16   `json:"rank"`                           // 总后台排序
	Sales         uint32  `json:"sales"`                          // 销量
	Price         float64 `binding:"required" json:"price"`       // 最低价格
	Cost          float64 `binding:"required" json:"cost"`        // 成本价
	OtPrice       float64 `binding:"required" json:"otPrice"`     // 原价
	Stock         uint    `binding:"required" json:"stock"`       // 总库存
	IsHot         uint8   `json:"isHot"`                          // 是否热卖
	IsBenefit     uint8   `json:"isBenefit"`                      // 促销单品
	IsBest        uint8   `json:"isBest"`                         // 是否精品
	IsNew         uint8   `json:"isNew"`                          // 是否新品
	IsGood        int     `json:"isGood"`                         // 是否优品推荐
	ProductType   int32   `binding:"required" json:"productType"` // 0.普通商品 1.秒杀商品,2.预售商品，3.助力商品
	Ficti         int32   `json:"ficti"`                          // 虚拟销量
	Browse        int     `json:"browse"`                         // 浏览量
	CodePath      string  `json:"codePath"`                       // 产品二维码地址(用户小程序海报)
	VideoLink     string  `json:"videoLink"`                      // 主图视频链接
	SpecType      int     `binding:"required" json:"specType"`    // 规格 0单 1多
	ExtensionType int     `json:"extensionType"`                  // 佣金比例 0.系统，1.自定义
	Refusal       string  `json:"refusal"`                        // 审核拒绝理由
	Rate          float64 `json:"rate"`                           // 评价分数
	ReplyCount    uint    `json:"replyCount"`                     // 评论数
	GiveCouponIDs string  `json:"giveCouponIds"`                  // 赠送优惠券
	IsGiftBag     int     `json:"isGiftBag"`                      // 是否为礼包
	CareCount     int     `json:"careCount"`                      // 收藏数

	Image       string `json:"image"`       // 商品图片
	SliderImage string `json:"sliderImage"` // 轮播图

	OldID             int `json:"oldId"`                                // 原商品ID
	TempID            int `json:"tempId"`                               // 运费模板ID
	SysTenancyID      int `json:"sysTenancyId"`                         // 商户 id
	SysBrandID        int `binding:"required" json:"sysBrandId"`        // 品牌 id
	TenancyCategoryID int `binding:"required" json:"tenancyCategoryId"` // 分类id
}
