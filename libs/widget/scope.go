package widget

import "GoTenancy/libs/qor/utils"

var registeredScopes []*Scope

// Scope widget scope
type Scope struct {
	Name    string
	Param   string
	Visible func(*Context) bool
}

// ToParam generate param for scope
func (scope *Scope) ToParam() string {
	if scope.Param != "" {
		return scope.Param
	}
	return utils.ToParamString(scope.Name)
}

// RegisterScope register scope for widget
func (widgets *Widgets) RegisterScope(scope *Scope) {
	registeredScopes = append(registeredScopes, scope)
}
