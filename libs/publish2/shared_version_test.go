package publish2_test

import (
	"testing"

	"GoTenancy/libs/publish2"
	"github.com/jinzhu/gorm"
)

type SharedVersionProduct struct {
	gorm.Model
	Name            string
	ColorVariations []SharedVersionColorVariation
	publish2.Version
}

type SharedVersionColorVariation struct {
	gorm.Model
	Name                   string
	SharedVersionProductID uint
	SizeVariations         []SharedVersionSizeVariation
	publish2.SharedVersion
}

type SharedVersionSizeVariation struct {
	gorm.Model
	Name                          string
	SharedVersionColorVariationID uint
	publish2.SharedVersion
}

func prepareSharedVersionProduct() *SharedVersionProduct {
	product := SharedVersionProduct{
		Name: "shared product 1",
		ColorVariations: []SharedVersionColorVariation{
			{
				Name: "cv1",
			},
			{
				Name: "cv2",
			},
		},
	}
	DB.Create(&product)

	product.SetVersionName("v1")
	product.ColorVariations[0].SetSharedVersionName("v1")
	DB.Save(&product)

	product.SetVersionName("v2")
	product.ColorVariations[0].SetSharedVersionName("")
	colorVariation := SharedVersionColorVariation{
		Name: "cv3",
	}
	colorVariation.SetSharedVersionName("v2")
	product.ColorVariations = append(product.ColorVariations, colorVariation)
	DB.Save(&product)

	return &product
}

func TestSharedVersions(t *testing.T) {
	product1 := prepareSharedVersionProduct()
	product2 := prepareSharedVersionProduct()

	var product1V1 SharedVersionProduct
	DB.Set(publish2.VersionNameMode, "v1").Preload("ColorVariations").Find(&product1V1, "id = ?", product1.ID)

	if len(product1V1.ColorVariations) != 2 {
		t.Errorf("Preload: Should have 2 color variations for product v1, but got %v", len(product1V1.ColorVariations))
	}

	var colorVariations1V1 []SharedVersionColorVariation
	DB.Model(&product1V1).Related(&colorVariations1V1)
	if len(colorVariations1V1) != 2 {
		t.Errorf("Related: Should have 2 color variations for product v1, but got %v", len(colorVariations1V1))
	}

	var product1V2 SharedVersionProduct
	DB.Set(publish2.VersionNameMode, "v2").Preload("ColorVariations").Find(&product1V2, "id = ?", product1.ID)

	if len(product1V2.ColorVariations) != 3 {
		t.Errorf("Preload: Should have 3 color variations for product v2, but got %v", len(product1V2.ColorVariations))
	}

	var products []SharedVersionProduct
	DB.Preload("ColorVariations").Find(&products)

	var product2V2 SharedVersionProduct
	for _, p := range products {
		if p.ID == product2.ID && p.VersionName == "v2" {
			product2V2 = p
		}
	}

	if len(product2V2.ColorVariations) != 3 {
		t.Errorf("Preload: Should have 3 color variations for product v2, but got %v", len(product2V2.ColorVariations))
	}
}
