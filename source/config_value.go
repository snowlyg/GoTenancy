package source

import (
	"github.com/gookit/color"
	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"gorm.io/gorm"
)

var SysConfigValue = new(configValue)

type configValue struct{}

var configValues = []model.SysConfigValue{
	{ConfigKey: "site_url", Value: "http://mer.crmeb.net/", SysTenancyID: 0},
	// 2	site_name	"CRMEB\u591a\u5546\u6237\u5546\u57ce"	0	2020-04-22 17:39:37
	// 3	site_open	"1"	0	2020-04-22 17:39:37
	// 4	wechat_name	"crmeb \u591a\u5546\u6237"	0	2020-04-22 18:18:20
	// 5	set_phone	18741523695	55	2020-04-23 18:17:26
	// 7	set_email	"mkpmkmp"	55	2020-04-23 18:27:29
	// 14	wechat_qrcode	""	0	2020-04-26 12:13:13
	// 15	wechat_avatar	""	0	2020-04-26 12:13:13
	// 16	wechat_share_img	""	0	2020-04-26 12:13:13
	// 17	wechat_share_title	""	0	2020-04-26 12:13:13
	// 18	wechat_share_synopsis	""	0	2020-04-26 12:13:13
	// 19	wechat_encode	"0"	0	2020-04-26 14:04:28
	// 20	extension_one_rate	0.02	0	2020-05-09 11:04:26
	// 21	extension_two_rate	0.01	0	2020-05-09 11:04:26
	// 22	extension_status	1	0	2020-05-09 11:04:27
	// 23	mer_store_stock	10	55	2020-05-18 11:35:27
	// 69	sms_user_pay_status	"1"	0	2020-05-25 15:26:05
	// 70	sms_user_postage_status	"1"	0	2020-05-25 15:26:05
	// 71	sms_user_take_status	"1"	0	2020-05-25 15:26:05
	// 72	sms_admin_order_status	"1"	0	2020-05-25 15:26:05
	// 73	sms_admin_pay_status	"1"	0	2020-05-25 15:26:06
	// 74	sms_admin_refund_status	"1"	0	2020-05-25 15:26:06
	// 75	sms_admin_take_status	"1"	0	2020-05-25 15:26:06
	// 76	sms_user_change_order_status	"1"	0	2020-05-25 15:26:06
	// 79	recharge_attention	"\u5145\u503c\u540e\u5e10\u6237\u7684\u91d1\u989d\u4e0d\u80fd\u63d0\u73b0\uff0c\u53ef\u7528\u4e8e\u5546\u57ce\u6d88\u8d39\u4f7f\u7528\n\u4f63\u91d1\u5bfc\u5165\u8d26\u6237\u4e4b\u540e\u4e0d\u80fd\u518d\u6b21\u5bfc\u51fa\u3001\u4e0d\u53ef\u63d0\u73b0\n\u8d26\u6237\u5145\u503c\u51fa\u73b0\u95ee\u9898\u53ef\u8054\u7cfb\u5546\u57ce\u5ba2\u670d\uff0c\u4e5f\u53ef\u62e8\u6253\u5546\u57ce\u5ba2\u670d\u70ed\u7ebf\uff1a4008888888"	0	2020-06-03 09:49:16
	// 89	auto_close_order_timer	0	0	2020-06-12 16:35:14
	// 90	refund_message	"\u6536\u8d27\u5730\u5740\u586b\u9519\u4e86\n\u4e0e\u63cf\u8ff0\u4e0d\u7b26\n\u4fe1\u606f\u586b\u9519\u4e86\uff0c\u91cd\u65b0\u62cd\n\u6536\u5230\u5546\u54c1\u635f\u574f\u4e86\n\u672a\u6309\u9884\u5b9a\u65f6\u95f4\u53d1\u8d27\n\u5176\u5b83\u539f\u56e0"	0	2020-06-12 16:35:14
	// 91	mer_refund_order_agree	7	0	2020-06-13 15:10:41
	// 92	mer_refund_address	"三亚噢"	56	2020-06-13 15:11:44
	// 93	mer_refund_user	"王二小"	55	2020-06-13 15:12:23
	// 94	bank	"天地银行"	55	2020-06-15 12:15:34
	// 95	bank_name	"少林寺驻武当山办事处大神父王朗"	55	2020-06-15 12:27:14
	// 96	bank_number	32342354353	55	2020-06-15 12:27:37
	// 97	bank_address	"是大是大非地方"	55	2020-06-15 12:28:13
	// 98	user_extract_min	100	0	2020-06-16 14:56:45
	// 99	lock_brokerage_timer	0	0	2020-06-17 18:33:09
	// 127	recharge_switch	"0"	0	2020-06-19 15:50:12
	// 128	store_user_min_recharge	100	0	2020-06-19 15:50:12
	// 129	balance_func_status	"0"	0	2020-06-19 15:54:26
	// 130	yue_pay_status	"1"	0	2020-06-19 16:02:30
	// 131	home_ad_pic	"http:\/\/mer.crmeb.net\/uploads\/def\/20200908\/87f2c4788665d03417d15074b2597e31.jpg"	0	2020-06-25 10:05:51
	// 132	home_ad_url	"\/pages\/users\/user_coupon\/index"	0	2020-06-25 10:05:51
	// 133	promoter_ explain	"\u963f\u8d3e\u514b\u65af"	0	2020-06-25 15:49:03
	// 134	promoter_bag_number	"2"	0	2020-06-25 15:49:03
	// 135	promoter_explain	"\u5145\u503c\u540e\u5e10\u6237\u7684\u91d1\u989d\u4e0d\u80fd\u63d0\u73b0\uff0c\u53ef\u7528\u4e8e\u5546\u57ce\u6d88\u8d39\u4f7f\u7528\n\u4f63\u91d1\u5bfc\u5165\u8d26\u6237\u4e4b\u540e\u4e0d\u80fd\u518d\u6b21\u5bfc\u51fa\u3001\u4e0d\u53ef\u63d0\u73b0\n\u8d26\u6237\u5145\u503c\u51fa\u73b0\u95ee\u9898\u53ef\u8054\u7cfb\u5546\u57ce\u5ba2\u670d\uff0c\u4e5f\u53ef\u62e8\u6253\u5546\u57ce\u5ba2\u670d\u70ed\u7ebf\uff1a4008888888"	0	2020-06-26 15:33:30
	// 136	max_bag_number	20	0	2020-06-26 15:33:30
	// 138	site_logo	"http:\/\/mer.crmeb.net\/uploads\/def\/20200816\/6c50374375d5fd6b2a8c40d49884daf6.png"	0	2020-06-30 16:17:55
	// 139	 share_title	"titletitletitletitle"	0	2020-06-30 16:17:55
	// 140	share_info	"CRMEB\u591a\u5546\u6237\u5546\u57ce"	0	2020-06-30 16:17:55
	// 141	share_pic	"http:\/\/mer.crmeb.net\/uploads\/def\/20200816\/67f5193c55d7b757f28c0204d3a5d07a.png"	0	2020-06-30 16:17:55
	// 142	share_title	"CRMEB\u591a\u5546\u6237\u5546\u57ce"	0	2020-06-30 16:21:08
	// 153	sms_fahuo_status	"1"	0	2020-07-01 16:57:18
	// 154	sms_take_status	"1"	0	2020-07-01 16:57:18
	// 155	sms_pay_status	"1"	0	2020-07-01 16:57:18
	// 156	sms_revision_status	"1"	0	2020-07-01 16:57:18
	// 157	sms_pay_false_status	"1"	0	2020-07-01 16:57:18
	// 158	sms_refund_fail_status	"1"	0	2020-07-01 16:57:18
	// 159	sms_refund_success_status	"1"	0	2020-07-01 16:57:18
	// 160	sms_refund_confirm_status	"1"	0	2020-07-01 16:57:18
	// 161	sms_admin_return_status	"1"	0	2020-07-01 16:57:18
	// 162	sms_admin_postage_status	"1"	0	2020-07-01 16:57:18
	// 163	sms_account	""	0	2020-07-02 19:00:51
	// 164	sms_token	""	0	2020-07-02 19:00:51
	// 165	sys_login_logo	"http:\/\/mer.crmeb.net\/uploads\/def\/20200816\/6c50374375d5fd6b2a8c40d49884daf6.png"	0	2020-07-07 15:23:07
	// 166	set_phone	"15109234132"	65	2020-07-08 17:48:21
	// 167	set_email	"78532941@qq.com"	65	2020-07-08 17:48:21
	// 168	mer_store_stock	"5"	65	2020-07-08 17:48:21
	// 169	mer_refund_address	"\u9655\u897f\u7701\u897f\u5b89\u5e02\u5317\u5927\u885775\u53f7"	65	2020-07-08 17:48:21
	// 170	mer_refund_user	"\u90d1\u6b63"	65	2020-07-08 17:48:21
	// 171	bank	"\u4e2d\u56fd\u519c\u4e1a\u94f6\u884c"	65	2020-07-08 17:48:21
	// 172	bank_number	"4214512365015841214"	65	2020-07-08 17:48:21
	// 173	bank_name	"\u90d1\u8def"	65	2020-07-08 17:48:21
	// 174	bank_address	"\u5317\u5927\u8857\u652f\u884c"	65	2020-07-08 17:48:21
	// 176	sys_menu_logo	"http:\/\/mer.crmeb.net\/uploads\/def\/20200816\/6c50374375d5fd6b2a8c40d49884daf6.png"	0	2020-07-09 20:55:28
	// 177	sys_menu_slogo	"http:\/\/mer.crmeb.net\/uploads\/def\/20200816\/e11a0f712ca67edff15e900858d690fa.png"	0	2020-07-09 20:55:28
	// 178	sys_login_title	""	0	2020-07-14 12:04:53
	// 179	express_app_code	""	0	2020-08-18 19:30:24
	// 180	sys_intention_agree	""	0	2020-08-18 19:30:24
	// 181	mer_intention_open	""	0	2020-08-18 19:30:24
	// 182	sms_time	30	0	2020-09-08 11:17:14
}

func (m *configValue) Init() error {
	return g.TENANCY_DB.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1, 2}).Find(&[]model.SysConfigValue{}).RowsAffected == 2 {
			color.Danger.Println("\n[Mysql] --> config_values 表的初始数据已存在!")
			return nil
		}
		if err := tx.Create(&configValues).Error; err != nil { // 遇到错误时回滚事务
			return err
		}
		color.Info.Println("\n[Mysql] --> config_values 表初始数据成功!")
		return nil
	})
}
