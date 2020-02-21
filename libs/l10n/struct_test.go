package l10n_test

import (
	"time"

	"GoTenancy/libs/l10n"
	"GoTenancy/libs/qor/test/utils"
	"github.com/jinzhu/gorm"
)

type Product struct {
	ID              int    `gorm:"primary_key"`
	Code            string `l10n:"sync"`
	Quantity        uint   `l10n:"sync"`
	Name            string
	DeletedAt       *time.Time
	ColorVariations []ColorVariation
	BrandID         uint `l10n:"sync"`
	Brand           Brand
	Tags            []Tag      `gorm:"many2many:product_tags"`
	Categories      []Category `gorm:"many2many:product_categories;ForeignKey:id;AssociationForeignKey:id"`
	l10n.Locale
}

// func (Product) LocaleCreatable() {}

type ColorVariation struct {
	ID       int `gorm:"primary_key"`
	Quantity int
	Color    Color
}

type Color struct {
	ID   int `gorm:"primary_key"`
	Code string
	Name string
	l10n.Locale
}

type Brand struct {
	ID   int `gorm:"primary_key"`
	Name string
	l10n.Locale
}

type Tag struct {
	ID   int `gorm:"primary_key"`
	Name string
	l10n.Locale
}

type Category struct {
	ID   int `gorm:"primary_key"`
	Name string
	l10n.Locale
}

var dbGlobal, dbCN, dbEN *gorm.DB

func init() {
	db := utils.TestDB()
	l10n.RegisterCallbacks(db)

	db.DropTableIfExists(&Product{})
	db.DropTableIfExists(&Brand{})
	db.DropTableIfExists(&Tag{})
	db.DropTableIfExists(&Category{})
	db.Exec("drop table product_tags;")
	db.Exec("drop table product_categories;")
	db.AutoMigrate(&Product{}, &Brand{}, &Tag{}, &Category{})

	dbGlobal = db
	dbCN = dbGlobal.Set("l10n:locale", "zh")
	dbEN = dbGlobal.Set("l10n:locale", "en")
}
