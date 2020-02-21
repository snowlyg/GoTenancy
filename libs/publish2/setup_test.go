package publish2_test

import (
	"GoTenancy/libs/l10n"
	"GoTenancy/libs/publish2"
	"GoTenancy/libs/qor/test/utils"
)

var DB = utils.TestDB()

func init() {
	models := []interface{}{
		&Wiki{}, &Post{}, &Article{}, &Discount{}, &User{}, &Campaign{},
		&Product{}, &L10nProduct{}, &SharedVersionProduct{}, &SharedVersionColorVariation{}, &SharedVersionSizeVariation{},
	}

	DB.DropTableIfExists(models...)
	DB.AutoMigrate(models...)
	publish2.RegisterCallbacks(DB)
	l10n.RegisterCallbacks(DB)
}
