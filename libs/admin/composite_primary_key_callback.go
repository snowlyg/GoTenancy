package admin

import (
	"fmt"
	"regexp"

	"github.com/jinzhu/gorm"
)

var primaryKeyRegexp = regexp.MustCompile(`primary_key\[.+_.+\]`)

func (admin Admin) registerCompositePrimaryKeyCallback() {
	if db := admin.DB; db != nil {
		// register middleware
		router := admin.GetRouter()
		router.Use(&Middleware{
			Name: "composite primary key filter",
			Handler: func(context *Context, middleware *Middleware) {
				db := context.GetDB()
				for key, value := range context.Request.URL.Query() {
					if primaryKeyRegexp.MatchString(key) {
						db = db.Set(key, value)
					}
				}
				context.SetDB(db)

				middleware.Next(context)
			},
		})

		callbackProc := db.Callback().Query().Before("gorm:query")
		callbackName := "qor_admin:composite_primary_key"
		if callbackProc.Get(callbackName) == nil {
			callbackProc.Register(callbackName, compositePrimaryKeyQueryCallback)
		}

		callbackProc = db.Callback().RowQuery().Before("gorm:row_query")
		if callbackProc.Get(callbackName) == nil {
			callbackProc.Register(callbackName, compositePrimaryKeyQueryCallback)
		}
	}
}

// DisableCompositePrimaryKeyMode disable composite primary key mode
var DisableCompositePrimaryKeyMode = "composite_primary_key:query:disable"

func compositePrimaryKeyQueryCallback(scope *gorm.Scope) {
	if value, ok := scope.Get(DisableCompositePrimaryKeyMode); ok && value != "" {
		return
	}

	tableName := scope.TableName()
	for _, primaryField := range scope.PrimaryFields() {
		if value, ok := scope.Get(fmt.Sprintf("primary_key[%v_%v]", tableName, primaryField.DBName)); ok && value != "" {
			scope.Search.Where(fmt.Sprintf("%v = ?", scope.Quote(primaryField.DBName)), value)
		}
	}
}
