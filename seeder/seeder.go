// build ignore

package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/azumads/faker"
	gormadapter "github.com/casbin/gorm-adapter/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/configor"
	"github.com/jinzhu/gorm"
	"github.com/snowlyg/go-tenancy/lib"
	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

var Fake *faker.Faker
var Seeds = struct {
	Perms []struct {
		Type        int8   `json:"type"`
		Title       string `json:"title"`
		Href        string `json:"href"`
		Icon        string `json:"icon"`
		Target      string `json:"target"`
		OrderNumber int64  `json:"order_number"`
		Method      string `json:"method"`
		Checked     int8   `json:"checked"`
		IsMenu      int8   `json:"is_menu"`
		Child       []struct {
			Type        int8   `json:"type"`
			Title       string `json:"title"`
			Href        string `json:"href"`
			Icon        string `json:"icon"`
			Target      string `json:"target"`
			OrderNumber int64  `json:"order_number"`
			Method      string `json:"method"`
			Checked     int8   `json:"checked"`
			IsMenu      int8   `json:"is_menu"`
			Child       []struct {
				Type        int8   `json:"type"`
				Title       string `json:"title"`
				Href        string `json:"href"`
				Icon        string `json:"icon"`
				Target      string `json:"target"`
				OrderNumber int64  `json:"order_number"`
				Method      string `json:"method"`
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
			Type:        m.Type,
			Title:       m.Title,
			Href:        m.Href,
			Icon:        m.Icon,
			Target:      m.Target,
			OrderNumber: m.OrderNumber,
			Method:      m.Method,
			Checked:     m.Checked,
			IsMenu:      m.IsMenu,
		}

		var menuchilds []*models.Perm
		if len(m.Child) > 0 {
			for _, mchild := range m.Child {
				menuchild := &models.Perm{
					Model:       gorm.Model{CreatedAt: time.Now()},
					Type:        mchild.Type,
					Title:       mchild.Title,
					Href:        mchild.Href,
					Icon:        mchild.Icon,
					Target:      mchild.Target,
					OrderNumber: mchild.OrderNumber,
					Method:      mchild.Method,
					Checked:     mchild.Checked,
					IsMenu:      mchild.IsMenu,
				}

				var mmenuchilds []*models.Perm
				if len(mchild.Child) > 0 {
					for _, mmchild := range mchild.Child {
						mmenuchild := &models.Perm{
							Model:       gorm.Model{CreatedAt: time.Now()},
							Type:        mmchild.Type,
							Title:       mmchild.Title,
							Href:        mmchild.Href,
							Icon:        mmchild.Icon,
							Target:      mmchild.Target,
							OrderNumber: mmchild.OrderNumber,
							Method:      mmchild.Method,
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

// CreateAdminRole 新建管理角色
func CreateAdminRole() {
	role := &models.Role{
		Name:        "超级管理员",
		DisplayName: "超级管理员",
		Rmk:         "超级管理员",
		IsAdmin:     sql.NullBool{Bool: true, Valid: true},
		Model:       gorm.Model{CreatedAt: time.Now()},
	}

	var permIds []string
	_, perms := sysinit.PermService.GetAll(map[string]interface{}{}, []string{"1", "2"}, false)
	for _, perm := range perms {
		if len(perm.Href) > 0 {
			permId := strconv.FormatUint(uint64(perm.ID), 10)
			permIds = append(permIds, permId)
		}
	}

	role.PermIds = permIds

	if err := sysinit.RoleService.Create(role); err != nil {
		panic(fmt.Sprintf("管理角色填充错误：%v", err))
	}
}

// CreateAdminUser 新建管理员
func CreateAdminUser() {
	adminrole, found := sysinit.RoleService.GetAdmin()
	if !found {
		panic("管理角色不存在")
	}
	admin := &models.User{
		Username: "username",
		Name:     "超级管理员",
		Email:    "admin@admin.com",
		Telphone: "13800138000",
		IsAdmin:  sql.NullBool{Bool: true, Valid: true},
		Model:    gorm.Model{CreatedAt: time.Now()},
		RoleIds:  []uint{adminrole.ID},
	}

	if err := sysinit.UserService.Create("password", admin); err != nil {
		panic(fmt.Sprintf("管理员填充错误：%v", err))
	}
}

// CreateTenantRoleAndUser 新建商户角色账号
func CreateTenantRoleAndUser() {
	tenantrole := &models.Role{
		Name:        "商户管理",
		DisplayName: "商户管理",
		Rmk:         "商户管理",
		IsAdmin:     sql.NullBool{Bool: false, Valid: true},
		Model:       gorm.Model{CreatedAt: time.Now()},
	}

	var permIds []string
	_, perms := sysinit.PermService.GetAll(map[string]interface{}{}, []string{"1", "3"}, false)
	for _, perm := range perms {
		if len(perm.Href) > 0 {
			permId := strconv.FormatUint(uint64(perm.ID), 10)
			permIds = append(permIds, permId)
		}
	}

	tenantrole.PermIds = permIds

	if err := sysinit.RoleService.Create(tenantrole); err != nil {
		panic(fmt.Sprintf("商户角色填充错误：%v", err))
	}

	tenantuser := &models.User{
		Username: "tenantname",
		Name:     "商户管理员",
		Email:    "tenant@admin.com",
		Telphone: "13800138001",
		IsAdmin:  sql.NullBool{Bool: false, Valid: true},
		Model:    gorm.Model{CreatedAt: time.Now()},
		RoleIds:  []uint{tenantrole.ID},
	}

	if err := sysinit.UserService.Create("password", tenantuser); err != nil {
		panic(fmt.Sprintf("商户账号填充错误：%v", err))
	}
}

// CreateRoles 新建用户
func CreateRoles() {
	for i := 0; i < 50; i++ {
		role := &models.Role{
			Name:        Fake.Name(),
			DisplayName: Fake.UserName(),
			Rmk:         Fake.Paragraph(1, true),
			IsAdmin:     sql.NullBool{Bool: false, Valid: true},
			Model:       gorm.Model{CreatedAt: time.Now()},
		}

		if err := sysinit.RoleService.Create(role); err != nil {
			panic(fmt.Sprintf("角色填充错误：%v", err))
		}
	}
}

// CreateUsers 新建用户
func CreateUsers() {
	// 最新手机正则 ^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$
	emailRegexp := regexp.MustCompile(".*(@.*)")
	totalCount := 50
	for i := 0; i < totalCount; i++ {
		user := &models.User{
			Username: Fake.UserName(),
			Name:     Fake.Name(),
			Email:    emailRegexp.ReplaceAllString(Fake.Email(), strings.Replace(strings.ToLower(Fake.UserName()), " ", "_", -1)+"@example.com"),
			Telphone: lib.CreatePhoneNumber(),
			IsAdmin:  sql.NullBool{Bool: false, Valid: true},
			Model:    gorm.Model{CreatedAt: time.Now()},
			RoleIds:  []uint{},
		}

		if err := sysinit.UserService.Create("password", user); err != nil {
			panic(fmt.Sprintf("用户填充错误：%v", err))
		}
	}
}

// CreateTenants 新建用户
func CreateTenants() {
	// 最新手机正则 ^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$
	emailRegexp := regexp.MustCompile(".*(@.*)")
	totalCount := 50
	for i := 0; i < totalCount; i++ {
		newUUID := uuid.New()
		tenant := &models.Tenant{
			UId:      newUUID.ID(),
			FullName: Fake.CompanyName(),
			Name:     Fake.CompanyName(),
			Email:    emailRegexp.ReplaceAllString(Fake.Email(), strings.Replace(strings.ToLower(Fake.CompanyName()), " ", "_", -1)+"@example.com"),
			Telphone: lib.CreatePhoneNumber(),
			Model:    gorm.Model{CreatedAt: time.Now()},
			Rmk:      Fake.Paragraph(1, true),
		}

		if err := sysinit.TenantService.Create(tenant); err != nil {
			panic(fmt.Sprintf("商户填充错误：%v", err))
		}
	}
}

/*
	AutoMigrates 重置数据表
	sysinit.Db.DropTableIfExists 删除存在数据表
	sysinit.Db.AutoMigrate 重建数据表
*/
func AutoMigrates() {
	sysinit.Db.DropTableIfExists("users", "perms", "roles", "tenants", "casbin_rule")

	sysinit.Db.AutoMigrate(
		&models.User{},
		&models.Role{},
		&models.Perm{},
		&models.Tenant{},
		&gormadapter.CasbinRule{},
	)
}
