package database

import (
	"fmt"
	"reflect"
	"strconv"

	"GoTenancy/libs/notification"
	"GoTenancy/libs/qor"
	"github.com/jinzhu/gorm"
)

type Config struct {
	DB *gorm.DB
}

func New(config *Config) *Database {
	if config.DB != nil {
		config.DB.AutoMigrate(&notification.QorNotification{})
	} else {
		fmt.Println("Need to have gorm DB in the configuration in order to run migrations")
	}
	return &Database{Config: config}
}

type Database struct {
	Config *Config
}

func (database *Database) Send(message *notification.Message, context *qor.Context) error {
	notice := notification.QorNotification{
		From:        database.getUserID(message.From, context),
		To:          database.getUserID(message.To, context),
		Title:       message.Title,
		Body:        message.Body,
		MessageType: message.MessageType,
		ResolvedAt:  message.ResolvedAt,
	}

	return context.GetDB().Save(&notice).Error
}

func (database *Database) GetNotifications(user interface{}, results *notification.NotificationsResult, _ *notification.Notification, context *qor.Context) error {
	var to = database.getUserID(user, context)
	var db = context.GetDB()

	var currentPage, perPage int

	if context.Request != nil {
		if p, err := strconv.Atoi(context.Request.URL.Query().Get("page")); err == nil {
			currentPage = p
		}

		if p, err := strconv.Atoi(context.Request.URL.Query().Get("per_page")); err == nil {
			perPage = p
		}
	}

	if perPage == 0 {
		perPage = 10
	}
	offset := currentPage * perPage

	commonDB := db.Order("created_at DESC").Where(fmt.Sprintf("%v = ?", db.Dialect().Quote("to")), to)

	// get unresolved notifications
	if err := commonDB.Offset(offset).Limit(perPage).Find(&results.Notifications, fmt.Sprintf("%v IS NULL", db.Dialect().Quote("resolved_at"))).Error; err != nil {
		return err
	}

	if len(results.Notifications) == perPage {
		return nil
	}

	if len(results.Notifications) == 0 {
		var unreadedCount int
		commonDB.Model(&notification.QorNotification{}).Where(fmt.Sprintf("%v IS NULL", db.Dialect().Quote("resolved_at"))).Count(&unreadedCount)
		offset -= unreadedCount
	} else if len(results.Notifications) < perPage {
		offset = 0
		perPage -= len(results.Notifications)
	}

	// get resolved notifications
	return commonDB.Offset(offset).Limit(perPage).Find(&results.Resolved, fmt.Sprintf("%v IS NOT NULL", db.Dialect().Quote("resolved_at"))).Error
}

func (database *Database) GetUnresolvedNotificationsCount(user interface{}, _ *notification.Notification, context *qor.Context) uint {
	var to = database.getUserID(user, context)
	var db = context.GetDB()

	var result uint
	db.Model(&notification.QorNotification{}).Where(fmt.Sprintf("%v = ? AND %v IS NULL", db.Dialect().Quote("to"), db.Dialect().Quote("resolved_at")), to).Count(&result)
	return result
}

func (database *Database) GetNotification(user interface{}, notificationID string, _ *notification.Notification, context *qor.Context) (*notification.QorNotification, error) {
	var (
		notice notification.QorNotification
		to     = database.getUserID(user, context)
		db     = context.GetDB()
	)

	err := db.First(&notice, fmt.Sprintf("%v = ? AND %v = ?", db.Dialect().Quote("to"), db.Dialect().Quote("id")), to, notificationID).Error
	return &notice, err
}

func (database *Database) getUserID(user interface{}, context *qor.Context) string {
	var (
		userID string
		scope  = context.GetDB().NewScope(user)
	)

	if scope.IndirectValue().Kind() == reflect.Struct {
		userID = fmt.Sprint(scope.PrimaryKeyValue())
	} else {
		userID = fmt.Sprint(user)
	}

	return userID
}
