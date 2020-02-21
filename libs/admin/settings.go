package admin

import (
	"encoding/json"
	"fmt"

	"GoTenancy/libs/qor"
	"github.com/jinzhu/gorm"
)

// SettingsStorageInterface settings storage interface
type SettingsStorageInterface interface {
	Get(key string, value interface{}, context *Context) error
	Save(key string, value interface{}, res *Resource, user qor.CurrentUser, context *Context) error
}

func newSettings(db *gorm.DB) SettingsStorageInterface {
	if db != nil {
		db.AutoMigrate(&QorAdminSetting{})
	}
	return settings{}
}

// QorAdminSetting admin settings
type QorAdminSetting struct {
	gorm.Model
	Key      string
	Resource string
	UserID   string
	Value    string `gorm:"size:65532"`
}

type settings struct{}

// Get load admin settings
func (settings) Get(key string, value interface{}, context *Context) error {
	var (
		settings  = []QorAdminSetting{}
		tx        = context.GetDB().New()
		resParams = ""
		userID    = ""
	)
	sqlCondition := fmt.Sprintf("%v = ? AND (resource = ? OR resource = ?) AND (user_id = ? OR user_id = ?)", tx.NewScope(nil).Quote("key"))

	if context.Resource != nil {
		resParams = context.Resource.ToParam()
	}

	if context.CurrentUser != nil {
		userID = ""
	}

	tx.Where(sqlCondition, key, resParams, "", userID, "").Order("user_id DESC, resource DESC, id DESC").Find(&settings)

	for _, setting := range settings {
		if err := json.Unmarshal([]byte(setting.Value), value); err != nil {
			return err
		}
	}

	return nil
}

// Save save admin settings
func (settings) Save(key string, value interface{}, res *Resource, user qor.CurrentUser, context *Context) error {
	var (
		tx          = context.GetDB().New()
		result, err = json.Marshal(value)
		resParams   = ""
		userID      = ""
	)

	if err != nil {
		return err
	}

	if res != nil {
		resParams = res.ToParam()
	}

	if user != nil {
		userID = ""
	}

	err = tx.Where(QorAdminSetting{
		Key:      key,
		UserID:   userID,
		Resource: resParams,
	}).Assign(QorAdminSetting{Value: string(result)}).FirstOrCreate(&QorAdminSetting{}).Error

	return err
}
