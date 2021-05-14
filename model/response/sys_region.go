package response

type SysRegion struct {
	Code       int         `json:"code"`
	PCode      int         `json:"pCode"`
	Name       string      `json:"name"`
	SubRegions []SysRegion `json:"subRegions" gorm:"foreignKey:PCode;references:Code;comment:子区域"`
}
