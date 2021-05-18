package response

type SysConfig struct {
	TenancyResponse
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
