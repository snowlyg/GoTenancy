package request

type GetSysConfig struct {
	Type string `json:"type" form:"type"  binding:"required"`
	Name string `json:"name" form:"name"  binding:"required"`
}
