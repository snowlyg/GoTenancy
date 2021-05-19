package response

type SysConfig struct {
	TenancyResponse
	Type  string `json:"type" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}
