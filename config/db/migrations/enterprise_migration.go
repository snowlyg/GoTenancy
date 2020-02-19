// +build enterprise

package migrations

import "GoTenancy/app/enterprise"

func init() {
	AutoMigrate(&enterprise.QorMicroSite{})
}
