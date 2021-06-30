package source

import (
	"time"

	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var TenancyProduct = new(product)

type product struct{}

// 出售中 1: is_show' => 1, 'status' => 1
// 仓库中 2:'is_show' => 2, 'status' => 1
// 3,4,5 商户才有
// 已售罄 3:'is_show' => 1, 'stock' => 0, 'status' => 1
// 警戒库存 4:'stock' => $stock ? $stock : 0, 'status' => 1
// 回收站 5:'deleted_at' => not null
// 待审核 6:'status' => 2
// 审核未通过 7:'status' => 3
var products = []model.TenancyProduct{
	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "领立裁腰带短袖连衣裙", StoreInfo: "短袖连衣裙", Keyword: "连衣裙", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 1, Price: 80.00, Cost: 50.00, OtPrice: 100.00, Stock: 399, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 1, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "", SliderImage: ""}},

	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "纯棉珠地撞色领polo裙", StoreInfo: "polo裙", Keyword: "polo裙", BarCode: "", SysBrandID: 2, IsShow: 2, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 1, Price: 160.00, Cost: 50.00, OtPrice: 180.00, Stock: 99, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "", SliderImage: ""}},

	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "精梳棉修身短袖T恤（圆/V领）", StoreInfo: "精梳", Keyword: "T恤", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 40, Rank: 0, Sales: 2, Price: 40.00, Cost: 20.00, OtPrice: 58.00, Stock: 0, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 102, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "", SliderImage: ""}},

	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "素湃黑科技纯棉疏水抗污短袖T恤", StoreInfo: "黑科技", Keyword: "T恤", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 2, ProductCategoryID: 162, UnitName: "件", Sort: 0, Rank: 0, Sales: 1, Price: 80.00, Cost: 60.00, OtPrice: 100.00, Stock: 99, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "", SliderImage: ""}},

	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "智能定制休闲单西 布雷泽海军蓝轻薄斜纹", StoreInfo: "西装定制", Keyword: "西装", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 3, ProductCategoryID: 162, UnitName: "件", Sort: 70, Rank: 0, Sales: 3, Price: 880.00, Cost: 500.00, OtPrice: 1680.00, Stock: 97, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 99, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 0, OldID: 0, Image: "", SliderImage: ""}},

	{TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now(), Valid: true}}, BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "梅湾街复古雪纺翻领上衣", StoreInfo: "雪纺", Keyword: "上衣", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 162, UnitName: "件", Sort: 56, Rank: 0, Sales: 1, Price: 88.00, Cost: 100.00, OtPrice: 200.00, Stock: 134, IsHot: 2, IsBenefit: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 96, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 4, OldID: 0, Image: "", SliderImage: ""}},
	// 8	64	女式纯棉条纹单兜衬衫	衬衫	衬衫		125	1	1	0	1	247	件	65	0	2	88.00	100.00	150.00	98	0	0	0	0	1	0	100	0			96	0	0		5.0	0	4,5	0	2020-07-09 20:21:25	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/f750c202007092020194655.jpg
	// 9	64	伯希和男款短袖T恤（情侣款）	短袖	短袖		125	1	1	0	1	172	件	0	0	2	100.00	50.00	150.00	98	0	0	0	0	1	0	100	0			96	0	0		5.0	0	4,5	0	2020-07-09 20:23:19	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/192e7202007092022294164.jpg
	// 10	55	腾讯黑鲨游戏手机3 Pro	小米黑鲨	黑鲨		125	1	-2	0	1	212	件	80	0	0	4988.00	4700.00	4999.00	20	0	0	0	0	1	0	100	0			42	0	0	1	5.0	0	6,7	0	2020-07-09 20:30:52	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/ecce9202007092030092652.jpg
	// 11	55	Redmi K30 Pro 变焦版	Redmi K30 Pro 变焦版	红米		125	1	1	0	1	212	个	70	0	0	3599.00	3299.00	3999.00	25	1	1	1	1	1	0	100	0			44	0	0		5.0	0	6,7	0	2020-07-09 20:32:59	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/3a1e7202007092031563735.jpg
	// 12	55	多亲AI助手/4G电话 内置小爱同学	备用机	老年机		125	1	1	0	1	213	件	10	0	0	699.00	799.00	799.00	100	0	0	0	0	1	0	100	0			48	0	0		5.0	0	6,7	0	2020-07-09 20:34:58	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/891ba202007092033471559.jpg
	// 13	55	佳能 PowerShot SX620 HS 数码相机 720	佳能	佳能		125	1	1	0	1	213	件	23	0	0	1999.00	1599.00	2299.00	102	0	0	0	0	0	0	100	0			42	0	0		5.0	0	6,7	0	2020-07-09 20:36:30	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/185c0202007092035275786.jpg
	// 15	55	小米MIX Alpha	小米MIX Alpha	小米		125	1	1	0	1	211	个	0	0	0	19999.00	159999.00	299999.00	10	0	0	0	0	0	0	100	0			44	0	0		5.0	0	6,7	0	2020-07-09 21:39:02	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/1204e202007092138295853.png
	// 16	66	大白菜	大白菜。好吃有可爱	白菜		118	1	0	0	1	165	kg	0	0	0	100.00	101.00	103.00	10000	0	0	0	0	1	0	100	0			100	0	1		5.0	0		1	2020-07-09 21:40:52	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/6bdcb202007092137241306.jpg
	// 17	66	354345353	543	5343		116	1	0	0	1	165	3	0	0	0	5.00	5.00	5.00	100	0	0	0	0	1	0	100	0			100	0	1		5.0	0		1	2020-07-09 21:49:45	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/6bdcb202007092137241306.jpg,http://cremb.oss-cn-beijing.aliyuncs.com/39c09202007092136174520.jpg
	// 18	66	多商户测试产品	多商户测试产品	测试产品		125	0	0	0	1	245	个	0	0	2	0.10	0.10	100.00	98	0	0	0	0	1	0	100	0			100	0	0		5.0	1	9,12	0	2020-07-09 21:57:45	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/f4b24202007092157105883.jpg
	// 19	66	测试商品1号				121	1	0	0	1	165	瓶	0	0	0	999.00		9999.00	299997	0	0	0	0	0	0	100	0			100	1	0		5.0	0	9,12	0	2020-07-09 21:58:16	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/5aa28202007092152419557.jpg
	// 20	66	111	11111	1111		122	0	0	0	1	172	坨	0	0	0	11.00	10.00	9.00	600	0	0	0	0	1	0	100	0			100	1	0		5.0	0	9,12	0	2020-07-09 21:59:11	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/f4b24202007092157105883.jpg
	// 21	64	保罗短袖polo衫t恤男 衬衫领2020新款夏季韩版潮流				121	0	0	0	1	172	件	0	0	0	388.00	288.00	450.00	999	0	0	0	0	1	0	100	0			96	0	0		5.0	0	4,5,4,5	0	2020-07-10 09:14:10	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/278e9202007100912266462.jpg,http://cremb.oss-cn-beijing.aliyuncs.com/9d3e3202007100912265638.jpg,http://cremb.oss-cn-beijing.aliyuncs.com/1f0f7202007100912263207.jpg,http://cremb.oss-cn-beijing.aliyuncs.com/28c50202007100912258175.jpg
	// 22	66	111				120	1	0	0	1	165	个	0	0	0	0.00	0.00	0.00	0	0	0	0	0	0	0	100	0			100	1	0		5.0	0	13	0	2020-07-10 18:00:15	0	1	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	http://cremb.oss-cn-beijing.aliyuncs.com/18918202007101054526973.jpg
	{BaseTenancyProduct: model.BaseTenancyProduct{SysTenancyID: 1, StoreName: "天丝一字领露肩上衣", StoreInfo: "天丝一字领露肩上衣", Keyword: "测试", BarCode: "", SysBrandID: 2, IsShow: 1, Status: 1, ProductCategoryID: 1, UnitName: "件", Sort: 10, Rank: 0, Sales: 1, Price: 50.00, Cost: 20.00, OtPrice: 100.00, Stock: 99, IsHot: 2, IsBest: 2, IsNew: 2, IsGood: 1, ProductType: 2, Ficti: 100, Browse: 0, CodePath: "", VideoLink: "", TempID: 102, SpecType: 2, ExtensionType: 2, Refusal: "", Rate: 5.0, ReplyCount: 0, GiveCouponIDs: "", IsGiftBag: 0, CareCount: 1, OldID: 0, Image: "", SliderImage: ""}, TENANCY_MODEL: g.TENANCY_MODEL{DeletedAt: gorm.DeletedAt{Time: time.Now()}}},
	// 191	55	简约时尚背包	简约时尚背包			125	1	1	0	1	348	个	0	0	0	99.00	50.00	129.00	10000	0	0	0	0	1	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:32:18	1	1	188	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg,http://mer.crmeb.net/uploads/def/20200912/3085e68f769e87a9d1d0ef308db70b8c.jpg,http://mer.crmeb.net/uploads/def/20200912/1809423b9139d4a993baef26735a79e1.jpg
	// 192	55	行李箱	行李箱			116	1	1	0	1	349	个	10	0	0	599.00	300.00	799.00	10000	0	0	0	0	0	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:33:45	0	1	190	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg,http://mer.crmeb.net/uploads/def/20200912/091e9b3cfa4d343937eede03d296ac58.jpg
	// 193	55	保温杯	保温杯			118	1	1	0	1	205	个	5	0	0	199.00	99.00	229.00	10000	0	0	0	0	0	1	100	0			49	0	0		5.0	0		0	2020-09-12 15:34:50	1	1	189	http://mer.crmeb.net/uploads/def/20200912/a89d6432f9424e76c673325885ab4706.jpg	http://mer.crmeb.net/uploads/def/20200912/c13462b94696adbfae14329b3e5aac6d.png,http://mer.crmeb.net/uploads/def/20200912/5c388b8c13eece7347baa7acfcc4b40f.png,http://mer.crmeb.net/uploads/def/20200912/a89d6432f9424e76c673325885ab4706.jpg
	// 194	55	智能音箱	智能音箱			118	1	1	0	1	211	个	10	0	0	299.00	199.00	399.00	994	0	0	0	0	0	1	100	0			42	0	0		5.0	0		0	2020-09-12 15:43:56	0	1	120	http://mer.crmeb.net/uploads/def/20200829/b5909ffa551317a0d92453114acb3d98.jpg	http://mer.crmeb.net/uploads/def/20200829/fd3a5f3f2ffe8e85a2761f979af63cb3.jpg,http://mer.crmeb.net/uploads/def/20200829/b5909ffa551317a0d92453114acb3d98.jpg
	// 195	55	简约时尚背包	简约时尚背包			125	1	1	0	1	348	个	0	0	0	99.00	50.00	129.00	10000	0	0	0	0	1	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:44:20	0	1	188	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg,http://mer.crmeb.net/uploads/def/20200912/3085e68f769e87a9d1d0ef308db70b8c.jpg,http://mer.crmeb.net/uploads/def/20200912/1809423b9139d4a993baef26735a79e1.jpg
	// 196	55	行李箱	行李箱			116	1	1	0	1	349	个	10	0	0	599.00	300.00	799.00	10000	0	0	0	0	0	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:46:08	0	1	190	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg,http://mer.crmeb.net/uploads/def/20200912/091e9b3cfa4d343937eede03d296ac58.jpg
	// 197	55	保温杯	保温杯			118	1	1	0	1	205	个	5	0	0	199.00	99.00	229.00	10000	0	0	0	0	0	1	100	0			49	0	0		5.0	0		0	2020-09-12 15:46:40	0	1	189	http://mer.crmeb.net/uploads/def/20200912/a89d6432f9424e76c673325885ab4706.jpg	http://mer.crmeb.net/uploads/def/20200912/c13462b94696adbfae14329b3e5aac6d.png,http://mer.crmeb.net/uploads/def/20200912/5c388b8c13eece7347baa7acfcc4b40f.png,http://mer.crmeb.net/uploads/def/20200912/a89d6432f9424e76c673325885ab4706.jpg
	// 198	55	简约时尚背包	简约时尚背包			125	1	1	0	1	348	个	0	0	0	99.00	50.00	129.00	10000	0	0	0	0	1	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:47:12	0	1	188	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg,http://mer.crmeb.net/uploads/def/20200912/3085e68f769e87a9d1d0ef308db70b8c.jpg,http://mer.crmeb.net/uploads/def/20200912/1809423b9139d4a993baef26735a79e1.jpg
	// 199	55	行李箱	行李箱			116	1	1	0	1	349	个	10	0	0	599.00	300.00	799.00	10000	0	0	0	0	0	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:47:43	0	1	190	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg,http://mer.crmeb.net/uploads/def/20200912/091e9b3cfa4d343937eede03d296ac58.jpg
	// 200	55	智能音箱	音箱			118	1	1	0	1	212	个	100	0	0	299.00	199.00	399.00	1000	0	0	0	0	1	1	100	0			49	0	0		5.0	0		0	2020-09-12 15:49:15	0	1	145	http://mer.crmeb.net/uploads/def/20200829/fd3a5f3f2ffe8e85a2761f979af63cb3.jpg	http://mer.crmeb.net/uploads/def/20200829/fd3a5f3f2ffe8e85a2761f979af63cb3.jpg,http://mer.crmeb.net/uploads/def/20200829/17818707ee11472f2d84a6830ce05763.png
	// 201	55	简约时尚背包	简约时尚背包			125	1	1	0	1	348	个	0	0	0	99.00	50.00	129.00	10000	0	0	0	0	1	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:49:37	0	1	188	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg	http://mer.crmeb.net/uploads/def/20200912/c7a0673c0750cc55db1ff23bafb9a758.jpg,http://mer.crmeb.net/uploads/def/20200912/3085e68f769e87a9d1d0ef308db70b8c.jpg,http://mer.crmeb.net/uploads/def/20200912/1809423b9139d4a993baef26735a79e1.jpg
	// 202	55	行李箱	行李箱			116	1	1	0	1	349	个	10	0	0	599.00	300.00	799.00	10000	0	0	0	0	0	1	100	0			48	0	0		5.0	0		0	2020-09-12 15:49:58	0	1	190	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg	http://mer.crmeb.net/uploads/def/20200912/f38e0d2fe50db1e764e6400e3984dce5.jpg,http://mer.crmeb.net/uploads/def/20200912/091e9b3cfa4d343937eede03d296ac58.jpg
}

func (m *product) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.TenancyProduct{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_products 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&products).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_products 表初始数据成功!")
		return nil
	})
}
