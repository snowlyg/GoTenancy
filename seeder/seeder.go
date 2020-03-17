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
	Menus []struct {
		Title  string `json:"title"`
		Href   string `json:"href"`
		Icon   string `json:"icon"`
		Target string `json:"target"`
		Child  []struct {
			Title  string `json:"title"`
			Href   string `json:"href"`
			Icon   string `json:"icon"`
			Target string `json:"target"`
			Child  []struct {
				Title  string `json:"title"`
				Href   string `json:"href"`
				Icon   string `json:"icon"`
				Target string `json:"target"`
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
		Username:  "username",
		Firstname: "超级管理员",
		Model:     gorm.Model{CreatedAt: time.Now()},
	}

	if err := sysinit.UserService.Create("password", admin); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}
}

// CreateMenus 新建菜单数据
func CreateMenus() {
	for _, m := range Seeds.Menus {
		menu := &models.Menu{
			Model:  gorm.Model{CreatedAt: time.Now()},
			Title:  m.Title,
			Href:   m.Href,
			Icon:   m.Icon,
			Target: m.Target,
		}

		var menuchilds []*models.Menu
		if len(m.Child) > 0 {
			for _, mchild := range m.Child {
				menuchild := &models.Menu{
					Model:  gorm.Model{CreatedAt: time.Now()},
					Title:  mchild.Title,
					Href:   mchild.Href,
					Icon:   mchild.Icon,
					Target: mchild.Target,
				}

				var mmenuchilds []*models.Menu
				if len(mchild.Child) > 0 {
					for _, mmchild := range mchild.Child {
						mmenuchild := &models.Menu{
							Model:  gorm.Model{CreatedAt: time.Now()},
							Title:  mmchild.Title,
							Href:   mmchild.Href,
							Icon:   mmchild.Icon,
							Target: mmchild.Target,
						}
						mmenuchilds = append(mmenuchilds, mmenuchild)
					}
				}
				menuchild.Child = mmenuchilds

				menuchilds = append(menuchilds, menuchild)
			}
			menu.Child = menuchilds
		}

		if err := sysinit.MenuService.Create(menu); err != nil {
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
	sysinit.Db.DropTableIfExists("users", "menus")

	sysinit.Db.AutoMigrate(
		&models.User{},
		&models.Menu{},
	)
}
