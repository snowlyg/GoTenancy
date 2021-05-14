package request

type SetRegionCode struct {
	Id            float64 `json:"id" form:"id"`
	SysRegionCode int     `json:"sysRegionCode"`
}
