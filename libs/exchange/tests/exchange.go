package tests

import (
	"GoTenancy/libs/exchange"
	"GoTenancy/libs/qor/test/utils"
)

var (
	// DB db used for testing
	DB = utils.TestDB()
	// ProductExchangeResource import product exchange definition
	ProductExchangeResource = exchange.NewResource(&Product{}, exchange.Config{PrimaryField: "Code"})
)

func init() {
	DB.DropTable(&Product{}, &Category{})
	DB.AutoMigrate(&Product{}, &Category{})

	ProductExchangeResource.Meta(&exchange.Meta{Name: "Code", Header: "代码"})
	ProductExchangeResource.Meta(&exchange.Meta{Name: "Name"})
	ProductExchangeResource.Meta(&exchange.Meta{Name: "Price"})
	ProductExchangeResource.Meta(&exchange.Meta{Name: "Tag"})
	ProductExchangeResource.Meta(&exchange.Meta{Name: "Category.Name", Header: "Category"})
}
