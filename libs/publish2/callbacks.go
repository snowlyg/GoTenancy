package publish2

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"GoTenancy/libs/qor/utils"
	"github.com/jinzhu/gorm"
)

const (
	ModeOff             = "off"
	ModeReverse         = "reverse"
	VersionMode         = "publish:version:mode"
	VersionNameMode     = "publish:version:name"
	VersionMultipleMode = "multiple"

	ScheduleMode     = "publish:schedule:mode"
	ComingOnlineMode = "coming_online"
	GoingOfflineMode = "going_offline"
	ScheduledTime    = "publish:schedule:current"
	ScheduledStart   = "publish:schedule:start"
	ScheduledEnd     = "publish:schedule:end"

	VisibleMode = "publish:visible:mode"
)

func IsSchedulableModel(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(ScheduledInterface)
	}
	return
}

func IsVersionableModel(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(VersionableInterface)
	}
	return
}

func IsShareableVersionModel(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(ShareableVersionInterface)
	}
	return
}

func IsPublishReadyableModel(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(PublishReadyInterface)
	}
	return
}

func RegisterCallbacks(db *gorm.DB) {
	if db.Callback().Query().Get("publish:query") == nil {
		db.Callback().Query().Before("gorm:query").Register("publish:query", queryCallback)
	}
	if db.Callback().Query().Get("publish:fix_preload") == nil {
		db.Callback().Query().After("gorm:preload").Register("publish:fix_preload", fixPreloadCallback)
	}
	if db.Callback().RowQuery().Get("publish:query") == nil {
		db.Callback().RowQuery().Before("gorm:row_query").Register("publish:query", queryCallback)
	}

	if db.Callback().Create().Get("publish:versions") == nil {
		db.Callback().Create().Before("gorm:begin_transaction").Register("publish:versions", createCallback)
	}
	if db.Callback().Update().Get("publish:versions") == nil {
		db.Callback().Update().Before("gorm:begin_transaction").Register("publish:versions", updateCallback)
	}

	if db.Callback().Delete().Get("publish:versions") == nil {
		db.Callback().Delete().Before("gorm:begin_transaction").Register("publish:versions", deleteCallback)
	}
}

func queryCallback(scope *gorm.Scope) {
	var (
		isSchedulable      = IsSchedulableModel(scope.Value)
		isVersionable      = IsVersionableModel(scope.Value)
		isShareableVersion = IsShareableVersionModel(scope.Value)
		isPublishReadyable = IsPublishReadyableModel(scope.Value)
		conditions         []string
		conditionValues    []interface{}
	)

	if isSchedulable {
		var (
			scheduledStartTime, scheduledEndTime, scheduledCurrentTime *time.Time
			mode, _                                                    = scope.DB().Get(ScheduleMode)
			comingOnlineMode                                           = mode == ComingOnlineMode
			goingOfflineMode                                           = mode == GoingOfflineMode
			modeON                                                     = (mode != ModeOff) && !comingOnlineMode && !goingOfflineMode
		)

		if v, ok := scope.Get(ScheduledStart); ok {
			if t, ok := v.(*time.Time); ok {
				scheduledStartTime = t
			} else if t, ok := v.(time.Time); ok {
				scheduledStartTime = &t
			}

			if scheduledStartTime != nil {
				if comingOnlineMode {
					conditions = append(conditions, "scheduled_start_at >= ?")
					conditionValues = append(conditionValues, scheduledStartTime)
				} else if goingOfflineMode {
					conditions = append(conditions, "scheduled_end_at >= ?")
					conditionValues = append(conditionValues, scheduledStartTime)
				} else if modeON {
					conditions = append(conditions, "(scheduled_end_at IS NULL OR scheduled_end_at >= ?)")
					conditionValues = append(conditionValues, scheduledStartTime)
				}
			}
		}

		if v, ok := scope.Get(ScheduledEnd); ok {
			if t, ok := v.(*time.Time); ok {
				scheduledEndTime = t
			} else if t, ok := v.(time.Time); ok {
				scheduledEndTime = &t
			}

			if scheduledEndTime != nil {
				if comingOnlineMode {
					conditions = append(conditions, "scheduled_start_at <= ?")
					conditionValues = append(conditionValues, scheduledEndTime)
				} else if goingOfflineMode {
					conditions = append(conditions, "scheduled_end_at <= ?")
					conditionValues = append(conditionValues, scheduledEndTime)
				} else if modeON {
					conditions = append(conditions, "(scheduled_start_at IS NULL OR scheduled_start_at <= ?)")
					conditionValues = append(conditionValues, scheduledEndTime)
				}
			}
		}

		if len(conditions) == 0 {
			if v, ok := scope.Get(ScheduledTime); ok {
				if t, ok := v.(*time.Time); ok {
					scheduledCurrentTime = t
				} else if t, ok := v.(time.Time); ok {
					scheduledCurrentTime = &t
				}
			}

			if scheduledCurrentTime == nil {
				now := gorm.NowFunc()
				scheduledCurrentTime = &now
			}

			if comingOnlineMode {
				conditions = append(conditions, "scheduled_start_at >= ?")
				conditionValues = append(conditionValues, scheduledCurrentTime)
			} else if goingOfflineMode {
				conditions = append(conditions, "scheduled_end_at >= ?")
				conditionValues = append(conditionValues, scheduledCurrentTime)
			} else if modeON {
				conditions = append(conditions, "(scheduled_start_at IS NULL OR scheduled_start_at <= ?) AND (scheduled_end_at IS NULL OR scheduled_end_at >= ?)")
				conditionValues = append(conditionValues, scheduledCurrentTime, scheduledCurrentTime)
			}
		}
	}

	if isPublishReadyable {
		switch mode, _ := scope.DB().Get(VisibleMode); mode {
		case ModeOff:
		default:
			conditions = append(conditions, "publish_ready = ?")
			conditionValues = append(conditionValues, true)
		}
	}

	if isVersionable {
		switch mode, _ := scope.DB().Get(VersionMode); mode {
		case VersionMultipleMode:
			scope.Search.Where(strings.Join(conditions, " AND "), conditionValues...)
		default:
			if versionName, ok := scope.DB().Get(VersionNameMode); ok && versionName != "" {
				scope.Search.Where("version_name = ?", versionName)
			} else {
				var sql string
				var primaryKeys []string

				if scope.HasColumn("DeletedAt") {
					conditions = append(conditions, "deleted_at IS NULL")
				}

				for _, primaryField := range scope.PrimaryFields() {
					if primaryField.DBName != "version_name" {
						primaryKeys = append(primaryKeys, fmt.Sprintf("%v.%v", scope.TableName(), primaryField.DBName))
					}
				}

				primaryKeyCondition := strings.Join(primaryKeys, ",")
				if len(conditions) == 0 {
					sql = fmt.Sprintf("(%v, %v.version_priority) IN (SELECT %v, MAX(%v.version_priority) FROM %v GROUP BY %v)", primaryKeyCondition, scope.QuotedTableName(), primaryKeyCondition, scope.QuotedTableName(), scope.QuotedTableName(), primaryKeyCondition)
				} else {
					sql = fmt.Sprintf("(%v, %v.version_priority) IN (SELECT %v, MAX(%v.version_priority) FROM %v WHERE %v GROUP BY %v)", primaryKeyCondition, scope.QuotedTableName(), primaryKeyCondition, scope.QuotedTableName(), scope.QuotedTableName(), strings.Join(conditions, " AND "), primaryKeyCondition)
				}

				scope.Search.Where(sql, conditionValues...)
			}
		}

		quotedTableName := scope.QuotedTableName()
		scope.Search.Order(fmt.Sprintf("%v.%v, %v.version_priority DESC", quotedTableName, scope.Quote(scope.PrimaryKey()), quotedTableName))
	} else {
		if isShareableVersion {
			var versionName string
			if source, ok := scope.DB().Get("gorm:association:source"); ok {
				if versionable, ok := source.(VersionableInterface); ok {
					versionName = versionable.GetVersionName()
				}
			}

			if v, ok := scope.DB().Get(VersionNameMode); ok && fmt.Sprint(v) != "" {
				versionName = fmt.Sprint(v)
			}

			if versionName != "" {
				var primaryKeys []string
				for _, primaryField := range scope.PrimaryFields() {
					if primaryField.DBName != "version_name" {
						primaryKeys = append(primaryKeys, scope.Quote(primaryField.DBName))
					}
				}

				primaryKeyCondition := strings.Join(primaryKeys, ",")

				scope.Search.Where(
					fmt.Sprintf("version_name = ? OR (version_name = ? AND (%v) NOT IN (SELECT %v FROM %v WHERE version_name = ?))", primaryKeyCondition, primaryKeyCondition, scope.QuotedTableName()),
					versionName, "", versionName,
				)
			}
		}

		scope.Search.Where(strings.Join(conditions, " AND "), conditionValues...)
	}
}

func fixPreloadCallback(scope *gorm.Scope) {
	filterFilterValuesWithVersion := func(gormField *gorm.Field, versionName string) {
		indirectFieldValue := reflect.Indirect(gormField.Field)
		switch indirectFieldValue.Kind() {
		case reflect.Slice:
			resultsMap := map[string]int{}
			results := reflect.New(indirectFieldValue.Type()).Elem()

			for i := 0; i < indirectFieldValue.Len(); i++ {
				if shareableVersion, ok := indirectFieldValue.Index(i).Addr().Interface().(ShareableVersionInterface); ok {
					fieldPrimaryValue := fmt.Sprint(scope.New(shareableVersion).PrimaryKeyValue())
					idx, ok := resultsMap[fieldPrimaryValue]

					if !ok && (shareableVersion.GetSharedVersionName() == versionName || shareableVersion.GetSharedVersionName() == "") {
						resultsMap[fieldPrimaryValue] = results.Len()
						results = reflect.Append(results, indirectFieldValue.Index(i))
					} else if shareableVersion.GetSharedVersionName() == versionName {
						results.Index(idx).Set(indirectFieldValue.Index(i))
					}
				}
			}

			gormField.Set(results)
		case reflect.Struct:
			if shareableVersion, ok := indirectFieldValue.Interface().(ShareableVersionInterface); ok {
				if shareableVersion.GetSharedVersionName() != "" && shareableVersion.GetSharedVersionName() != versionName {
					gormField.Set(reflect.New(indirectFieldValue.Type()))
				}
			}
		}
	}

	fixSharedVersionRecords := func(value interface{}, fieldName string) {
		reflectValue := reflect.Indirect(reflect.ValueOf(value))

		switch reflectValue.Kind() {
		case reflect.Slice:
			for i := 0; i < reflectValue.Len(); i++ {
				v := reflectValue.Index(i)
				if versionable, ok := v.Addr().Interface().(VersionableInterface); ok {
					if fieldValue, ok := scope.New(v.Addr().Interface()).FieldByName(fieldName); ok {
						filterFilterValuesWithVersion(fieldValue, versionable.GetVersionName())
					}
				}
			}
		case reflect.Struct:
			if versionable, ok := value.(VersionableInterface); ok {
				if fieldValue, ok := scope.New(value).FieldByName(fieldName); ok {
					filterFilterValuesWithVersion(fieldValue, versionable.GetVersionName())
				}
			}
		}
	}

	if IsVersionableModel(scope.Value) {
		for _, field := range scope.Fields() {
			if IsShareableVersionModel(reflect.New(field.Struct.Type).Interface()) {
				fixSharedVersionRecords(scope.Value, field.Name)
			}
		}
	}
}

func createCallback(scope *gorm.Scope) {
	if IsVersionableModel(scope.Value) {
		if field, ok := scope.FieldByName("VersionName"); ok {
			if field.IsBlank {
				field.Set(DefaultVersionName)
			}
		}

		updateVersionPriority(scope)
	}

	if IsShareableVersionModel(scope.Value) {
		if field, ok := scope.FieldByName("VersionName"); ok {
			field.IsBlank = false
		}
	}
}

func updateCallback(scope *gorm.Scope) {
	if IsVersionableModel(scope.Value) {
		updateVersionPriority(scope)
	}
}

func deleteCallback(scope *gorm.Scope) {
	if versionName, ok := scope.DB().Get(VersionNameMode); ok && versionName != "" {
		if IsVersionableModel(scope.Value) || IsShareableVersionModel(scope.Value) {
			scope.Search.Where("version_name = ?", versionName)
		}
	}
}

func updateVersionPriority(scope *gorm.Scope) {
	if field, ok := scope.FieldByName("VersionPriority"); ok {
		var scheduledTime *time.Time
		var versionName string

		if scheduled, ok := scope.Value.(ScheduledInterface); ok {
			scheduledTime = scheduled.GetScheduledStartAt()
		}

		if scheduledTime == nil {
			unix := time.Unix(0, 0)
			scheduledTime = &unix
		}

		if versionable, ok := scope.Value.(VersionableInterface); ok {
			versionName = versionable.GetVersionName()
		}

		priority := fmt.Sprintf("%v_%v", scheduledTime.UTC().Format(time.RFC3339), versionName)
		field.Set(priority)
	}
}
