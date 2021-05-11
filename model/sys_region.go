package model

type SysRegion struct {
	Code       int         `json:"code" gorm:""`
	PCode      int         `json:"pCode" gorm:""`
	Name       string      `json:"name" gorm:""`
	SubRegions []SysRegion `json:"subRegions" gorm:"foreignKey:PCode;references:Code;comment:子区域"`
}
