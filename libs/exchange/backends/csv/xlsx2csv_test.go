package csv

import (
	"io/ioutil"
	"testing"
)

func TestGenerateCSVFromXLSXFile(t *testing.T) {
	reader, err := generateCSVFromXLSXFile("testdata/oos_sample.xlsx")
	if err != nil {
		t.Fatal(err)
	}

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	if csvData != string(data) {
		t.Fatalf("want %q but get %q", csvData, string(data))
	}
}

func TestIsXLSXFile(t *testing.T) {
	if isXLSXFile("testdata/oos_sample.xlsx") == false {
		t.Fatalf("want true but get false")
	}
	if isXLSXFile("testdata/oos_sample.csv") == true {
		t.Fatalf("want false but get true")
	}
}

var csvData = `Gender,Category 1,Category 2,Category 3,Number of SKU on EC,Inventory,Number of SKU with inventory at 0,% out of stock (= Number of SKU with inventory at 0/Number of SKU on EC)
UNISEX,アクセサリー,香水,,4,8,0,0
UNISEX,Total,,,4,6,1,0.25
MALE,ウェア,シャツ,short,56,34,35,0.625
`
