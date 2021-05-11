package response

type SysTenancy struct {
	TenancyResponse
	UUID         string `json:"uuid"`
	Name         string `json:"name"`
	Tele         string `json:"tele"`
	Address      string `json:"address"`
	BusinessTime string `json:"businessTime"`
}
