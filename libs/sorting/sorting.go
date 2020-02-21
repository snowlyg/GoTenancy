package sorting

import (
	"encoding/json"
	"fmt"
	"reflect"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/publish"
	"github.com/jinzhu/gorm"
)

type sortingInterface interface {
	GetPosition() int
	SetPosition(int)
}

type sortingDescInterface interface {
	GetPosition() int
	SetPosition(int)
	SortingDesc()
}

// Sorting ascending mode
type Sorting struct {
	Position int `sql:"DEFAULT:NULL"`
}

// GetPosition get current position
func (position Sorting) GetPosition() int {
	return position.Position
}

// SetPosition set position, only set field value, won't save
func (position *Sorting) SetPosition(pos int) {
	position.Position = pos
}

// SortingDESC descending mode
type SortingDESC struct {
	Sorting
}

// SortingDesc make your model sorting desc by default
func (SortingDESC) SortingDesc() {}

func init() {
	admin.RegisterViewPath("libs/sorting/views")
}

func newModel(value interface{}) interface{} {
	return reflect.New(reflect.Indirect(reflect.ValueOf(value)).Type()).Interface()
}

func move(db *gorm.DB, value sortingInterface, pos int) (err error) {
	var startedTransaction bool
	var tx = db.Set("publish:publish_event", true)
	if t := tx.Begin(); t.Error == nil {
		startedTransaction = true
		tx = t
	}

	scope := db.NewScope(value)
	for _, field := range scope.PrimaryFields() {
		if field.DBName != "id" {
			tx = tx.Where(fmt.Sprintf("%s = ?", field.DBName), field.Field.Interface())
		}
	}

	currentPos := value.GetPosition()

	var results *gorm.DB
	if pos > 0 {
		results = tx.Model(newModel(value)).
			Where("position > ? AND position <= ?", currentPos, currentPos+pos).
			UpdateColumn("position", gorm.Expr("position - ?", 1))
	} else {
		results = tx.Model(newModel(value)).
			Where("position < ? AND position >= ?", currentPos, currentPos+pos).
			UpdateColumn("position", gorm.Expr("position + ?", 1))
	}

	if err = results.Error; err == nil {
		var rowsAffected = int(results.RowsAffected)
		if pos < 0 {
			rowsAffected = -rowsAffected
		}
		value.SetPosition(currentPos + rowsAffected)
		err = tx.Model(value).UpdateColumn("position", gorm.Expr("position + ?", rowsAffected)).Error
	}

	// Create Publish Event
	createPublishEvent(tx, value)

	if startedTransaction {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return err
}

func createPublishEvent(db *gorm.DB, value interface{}) (err error) {
	// Create Publish Event in Draft Mode
	if publish.IsDraftMode(db) && publish.IsPublishableModel(value) {
		scope := db.NewScope(value)
		var sortingPublishEvent = changedSortingPublishEvent{
			Table: scope.TableName(),
		}
		for _, field := range scope.PrimaryFields() {
			sortingPublishEvent.PrimaryKeys = append(sortingPublishEvent.PrimaryKeys, field.DBName)
		}

		var result []byte
		if result, err = json.Marshal(sortingPublishEvent); err == nil {
			err = db.New().Where("publish_status = ?", publish.DIRTY).Where(map[string]interface{}{
				"name":     "changed_sorting",
				"argument": string(result),
			}).Attrs(map[string]interface{}{
				"publish_status": publish.DIRTY,
				"description":    "Changed sort order for " + scope.GetModelStruct().ModelType.Name(),
			}).FirstOrCreate(&publish.PublishEvent{}).Error
		}
	}
	return
}

// MoveUp move position up
func MoveUp(db *gorm.DB, value sortingInterface, pos int) error {
	return move(db, value, -pos)
}

// MoveDown move position down
func MoveDown(db *gorm.DB, value sortingInterface, pos int) error {
	return move(db, value, pos)
}

// MoveTo move position to
func MoveTo(db *gorm.DB, value sortingInterface, pos int) error {
	return move(db, value, pos-value.GetPosition())
}
