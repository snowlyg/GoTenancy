package request

type SetRegionCode struct {
	Id            float64 `json:"id" form:"id" binding:"required,gt=0"`
	SysRegionCode int     `json:"sysRegionCode" binding:"required"`
}
type CreateSysTenancy struct {
	Name          string `json:"name" form:"name" binding:"required"`
	Tele          string `json:"tele" form:"tele" `
	Address       string `json:"address" form:"address" `
	BusinessTime  string `json:"businessTime" form:"businessTime"`
	SysRegionCode int    `json:"sysRegionCode" form:"sysRegionCode" binding:"required"`
}
type UpdateSysTenancy struct {
	Id            uint   `json:"id" form:"id" binding:"required,gt=0"`
	Name          string `json:"name" form:"name" binding:"required"`
	Tele          string `json:"tele" form:"tele" `
	Address       string `json:"address" form:"address" `
	BusinessTime  string `json:"businessTime" form:"businessTime"`
	SysRegionCode int    `json:"sysRegionCode" form:"sysRegionCode" binding:"required"`
}
