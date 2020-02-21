package excel_test

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"testing"

	"GoTenancy/libs/exchange"
	"GoTenancy/libs/exchange/backends/excel"
	"GoTenancy/libs/exchange/tests"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/resource"
	"GoTenancy/libs/qor/test/utils"
	"github.com/360EntSecGroup-Skylar/excelize"
)

var db = utils.TestDB()
var product *exchange.Resource

func init() {
	db.DropTable(&tests.Product{}, &tests.Category{})
	db.AutoMigrate(&tests.Product{}, &tests.Category{})

	product = exchange.NewResource(&tests.Product{}, exchange.Config{PrimaryField: "Code"})
	product.Meta(&exchange.Meta{Name: "Code", Header: "代码"})
	product.Meta(&exchange.Meta{Name: "Name"})
	product.Meta(&exchange.Meta{Name: "Price"})
	product.Meta(&exchange.Meta{Name: "Tag"})
	product.Meta(&exchange.Meta{Name: "Category.Name", Header: "Category"})
}

func newContext() *qor.Context {
	return &qor.Context{DB: db}
}

func checkProduct(t *testing.T, filename string) {
	excelFile, _ := excelize.OpenFile(filename)
	activeSheetIndex := excelFile.GetActiveSheetIndex()
	if activeSheetIndex == 0 && excelFile.SheetCount > 0 {
		activeSheetIndex = 1
	}
	params, err := excelFile.GetRows(excelFile.GetSheetName(activeSheetIndex))
	if err != nil {
		t.Errorf("Failed to call the GetRows method, get error %v", err)
	}

	if len(params) == 0 {
		t.Errorf("No products found in the templates")
	}

	for index, param := range params {
		if index == 0 {
			continue
		}
		var count int
		if db.Model(&tests.Product{}).Where("code = ?", param[0]).Count(&count); count != 1 {
			t.Errorf("Found %v with code %v, but should find one (%v)", count, param[0], filename)
			break
		}

		if db.Model(&tests.Product{}).Where("code = ? AND name = ? AND price = ?", param[0], param[1], param[2]).Count(&count); count != 1 {
			t.Errorf("Found %v with params %v, but should find one (%v)", count, param, filename)
			break
		}

		var product tests.Product
		if len(param) == 4 {
			db.Preload("Category").Where("code = ? AND name = ? AND price = ?", param[0], param[1], param[2]).First(&product)
			if product.Category.Name != param[3] {
				t.Errorf("Category %v should not imported, but product's category is %#v (%v)", param[3], product.Category, filename)
			}
		}
	}
}

func TestImportExcel(t *testing.T) {
	if err := product.Import(excel.New("fixtures/products.xlsx"), newContext()); err != nil {
		t.Fatalf("Failed to import excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products.xlsx")

	if err := product.Import(excel.New("fixtures/products_update.xlsx"), newContext()); err != nil {
		t.Fatalf("Failed to import excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products_update.xlsx")
}

func TestImportExcelFromReader(t *testing.T) {
	reader, err := os.Open("fixtures/products.xlsx")
	if err != nil {
		t.Errorf("no error should happen when open products.xlsx")
	}

	if err := product.Import(excel.New(reader), newContext()); err != nil {
		t.Fatalf("Failed to import excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products.xlsx")

	updateReader, err := os.Open("fixtures/products_update.xlsx")
	if err := product.Import(excel.New(updateReader), newContext()); err != nil {
		t.Fatalf("Failed to import excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products_update.xlsx")
}

func TestExportExcel(t *testing.T) {
	product.Import(excel.New("fixtures/products.xlsx"), newContext())

	if err := product.Export(excel.New("fixtures/products2.xlsx"), newContext()); err != nil {
		t.Fatalf("Failed to export excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products2.xlsx")
}

func TestExportExcelToWriter(t *testing.T) {
	writerCloser, err := os.OpenFile("fixtures/products_out.xlsx", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Errorf("Failed to open products out")
	}

	product.Import(excel.New("fixtures/products.xlsx"), newContext())

	if err := product.Export(excel.New(writerCloser), newContext()); err != nil {
		t.Fatalf("Failed to export excel, get error %v", err)
	}

	checkProduct(t, "fixtures/products2.xlsx")
}

func TestImportWithInvalidData(t *testing.T) {
	product = exchange.NewResource(&tests.Product{}, exchange.Config{PrimaryField: "Code"})
	product.Meta(&exchange.Meta{Name: "Code", Header: "代码"})
	product.Meta(&exchange.Meta{Name: "Name"})
	product.Meta(&exchange.Meta{Name: "Price"})

	product.AddValidator(&resource.Validator{
		Handler: func(result interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			if f, err := strconv.ParseFloat(fmt.Sprint(metaValues.Get("Price").Value), 64); err == nil {
				if f == 0 {
					return errors.New("product's price can't be env")
				}
				return nil
			} else {
				return err
			}
		},
	})

	if err := product.Import(excel.New("fixtures/products.xlsx"), newContext()); err != nil {
		t.Errorf("Failed to import product, get error: %v", err)
	}

	if err := product.Import(excel.New("fixtures/invalid_price_products.xlsx"), newContext()); err == nil {
		t.Error("should get error when import products with invalid price")
	}
}

func TestProcessImportedData(t *testing.T) {
	product = exchange.NewResource(&tests.Product{}, exchange.Config{PrimaryField: "Code"})
	product.Meta(&exchange.Meta{Name: "Code", Header: "代码"})
	product.Meta(&exchange.Meta{Name: "Name"})
	product.Meta(&exchange.Meta{Name: "Price"})

	product.AddProcessor(&resource.Processor{
		Handler: func(result interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
			product := result.(*tests.Product)
			product.Price = float64(int(product.Price * 1.1)) // Add 10% Tax
			return nil
		},
	})

	if err := product.Import(excel.New("fixtures/products.xlsx"), newContext()); err != nil {
		t.Errorf("Failed to import product, get error: %v", err)
	}

	checkProduct(t, "fixtures/products_with_tax.xlsx")
}
