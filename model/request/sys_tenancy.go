package request

type SetRegionCode struct {
	Id            float64 `json:"id" form:"id" binding:"required,gt=0"`
	SysRegionCode int     `json:"sysRegionCode" binding:"required"`
}

type TenancyPageInfo struct {
	PageInfo
	Date    string `json:"date"`
	Status  string `json:"status"`
	Keyword string `json:"keyword"`
}
