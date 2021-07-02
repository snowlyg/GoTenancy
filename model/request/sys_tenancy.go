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

type UpdateClientTenancy struct {
	Avatar string `json:"avatar"`
	Banner string `json:"banner"`
	Info   string `json:"info"`
	State  int    `json:"state"`
	Tele   string `json:"tele"`
}

type SetCopyProductNum struct {
	CopyNum int `json:"copyNum"`
	Num     int `json:"num"`
	Type    int `json:"type" binding:"required"` // 1:+ ,2:-
}
