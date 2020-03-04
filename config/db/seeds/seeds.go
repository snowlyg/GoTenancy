package main

import (
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"
	//"github.com/qor/publish2"

	"go-tenancy/config/db"
)

var Fake *faker.Faker
var (
	Root, _ = os.Getwd()
	DraftDB = db.DB
	//.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)
)

var Seeds = struct {
	RabcPermissions []struct {
		Name        string
		DisplayName string
		Description string
		Act         string
	}
	RabcRoles []struct {
		Name        string
		DisplayName string
		Description string
	}
	RabcUsers []struct {
		Name     string
		Username string
		Password string
	}

	Tenants []struct {
		Name         string
		Mode         uint8
		Times        uint
		ExpireTime   *time.Time
		CreationTime *time.Time
		State        int8
		IsTop        uint8
		Order        int64
		IsDel        int8
		Amount       float32
		Price        float32
		Logo         string
		Tag          string
		AreaId       int64
		Province     string
		City         string
		County       string
		Addr         string
		Linkman      string
		Phone        string
		Lng          float64
		Lat          float64
		Appid        uint
		FullName     string
		CertifyPics  string
		Desc         string
		Pics         string
		Remark       string

		PermissionKey string
		TenantKey     string
		RabcUsers     []struct {
			Name     string
			Username string
			Password string
		}
	}

	Stores []struct {
		Name      string
		Phone     string
		Email     string
		Country   string
		Zip       string
		City      string
		Region    string
		Address   string
		Latitude  float64
		Longitude float64
	}
	Setting struct {
		ShippingFee     uint
		GiftWrappingFee uint
		CODFee          uint `gorm:"column:cod_fee"`
		TaxRate         int
		Address         string
		City            string
		Region          string
		Country         string
		Zip             string
		Latitude        float64
		Longitude       float64
	}
	Seo struct {
		SiteName    string
		DefaultPage struct {
			Title       string
			Description string
			Keywords    string
		}
		HomePage struct {
			Title       string
			Description string
			Keywords    string
		}
		ProductPage struct {
			Title       string
			Description string
			Keywords    string
		}
	}

	Slides []struct {
		Title    string
		SubTitle string
		Button   string
		Link     string
		Image    string
	}
	MediaLibraries []struct {
		Title string
		Image string
	}
	BannerEditorSettings []struct {
		ID    string
		Kind  string
		Value string
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob(filepath.Join("config", "db", "seeds", "data", "*.yml"))
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

func TruncateTables(tables ...interface{}) {
	for _, table := range tables {
		if err := DraftDB.DropTableIfExists(table).Error; err != nil {
			panic(err)
		}

		DraftDB.AutoMigrate(table)
	}
}
