package response

import "github.com/snowlyg/go-tenancy/model"

type SysTenancy struct {
	TenancyResponse
	model.BaseTenancy
}

type Counts struct {
	Invalid int
	Valid   int
}
