package oss_test

import (
	"fmt"
	"image"
	"image/gif"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"GoTenancy/libs/media"
	"GoTenancy/libs/media/oss"
	"GoTenancy/libs/oss/s3"
	"GoTenancy/libs/qor/test/utils"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
)

var (
	db       = utils.TestDB()
	S3Config = struct {
		AccessKeyID     string `env:"AWS_ACCESS_KEY_ID"`
		SecretAccessKey string `env:"AWS_SECRET_ACCESS_KEY"`
		Region          string `env:"AWS_Region"`
		S3Bucket        string `env:"AWS_Bucket"`
	}{}
)

func init() {
	configor.Load(&S3Config)
	if S3Config.AccessKeyID != "" {
		log.Println("Testing S3...")
		oss.Storage = s3.New(&s3.Config{
			AccessID:  S3Config.AccessKeyID,
			AccessKey: S3Config.SecretAccessKey,
			Region:    S3Config.Region,
			Bucket:    S3Config.S3Bucket,
		})
	} else {
		log.Println("Testing FileSystem, S3 Access Key need to be specfied to test S3...")
	}
}

type MyOSS struct {
	oss.OSS
}

func (MyOSS) GetSizes() map[string]*media.Size {
	return map[string]*media.Size{
		"small1": {20, 10, false},
		"small2": {20, 10, false},
		"square": {30, 30, false},
		"big":    {50, 50, false},
	}
}

type User struct {
	gorm.Model
	Name    string
	Avatar  MyOSS
	Avatar2 oss.OSS `sql:"size:4294967295;" media_library:"url:/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}"`
}

func init() {
	if err := db.DropTableIfExists(&User{}).Error; err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	media.RegisterCallbacks(db)
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

func getFile(filePath string) (*os.File, bool) {
	if strings.HasPrefix(filePath, "//") {
		filePath = fmt.Sprintf("http:%v", filePath)
	}

	if strings.HasPrefix(filePath, "http:") {
		resp, err := http.Get(filePath)
		if err != nil || resp.StatusCode != http.StatusOK {
			return nil, false
		}

		file, err := ioutil.TempFile(os.TempDir(), "media")
		_, err = io.Copy(file, resp.Body)
		file.Seek(0, 0)
		return file, err == nil
	}

	f, err := os.Open(filepath.Join("public", filePath))
	return f, err == nil
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

	if _, ok := getFile(user.Avatar.URL()); !ok {
		t.Errorf("%v is an invalid path", user.Avatar.URL())
	}

	styleCases := []struct {
		styles []string
	}{
		{[]string{"big"}},
		{[]string{"small1", "small2"}},
	}
	for _, c := range styleCases {
		filePath = user.Avatar.URL(c.styles...)
		if _, ok := getFile(filePath); !ok {
			t.Errorf("%v is an invalid path", user.Avatar.URL())
		}
		if strings.Split(path.Base(filePath), ".")[2] != c.styles[0] {
			t.Errorf(`media.Base#URL(%q) == %q, it's a wrong path`, strings.Join(c.styles, ","), filePath)
		}
	}
}

func TestSaveIntoOSS(t *testing.T) {
	var user = User{Name: "jinzhu"}
	if avatar, err := os.Open("test/logo.png"); err == nil {
		avatarStat, _ := avatar.Stat()
		user.Avatar.Scan(avatar)

		if err := db.Save(&user).Error; err == nil {
			if _, ok := getFile(user.Avatar.URL()); !ok {
				t.Errorf("should find saved user avatar")
			}

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar.Scan(`{"CropOptions": {"small1": {"X": 5, "Y": 5, "Height": 10, "Width": 20}, "small2": {"X": 0, "Y": 0, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar.URL() == user.Avatar.URL() {
				t.Errorf("url should be different after crop")
			}

			file, hasFile := getFile(newUser.Avatar.URL("small1"))
			if !hasFile {
				t.Errorf("Failed open croped image")
			}

			if image, _, err := image.DecodeConfig(file); err == nil {
				if image.Width != 20 || image.Height != 10 {
					t.Errorf("image should be croped successfully")
				}
			} else {
				t.Errorf("Failed to decode croped image, got %v", err)
			}

			originalFile, hasFile := getFile(newUser.Avatar.URL("original"))
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

func TestSaveGifIntoOSS(t *testing.T) {
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
			if _, ok := getFile(user.Avatar.URL()); !ok {
				t.Errorf("should find saved user avatar")
			}

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar.Scan(`{"CropOptions": {"small1": {"X": 5, "Y": 5, "Height": 10, "Width": 20}, "small2": {"X": 0, "Y": 0, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar.URL() == user.Avatar.URL() {
				t.Errorf("url should be different after crop")
			}

			file, hasFile := getFile(newUser.Avatar.URL("small1"))
			if !hasFile {
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

			originalFile, hasFile := getFile(newUser.Avatar.URL("original"))
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
			if _, hasFile := getFile(user.Avatar2.URL()); !hasFile {
				t.Errorf("should find saved user avatar")
			}

			var newUser User
			db.First(&newUser, user.ID)
			newUser.Avatar2.Scan(`{"CropOptions": {"original": {"X": 5, "Y": 5, "Height": 10, "Width": 20}}, "Crop": true}`)
			db.Save(&newUser)

			if newUser.Avatar2.URL() != user.Avatar2.URL() {
				t.Errorf("url should be same after crop")
			}

			file, hasFile := getFile(newUser.Avatar2.URL())
			if !hasFile {
				t.Errorf("Failed open croped image")
			}

			if image, _, err := image.DecodeConfig(file); err == nil {
				if image.Width != 20 || image.Height != 10 {
					t.Errorf("image should be croped successfully")
				}
			} else {
				t.Errorf("Failed to decode croped image")
			}

			originalFile, hasFile := getFile(newUser.Avatar2.URL("original"))
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
