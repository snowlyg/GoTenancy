package filesystem_test

import (
	"image"
	"image/gif"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"GoTenancy/libs/media"
	"GoTenancy/libs/media/filesystem"
	"GoTenancy/libs/qor/test/utils"
	"github.com/jinzhu/gorm"
)

var db = utils.TestDB()

type MyFileSystem struct {
	filesystem.FileSystem
}

func (MyFileSystem) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small1":   {Width: 20, Height: 10},
		"small2":   {Width: 20, Height: 10},
		"square":   {Width: 30, Height: 30},
		"big":      {Width: 50, Height: 50},
		"large":    {Width: 300, Height: 300},
		"slarge":   {Width: 400, Height: 0},
		"sslarge":  {Width: 0, Height: 500},
		"ssslarge": {Width: 0, Height: 0},
	}
}

type User struct {
	gorm.Model
	Name    string
	Avatar  MyFileSystem
	Avatar2 filesystem.FileSystem `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
}

func TestMain(m *testing.M) {
	os.RemoveAll("public")
	if err := db.DropTableIfExists(&User{}).Error; err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	media.RegisterCallbacks(db)

	m.Run()
}

func TestURLWithoutFile(t *testing.T) {
	user := User{Name: "jinzhu"}

	if got, want := user.Avatar.URL(), ""; got != want {
		t.Errorf(`media.Base#URL() == %q, want %q`, got, want)
	}
	if got, want := user.Avatar.URL("big"), ""; got != want {
		t.Errorf(`media.Base#URL("big") == %q, want %q`, got, want)
	}
	if got, want := user.Avatar.URL("small1", "small2"), ""; got != want {
		t.Errorf(`media.Base#URL("small1", "small2") == %q, want %q`, got, want)
	}
}

func TestURLWithFile(t *testing.T) {
	var filePath string
	user := User{Name: "jinzhu"}

	if avatar, err := os.Open("test/logo.png"); err != nil {
		panic("file doesn't exist")
	} else {
		user.Avatar.Scan(avatar)
	}
	if err := db.Save(&user).Error; err != nil {
		panic(err)
	}

	filePath = user.Avatar.URL()
	if _, err := os.Stat(filepath.Join("public", filePath)); err != nil {
		t.Errorf(`media.Base#URL() == %q, it's an invalid path`, filePath)
	}

	styleCases := []struct {
		styles []string
	}{
		{[]string{"big"}},
		{[]string{"small1", "small2"}},
	}
	for _, c := range styleCases {
		filePath = user.Avatar.URL(c.styles...)
		if _, err := os.Stat(filepath.Join("public", filePath)); err != nil {
			t.Errorf(`media.Base#URL(%q) == %q, it's an invalid path`, strings.Join(c.styles, ","), filePath)
		}
		if strings.Split(path.Base(filePath), ".")[2] != c.styles[0] {
			t.Errorf(`media.Base#URL(%q) == %q, it's a wrong path`, strings.Join(c.styles, ","), filePath)
		}
	}
}

func checkUserAvatar(user *User, t *testing.T) {
	for name, size := range user.Avatar.GetSizes() {
		file, err := os.Open(filepath.Join("public", user.Avatar.URL(name)))
		if err != nil {
			t.Errorf("Failed open croped image")
		}

		if image, _, err := image.DecodeConfig(file); err == nil {
			if (size.Width != 0 && image.Width != size.Width) || (image.Width == 0) {
				t.Errorf("image's width is not cropped correctly")
			}
			if (size.Height != 0 && image.Height != size.Height) || (image.Height == 0) {
				t.Errorf("image's height is not cropped correctly")
			}
		} else {
			t.Errorf("Failed to decode croped image, got err %v when decoding %v", err, user.Avatar.URL(name))
		}
	}
}

func TestSaveIntoFileSystem(t *testing.T) {
	var user = User{Name: "jinzhu"}
	if avatar, err := os.Open("test/logo.png"); err == nil {
		avatarStat, _ := avatar.Stat()
		user.Avatar.Scan(avatar)

		if err := db.Save(&user).Error; err == nil {
			if _, err := os.Stat(filepath.Join("public", user.Avatar.URL())); err != nil {
				t.Errorf("should find saved user avatar")
			}

			checkUserAvatar(&user, t)

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar.Scan(`{"CropOptions": {"small1": {"X": 5, "Y": 5, "Height": 10, "Width": 20}, "small2": {"X": 0, "Y": 0, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar.URL() == user.Avatar.URL() {
				t.Errorf("url should be different after crop")
			}

			checkUserAvatar(&newUser, t)

			originalFile, err := os.Open(filepath.Join("public", newUser.Avatar.URL("original")))
			if err != nil {
				t.Errorf("Failed open original image")
			}

			if stat, err := originalFile.Stat(); err != nil {
				t.Errorf("original file should be there")
			} else if avatarStat.Size() != stat.Size() {
				t.Errorf("Original file should not be changed after crop")
			}
		} else {
			t.Errorf("should saved user successfully")
		}
	} else {
		panic("file doesn't exist")
	}
}

func TestSaveGifIntoFileSystem(t *testing.T) {
	var user = User{Name: "jinzhu"}
	if avatar, err := os.Open("test/test.gif"); err == nil {
		avatarStat, _ := avatar.Stat()
		var frames int
		if g, err := gif.DecodeAll(avatar); err == nil {
			frames = len(g.Image)
		}

		avatar.Seek(0, 0)
		user.Avatar.Scan(avatar)
		if err := db.Save(&user).Error; err == nil {
			if _, err := os.Stat(filepath.Join("public", user.Avatar.URL())); err != nil {
				t.Errorf("should find saved user avatar")
			}

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar.Scan(`{"CropOptions": {"small1": {"X": 5, "Y": 5, "Height": 10, "Width": 20}, "small2": {"X": 0, "Y": 0, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar.URL() == user.Avatar.URL() {
				t.Errorf("url should be different after crop")
			}

			file, err := os.Open(filepath.Join("public", newUser.Avatar.URL("small1")))
			if err != nil {
				t.Errorf("Failed open croped image")
			}

			if g, err := gif.DecodeAll(file); err == nil {
				if g.Config.Width != 20 || g.Config.Height != 10 {
					t.Errorf("gif should be croped successfully")
				}

				for _, image := range g.Image {
					if image.Rect.Dx() != 20 || image.Rect.Dy() != 10 {
						t.Errorf("gif's frames should be croped successfully, but it is %vx%v", image.Rect.Dx(), image.Rect.Dy())
					}
				}
				if frames != len(g.Image) || frames == 0 {
					t.Errorf("Gif's frames should be same")
				}
			} else {
				t.Errorf("Failed to decode croped gif image")
			}

			originalFile, err := os.Open(filepath.Join("public", newUser.Avatar.URL("original")))
			if stat, err := originalFile.Stat(); err != nil {
				t.Errorf("original file should be there")
			} else if avatarStat.Size() != stat.Size() {
				t.Errorf("Original file should not be changed after crop")
			}
		} else {
			t.Errorf("should saved user successfully")
		}
	} else {
		panic("file doesn't exist")
	}
}

func TestCropFileWithSameName(t *testing.T) {
	var user = User{Name: "jinzhu"}
	if avatar, err := os.Open("test/logo.png"); err == nil {
		avatarStat, _ := avatar.Stat()
		user.Avatar2.Scan(avatar)

		if err := db.Save(&user).Error; err == nil {
			if _, err := os.Stat(filepath.Join("public", user.Avatar2.URL())); err != nil {
				t.Errorf("should find saved user avatar")
			}

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar2.Scan(`{"CropOptions": {"original": {"X": 5, "Y": 5, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar2.URL() != user.Avatar2.URL() {
				t.Errorf("url should be same after crop")
			}

			file, err := os.Open(filepath.Join("public", newUser.Avatar2.URL()))
			if err != nil {
				t.Errorf("Failed open croped image")
			}

			if image, _, err := image.DecodeConfig(file); err == nil {
				if image.Width != 20 || image.Height != 10 {
					t.Errorf("image should be croped successfully")
				}
			} else {
				t.Errorf("Failed to decode croped image")
			}

			originalFile, err := os.Open(filepath.Join("public", newUser.Avatar2.URL("original")))
			if stat, err := originalFile.Stat(); err != nil {
				t.Errorf("original file should be there")
			} else if avatarStat.Size() != stat.Size() {
				t.Errorf("Original file should not be changed after crop")
			}
		} else {
			t.Errorf("should saved user successfully")
		}
	} else {
		panic("file doesn't exist")
	}
}
