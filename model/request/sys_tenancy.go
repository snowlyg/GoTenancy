package request

type SetRegionCode struct {
	Id            float64 `json:"id" form:"id" validate:"required,gt=0"`
	SysRegionCode int     `json:"sysRegionCode" validate:"required"`
}
type CreateSysTenancy struct {
	Name          string `json:"name" form:"name" validate:"required"`
	Tele          string `json:"tele" form:"tele" `
	Address       string `json:"address" form:"address" `
	BusinessTime  string `json:"businessTime" form:"businessTime"`
	SysRegionCode int    `json:"sysRegionCode" form:"sysRegionCode" validate:"required"`
}
type UpdateSysTenancy struct {
	Id            uint   `json:"id" form:"id" validate:"required,gt=0"`
	Name          string `json:"name" form:"name" validate:"required"`
	Tele          string `json:"tele" form:"tele" `
	Address       string `json:"address" form:"address" `
	BusinessTime  string `json:"businessTime" form:"businessTime"`
	SysRegionCode int    `json:"sysRegionCode" form:"sysRegionCode" validate:"required"`
}
