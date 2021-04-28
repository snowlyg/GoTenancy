package response

import "github.com/snowlyg/go-tenancy/model/request"

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
