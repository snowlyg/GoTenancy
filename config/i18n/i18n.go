package i18n

import (
	"path/filepath"

	"GoTenancy/libs/i18n"
	"GoTenancy/libs/i18n/backends/database"
	"GoTenancy/libs/i18n/backends/yaml"

	"GoTenancy/config"
	"GoTenancy/config/db"
)

var I18n *i18n.I18n

func init() {
	I18n = i18n.New(database.New(db.DB), yaml.New(filepath.Join(config.Root, "config/locales")))
}
