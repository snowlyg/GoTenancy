package request

type GetSysConfig struct {
	Type string `json:"type" form:"type"`
	Name string `json:"name" form:"name"`
}
