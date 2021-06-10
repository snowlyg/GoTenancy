package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var Config = new(config)

type config struct{}

var configs = []model.SysConfig{
	{SysConfigCategoryID: 2, ConfigName: "网站域名", ConfigKey: "site_url", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 2, ConfigName: "网站名称", ConfigKey: "site_name", ConfigType: "input", ConfigRule: "", Required: 1, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 2, ConfigName: "网站开启", ConfigKey: "site_open", ConfigType: "radio", ConfigRule: "0:关闭1:开启", Required: 1, Info: "", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号名称", ConfigKey: "wechat_name", ConfigType: "input", ConfigRule: "", Required: 2, Info: "设置公众号名称", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "联系电话", ConfigKey: "set_phone", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 5, ConfigName: "联系邮箱", ConfigKey: "set_email", ConfigType: "input", ConfigRule: "", Required: 2, Info: "", Sort: 0, UserType: 1, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信号", ConfigKey: "wechat_id", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号原始id", ConfigKey: "wechat_sourceid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号原始id", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号AppID", ConfigKey: "wechat_appid", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号AppID", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号AppSecret", ConfigKey: "wechat_appsecret", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号AppSecret", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信验证TOKEN", ConfigKey: "wechat_token", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信验证TOKEN", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信EncodingAESKey", ConfigKey: "wechat_encodingaeskey", ConfigType: "input", ConfigRule: "", Required: 2, Info: "公众号消息加解密Key,在使用安全模式情况下要填写该值，请先在管理中心修改，然后填写该值，仅限服务号和认证订阅号", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号二维码", ConfigKey: "wechat_qrcode", ConfigType: "image", ConfigRule: "", Required: 2, Info: "公众号二维码", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "公众号logo", ConfigKey: "wechat_avatar", ConfigType: "image", ConfigRule: "", Required: 2, Info: "公众号logo", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享图片", ConfigKey: "wechat_share_img", ConfigType: "image", ConfigRule: "", Required: 2, Info: "若填写此图片地址，则分享网页出去时会分享此图片。可有效防止分享图片变形", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享标题", ConfigKey: "wechat_share_title", ConfigType: "input", ConfigRule: "", Required: 2, Info: "微信分享标题", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "微信分享简介", ConfigKey: "wechat_share_synopsis", ConfigType: "textarea", ConfigRule: "", Required: 2, Info: "微信分享简介", Sort: 0, UserType: 2, Status: 1},
	{SysConfigCategoryID: 3, ConfigName: "消息加解密方式", ConfigKey: "wechat_encode", ConfigType: "radio", ConfigRule: "0:明文模式1:兼容模式2:安全模式", Required: 1, Info: "如需使用安全模式请在管理中心修改，仅限服务号和认证订阅号", Sort: 1, UserType: 2, Status: 1},
	// 194	0	一级分销比例	extension_one_rate	input		0	设置分销员一级分销比例, 最大为1(1=100%)	0	0	1	2020-05-08 15:38:03
	// 195	0	二级分销比例	extension_two_rate	input		0	设置分销员二级分销比例, 最大为1(1=100%)	0	0	1	2020-05-08 15:39:03
	// 196	0	开启分销	extension_status	radio	1:开启
	// 0:关闭	1	是否开启分销,关闭后不进行返佣	1	0	1	2020-05-08 15:41:04
	// 197	5	警戒库存	mer_store_stock	input		0	设置商品的警戒库存	0	1	1	2020-05-18 11:23:43
	// 198	0	短信平台账号	sms_account	input		0	设置短信平台账号	0	0	0	2020-05-18 15:33:58
	// 199	0	短信平台密码	sms_token	input		0	设置短信平台密码	0	0	0	2020-05-18 15:34:22
	// 208	6	Appid	pay_weixin_appid	input		0	微信公众号身份的唯一标识。审核通过后，在微信发送的邮件中查看。	0	0	1	2020-06-02 17:56:51
	// 209	6	Appsecret	pay_weixin_appsecret	input		0	JSAPI接口中获取openid，审核后在公众平台开启开发模式后可查看。	0	0	1	2020-06-02 17:57:23
	// 210	6	Mchid	pay_weixin_mchid	input		0	受理商ID，身份标识	0	0	1	2020-06-02 17:57:49
	// 211	6	微信支付证书	pay_weixin_client_cert	file		0	微信支付证书，在微信商家平台中可以下载！文件名一般为apiclient_cert.pem	0	0	1	2020-06-02 17:58:33
	// 212	6	微信支付证书密钥	pay_weixin_client_key	file		0	微信支付证书密钥，在微信商家平台中可以下载！文件名一般为apiclient_key.pem	0	0	1	2020-06-02 17:58:59
	// 213	6	Key	pay_weixin_key	input		0	商户支付密钥Key。审核通过后，在微信发送的邮件中查看。	0	0	1	2020-06-02 17:59:24
	// 214	6	开启	pay_weixin_open	radio	0:关闭
	// 1:开启	0	是否启用微信支付	0	0	1	2020-06-02 18:00:04
	// 215	12	充值注意事项	recharge_attention	textarea		0	充值注意事项	0	0	1	2020-06-03 09:49:01
	// 216	1	订单自动关闭时间	auto_close_order_timer	number		0	订单自动关闭时间(单位:分钟)	0	0	1	2020-06-03 09:49:01
	// 217	5	默认退货收货地址	mer_refund_address	input		0	设置默认退货收货地址	0	1	1	2020-06-12 16:28:41
	// 218	5	默认退货收货人	mer_refund_user	input		0	设置默认退货收货人	0	1	1	2020-06-12 16:29:15
	// 219	1	退款理由	refund_message	textarea		0	设置退款理由	0	0	1	2020-06-12 16:34:51
	// 220	1	商户自动处理退款订单期限（天）	mer_refund_order_agree	number		1	申请退款的订单超过期限，将自动退款处理。	0	0	1	2020-06-13 14:59:35
	// 221	5	银行卡开户行	bank	input		1		0	1	1	2020-06-15 10:56:52
	// 222	5	银行卡卡号	bank_number	input		1		0	1	1	2020-06-15 10:57:20
	// 223	5	银行卡持卡人姓名	bank_name	input		1		0	1	1	2020-06-15 10:57:50
	// 224	5	银行卡开户行地址	bank_address	input		1		0	1	1	2020-06-15 10:58:16
	// 225	13	佣金最低提现金额	user_extract_min	number	佣金达到最低额才可以提现	1		0	0	1	2020-06-16 14:53:52
	// 226	13	佣金冻结时间	lock_brokerage_timer	number		0	设置佣金冻结时间(天)	0	0	1	2020-06-17 14:36:16
	// 227	1	快递查询密钥	express_app_code	input		0	阿里云快递查询接口密钥购买地址：https://market.aliyun.com/products/56928004/cmapi021863.html	0	0	1	2020-06-17 18:32:16
	// 228	7	空间域名 Domain	uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 10:21:19
	// 229	7	accessKey	accessKey	input		0	accessKey	0	0	1	2020-06-18 10:21:37
	// 230	7	secretKey	secretKey	input		0	secretKey	0	0	1	2020-06-18 10:22:40
	// 231	7	存储空间名称	storage_name	input		0	存储空间名称	0	0	1	2020-06-18 10:22:57
	// 232	7	所属地域	storage_region	input		0	所属地域	0	0	1	2020-06-18 10:23:21
	// 233	0	上传类型	upload_type	radio	1:本地存储
	// 2:七牛云存储
	// 3:阿里云OSS
	// 4:腾讯COS	0	文件上传的类型	0	0	1	2020-06-18 14:46:29
	// 237	8	空间域名 Domain	qiniu_uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 15:14:29
	// 238	8	accessKey	qiniu_accessKey	input		0	accessKey	0	0	1	2020-06-18 15:14:51
	// 239	8	secretKey	qiniu_secretKey	input		0	secretKey	0	0	1	2020-06-18 15:15:05
	// 240	8	存储空间名称	qiniu_storage_name	input		0	存储空间名称	0	0	1	2020-06-18 15:15:22
	// 241	8	所属地域	qiniu_storage_region	input		0	所属地域	0	0	1	2020-06-18 15:15:40
	// 242	9	空间域名 Domain	tengxun_uploadUrl	input		0	空间域名 Domain	0	0	1	2020-06-18 15:40:03
	// 243	9	accessKey	tengxun_accessKey	input		0	accessKey	0	0	1	2020-06-18 15:40:32
	// 244	9	secretKey	tengxun_secretKey	input		0	secretKey	0	0	1	2020-06-18 15:40:48
	// 245	9	存储空间名称	tengxun_storage_name	input		0	存储空间名称	0	0	1	2020-06-18 15:41:06
	// 246	9	所属地域	tengxun_storage_region	input		0	所属地域	0	0	1	2020-06-18 15:41:39
	// 247	10	appId	routine_appId	input		0	appId	0	0	1	2020-06-19 10:07:31
	// 248	10	小程序AppSecret	routine_appsecret	input		0	小程序AppSecret	0	0	1	2020-06-19 10:09:32
	// 249	10	小程序授权logo	routine_logo	image		0	小程序授权logo	0	0	1	2020-06-19 10:10:36
	// 250	10	小程序名称	routine_name	input		0	小程序名称	0	0	1	2020-06-19 10:11:07
	// 252	11	Appid	pay_routine_appid	input		0	小程序Appid	0	0	1	2020-06-19 10:42:10
	// 253	11	Appsecret	pay_routine_appsecret	input		0	小程序Appsecret	0	0	1	2020-06-19 10:46:34
	// 254	11	Mchid	pay_routine_mchid	input		0	商户号	0	0	1	2020-06-19 10:46:59
	// 255	11	Key	pay_routine_key	input		0	商户key	0	0	1	2020-06-19 10:47:21
	// 256	11	小程序支付证书	pay_routine_client_cert	file		0	小程序支付证书	0	0	1	2020-06-19 10:47:46
	// 257	11	小程序支付证书密钥	pay_routine_client_key	file		0	小程序支付证书密钥	0	0	1	2020-06-19 10:48:02
	// 258	12	余额充值开关	recharge_switch	radio	1:开启
	// 0:关闭	0	余额充值开关	0	0	1	2020-06-19 15:47:12
	// 259	12	用户最低充值金额	store_user_min_recharge	number		0	用户最低充值金额	0	0	1	2020-06-19 15:47:40
	// 260	12	余额功能启用	balance_func_status	radio	1:开启
	// 0:关闭	0	商城余额功能启用或者关闭	0	0	1	2020-06-19 15:54:16
	// 261	12	余额支付状态	yue_pay_status	radio	1:开启
	// 0:关闭	0	余额支付状态	0	0	1	2020-06-19 16:02:16
	// 262	1	首页广告图	home_ad_pic	image		0	设置首页广告图(750*164px)	0	0	1	2020-06-25 10:05:06
	// 263	1	首页广告链接	home_ad_url	input		0	设置首页广告链接	0	0	1	2020-06-25 10:05:32
	// 264	13	分销说明	promoter_explain	textarea		0		0	0	1	2020-06-25 15:32:21
	// 265	13	商户设置礼包最大数量	max_bag_number	number		0		0	0	1	2020-06-25 15:48:27
	// 266	2	商城 logo	site_logo	image		0	设置商城logo(254*90px)	0	0	1	2020-06-26 11:41:23
	// 268	2	商城分享标题	share_title	input		0	商城分享标题	0	0	1	2020-06-30 16:16:30
	// 269	2	商城分享简介	share_info	input		0	商城分享简介	0	0	1	2020-06-30 16:16:57
	// 270	2	商城分享图片	share_pic	image		0	商城分享图片	0	0	1	2020-06-30 16:17:23
	// 271	4	发货提醒	sms_fahuo_status	radio	0:关闭
	// 1:开启	0	发货提醒	0	0	1	2020-07-01 11:00:13
	// 272	4	确认收货短信提醒	sms_take_status	radio	0:关闭
	// 1:开启	0	确认收货短信提醒	0	0	1	2020-07-01 11:20:51
	// 273	4	用户下单通知提醒	sms_pay_status	radio	0:关闭
	// 1:开启	0	用户下单通知提醒	0	0	1	2020-07-01 11:25:32
	// 274	4	改价提醒	sms_revision_status	radio	0:关闭
	// 1:开启	0	改价提醒	0	0	1	2020-07-01 11:30:04
	// 275	4	提醒付款通知	sms_pay_false_status	radio	0:关闭
	// 1:开启	0	提醒付款通知	0	0	1	2020-07-01 11:42:41
	// 276	4	商家拒绝退款提醒	sms_refund_fail_status	radio	0:关闭
	// 1:开启	0	商家拒绝退款提醒	0	0	1	2020-07-01 11:55:05
	// 277	4	商家同意退款提醒	sms_refund_success_status	radio	0:关闭
	// 1:开启	0	商家同意退款提醒	0	0	1	2020-07-01 12:01:32
	// 278	4	退款确认提醒	sms_refund_confirm_status	radio	0:关闭
	// 1:开启	0	退款确认提醒	0	0	1	2020-07-01 12:06:14
	// 279	4	管理员下单提醒	sms_admin_pay_status	radio	0:关闭
	// 1:开启	0	管理员下单提醒	0	0	1	2020-07-01 12:17:42
	// 280	4	管理员退货提醒	sms_admin_return_status	radio	0:关闭
	// 1:开启	0	管理员退货提醒	0	0	1	2020-07-01 12:31:13
	// 281	4	管理员确认收货提醒	sms_admin_take_status	radio	0:关闭
	// 1:开启	0	管理员确认收货提醒	0	0	1	2020-07-01 12:32:01
	// 282	4	退货信息提醒	sms_admin_postage_status	radio	0:关闭
	// 1:开启	0	退货信息提醒	0	0	1	2020-07-01 14:04:27
	// 283	2	后台登录页logo	sys_login_logo	image		0	后台登录页logo	0	0	1	2020-07-06 15:20:18
	// 284	2	后台登录页标题	sys_login_title	input		0	后台登录页标题	0	0	1	2020-07-07 16:39:25
	// 285	1	菜单logo	sys_menu_logo	image		0	设置菜单顶部logo	0	0	1	2020-07-09 16:31:46
	// 286	1	菜单小logo	sys_menu_slogo	image		0	设置菜单顶部小logo	0	0	1	2020-07-09 16:31:46
	// 287	1	商户入驻协议	sys_intention_agree	textarea		0	商户入驻协议	0	0	0	2020-07-20 11:19:32
	// 288	1	开启商户入驻	mer_intention_open	radio	0:关闭
	// 1:开启	0	是否开启商户入驻功能	0	0	1	2020-07-27 14:47:45
	// 289	4	预售尾款支付通知	sms_pay_presell_status	radio	0:关闭
	// 1:开启	0		1	0	1	2020-11-30 17:46:45
	// 290	5	打印机终端号	terminal_number	input		0	打印机终端号	0	1	1	2020-07-29 15:47:41
	// 291	5	打印机应用ID	printing_client_id	input		0	打印机开发者用户ID	0	1	1	2020-07-29 15:55:13
	// 292	5	打印机用户ID	develop_id	input		0	打印机的应用ID	0	1	1	2020-07-29 15:56:06
	// 293	5	打印机密匙	printing_api_key	input		0	打印机应用密匙	0	1	1	2020-07-29 15:57:46
	// 294	1	开启直播免审核	broadcast_room_type	radio	0:关闭
	// 1:开启	0	是否开启直播免审核	0	0	1	2020-07-29 16:50:35
	// 295	1	开启复制第三方平台商品	copy_product_status	radio	0:关闭
	// 1:开启	0	是否开启复制商品功能	0	0	1	2020-07-30 15:49:01
	// 296	1	复制商品接口KEY	copy_product_apikey	input		0	接口key	0	0	1	2020-07-30 15:49:46
	// 297	1	开启直播商品免审核	broadcast_goods_type	radio	0:关闭
	// 1:开启	0	是否开启直播商品免审核	0	0	1	2020-07-30 16:00:53
	// 298	1	腾讯地图KEY	tx_map_key	input		0	腾讯地图KEY	0	0	1	2020-08-01 11:55:13
	// 299	14	开启门店自提	mer_take_status	radio	0:关闭
	// 1:开启	0	是否开启门店自提	0	1	1	2020-08-01 11:59:14
	// 300	14	自提点名称	mer_take_name	input		0	设置自提点名称	0	1	1	2020-08-01 12:01:08
	// 301	14	自提点手机号	mer_take_phone	input		0	设置自提点手机号	0	1	1	2020-08-01 12:02:04
	// 302	14	自提点地址	mer_take_address	input		0	设置自提点地址	0	1	1	2020-08-01 12:03:57
	// 303	14	店铺经纬度	mer_take_location	input		0	设置店铺经纬度	0	1	1	2020-08-01 12:16:18
	// 304	14	自提点营业日期	mer_take_day	input		0	设置自提点营业日期	0	1	1	2020-08-01 12:17:06
	// 305	14	自提点营业时间	mer_take_time	input		0	设置自提点营业时间	0	1	1	2020-08-01 12:17:39
	// 306	1	订单自动收货时间(天)	auto_take_order_timer	number		0	设置订单自动收货时间(天)	0	0	1	2020-08-04 14:57:23
	// 307	1	默认赠送复制次数	copy_product_defaul	number		0	默认给商户赠送可用次数	0	0	1	2020-08-06 12:16:20
	// 308	1	是否展示店铺	hide_mer_status	radio	1:关闭
	// 0:开启	0	是否展示店铺	0	0	1	2020-08-17 15:03:44
	// 309	4	直播审核通过主播通知	sms_broadcast_room_status	radio	0:关闭
	// 1:开启	0		0	0	1	2020-09-08 15:53:42
	// 310	4	验证码时效配置(分钟)	sms_time	number		0		0	0	1	2020-09-08 15:53:42
	// 311	5	打印机自动打印	printing_auto_status	radio	0:关闭
	// 1:开启	0	开启后订单支付成功后自动打印	0	1	1	2020-10-17 11:25:09
	// 312	15	支付宝支付状态	alipay_open	radio	0:关闭
	// 1:开启	0		0	0	1	2020-10-22 11:40:41
	// 313	15	支付宝app_id	alipay_app_id	input		0		0	0	1	2020-10-22 11:41:23
	// 314	15	支付宝公钥	alipay_public_key	input		0		0	0	1	2020-10-22 11:41:51
	// 315	15	支付密钥	alipay_private_key	input		0		0	0	1	2020-10-22 11:42:22
	// 316	5	打印机开启	printing_status	radio	0:关闭
	// 1:开启	0		0	1	1	2020-11-10 17:59:49
	// 317	5	开启发票	mer_open_receipt	radio	0:关闭
	// 1:开启	0	设置是否开启发票	0	1	1	2020-11-14 11:40:10
}

func (m *config) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.SysConfig{}).RowsAffected == 1 {
			color.Danger.Println("\n[Mysql] --> sys_configs 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configs).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> sys_configs 表初始数据成功!")
		return nil
	})
}
