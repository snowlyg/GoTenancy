package response

import "github.com/snowlyg/go-tenancy/model"

type SysTenancy struct {
	TenancyResponse
	model.BaseTenancy
}
type TenancySelect struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Counts struct {
	Invalid int
	Valid   int
}

type LoginTenancy struct {
	Admin SysTenancyUser `json:"admin"`
	Exp   int64          `json:"exp"`
	Token string         `json:"token"`
	Url   string         `json:"url"`
}

type TenancyInfo struct {
	Avatar string `json:"avatar"`
	Banner string `json:"banner"`
	Id     uint   `json:"id"`
	Info   string `json:"info"`
	Name   string `json:"name"`
}
