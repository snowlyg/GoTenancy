package response

import "github.com/snowlyg/go-tenancy/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
