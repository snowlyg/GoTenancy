package main

import (
	"fmt"

	"github.com/snowlyg/go-tenancy/models"
	"github.com/snowlyg/go-tenancy/sysinit"
)

func main() {

	AutoMigrates()

	fmt.Println(fmt.Sprintf("管理员填充开始"))
	createAdminUsers()
	fmt.Println(fmt.Sprintf("管理员填充完成"))
}

/*
	AutoMigrates 重置数据表

	sysinit.Db.DropTableIfExists 删除存在数据表
	sysinit.Db.AutoMigrate 重建数据表
*/

func AutoMigrates() {
	tableNames := []string{"users"}
	sysinit.Db.DropTableIfExists(tableNames)

	sysinit.Db.AutoMigrate(
		&models.User{},
	)
}
