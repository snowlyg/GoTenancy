package publish

import (
	"errors"
	"fmt"

	"GoTenancy/libs/qor"
	"github.com/jinzhu/gorm"
)

// EventInterface defined methods needs for a publish event
type EventInterface interface {
	Publish(db *gorm.DB, event PublishEventInterface) error
	Discard(db *gorm.DB, event PublishEventInterface) error
}

var events = map[string]EventInterface{}

// RegisterEvent register publish event
func RegisterEvent(name string, event EventInterface) {
	events[name] = event
}

// PublishEvent default publish event model
type PublishEvent struct {
	gorm.Model
	Name          string
	Description   string
	Argument      string `sql:"size:65532"`
	PublishStatus bool
	PublishedBy   string
}

func getCurrentUser(db *gorm.DB) (string, bool) {
	if user, hasUser := db.Get("qor:current_user"); hasUser {
		var currentUser string
		if primaryField := db.NewScope(user).PrimaryField(); primaryField != nil {
			currentUser = fmt.Sprintf("%v", primaryField.Field.Interface())
		} else {
			currentUser = fmt.Sprintf("%v", user)
		}

		return currentUser, true
	}

	return "", false
}

// Publish publish data
func (publishEvent *PublishEvent) Publish(db *gorm.DB) error {
	if event, ok := events[publishEvent.Name]; ok {
		err := event.Publish(db, publishEvent)
		if err == nil {
			var updateAttrs = map[string]interface{}{"PublishStatus": PUBLISHED}
			if user, hasUser := getCurrentUser(db); hasUser {
				updateAttrs["PublishedBy"] = user
			}
			err = db.Model(publishEvent).Update(updateAttrs).Error
		}
		return err
	}
	return errors.New("event not found")
}

// Discard discard data
func (publishEvent *PublishEvent) Discard(db *gorm.DB) error {
	if event, ok := events[publishEvent.Name]; ok {
		err := event.Discard(db, publishEvent)
		if err == nil {
			var updateAttrs = map[string]interface{}{"PublishStatus": PUBLISHED}
			if user, hasUser := getCurrentUser(db); hasUser {
				updateAttrs["PublishedBy"] = user
			}
			err = db.Model(publishEvent).Update(updateAttrs).Error
		}
		return err
	}
	return errors.New("event not found")
}

// VisiblePublishResource force to display publish event in publish drafts even it is hidden in the menus
func (PublishEvent) VisiblePublishResource(*qor.Context) bool {
	return true
}
