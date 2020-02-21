package publish

import (
	"fmt"
	"strings"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/utils"
	"GoTenancy/libs/worker"
	"github.com/jinzhu/gorm"

	"reflect"
)

const (
	// PUBLISHED publish status published
	PUBLISHED = false
	// DIRTY publish status dirty
	DIRTY = true

	publishDraftMode = "publish:draft_mode"
	publishEvent     = "publish:publish_event"
)

type publishInterface interface {
	GetPublishStatus() bool
	SetPublishStatus(bool)
}

// PublishEventInterface defined publish event itself's interface
type PublishEventInterface interface {
	Publish(*gorm.DB) error
	Discard(*gorm.DB) error
}

// Status publish status, need to be embedded in your models to get the publish feature
type Status struct {
	PublishStatus bool
}

// GetPublishStatus get publish status
func (s Status) GetPublishStatus() bool {
	return s.PublishStatus
}

// SetPublishStatus set publish status
func (s *Status) SetPublishStatus(status bool) {
	s.PublishStatus = status
}

// ConfigureQorResource configure qor resource for qor admin
func (s Status) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		if res.GetMeta("PublishStatus") == nil {
			res.IndexAttrs(res.IndexAttrs(), "-PublishStatus")
			res.NewAttrs(res.NewAttrs(), "-PublishStatus")
			res.EditAttrs(res.EditAttrs(), "-PublishStatus")
			res.ShowAttrs(res.ShowAttrs(), "-PublishStatus", false)
		}
	}
}

// Publish defined a publish struct
type Publish struct {
	DB              *gorm.DB
	SearchHandler   func(db *gorm.DB, context *qor.Context) *gorm.DB
	WorkerScheduler *worker.Worker
	logger          LoggerInterface
	deleteCallback  func(*gorm.Scope)
}

// IsDraftMode check if current db in draft mode
func IsDraftMode(db *gorm.DB) bool {
	if draftMode, ok := db.Get(publishDraftMode); ok {
		if isDraft, ok := draftMode.(bool); ok && isDraft {
			return true
		}
	}
	return false
}

// IsPublishEvent check if current model is a publish event model
func IsPublishEvent(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(PublishEventInterface)
	}
	return
}

// IsPublishableModel check if current model is a publishable
func IsPublishableModel(model interface{}) (ok bool) {
	if model != nil {
		_, ok = reflect.New(utils.ModelType(model)).Interface().(publishInterface)
	}
	return
}

var injectedJoinTableHandler = map[reflect.Type]bool{}

// New initialize a publish instance
func New(db *gorm.DB) *Publish {
	tableHandler := gorm.DefaultTableNameHandler
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		tableName := tableHandler(db, defaultTableName)

		if db != nil {
			if IsPublishableModel(db.Value) {
				// Set join table handler
				typ := utils.ModelType(db.Value)
				if !injectedJoinTableHandler[typ] {
					injectedJoinTableHandler[typ] = true
					scope := db.NewScope(db.Value)
					for _, field := range scope.GetModelStruct().StructFields {
						if many2many := utils.ParseTagOption(field.Tag.Get("gorm"))["MANY2MANY"]; many2many != "" {
							db.SetJoinTableHandler(db.Value, field.Name, &publishJoinTableHandler{})
							db.AutoMigrate(db.Value)
						}
					}
				}

				var forceDraftTable bool
				if forceDraft, ok := db.Get("publish:force_draft_table"); ok {
					if forceMode, ok := forceDraft.(bool); ok && forceMode {
						forceDraftTable = true
					}
				}

				if IsDraftMode(db) || forceDraftTable {
					return DraftTableName(tableName)
				}
			}
		}
		return tableName
	}

	db.AutoMigrate(&PublishEvent{})

	db.Callback().Create().Before("gorm:begin_transaction").Register("publish:set_table_to_draft", setTableAndPublishStatus(true))
	db.Callback().Create().Before("gorm:commit_or_rollback_transaction").
		Register("publish:sync_to_production_after_create", syncCreateFromProductionToDraft)
	db.Callback().Create().Before("gorm:commit_or_rollback_transaction").Register("gorm:create_publish_event", createPublishEvent)

	db.Callback().Delete().Before("gorm:begin_transaction").Register("publish:set_table_to_draft", setTableAndPublishStatus(true))
	deleteCallback := db.Callback().Delete().Get("gorm:delete")
	db.Callback().Delete().Replace("gorm:delete", deleteScope)
	db.Callback().Delete().Before("gorm:commit_or_rollback_transaction").
		Register("publish:sync_to_production_after_delete", syncDeleteFromProductionToDraft)
	db.Callback().Delete().Before("gorm:commit_or_rollback_transaction").Register("gorm:create_publish_event", createPublishEvent)

	db.Callback().Update().Before("gorm:begin_transaction").Register("publish:set_table_to_draft", setTableAndPublishStatus(true))
	db.Callback().Update().Before("gorm:commit_or_rollback_transaction").
		Register("publish:sync_to_production", syncUpdateFromProductionToDraft)
	db.Callback().Update().Before("gorm:commit_or_rollback_transaction").Register("gorm:create_publish_event", createPublishEvent)

	db.Callback().RowQuery().Before("gorm:row_query").Register("publish:set_table_in_draft_mode", setTableAndPublishStatus(false))
	db.Callback().Query().Before("gorm:query").Register("publish:set_table_in_draft_mode", setTableAndPublishStatus(false))

	searchHandler := func(db *gorm.DB, context *qor.Context) *gorm.DB {
		return db.Unscoped()
	}
	return &Publish{SearchHandler: searchHandler, DB: db, deleteCallback: deleteCallback, logger: Logger}
}

// DraftTableName get draft table name of passed in string
func DraftTableName(table string) string {
	return OriginalTableName(table) + "_draft"
}

// OriginalTableName get original table name of passed in string
func OriginalTableName(table string) string {
	return strings.TrimSuffix(table, "_draft")
}

// AutoMigrate run auto migrate in draft tables
func (pb *Publish) AutoMigrate(values ...interface{}) {
	for _, value := range values {
		tableName := pb.DB.NewScope(value).TableName()
		pb.DraftDB().Table(DraftTableName(tableName)).AutoMigrate(value)
	}
}

// ProductionDB get db in production mode
func (pb Publish) ProductionDB() *gorm.DB {
	return pb.DB.Set(publishDraftMode, false)
}

// DraftDB get db in draft mode
func (pb Publish) DraftDB() *gorm.DB {
	return pb.DB.Set(publishDraftMode, true)
}

// Logger set logger that used to print publish logs
func (pb Publish) Logger(l LoggerInterface) *Publish {
	return &Publish{
		WorkerScheduler: pb.WorkerScheduler,
		DB:              pb.DB,
		logger:          l,
		deleteCallback:  pb.deleteCallback,
	}
}

func (pb Publish) newResolver(records ...interface{}) *resolver {
	return &resolver{publish: pb, Records: records, DB: pb.DB, Dependencies: map[string]*dependency{}}
}

// Publish publish records
func (pb Publish) Publish(records ...interface{}) {
	pb.newResolver(records...).Publish()
}

// Discard discard records
func (pb Publish) Discard(records ...interface{}) {
	pb.newResolver(records...).Discard()
}

func (pb Publish) search(db *gorm.DB, res *admin.Resource, ids [][]string) *gorm.DB {
	var primaryKeys []string
	var primaryValues [][][]interface{}
	var scope = db.NewScope(res.Value)

	for _, primaryField := range scope.PrimaryFields() {
		primaryKeys = append(primaryKeys, fmt.Sprintf("%v.%v", scope.TableName(), primaryField.DBName))
	}

	for _, id := range ids {
		var primaryValue [][]interface{}
		for idx, value := range id {
			primaryValue = append(primaryValue, []interface{}{primaryKeys[idx], value})
		}
		primaryValues = append(primaryValues, primaryValue)
	}

	sql := fmt.Sprintf("%v IN (%v)", toQueryCondition(scope, primaryKeys), toQueryMarks(primaryValues))
	return pb.SearchHandler(db, nil).Where(sql, toQueryValues(primaryValues)...)
}

func (pb Publish) searchWithPublishIDs(db *gorm.DB, Admin *admin.Admin, publishIDs []string) (results []interface{}) {
	var values = map[string][][]string{}

	for _, publishID := range publishIDs {
		if primaryValues := strings.Split(publishID, "__"); len(primaryValues) >= 2 {
			name := primaryValues[0]
			values[name] = append(values[name], primaryValues[1:])
		}
	}

	for name, value := range values {
		res := Admin.GetResource(name)
		result := res.NewSlice()
		if pb.search(db, res, value).Find(result).Error == nil {
			resultValues := reflect.Indirect(reflect.ValueOf(result))
			for i := 0; i < resultValues.Len(); i++ {
				results = append(results, resultValues.Index(i).Interface())
			}
		}
	}

	return
}
