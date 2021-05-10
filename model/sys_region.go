package model

type SysRegion struct {
	Code       int         `json:"code" gorm:""`
	PCode      int         `json:"p_code" gorm:""`
	Name       string      `json:"name" gorm:""`
	SubRegions []SysRegion `json:"sub_regions" gorm:"foreignKey:PCode;references:Code;comment:子区域"`
}
