package admin

import (
	"GoTenancy/libs/qor"
	"github.com/jinzhu/gorm"
)

// Scope register scope for qor resource
func (res *Resource) Scope(scope *Scope) {
	if scope.Label == "" {
		scope.Label = scope.Name
	}
	res.scopes = append(res.scopes, scope)
}

// Scope scope definiation
type Scope struct {
	Name    string
	Label   string
	Group   string
	Visible func(context *Context) bool
	Handler func(*gorm.DB, *qor.Context) *gorm.DB
	Default bool
}
