package memory

import (
	"reflect"
	"testing"

	"GoTenancy/libs/cache"
)

var client cache.CacheStoreInterface

func init() {
	client = New()
}

func TestPlainText(t *testing.T) {
	if err := client.Set("hello_world", "Hello World"); err != nil {
		t.Errorf("No error should happen when saving plain text into client")
	}

	if value, err := client.Get("hello_world"); err != nil || value != "Hello World" {
		t.Errorf("found value: %v", value)
	}

	if err := client.Set("hello_world", "Hello World2"); err != nil {
		t.Errorf("No error should happen when updating saved value")
	}

	if value, err := client.Get("hello_world"); err != nil || value != "Hello World2" {
		t.Errorf("value should been updated: %v", value)
	}

	if err := client.Delete("hello_world"); err != nil {
		t.Errorf("failed to delete value: %v", err)
	}

	if _, err := client.Get("hello_world"); err == nil {
		t.Errorf("the key should been deleted")
	}
}

func TestUnmarshal(t *testing.T) {
	type result struct {
		Name  string
		Value string
	}

	r1 := result{Name: "result_name_1", Value: "result_value_1"}
	if err := client.Set("unmarshal", r1); err != nil {
		t.Errorf("No error should happen when saving struct into client")
	}

	var r2 result
	if err := client.Unmarshal("unmarshal", &r2); err != nil || !reflect.DeepEqual(r1, r2) {
		t.Errorf("found value: %#v", r2)
	}

	if err := client.Delete("unmarshal"); err != nil {
		t.Errorf("failed to delete value: %v", err)
	}

	if err := client.Unmarshal("unmarshal", &r2); err == nil {
		t.Errorf("the key should been deleted")
	}
}

func TestFetch(t *testing.T) {
	var result int
	var fc = func() interface{} {
		result++
		return result
	}

	if value, err := client.Fetch("fetch", fc); err != nil || value != "1" {
		t.Errorf("Should get result from func if key not found")
	}

	if value, err := client.Fetch("fetch", fc); err != nil || value != "1" {
		t.Errorf("Should lookup result from cache store if key is existing")
	}
}
