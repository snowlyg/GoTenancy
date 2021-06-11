package request

// pay_weixin_appid: "sdfsdfas"
// pay_weixin_appsecret: "sadfsadf"
// pay_weixin_client_cert: ""
// pay_weixin_client_key: ""
// pay_weixin_key: ""
// pay_weixin_mchid: "asfsdafa"
// pay_weixin_open: ""
type PayWeiXin struct {
	PayWeixinAppid      string `json:"pay_weixin_appid" uri:"pay_weixin_appid" form:"pay_weixin_appid"`
	PayWeixinAppsecret  string `json:"pay_weixin_appsecret" uri:"pay_weixin_appsecret" form:"pay_weixin_appsecret"`
	PayWeixinClientCert string `json:"pay_weixin_client_cert" uri:"pay_weixin_client_cert" form:"pay_weixin_client_cert"`
	PayWeixinClientKey  string `json:"pay_weixin_client_key" uri:"pay_weixin_client_key" form:"pay_weixin_client_key"`
	PayWeixinKey        string `json:"pay_weixin_key" uri:"pay_weixin_key" form:"pay_weixin_key"`
	PayWeixinMchid      string `json:"pay_weixin_mchid" uri:"pay_weixin_mchid" form:"pay_weixin_mchid"`
	PayWeixinOpen       string `json:"pay_weixin_open" uri:"pay_weixin_open" form:"pay_weixin_open"`
}
