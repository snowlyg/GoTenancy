package publish2_test

import (
	"testing"
	"time"

	"GoTenancy/libs/l10n"
	"GoTenancy/libs/publish2"
	"github.com/jinzhu/gorm"
)

var dbGlobal, dbCN, dbEN *gorm.DB

func init() {
	dbGlobal = DB
	dbCN = dbGlobal.Set("l10n:locale", "zh")
	dbEN = dbGlobal.Set("l10n:locale", "en")
}

type L10nProduct struct {
	gorm.Model
	Name string
	Body string

	publish2.Version
	publish2.Schedule
	publish2.Visible

	l10n.Locale
}

func TestL10nWithVersions(t *testing.T) {
	name := "product 1"
	now := time.Now()
	oneDayAgo := now.Add(-24 * time.Hour)
	oneDayLater := now.Add(24 * time.Hour)

	product := L10nProduct{Name: name}
	product.SetPublishReady(true)
	dbGlobal.Create(&product)
	dbEN.Create(&product)
	dbCN.Create(&product)

	var productCN, productEN L10nProduct

	dbEN.First(&productEN, "id = ?", product.ID)
	dbCN.First(&productCN, "id = ?", product.ID)

	productEN.SetVersionName("v1-en")
	productEN.Body = name + " - v1-EN"
	productEN.SetScheduledStartAt(&oneDayAgo)
	productEN.SetScheduledEndAt(&now)
	dbEN.Save(&productEN)

	productEN.SetVersionName("v2-en")
	productEN.Body = name + " - v2-EN"
	productEN.SetPublishReady(false)
	productEN.SetScheduledStartAt(&oneDayAgo)
	productEN.SetScheduledEndAt(&oneDayLater)
	dbEN.Save(&productEN)

	productEN.SetVersionName("v3-en")
	productEN.Body = name + " - v3-EN"
	productEN.SetPublishReady(true)
	productEN.SetScheduledStartAt(&now)
	productEN.SetScheduledEndAt(&oneDayLater)
	dbEN.Save(&productEN)

	productCN.SetVersionName("v1-cn")
	productCN.Body = name + " - v1-CN"
	productCN.SetScheduledStartAt(&oneDayAgo)
	productCN.SetScheduledEndAt(&now)
	dbCN.Save(&productCN)

	productCN.SetVersionName("v2-cn")
	productCN.Body = name + " - v2-CN"
	productCN.SetPublishReady(false)
	productCN.SetScheduledStartAt(&oneDayAgo)
	productCN.SetScheduledEndAt(&oneDayLater)
	dbCN.Save(&productCN)

	var count int
	dbEN.Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 1 {
		t.Errorf("Should only find one valid product, but got %v", count)
	}

	dbEN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(-time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("EN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbCN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(-time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("CN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbEN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("EN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbCN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 1 {
		t.Errorf("CN: Should only find one valid product when scheduled time, but got %v", count)
	}

	dbEN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("EN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbCN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 1 {
		t.Errorf("CN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbEN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(-time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 3 {
		t.Errorf("EN: Should only find two valid product when scheduled time, but got %v", count)
	}

	dbCN.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(-time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&L10nProduct{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("CN: Should only find two valid product when scheduled time, but got %v", count)
	}
}
