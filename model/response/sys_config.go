package response

type SysConfig struct {
	TenancyResponse
	Type  string `json:"type" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Value string `json:"value" binding:"required"`
}
