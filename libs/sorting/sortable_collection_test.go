package sorting_test

import (
	"fmt"
	"reflect"
	"testing"

	"GoTenancy/libs/sorting"
	"github.com/jinzhu/gorm"
)

type ColorVariation struct {
	gorm.Model
	Code string
}

func checkOrder(results interface{}, order []string) error {
	values := reflect.Indirect(reflect.ValueOf(results))
	for idx, o := range order {
		value := values.Index(idx)
		primaryValue := fmt.Sprint(reflect.Indirect(value).FieldByName("ID").Interface())
		if primaryValue != o {
			return fmt.Errorf("#%v of values's primary key is %v, but should be %v", idx+1, primaryValue, o)
		}
	}
	return nil
}

func TestSortSlice(t *testing.T) {
	colorVariations := []ColorVariation{
		{Model: gorm.Model{ID: 1}, Code: "1"},
		{Model: gorm.Model{ID: 2}, Code: "2"},
		{Model: gorm.Model{ID: 3}, Code: "3"},
	}

	collectionSorting := sorting.SortableCollection{PrimaryKeys: []string{"3", "1", "2"}}
	collectionSorting.Sort(colorVariations)

	if err := checkOrder(colorVariations, []string{"3", "1", "2"}); err != nil {
		t.Error(err)
	}
}

func TestSort(t *testing.T) {
	colorVariations := &[]ColorVariation{
		{Model: gorm.Model{ID: 1}, Code: "1"},
		{Model: gorm.Model{ID: 2}, Code: "2"},
		{Model: gorm.Model{ID: 3}, Code: "3"},
	}

	collectionSorting := sorting.SortableCollection{PrimaryKeys: []string{"3", "1", "2"}}
	collectionSorting.Sort(colorVariations)

	if err := checkOrder(colorVariations, []string{"3", "1", "2"}); err != nil {
		t.Error(err)
	}
}

func TestSortPointer(t *testing.T) {
	colorVariations := &[]*ColorVariation{
		{Model: gorm.Model{ID: 1}, Code: "1"},
		{Model: gorm.Model{ID: 2}, Code: "2"},
		{Model: gorm.Model{ID: 3}, Code: "3"},
	}

	collectionSorting := sorting.SortableCollection{PrimaryKeys: []string{"3", "1", "2"}}
	collectionSorting.Sort(colorVariations)

	if err := checkOrder(colorVariations, []string{"3", "1", "2"}); err != nil {
		t.Error(err)
	}
}

func TestSortWithSomePrimaryKeys(t *testing.T) {
	colorVariations := &[]ColorVariation{
		{Model: gorm.Model{ID: 1}, Code: "1"},
		{Model: gorm.Model{ID: 2}, Code: "2"},
		{Model: gorm.Model{ID: 3}, Code: "3"},
		{Model: gorm.Model{ID: 4}, Code: "4"},
	}

	collectionSorting := sorting.SortableCollection{PrimaryKeys: []string{"3", "1"}}
	collectionSorting.Sort(colorVariations)

	if err := checkOrder(colorVariations, []string{"3", "1", "2", "4"}); err != nil {
		t.Error(err)
	}
}

func TestSortPointerWithSomePrimaryKeys(t *testing.T) {
	colorVariations := &[]*ColorVariation{
		{Model: gorm.Model{ID: 1}, Code: "1"},
		{Model: gorm.Model{ID: 2}, Code: "2"},
		{Model: gorm.Model{ID: 3}, Code: "3"},
		{Model: gorm.Model{ID: 4}, Code: "4"},
	}

	collectionSorting := sorting.SortableCollection{PrimaryKeys: []string{"3", "1"}}
	collectionSorting.Sort(colorVariations)

	if err := checkOrder(colorVariations, []string{"3", "1", "2", "4"}); err != nil {
		t.Error(err)
	}
}
