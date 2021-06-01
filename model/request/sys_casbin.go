package request

// Casbin info structure
type CasbinInfo struct {
	Path   string `json:"path"  binding:"required"`
	Method string `json:"method"  binding:"required"`
}

// Casbin structure for input parameters
type CasbinInReceive struct {
	AuthorityId string       `json:"authorityId"  binding:"required"`
	CasbinInfos []CasbinInfo `json:"casbinInfos" `
}
