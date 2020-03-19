// build ignore

package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/azumads/faker"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

var Fake *faker.Faker
var Seeds = struct {
	Perms []struct {
		Title       string `json:"title"`
		Href        string `json:"href"`
		Icon        string `json:"icon"`
		Target      string `json:"target"`
		OrderNumber int64  `json:"order_number"`
		Authority   string `json:"authority"`
		Checked     int8   `json:"checked"`
		IsMenu      int8   `json:"is_menu"`
		Child       []struct {
			Title       string `json:"title"`
			Href        string `json:"href"`
			Icon        string `json:"icon"`
			Target      string `json:"target"`
			OrderNumber int64  `json:"order_number"`
			Authority   string `json:"authority"`
			Checked     int8   `json:"checked"`
			IsMenu      int8   `json:"is_menu"`
			Child       []struct {
				Title       string `json:"title"`
				Href        string `json:"href"`
				Icon        string `json:"icon"`
				Target      string `json:"target"`
				OrderNumber int64  `json:"order_number"`
				Authority   string `json:"authority"`
				Checked     int8   `json:"checked"`
				IsMenu      int8   `json:"is_menu"`
			}
		}
	}
}{}

func init() {
	Fake, _ = faker.New("en")
	Fake.Rand = rand.New(rand.NewSource(42))
	rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob(filepath.Join("seeder", "data", "*.yml"))
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

// CreatePerms 新建菜单
func CreatePerms() {
	for _, m := range Seeds.Perms {
		menu := &models.Perm{
			Model:       gorm.Model{CreatedAt: time.Now()},
			Title:       m.Title,
			Href:        m.Href,
			Icon:        m.Icon,
			Target:      m.Target,
			OrderNumber: m.OrderNumber,
			Authority:   m.Authority,
			Checked:     m.Checked,
			IsMenu:      m.IsMenu,
		}

		var menuchilds []*models.Perm
		if len(m.Child) > 0 {
			for _, mchild := range m.Child {
				menuchild := &models.Perm{
					Model:       gorm.Model{CreatedAt: time.Now()},
					Title:       mchild.Title,
					Href:        mchild.Href,
					Icon:        mchild.Icon,
					Target:      mchild.Target,
					OrderNumber: mchild.OrderNumber,
					Authority:   mchild.Authority,
					Checked:     mchild.Checked,
					IsMenu:      mchild.IsMenu,
				}

				var mmenuchilds []*models.Perm
				if len(mchild.Child) > 0 {
					for _, mmchild := range mchild.Child {
						mmenuchild := &models.Perm{
							Model:       gorm.Model{CreatedAt: time.Now()},
							Title:       mmchild.Title,
							Href:        mmchild.Href,
							Icon:        mmchild.Icon,
							Target:      mmchild.Target,
							OrderNumber: mmchild.OrderNumber,
							Authority:   mmchild.Authority,
							Checked:     mmchild.Checked,
							IsMenu:      mmchild.IsMenu,
						}
						mmenuchilds = append(mmenuchilds, mmenuchild)
					}
				}
				menuchild.Child = mmenuchilds

				menuchilds = append(menuchilds, menuchild)
			}
			menu.Child = menuchilds
		}

		if err := sysinit.PermService.Create(menu); err != nil {
			panic(fmt.Sprintf("菜单填充错误：%v", err))
		}
	}

}

// CreateAdminRoles 新建管理角色
func CreateAdminRoles() {
	role := &models.Role{
		Name:        "超级管理员",
		DisplayName: "超级管理员",
		Rmk:         "超级管理员",
		IsAdmin:     sql.NullBool{Bool: true, Valid: true},
		Model:       gorm.Model{CreatedAt: time.Now()},
	}

	if err := sysinit.RoleService.Create(role); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}
}

// CreateAdminUsers 新建管理员
func CreateAdminUsers() {
	admin := &models.User{
		Username: "username",
		Name:     "超级管理员",
		Email:    "admin@admin.com",
		Telphone: "13800138000",
		IsAdmin:  sql.NullBool{Bool: true, Valid: true},
		Model:    gorm.Model{CreatedAt: time.Now()},
	}

	if err := sysinit.UserService.Create("password", admin); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}
}

// CreateUsers 新建用户
func CreateUsers() {
	emailRegexp := regexp.MustCompile(".*(@.*)")
	// 最新手机正则 ^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$
	totalCount := 50
	for i := 0; i < totalCount; i++ {
		name := Fake.Name()
		admin := &models.User{
			Username: name,
			Name:     name,
			Email:    emailRegexp.ReplaceAllString(Fake.Email(), strings.Replace(strings.ToLower(name), " ", "_", -1)+"@example.com"),
			Telphone: createPhoneNumber(i),
			IsAdmin:  sql.NullBool{Bool: false, Valid: true},
			Model:    gorm.Model{CreatedAt: time.Now()},
		}

		if err := sysinit.UserService.Create("password", admin); err != nil {
			panic(fmt.Sprintf("管理员填充错误：%v", err))
		}
	}
}

// 生成随机手机号码
func createPhoneNumber(i int) string {
	prelist := []string{"130", "131", "132", "133", "134", "135", "136", "137", "138", "139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186", "187", "188"}

	rd := rand.New(rand.NewSource(time.Now().UnixNano()))

	phoneNumber := prelist[rd.Intn(23)] + fmt.Sprintf("%08v", rd.Int31n(100000000))
	return phoneNumber
}

/*
	AutoMigrates 重置数据表

	sysinit.Db.DropTableIfExists 删除存在数据表
	sysinit.Db.AutoMigrate 重建数据表
*/
func AutoMigrates() {
	sysinit.Db.DropTableIfExists("users", "perms", "roles")

	sysinit.Db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Perm{},
	)
}
