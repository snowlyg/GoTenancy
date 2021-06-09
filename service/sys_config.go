package service

import (
	"errors"

	"github.com/snowlyg/go-tenancy/g"
	"github.com/snowlyg/go-tenancy/model"
	"github.com/snowlyg/go-tenancy/model/request"
	"github.com/snowlyg/go-tenancy/model/response"
	"gorm.io/gorm"
)

// getFrom() {
// 	if (this.$route.params.key === "upload") {
// 		this.FromData = {
// 			rule: [
// 				{
// 					type: "input",
// 					field: "pay_routine_appid",
// 					value: "",
// 					title: "Appid",
// 					info: "小程序Appid",
// 					props: { type: "text", placeholder: "请输入Appid" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_appsecret",
// 					value: "",
// 					title: "Appsecret",
// 					info: "小程序Appsecret",
// 					props: { type: "text", placeholder: "请输入Appsecret" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_mchid",
// 					value: "",
// 					title: "Mchid",
// 					info: "商户号",
// 					props: { type: "text", placeholder: "请输入Mchid" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_key",
// 					value: "",
// 					title: "Key",
// 					info: "商户key",
// 					props: { type: "text", placeholder: "请输入Key" },
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_routine_client_cert",
// 					value: "",
// 					title: "小程序支付证书",
// 					name: "pay_routine_client_cert",
// 					info: "小程序支付证书",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + this.token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_routine_client_key",
// 					value: "",
// 					title: "小程序支付证书密钥",
// 					name: "pay_routine_client_key",
// 					info: "小程序支付证书密钥",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + this.token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 			],
// 			action: "/sys/config/save/routine_pay.html",
// 			method: "POST",
// 			title: "小程序支付配置",
// 			config: {},
// 		};
// 	} else if (this.$route.params.key === "routine_pay") {
// 		this.FromData = {
// 			rule: [
// 				{
// 					type: "input",
// 					field: "pay_routine_appid",
// 					value: "",
// 					title: "Appid",
// 					info: "小程序Appid",
// 					props: { type: "text", placeholder: "请输入Appid" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_appsecret",
// 					value: "",
// 					title: "Appsecret",
// 					info: "小程序Appsecret",
// 					props: { type: "text", placeholder: "请输入Appsecret" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_mchid",
// 					value: "",
// 					title: "Mchid",
// 					info: "商户号",
// 					props: { type: "text", placeholder: "请输入Mchid" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_routine_key",
// 					value: "",
// 					title: "Key",
// 					info: "商户key",
// 					props: { type: "text", placeholder: "请输入Key" },
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_routine_client_cert",
// 					value: "",
// 					title: "小程序支付证书",
// 					name: "pay_routine_client_cert",
// 					info: "小程序支付证书",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + this.token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_routine_client_key",
// 					value: "",
// 					title: "小程序支付证书密钥",
// 					name: "pay_routine_client_key",
// 					info: "小程序支付证书密钥",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + this.token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 			],
// 			action: "/sys/config/save/routine_pay.html",
// 			method: "POST",
// 			title: "小程序支付配置",
// 			config: {},
// 		};
// 	} else if (this.$route.params.key === "wechat_payment") {
// 		this.FromData = {
// 			rule: [
// 				{
// 					type: "input",
// 					field: "pay_weixin_appid",
// 					value: "sdfsdfas",
// 					title: "Appid",
// 					info: "微信公众号身份的唯一标识。审核通过后，在微信发送的邮件中查看。",
// 					props: { type: "text", placeholder: "请输入Appid" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_weixin_appsecret",
// 					value: "sadfsadf",
// 					title: "Appsecret",
// 					info: "JSAPI接口中获取openid，审核后在公众平台开启开发模式后可查看。",
// 					props: { type: "text", placeholder: "请输入Appsecret" },
// 				},
// 				{
// 					type: "input",
// 					field: "pay_weixin_mchid",
// 					value: "asfsdafa",
// 					title: "Mchid",
// 					info: "受理商ID，身份标识",
// 					props: { type: "text", placeholder: "请输入Mchid" },
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_weixin_client_cert",
// 					value: "",
// 					title: "微信支付证书",
// 					name: "pay_weixin_client_cert",
// 					info: "微信支付证书，在微信商家平台中可以下载！文件名一般为apiclient_cert.pem",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 				{
// 					type: "upload",
// 					field: "pay_weixin_client_key",
// 					value: "",
// 					title: "微信支付证书密钥",
// 					name: "pay_weixin_client_key",
// 					info: "微信支付证书密钥，在微信商家平台中可以下载！文件名一般为apiclient_key.pem",
// 					props: {
// 						limit: 1,
// 						uploadType: "file",
// 						headers: {
// 							"Authorization": "Bearer " + this.token,
// 						},
// 						data: [],
// 						action: "/sys/config/setting/upload_file/file.html",
// 					},
// 				},
// 				{
// 					type: "input",
// 					field: "pay_weixin_key",
// 					value: "",
// 					title: "Key",
// 					info: "商户支付密钥Key。审核通过后，在微信发送的邮件中查看。",
// 					props: { type: "text", placeholder: "请输入Key" },
// 				},
// 				{
// 					type: "radio",
// 					field: "pay_weixin_open",
// 					value: "",
// 					title: "开启",
// 					info: "是否启用微信支付",
// 					props: {},
// 					options: [
// 						{ value: "0", label: "关闭" },
// 						{ value: "1", label: "开启" },
// 					],
// 				},
// 			],
// 			action: "/sys/config/save/wechat_payment.html",
// 			method: "POST",
// 			title: "公众号支付配置",
// 			config: {},
// 		};
// 	} else if (this.$route.params.key === "alipay") {
// 		this.FromData = {
// 			rule: [
// 				{
// 					type: "radio",
// 					field: "alipay_open",
// 					value: "",
// 					title: "支付宝支付状态",
// 					info: "",
// 					props: {},
// 					options: [
// 						{ value: "0", label: "关闭" },
// 						{ value: "1", label: "开启" },
// 					],
// 				},
// 				{
// 					type: "input",
// 					field: "alipay_app_id",
// 					value: "",
// 					title: "支付宝app_id",
// 					info: "",
// 					props: { type: "text", placeholder: "请输入支付宝app_id" },
// 				},
// 				{
// 					type: "input",
// 					field: "alipay_public_key",
// 					value: "",
// 					title: "支付宝公钥",
// 					info: "",
// 					props: { type: "text", placeholder: "请输入支付宝公钥" },
// 				},
// 				{
// 					type: "input",
// 					field: "alipay_private_key",
// 					value: "",
// 					title: "支付密钥",
// 					info: "",
// 					props: { type: "text", placeholder: "请输入支付密钥" },
// 				},
// 			],
// 			action: "/sys/config/save/alipay.html",
// 			method: "POST",
// 			title: "支付宝支付配置",
// 			config: {},
// 		};
// 	}
// },
// CreateConfig
func CreateConfig(m model.SysConfig) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("name = ?", m.Name).Where("type = ?", m.Type).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, errors.New("设置名称已经使用")
	}
	err = g.TENANCY_DB.Create(&m).Error
	return m, err
}

// GetConfigByName
func GetConfigByName(name, style string) (model.SysConfig, error) {
	var config model.SysConfig
	err := g.TENANCY_DB.Where("name = ?", name).Where("type = ?", style).First(&config).Error
	return config, err
}

// GetConfigInfoList
func GetConfigInfoList(info request.PageInfo) ([]response.SysConfig, int64, error) {
	var configList []response.SysConfig
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := g.TENANCY_DB.Model(&model.SysConfig{})
	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return configList, total, err
	}
	err = db.Limit(limit).Offset(offset).Find(&configList).Error
	return configList, total, err
}

// UpdateConfig
func UpdateConfig(m model.SysConfig) (model.SysConfig, error) {
	err := g.TENANCY_DB.Where("name = ?", m.Name).Where("id <> ?", m.ID).Where("type = ?", m.Type).First(&model.SysConfig{}).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return m, errors.New("设置名称已经使用")
	}
	err = g.TENANCY_DB.Updates(&m).Error
	return m, err
}

// DeleteConfig
func DeleteConfig(id float64) error {
	var config model.SysConfig
	return g.TENANCY_DB.Where("id = ?", id).Delete(&config).Error
}
