package model

import (
	"github.com/snowlyg/go-tenancy/g"
)

type JwtBlacklist struct {
	g.TENANCY_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
