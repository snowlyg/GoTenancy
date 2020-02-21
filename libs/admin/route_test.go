package admin

import (
	"strings"
	"testing"
)

func TestSortRouter(t *testing.T) {
	router := newRouter()
	router.Get("/api/orders/:order_id", nil)
	router.Get("/api/orders/new", nil)
	router.Get("/api/orders/:order_id/order_items", nil)
	router.Get("/api/orders/:order_id/order_items/:order_item_id", nil)
	router.Get("/api/orders/:order_id/order_items/new", nil)
	router.Get("/api/orders/:order_id/order_items/:order_item_id/edit", nil)

	paths := []string{}
	for _, r := range router.routers["GET"] {
		paths = append(paths, r.Path)
	}

	if strings.Join(paths, ",") != strings.Join([]string{"/api/orders/new", "/api/orders/:order_id/order_items/new", "/api/orders/:order_id/order_items", "/api/orders/:order_id", "/api/orders/:order_id/order_items/:order_item_id/edit", "/api/orders/:order_id/order_items/:order_item_id"}, ",") {
		t.Errorf("Sorted path is not equal, got %v", strings.Join(paths, ", "))
	}
}
