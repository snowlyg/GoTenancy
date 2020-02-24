package database

import (
	"GoTenancy/libs/i18n"
	"github.com/jinzhu/gorm"
)

// Translation is a struct used to save translations into databae
type Translation struct {
	Locale string `sql:"size:12;"`
	Key    string `sql:"size:4294967295;"`
	Value  string `sql:"size:4294967295;"`
}

// New new DB backend for I18n
func New(db *gorm.DB) i18n.Backend {
	db.AutoMigrate(&Translation{})
	//if err := db.Model(&Translation{}).AddUniqueIndex("idx_translations_key_with_locale", "locale").Error; err != nil {
	//	fmt.Printf("Failed to create unique index for translations key & locale, got: %v\n", err.Error())
	//}
	return &Backend{DB: db}
}

// Backend DB backend
type Backend struct {
	DB *gorm.DB
}

// LoadTranslations load translations from DB backend
func (backend *Backend) LoadTranslations() (translations []*i18n.Translation) {
	backend.DB.Find(&translations)
	return translations
}

// SaveTranslation save translation into DB backend
func (backend *Backend) SaveTranslation(t *i18n.Translation) error {
	return backend.DB.Where(Translation{Key: t.Key, Locale: t.Locale}).
		Assign(Translation{Value: t.Value}).
		FirstOrCreate(&Translation{}).Error
}

// DeleteTranslation delete translation into DB backend
func (backend *Backend) DeleteTranslation(t *i18n.Translation) error {
	return backend.DB.Where(Translation{Key: t.Key, Locale: t.Locale}).Delete(&Translation{}).Error
}
