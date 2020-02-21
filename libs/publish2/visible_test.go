package publish2_test

import (
	"testing"

	"GoTenancy/libs/publish2"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name string
	publish2.Visible
}

func TestPublishReady(t *testing.T) {
	user := User{Name: "user"}
	DB.Create(&user)

	if !DB.First(&User{}, "id = ?", user.ID).RecordNotFound() {
		t.Errorf("Should not able to find created record when publish not ready")
	}

	if DB.Set(publish2.VisibleMode, publish2.ModeOff).First(&User{}, "id = ?", user.ID).RecordNotFound() {
		t.Errorf("Should be able to find created record with visible mode `all`")
	}

	user.SetPublishReady(true)
	DB.Save(&user)

	if DB.First(&User{}, "id = ?", user.ID).RecordNotFound() {
		t.Errorf("Should be able to find created record when publish is ready")
	}
}
