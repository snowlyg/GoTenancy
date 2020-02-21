package admin

import (
	"html/template"
	"path/filepath"
	"reflect"

	"GoTenancy/libs/assetfs"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/roles"
)

var (
	globalViewPaths []string
	globalAssetFSes []assetfs.Interface
)

// HasPermissioner has permission interface
type HasPermissioner interface {
	HasPermission(roles.PermissionMode, *qor.Context) bool
}

// ResourceNamer is an interface for models that defined method `ResourceName`
type ResourceNamer interface {
	ResourceName() string
}

// I18n define admin's i18n interface
type I18n interface {
	Scope(scope string) I18n
	Default(value string) I18n
	T(locale string, key string, args ...interface{}) template.HTML
}

// RegisterViewPath register view path for all assetfs
func RegisterViewPath(pth string) {
	globalViewPaths = append(globalViewPaths, pth)

	for _, assetFS := range globalAssetFSes {
		if assetFS.RegisterPath(filepath.Join(utils.AppRoot, "vendor", pth)) != nil {
			for _, gopath := range utils.GOPATH() {
				if assetFS.RegisterPath(filepath.Join(gopath, "src", pth)) == nil {
					break
				}
			}
		}
	}
}

func equal(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}
