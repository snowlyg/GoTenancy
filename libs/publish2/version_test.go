package publish2_test

import (
	"testing"
	"time"

	"GoTenancy/libs/publish2"
	"github.com/jinzhu/gorm"
)

type Wiki struct {
	gorm.Model
	Title string
	Body  string
	publish2.Version
}

func TestVersions(t *testing.T) {
	var wiki = Wiki{Title: "wiki 1", Body: "wiki 1"}
	DB.Create(&wiki)

	wiki.SetVersionName("v1")
	wiki.Body = "wiki 1 - v1"
	DB.Save(&wiki)

	wiki.SetVersionName("v2")
	wiki.Body = "wiki 1 - v2"
	DB.Save(&wiki)

	var count int
	DB.Model(&Wiki{}).Where("id = ?", wiki.ID).Count(&count)
	if count != 1 {
		t.Errorf("Should only find one version for wiki, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Model(&Wiki{}).Where("id = ?", wiki.ID).Count(&count)
	if count != 3 {
		t.Errorf("Should find all versions for wiki when with multiple mode, but got %v", count)
	}

	DB.Set(publish2.VersionNameMode, "v2").Delete(&Wiki{}, "id = ?", wiki.ID)
	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Model(&Wiki{}).Where("id = ?", wiki.ID).Count(&count)
	if count != 2 {
		t.Errorf("After delete version v2, should only have 2 records left, but got %v", count)
	}
}

type Post struct {
	gorm.Model
	Title string
	Body  string
	publish2.Version
	publish2.Schedule
}

func TestVersionsWithSchedule(t *testing.T) {
	now := time.Now()
	oneDayLater := now.Add(24 * time.Hour)

	post := Post{Title: "post 1", Body: "post 1"}
	DB.Create(&post)

	post.SetVersionName("v1")
	post.Body = "post 1 - v1"
	post.SetScheduledStartAt(&now)
	post.SetScheduledEndAt(&oneDayLater)
	DB.Save(&post)

	post.SetVersionName("v2")
	post.Body = "post 1 - v2"
	post.SetScheduledStartAt(&oneDayLater)
	post.SetScheduledEndAt(nil)
	DB.Save(&post)

	var count int
	DB.Model(&Post{}).Where("id = ?", post.ID).Count(&count)
	if count != 1 {
		t.Errorf("Should have one available post")
	}

	var post1, post2, post3 Post
	DB.Set(publish2.ScheduledTime, now.Add(-24*time.Hour)).Model(&Post{}).Where("id = ?", post.ID).First(&post1)
	if post1.Body != "post 1" {
		t.Errorf("should find default version, but got %v", post1.Body)
	}

	DB.Set(publish2.ScheduledTime, now.Add(6*time.Hour)).Model(&Post{}).Where("id = ?", post.ID).First(&post2)
	if post2.Body != "post 1 - v1" {
		t.Errorf("should find first version, but got %v", post2.Body)
	}

	DB.Set(publish2.ScheduledTime, now.Add(25*time.Hour)).Model(&Post{}).Where("id = ?", post.ID).First(&post3)
	if post3.Body != "post 1 - v2" {
		t.Errorf("should find second version, but got %v", post3.Body)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(6*time.Hour)).Model(&Post{}).Where("id = ?", post.ID).Count(&count)
	if count != 2 {
		t.Errorf("Should find two valid versions for posts that match current schedule, but got %v", count)
	}
}

func TestVersionsWithOverlappedSchedule(t *testing.T) {
	now := time.Now()
	postV1 := prepareOverlappedPost("post 2 - 1")
	postV2 := prepareOverlappedPost("post 3 - 2")

	var post1, post2, post3 Post
	DB.Set(publish2.ScheduledTime, now.Add(-36*time.Hour)).Model(&Post{}).Where("id = ?", postV1.ID).First(&post1)
	if post1.Body != postV1.Title {
		t.Errorf("should find default version, but got %v", post1.Body)
	}

	DB.Set(publish2.ScheduledTime, now.Add(6*time.Hour)).Model(&Post{}).Where("id = ?", postV1.ID).First(&post2)
	if post2.Body != postV1.Title+" - v2" {
		t.Errorf("should find first version, but got %v", post2.Body)
	}

	DB.Set(publish2.ScheduledTime, now.Add(25*time.Hour)).Model(&Post{}).Where("id = ?", postV1.ID).First(&post3)
	if post3.Body != postV1.Title+" - v1" {
		t.Errorf("should find second version, but got %v", post3.Body)
	}

	var count uint
	var postIDs = []uint{postV1.ID, postV2.ID}
	DB.Set(publish2.ScheduledTime, now.Add(-36*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 2 {
		t.Errorf("should only find 2 valid versions, but got %v", count)
	}

	DB.Set(publish2.ScheduledTime, now.Add(6*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 2 {
		t.Errorf("should only find 2 valid versions, but got %v", count)
	}

	DB.Set(publish2.ScheduledTime, now.Add(25*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 2 {
		t.Errorf("should only find 2 valid versions, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(25*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 4 {
		t.Errorf("Should find 4 valid versions for posts that match current schedule, but got %v", count)
	}
}

func TestVersionsWithScheduleRange(t *testing.T) {
	now := time.Now()
	postV1 := prepareOverlappedPost("post 5 - 1")
	postV2 := prepareOverlappedPost("post 5 - 2")

	var count uint
	var postIDs = []uint{postV1.ID, postV2.ID}
	DB.Set(publish2.ScheduledStart, now.Add(-36*time.Hour)).Set(publish2.ScheduledEnd, now.Add(-6*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 2 {
		t.Errorf("should only find 2 valid versions in scheduled range, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(-36*time.Hour)).Set(publish2.ScheduledEnd, now.Add(-6*time.Hour)).Model(&Post{}).Where("id IN (?)", postIDs).Count(&count)
	if count != 4 {
		t.Errorf("should only find 4 valid versions in scheduled range with multiple mode, but got %v", count)
	}
}

func prepareOverlappedPost(name string) *Post {
	now := time.Now()
	oneDayAgo := now.Add(-24 * time.Hour)
	oneDayLater := now.Add(24 * time.Hour)

	post := Post{Title: name, Body: name}
	DB.Create(&post)

	post.SetVersionName("v1")
	post.Body = name + " - v1"
	post.SetScheduledStartAt(&oneDayAgo)
	post.SetScheduledEndAt(nil)
	DB.Save(&post)

	post.SetVersionName("v2")
	post.Body = name + " - v2"
	post.SetScheduledStartAt(&now)
	post.SetScheduledEndAt(&oneDayLater)
	DB.Save(&post)

	return &post
}

type Article struct {
	gorm.Model
	Title string
	Body  string
	publish2.Version
	publish2.Visible
}

func TestVersionsWithPublishReady(t *testing.T) {
	articleV1 := Article{Title: "article 1", Body: "article 1"}
	articleV1.PublishReady = true
	DB.Create(&articleV1)

	articleV1.SetVersionName("v1")
	articleV1.PublishReady = false
	DB.Save(&articleV1)

	articleV1.SetVersionName("v2")
	articleV1.PublishReady = true
	DB.Save(&articleV1)

	articleV2 := Article{Title: "article 2", Body: "article 2"}
	articleV2.PublishReady = true
	DB.Create(&articleV2)

	articleV2.SetVersionName("v1")
	articleV2.PublishReady = false
	DB.Save(&articleV2)

	articleV2.SetVersionName("v2")
	articleV2.PublishReady = false
	DB.Save(&articleV2)

	var count int
	DB.Model(&Article{}).Where("id IN (?)", []uint{articleV1.ID, articleV2.ID}).Count(&count)
	if count != 2 {
		t.Errorf("Should find two articles, but got %v", count)
	}

	var article1, article2 Article
	DB.Model(&Article{}).Where("id = ?", articleV1.ID).First(&article1)
	if article1.VersionName != "v2" {
		t.Errorf("Should find article v2 as it is latest versible version")
	}

	DB.Model(&Article{}).Where("id = ?", articleV2.ID).First(&article2)
	if article2.VersionName != publish2.DefaultVersionName {
		t.Errorf("Should find article w/o version name as no other versions is visible")
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Model(&Article{}).Where("id IN (?)", []uint{articleV1.ID, articleV2.ID}).Count(&count)
	if count != 3 {
		t.Errorf("Should find 3 visible versions for article, but got %v", count)
	}
}

type Product struct {
	gorm.Model
	Name string
	Body string
	publish2.Version
	publish2.Schedule
	publish2.Visible
}

func TestProductWithVersionAndScheduleAndPublishReady(t *testing.T) {
	name := "product 1"
	now := time.Now()
	oneDayAgo := now.Add(-24 * time.Hour)
	oneDayLater := now.Add(24 * time.Hour)

	product := Product{Name: name}
	product.SetPublishReady(true)
	DB.Create(&product)

	product.SetVersionName("v1")
	product.Body = name + " - v1"
	product.SetScheduledStartAt(&oneDayAgo)
	product.SetScheduledEndAt(&now)
	DB.Save(&product)

	product.SetVersionName("v2")
	product.Body = name + " - v2"
	product.SetPublishReady(false)
	product.SetScheduledStartAt(&oneDayAgo)
	product.SetScheduledEndAt(&oneDayLater)
	DB.Save(&product)

	product.SetVersionName("v3")
	product.Body = name + " - v3"
	product.SetPublishReady(true)
	product.SetScheduledStartAt(&now)
	product.SetScheduledEndAt(&oneDayLater)
	DB.Save(&product)

	var count int
	DB.Model(&Product{}).Where("id = ?", product.ID).Count(&count)
	if count != 1 {
		t.Errorf("Should only find one valid product, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(-time.Hour)).Model(&Product{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("Should only find two valid product when scheduled time, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledTime, now.Add(time.Hour)).Model(&Product{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("Should only find two valid product when scheduled time, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&Product{}).Where("id = ?", product.ID).Count(&count)
	if count != 2 {
		t.Errorf("Should only find two valid product when scheduled time, but got %v", count)
	}

	DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.ScheduledStart, now.Add(-time.Hour)).Set(publish2.ScheduledEnd, now.Add(24*time.Hour)).Model(&Product{}).Where("id = ?", product.ID).Count(&count)
	if count != 3 {
		t.Errorf("Should only find two valid product when scheduled time, but got %v", count)
	}
}
