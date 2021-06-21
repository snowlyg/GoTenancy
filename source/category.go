package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Category = new(category)

type category struct{}

var categories = []model.ProductCategory{
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 173}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "品牌服饰", Path: "/", Sort: 2, Status: g.StatusTrue, Level: 0}, SysTenancyID: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 179}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "美容美发", Path: "/", Sort: 10, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 180}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "生鲜食品", Path: "/", Sort: 1, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 181}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "美容彩妆", Path: "/", Sort: 4, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 182}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "母婴专区", Path: "/", Sort: 2, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 183}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "食品饮料", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 184}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "数码家电", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 185}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "营养保健", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 186}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "精品服装", Path: "/", Sort: 20, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 187}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "鲜花预定", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 188}, BaseProductCategory: model.BaseProductCategory{Pid: 0, CateName: "教育培训", Path: "/", Sort: 0, Pic: "", Status: g.StatusTrue, Level: 0}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 162}, BaseProductCategory: model.BaseProductCategory{Pid: 186, CateName: "男士上衣", Path: "/186/", Sort: 9, Pic: "http://qmplusimg.henrongyi.top/head.png", Level: 1, Status: g.StatusTrue}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 165}, BaseProductCategory: model.BaseProductCategory{Pid: 171, CateName: "祛斑祛痘", Path: "/181/171/", Sort: 0, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 171}, BaseProductCategory: model.BaseProductCategory{Pid: 181, CateName: "普通化妆", Path: "/181/", Sort: 0, Status: g.StatusTrue, Level: 1}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 172}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "T恤/POLO", Path: "/186/162/", Sort: 9, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 174}, BaseProductCategory: model.BaseProductCategory{Pid: 173, CateName: "时尚女装", Path: "/173/", Sort: 0, Status: g.StatusTrue, Level: 1}, SysTenancyID: 1},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 175}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "商务衬衫", Path: "/186/162/", Sort: 6, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 176}, BaseProductCategory: model.BaseProductCategory{Pid: 162, CateName: "休闲衬衫", Path: "/186/162/", Sort: 6, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	{TENANCY_MODEL: g.TENANCY_MODEL{ID: 177}, BaseProductCategory: model.BaseProductCategory{Pid: 171, CateName: "补水保湿", Path: "	/181/171/", Sort: 0, Pic: "http://qmplusimg.henrongyi.top/head.png", Status: g.StatusTrue, Level: 2}, SysTenancyID: 0},
	// 178	171	防晒保湿	/181/171/	0	http://mer.crmeb.net/uploads/def/20200816/7021ecaa473f2497d5cf86f50f341cef.jpg	1	2	0	2020-08-16 02:28:29

	// 189	188	网上课程	/188/	0	http://mer.crmeb.net/uploads/def/20200816/077f8b9ec5f31cdfd8f948fb1a4d015b.jpg	1	1	0	2020-09-08 17:09:56
	// 190	188	线下课程	/188/	0	http://mer.crmeb.net/uploads/def/20200816/761aaa9878d571bfa6c4a62b214b7d08.jpg	1	1	0	2020-09-08 17:09:57
	// 191	189	相关书籍	/188/189/	0	http://mer.crmeb.net/uploads/def/20200816/3208fff8f6145ca1ca55d7424b1acf01.jpg	1	2	0	2020-08-16 02:18:36
	// 192	189	视频教程	/188/189/	0	http://mer.crmeb.net/uploads/def/20200816/3870d7bc52833f7440906a85f06f9a18.png	1	2	0	2020-08-16 02:18:45
	// 193	179	实体店铺	/179/	0		1	1	0	2020-09-08 17:09:34
	// 194	236	西安周边	/179/236/	0	http://mer.crmeb.net/uploads/def/20200816/583665401c35d066afc23b1268cc8014.jpg	1	2	0	2020-08-16 02:26:13
	// 195	236	北京周边	/179/236/	0	http://mer.crmeb.net/uploads/def/20200816/ea808dda99c4f9c56ca20ae5cab33639.jpg	1	2	0	2020-08-16 02:26:22
	// 196	236	上海周边	/179/236/	0	http://mer.crmeb.net/uploads/def/20200816/3e0a80efb31d583a4b9339adadab78fa.jpg	1	2	0	2020-08-16 02:26:35
	// 197	189	在线教学	/188/189/	0	http://mer.crmeb.net/uploads/def/20200816/ccfc669c9ee88d18f5fdc9959188d732.jpg	1	2	0	2020-08-16 02:18:54
	// 198	180	新鲜蔬菜	/180/	0		1	1	0	2020-09-08 17:09:44
	// 199	198	新鲜辣椒	/180/198/	0	http://mer.crmeb.net/uploads/def/20200816/08295c5aa6fae9e11f356c1e23061183.png	1	2	0	2020-08-16 02:31:30
	// 200	198	新鲜青菜	/180/198/	0	http://mer.crmeb.net/uploads/def/20200816/19026c7843b912b5c43b24316a9e230a.png	1	2	0	2020-08-16 02:31:36
	// 201	198	新鲜茄子	/180/198/	0	http://mer.crmeb.net/uploads/def/20200816/2c85030716c713491b3e41e0a5332f9e.png	1	2	0	2020-08-16 02:31:43
	// 202	182	母婴产品	/182/	0		1	1	0	2020-09-08 17:09:41
	// 203	202	优质奶粉	/182/202/	0	http://mer.crmeb.net/uploads/def/20200816/e67845c5cf12c018a21de4459f20ae0b.jpg	1	2	0	2020-08-16 02:28:41
	// 204	202	高档奶嘴	/182/202/	0	http://mer.crmeb.net/uploads/def/20200816/681f21be4be51e986819d4bb59b62dac.jpg	1	2	0	2020-08-16 02:29:40
	// 205	202	精品奶壶	/182/202/	0	http://mer.crmeb.net/uploads/def/20200816/b930aa9f7f1fe86457edfbb75b9da597.jpg	1	2	0	2020-08-16 02:29:54
	// 206	183	碳酸饮料	/183/	0		1	1	0	2020-09-08 17:09:48
	// 207	206	可口可乐	/183/206/	0	http://mer.crmeb.net/uploads/def/20200816/cc819f999c4f6b081e5ab19b418eea44.jpg	1	2	0	2020-08-16 02:31:50
	// 208	206	大瓶雪碧	/183/206/	0	http://mer.crmeb.net/uploads/def/20200816/65fff6aa36ccb22cb57412418f004198.jpg	1	2	0	2020-08-16 02:31:56
	// 209	206	健力宝	/183/206/	0	http://mer.crmeb.net/uploads/def/20200816/cb79a38d07e822ce5404b557ac8acd84.jpg	1	2	0	2020-08-16 02:32:03
	// 210	184	家用电器	/184/	0		1	1	0	2020-07-03 17:50:22
	// 211	210	液晶电视	/184/210/	0	http://mer.crmeb.net/uploads/def/20200816/3993e562cb9340713b6c0b28c5a8a33e.jpg	1	2	0	2020-08-16 02:32:18
	// 212	210	洗衣机	/184/210/	0	http://mer.crmeb.net/uploads/def/20200816/30270694cd6e44cba62a22970b9a09b6.jpg	1	2	0	2020-08-16 02:32:28
	// 213	210	电冰箱	/184/210/	0	http://mer.crmeb.net/uploads/def/20200816/06cc4e50c126261dd71a8dbbe0ee5ba0.jpg	1	2	0	2020-08-16 02:32:38
	// 214	185	营养食品	/185/	0		1	1	0	2020-09-08 17:09:51
	// 215	214	东北人参	/185/214/	0	http://mer.crmeb.net/uploads/def/20200816/ff171a16dfcb7c66a6caf38be98335e5.jpg	1	2	0	2020-08-16 02:32:50
	// 216	214	百年灵芝	/185/214/	0	http://mer.crmeb.net/uploads/def/20200816/097578db9425d41eed48fbfd0b28bb67.jpg	1	2	0	2020-08-16 02:33:02
	// 217	214	土鸡蛋	/185/214/	0	http://mer.crmeb.net/uploads/def/20200816/509c5c8395e9978b04245b0759ffde8f.jpg	1	2	0	2020-08-16 02:33:16
	// 218	187	各类鲜花	/187/	0		1	1	0	2020-09-08 17:09:53
	// 219	218	玫瑰花	/187/218/	0	http://mer.crmeb.net/uploads/def/20200816/f8fb5ff8bf25f86494cf246efedd2d7e.jpg	1	2	0	2020-08-16 02:33:32
	// 220	218	百合花	/187/218/	0	http://mer.crmeb.net/uploads/def/20200816/6d43b54cead66fb265ec891cf57c84e3.jpg	1	2	0	2020-08-16 02:34:49
	// 221	218	仙人掌	/187/218/	0	http://mer.crmeb.net/uploads/def/20200816/92d2d477a247a449342874e212375fee.jpg	1	2	0	2020-08-16 02:35:37
	// 222	0	手机	/	0		1	0	56	2020-07-04 09:34:02
	// 223	222	老人机	/222/	0		1	1	56	2020-07-04 09:34:11
	// 224	222	智能机	/222/	0		1	1	56	2020-07-04 09:34:19
	// 225	190	线下书籍	/188/190/	0	http://mer.crmeb.net/uploads/def/20200816/ccfc669c9ee88d18f5fdc9959188d732.jpg	1	2	0	2020-08-16 02:19:24
	// 226	190	线下课堂	/188/190/	0	http://mer.crmeb.net/uploads/def/20200816/8b6c719f1ec70dd00db3d59f7ca844a2.jpg	1	2	0	2020-08-16 02:19:43
	// 227	190	视频教材	/188/190/	0	http://mer.crmeb.net/uploads/def/20200816/8b6c719f1ec70dd00db3d59f7ca844a2.jpg	1	2	0	2020-08-16 02:19:59
	// 228	237	兰州周边	/179/237/	0	http://mer.crmeb.net/uploads/def/20200816/4815ba391c92f89a23d80147b7c87ae0.jpg	1	2	0	2020-08-16 02:26:49
	// 229	237	榆林周边	/179/237/	0	http://mer.crmeb.net/uploads/def/20200816/d0cbc7bc58992a477d6526327ba3725c.jpg	1	2	0	2020-08-16 02:27:10
	// 230	237	咸阳周边	/179/237/	0	http://mer.crmeb.net/uploads/def/20200816/20f9e468f08fea2953eecba133b41287.jpg	1	2	0	2020-08-16 02:27:26
	// 231	193	宝鸡周边	/179/193/	0	http://mer.crmeb.net/uploads/def/20200816/20f9e468f08fea2953eecba133b41287.jpg	1	2	0	2020-08-16 02:25:17
	// 232	193	天水周边	/179/193/	0	http://mer.crmeb.net/uploads/def/20200816/073518144d9b75fae89a9e24341087b0.jpg	1	2	0	2020-08-16 02:25:38
	// 233	193	陇西周边	/179/193/	0	http://mer.crmeb.net/uploads/def/20200816/7b5a0d536bcf21b7ee4791e7503e9131.jpg	1	2	0	2020-08-16 02:25:47
	// 236	179	周边店铺	/179/	0		1	1	0	2020-09-08 17:09:35
	// 237	179	品牌美发	/179/	0		1	1	0	2020-09-08 17:09:36
	// 238	186	男士裤装	/186/	8		1	1	0	2020-07-09 17:00:07
	// 239	186	女装	/186/	7		1	1	0	2020-07-09 17:00:12
	// 240	186	首饰	/186/	6		1	1	0	2020-07-09 17:00:16
	// 241	186	配饰	/186/	5		1	1	0	2020-07-09 17:00:49
	// 244	238	牛仔裤	/186/238/	8	http://mer.crmeb.net/uploads/def/20200816/a4640080bc06d5e36fdbce11f5f4b37c.jpg	1	2	0	2020-08-16 02:23:15
	// 245	238	短裤	/186/238/	7	http://mer.crmeb.net/uploads/def/20200816/12586427d1132e6fb8f5d6ee02633f39.jpg	1	2	0	2020-08-16 02:23:23
	// 247	239	T恤/卫衣	/186/239/	9	http://mer.crmeb.net/uploads/def/20200816/adbd7f096553912f6cf7dc9757bcf550.jpg	1	2	0	2020-08-16 02:38:28
	// 248	239	毛衫/针织衫	/186/239/	8	http://mer.crmeb.net/uploads/def/20200816/0c3cf51ec94a35a6e01e1def682d6ec8.jpg	1	2	0	2020-08-16 02:39:39
	// 249	239	衬衫	/186/239/	7	http://mer.crmeb.net/uploads/def/20200816/43c72a7ea5c887965c076cf4356a4af9.jpg	1	2	0	2020-08-16 02:38:42
	// 250	241	皮带	/186/241/	9	http://mer.crmeb.net/uploads/def/20200816/d6d8f024c73a97966b0b998bb1dea8e6.jpg	1	2	0	2020-08-16 02:47:33
	// 251	241	帽子	/186/241/	8	http://mer.crmeb.net/uploads/def/20200816/4a4c64f66badd064474f30bcbdd7a4f7.jpg	1	2	0	2020-08-16 02:47:39
	// 252	241	围巾手套	/186/241/	7	http://mer.crmeb.net/uploads/def/20200816/488e66cde10731d76155b1d573476e12.jpg	1	2	0	2020-08-16 02:47:46
	// 253	240	项链	/186/240/	9	http://mer.crmeb.net/uploads/def/20200816/dd5acd0ecd2af873ea1959c4423c0380.jpg	1	2	0	2020-08-16 02:41:32
	// 254	240	戒指	/186/240/	8	http://mer.crmeb.net/uploads/def/20200816/cb02dc55f71956cb740417543c9aa342.jpg	1	2	0	2020-08-16 02:41:38
	// 255	240	耳饰	/186/240/	7	http://mer.crmeb.net/uploads/def/20200816/4a8804ad2180e4a57086962bb8cc5ca6.jpg	1	2	0	2020-08-16 02:41:45
	// 256	179	高级美容	/179/	0		1	1	0	2020-09-08 17:09:33
	// 257	256	洗面奶	/179/256/	0	http://mer.crmeb.net/uploads/def/20200816/2cd0c0cafdfb359bb9875815f8e70cdd.jpg	1	2	0	2020-08-16 02:24:16
	// 258	256	美白霜	/179/256/	0	http://mer.crmeb.net/uploads/def/20200816/4df411dd11a169c4d9d655dc6eb9e532.jpg	1	2	0	2020-08-16 02:24:52
	// 259	256	防晒霜	/179/256/	0	http://mer.crmeb.net/uploads/def/20200816/7021ecaa473f2497d5cf86f50f341cef.jpg	1	2	0	2020-08-16 02:25:05
	// 260	238	休闲裤	/186/238/	9	http://mer.crmeb.net/uploads/def/20200816/572ac105d3637cf3809542c784adc042.jpg	1	2	0	2020-08-16 02:23:06
	// 261	0	服饰	/	0	http://cremb.oss-cn-beijing.aliyuncs.com/868ec202007131729001746.jpg	1	0	65	2020-07-15 09:12:32
	// 262	261	女装	/261/	0	http://cremb.oss-cn-beijing.aliyuncs.com/66de1202007081930314002.jpg	1	1	65	2020-07-09 18:27:59
	// 263	0	饰品	/	0	http://cremb.oss-cn-beijing.aliyuncs.com/c099c202007081709595845.jpg	1	0	65	2020-07-09 18:28:54
	// 264	186	反季清仓	/186/	4		1	1	0	2020-07-09 17:01:24
	// 265	162	夹克/风衣	/186/162/	8	http://mer.crmeb.net/uploads/def/20200816/5e32bed2a18ad172b0035d43d0100cdd.jpg	1	2	0	2020-08-16 02:21:45
	// 266	162	西服西裤	/186/162/	5	http://mer.crmeb.net/uploads/def/20200816/f3aee4dab8b638a6012830cd3fdf264f.jpg	1	2	0	2020-08-16 02:22:40
	// 267	162	卫衣	/186/162/	7	, Pic: "http://qmplusimg.henrongyi.top/head.png"	1	2	0	2020-08-16 02:22:03
	// 268	239	西服/夹克	/186/239/	6	http://mer.crmeb.net/uploads/def/20200816/eb2ab9a853ecaa08016fe07d0b36e4ec.jpg	1	2	0	2020-08-16 02:38:49
	// 269	239	连衣裙	/186/239/	5	http://mer.crmeb.net/uploads/def/20200816/adc01b6b827892c09c1dc00b260dbeaa.jpg	1	2	0	2020-08-16 02:38:57
	// 270	239	 半身裙	/186/239/	4	http://mer.crmeb.net/uploads/def/20200816/8775b0d085e91d20bc581ea0dacee9e1.jpg	1	2	0	2020-08-16 02:39:04
	// 271	240	手链	/186/240/	5	http://mer.crmeb.net/uploads/def/20200816/d9ea1ada6cdc7374986fcccb898d620c.jpg	1	2	0	2020-08-16 02:41:52
	// 272	240	轻定制	/186/240/	4	http://mer.crmeb.net/uploads/def/20200816/d6d8f024c73a97966b0b998bb1dea8e6.jpg	1	2	0	2020-08-16 02:41:59
	// 273	264	羽绒服/棉服	/186/264/	9	http://mer.crmeb.net/uploads/def/20200816/f6d59b5ee38f47a29ec78ac5e37f383c.jpg	1	2	0	2020-08-16 02:49:36
	// 274	264	毛衫/针织衫	/186/264/	8	http://mer.crmeb.net/uploads/def/20200816/2b55a4c7f82de796eaab561624125a75.jpg	1	2	0	2020-08-16 02:49:42
	// 275	261	男装	/261/	0	http://cremb.oss-cn-beijing.aliyuncs.com/33935202007081901059221.jpg	1	1	65	2020-07-09 18:28:20
	// 276	263	手表	/263/	0	http://cremb.oss-cn-beijing.aliyuncs.com/c099c202007081709595845.jpg	1	1	65	2020-07-09 18:29:07
	// 277	173	大牌童装	/173/	0		1	1	64	2020-07-09 20:12:41
	// 278	173	亲子乐园	/173/	0		1	1	64	2020-07-09 20:12:53
	// 279	173	商务男装	/173/	0		1	1	64	2020-07-09 20:13:11
	// 280	0	应季夏装	/	0		1	0	64	2020-07-09 20:13:46
	// 281	280	雪纺	/280/	0		1	1	64	2020-07-09 20:14:08
	// 282	280	真丝	/280/	0		1	1	64	2020-07-09 20:14:22
	// 283	0	手机	/	0		1	0	55	2020-07-09 20:27:06
	// 284	283	游戏手机	/283/	0		1	1	55	2020-07-09 20:27:13
	// 285	283	旗舰机	/283/	0		1	1	55	2020-07-09 20:27:26
	// 286	283	拍照手机	/283/	0		1	1	55	2020-07-09 20:27:33
	// 287	283	音乐手机	/283/	0		1	1	55	2020-07-09 20:27:53
	// 288	0	数码产品	/	0		1	0	55	2020-07-09 20:28:00
	// 289	288	电脑	/288/	0		1	1	55	2020-07-09 20:28:07
	// 290	288	数码周边	/288/	0		1	1	55	2020-07-09 20:28:15
	// 291	288	耳机	/288/	0		1	1	55	2020-07-09 20:28:24
	// 292	288	充电器	/288/	0		1	1	55	2020-07-09 20:28:30
	// 293	0	1	/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	0	66	2020-08-20 09:58:25
	// 294	0	生鲜果蔬	/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	0	66	2020-08-20 09:58:25
	// 295	294	111	/294/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	1	66	2020-08-20 09:58:25
	// 296	0	一级分类	/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	0	66	2020-08-20 09:58:25
	// 297	296	二级分类	/296/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	1	66	2020-08-20 09:58:25
	// 298	0	在线教育	/	1	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	0	66	2020-08-20 09:58:25
	// 299	0	鞋包	/	0	http://cremb.oss-cn-beijing.aliyuncs.com/1446d202007092004219739.jpg	1	0	65	2020-07-15 09:12:53
	// 300	240	珠宝首饰	/186/240/	0	http://mer.crmeb.net/uploads/def/20200816/9a6a2e1231fb19517ed1de71206a0657.jpg	1	2	0	2020-08-20 09:58:25
}

func (m *category) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.ProductCategory{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> categories 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&categories).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> categories 表初始数据成功!")
		return nil
	})
}
