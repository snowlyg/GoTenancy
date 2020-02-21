package dummy

import (
	"fmt"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/media"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
)

// NewDummyAdmin generate admin for dummy app
func NewDummyAdmin(keepData ...bool) *admin.Admin {
	var (
		db     = utils.TestDB()
		models = []interface{}{&User{}, &CreditCard{}, &Address{}, &Language{}, &Profile{}, &Phone{}, &Company{}}
		Admin  = admin.New(&qor.Config{DB: db})
	)

	media.RegisterCallbacks(db)

	for _, value := range models {
		if len(keepData) == 0 {
			db.DropTableIfExists(value)
		}
		db.AutoMigrate(value)
	}

	Admin.AddResource(&Company{})
	Admin.AddResource(&Language{}, &admin.Config{Name: "语种 & 语言", Priority: -1})
	user := Admin.AddResource(&User{})
	user.Meta(&admin.Meta{
		Name: "CreditCard",
		Type: "single_edit",
	})
	user.Meta(&admin.Meta{
		Name: "Languages",
		Type: "select_many",
		Collection: func(resource interface{}, context *qor.Context) (results [][]string) {
			if languages := []Language{}; !context.GetDB().Find(&languages).RecordNotFound() {
				for _, language := range languages {
					results = append(results, []string{fmt.Sprint(language.ID), language.Name})
				}
			}
			return
		},
	})

	return Admin
}
