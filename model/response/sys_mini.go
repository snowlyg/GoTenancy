package response

type SysMini struct {
	TenancyResponse
	UUID      string `json:"uuid"`
	Name      string `json:"name"`
	AppID     string `json:"appId"`
	AppSecret string `json:"appSecret"`
	Remark    string `json:"remark"`
}
