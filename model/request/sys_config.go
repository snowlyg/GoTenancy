package request

type GetSysConfig struct {
	Type string `json:"type" form:"type"  validate:"required"`
	Name string `json:"name" form:"name"  validate:"required"`
}
