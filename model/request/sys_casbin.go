package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"  validate:"required"`
	Method string `json:"method"  validate:"required"`
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"  validate:"required"`
	CasbinInfos []CasbinInfo `json:"casbinInfos" `
}
