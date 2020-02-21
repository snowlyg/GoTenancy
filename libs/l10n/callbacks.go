package l10n

import (
	"fmt"
	"reflect"

	"GoTenancy/libs/qor/utils"
	"github.com/jinzhu/gorm"
)

func beforeQuery(scope *gorm.Scope) {
	if IsLocalizable(scope) {
		quotedTableName := scope.QuotedTableName()
		quotedPrimaryKey := scope.Quote(scope.PrimaryKey())
		_, hasDeletedAtColumn := scope.FieldByName("deleted_at")

		locale, isLocale := getQueryLocale(scope)
		switch mode, _ := scope.DB().Get("l10n:mode"); mode {
		case "unscoped":
		case "global":
			scope.Search.Where(fmt.Sprintf("%v.language_code = ?", quotedTableName), Global)
		case "locale":
			scope.Search.Where(fmt.Sprintf("%v.language_code = ?", quotedTableName), locale)
		case "reverse":
			if !scope.Search.Unscoped && hasDeletedAtColumn {
				scope.Search.Where(fmt.Sprintf(
					"(%v.%v NOT IN (SELECT DISTINCT(%v) FROM %v t2 WHERE t2.language_code = ? AND t2.deleted_at IS NULL) AND %v.language_code = ?)", quotedTableName, quotedPrimaryKey, quotedPrimaryKey, quotedTableName, quotedTableName), locale, Global)
			} else {
				scope.Search.Where(fmt.Sprintf("(%v.%v NOT IN (SELECT DISTINCT(%v) FROM %v t2 WHERE t2.language_code = ?) AND %v.language_code = ?)", quotedTableName, quotedPrimaryKey, quotedPrimaryKey, quotedTableName, quotedTableName), locale, Global)
			}
		case "fallback":
			fallthrough
		default:
			if isLocale {
				if !scope.Search.Unscoped && hasDeletedAtColumn {
					scope.Search.Where(fmt.Sprintf("((%v.%v NOT IN (SELECT DISTINCT(%v) FROM %v t2 WHERE t2.language_code = ? AND t2.deleted_at IS NULL) AND %v.language_code = ?) OR %v.language_code = ?) AND %v.deleted_at IS NULL", quotedTableName, quotedPrimaryKey, quotedPrimaryKey, quotedTableName, quotedTableName, quotedTableName, quotedTableName), locale, Global, locale)
				} else {
					scope.Search.Where(fmt.Sprintf("(%v.%v NOT IN (SELECT DISTINCT(%v) FROM %v t2 WHERE t2.language_code = ?) AND %v.language_code = ?) OR (%v.language_code = ?)", quotedTableName, quotedPrimaryKey, quotedPrimaryKey, quotedTableName, quotedTableName, quotedTableName), locale, Global, locale)
				}
				scope.Search.Order(gorm.Expr(fmt.Sprintf("%v.language_code = ? DESC", quotedTableName), locale))
			} else {
				scope.Search.Where(fmt.Sprintf("%v.language_code = ?", quotedTableName), Global)
			}
		}
	}
}

func beforeCreate(scope *gorm.Scope) {
	if IsLocalizable(scope) {
		if locale, ok := getLocale(scope); ok { // is locale
			if isLocaleCreatable(scope) || !scope.PrimaryKeyZero() {
				setLocale(scope, locale)
			} else {
				err := fmt.Errorf("the resource %v cannot be created in %v", scope.GetModelStruct().ModelType.Name(), locale)
				scope.Err(err)
			}
		} else {
			setLocale(scope, Global)
		}
	}
}

func beforeUpdate(scope *gorm.Scope) {
	if IsLocalizable(scope) {
		locale, isLocale := getLocale(scope)

		switch mode, _ := scope.DB().Get("l10n:mode"); mode {
		case "unscoped":
		default:
			scope.Search.Where(fmt.Sprintf("%v.language_code = ?", scope.QuotedTableName()), locale)
			setLocale(scope, locale)
		}

		if isLocale {
			scope.Search.Omit(syncColumns(scope)...)
		}
	}
}

func afterUpdate(scope *gorm.Scope) {
	if !scope.HasError() {
		if IsLocalizable(scope) {
			if locale, ok := getLocale(scope); ok {
				if scope.DB().RowsAffected == 0 && !scope.PrimaryKeyZero() { //is locale and nothing updated
					var count int
					var query = fmt.Sprintf("%v.language_code = ? AND %v.%v = ?", scope.QuotedTableName(), scope.QuotedTableName(), scope.PrimaryKey())

					// if enabled soft delete, delete soft deleted records
					if scope.HasColumn("DeletedAt") {
						scope.NewDB().Unscoped().Where("deleted_at is not null").Where(query, locale, scope.PrimaryKeyValue()).Delete(scope.Value)
					}

					// if no localized records exist, localize it
					if scope.NewDB().Table(scope.TableName()).Where(query, locale, scope.PrimaryKeyValue()).Count(&count); count == 0 {
						scope.DB().RowsAffected = scope.DB().Create(scope.Value).RowsAffected
					}
				}
			} else if syncColumns := syncColumns(scope); len(syncColumns) > 0 { // is global
				if mode, _ := scope.DB().Get("l10n:mode"); mode != "unscoped" {
					if scope.DB().RowsAffected > 0 {
						var primaryField = scope.PrimaryField()
						var syncAttrs = map[string]interface{}{}

						if updateAttrs, ok := scope.InstanceGet("gorm:update_attrs"); ok {
							for key, value := range updateAttrs.(map[string]interface{}) {
								for _, syncColumn := range syncColumns {
									if syncColumn == key {
										syncAttrs[syncColumn] = value
										break
									}
								}
							}
						} else {
							for _, syncColumn := range syncColumns {
								if field, ok := scope.FieldByName(syncColumn); ok && field.IsNormal {
									syncAttrs[syncColumn] = field.Field.Interface()
								}
							}
						}

						if len(syncAttrs) > 0 {
							db := scope.DB().Model(reflect.New(utils.ModelType(scope.Value)).Interface()).Set("l10n:mode", "unscoped").Where("language_code <> ?", Global)
							if !primaryField.IsBlank {
								db = db.Where(fmt.Sprintf("%v = ?", primaryField.DBName), primaryField.Field.Interface())
							}
							scope.Err(db.UpdateColumns(syncAttrs).Error)
						}
					}
				}
			}
		}
	}
}

func beforeDelete(scope *gorm.Scope) {
	if IsLocalizable(scope) {
		if locale, ok := getQueryLocale(scope); ok { // is locale
			scope.Search.Where(fmt.Sprintf("%v.language_code = ?", scope.QuotedTableName()), locale)
		}
	}
}

// RegisterCallbacks 注册回调函数到 GORM DB
func RegisterCallbacks(db *gorm.DB) {
	callback := db.Callback()

	if callback.Create().Get("l10n:before_create") == nil {
		callback.Create().Before("gorm:before_create").Register("l10n:before_create", beforeCreate)
	}

	if callback.Update().Get("l10n:before_update") == nil {
		callback.Update().Before("gorm:before_update").Register("l10n:before_update", beforeUpdate)
	}
	if callback.Update().Get("l10n:after_update") == nil {
		callback.Update().After("gorm:after_update").Register("l10n:after_update", afterUpdate)
	}

	if callback.Delete().Get("l10n:before_delete") == nil {
		callback.Delete().Before("gorm:before_delete").Register("l10n:before_delete", beforeDelete)
	}

	if callback.RowQuery().Get("l10n:before_query") == nil {
		callback.RowQuery().Before("gorm:row_query").Register("l10n:before_query", beforeQuery)
	}
	if callback.Query().Get("l10n:before_query") == nil {
		callback.Query().Before("gorm:query").Register("l10n:before_query", beforeQuery)
	}
}
