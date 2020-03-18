// build ignore

package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

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
	//Fake, _ = faker.New("en")
	//Fake.Rand = rand.New(rand.NewSource(42))
	//rand.Seed(time.Now().UnixNano())

	filepaths, _ := filepath.Glob(filepath.Join("seeder", "data", "*.yml"))
	if err := configor.Load(&Seeds, filepaths...); err != nil {
		panic(err)
	}
}

// CreateAdminUsers 新建管理员数据
func CreateAdminUsers() {
	admin := &models.User{
		Username: "username",
		Name:     "超级管理员",
		Email:    "admin@admin.com",
		Telphone: "13800138000",
		City:     "东莞",
		Model:    gorm.Model{CreatedAt: time.Now()},
	}

	if err := sysinit.UserService.Create("password", admin); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}
}

// CreatePerms 新建菜单数据
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

/*
	AutoMigrates 重置数据表

	sysinit.Db.DropTableIfExists 删除存在数据表
	sysinit.Db.AutoMigrate 重建数据表
*/
func AutoMigrates() {
	sysinit.Db.DropTableIfExists("users", "perms")

	sysinit.Db.AutoMigrate(
		&models.User{},
		&models.Perm{},
	)
}
