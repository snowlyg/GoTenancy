// build ignore

package main

import (
	"fmt"
)

func main() {

	AutoMigrates()

	fmt.Println(fmt.Sprintf("管理员填充开始！！"))
	CreateAdminUsers()
	fmt.Println(fmt.Sprintf("管理员填充完成！！"))

	fmt.Println(fmt.Sprintf("权限填充开始！！"))
	CreatePerms()
	fmt.Println(fmt.Sprintf("权限填充完成！！"))
}
