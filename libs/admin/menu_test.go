package admin

import (
	"testing"

	"GoTenancy/libs/qor"
)

func generateResourceMenu(resource *Resource) *Menu {
	return &Menu{RelativePath: resource.ToParam(), Name: resource.Name}
}

func TestMenu(t *testing.T) {
	admin := New(&qor.Config{})
	admin.router.Prefix = "/admin"

	menu := &Menu{Name: "Dashboard", Link: "/link1"}
	admin.AddMenu(menu)

	if menu.URL() != "/link1" {
		t.Errorf("menu's URL should be correct")
	}

	if admin.GetMenu("Dashboard") == nil {
		t.Errorf("menu %v not added", "Dashboard")
	}

	menu2 := &Menu{Name: "Dashboard", RelativePath: "/link2"}
	admin.AddMenu(menu2)
	if menu2.URL() != "/admin/link2" {
		t.Errorf("menu's URL should be correct")
	}

	type Res struct{}
	admin.AddResource(&Res{})

	if menu := admin.GetMenu("Res"); menu == nil {
		t.Errorf("menu %v not added", "Res")
	} else if menu.URL() != "/admin/res" {
		t.Errorf("menu %v' URL should be correct, got %v", "Res", menu.URL())
	}

	admin.AddResource(&Res{}, &Config{Name: "Res2", Menu: []string{"management"}})

	if menu := admin.GetMenu("Res2"); menu == nil {
		t.Errorf("menu %v not added", "Res2")
	} else if menu.URL() != "/admin/res2" {
		t.Errorf("menu %v' URL should be correct, got %v", "Res2", menu.URL())
	} else if len(menu.Ancestors) != 1 || menu.Ancestors[0] != "management" {
		t.Errorf("menu %v' ancestors should be correct", "Res2")
	}

	if menu := admin.GetMenu("management", "Res2"); menu == nil {
		t.Errorf("menu %v not added", "Res2")
	} else if menu.URL() != "/admin/res2" {
		t.Errorf("menu %v' URL should be correct, got %v", "Res2", menu.URL())
	} else if len(menu.Ancestors) != 1 || menu.Ancestors[0] != "management" {
		t.Errorf("menu %v' ancestors should be correct", "Res2")
	}

	if menu := admin.GetMenu("management", "Res"); menu != nil {
		t.Errorf("menu management>Res should not found")
	}
}

func TestMenuPriority(t *testing.T) {
	admin := New(&qor.Config{})
	admin.router.Prefix = "/admin"

	admin.AddMenu(&Menu{Name: "Name1", Priority: 2})
	admin.AddMenu(&Menu{Name: "Name2", Priority: -1})
	admin.AddMenu(&Menu{Name: "Name3", Priority: 3})
	admin.AddMenu(&Menu{Name: "Name4", Priority: 4})
	admin.AddMenu(&Menu{Name: "Name5", Priority: 1})
	admin.AddMenu(&Menu{Name: "Name6", Priority: 0})
	admin.AddMenu(&Menu{Name: "Name7", Priority: -2})
	admin.AddMenu(&Menu{Name: "SubName1", Ancestors: []string{"Name5"}, Priority: 1})
	admin.AddMenu(&Menu{Name: "SubName2", Ancestors: []string{"Name5"}, Priority: 3})
	admin.AddMenu(&Menu{Name: "SubName3", Ancestors: []string{"Name5"}, Priority: -1})
	admin.AddMenu(&Menu{Name: "SubName4", Ancestors: []string{"Name5"}, Priority: 4})
	admin.AddMenu(&Menu{Name: "SubName5", Ancestors: []string{"Name5"}, Priority: -1})
	admin.AddMenu(&Menu{Name: "SubName1", Ancestors: []string{"Name1"}})
	admin.AddMenu(&Menu{Name: "SubName2", Ancestors: []string{"Name1"}, Priority: 2})
	admin.AddMenu(&Menu{Name: "SubName3", Ancestors: []string{"Name1"}, Priority: -2})
	admin.AddMenu(&Menu{Name: "SubName4", Ancestors: []string{"Name1"}, Priority: 1})
	admin.AddMenu(&Menu{Name: "SubName5", Ancestors: []string{"Name1"}, Priority: -1})

	menuNames := []string{"Name5", "Name1", "Name3", "Name4", "Name6", "Name7", "Name2"}
	submenuNames := []string{"SubName1", "SubName2", "SubName4", "SubName3", "SubName5"}
	submenuNames2 := []string{"SubName4", "SubName2", "SubName1", "SubName3", "SubName5"}
	for idx, menu := range admin.GetMenus() {
		if menuNames[idx] != menu.Name {
			t.Errorf("#%v menu should be %v, but got %v", idx, menuNames[idx], menu.Name)
		}

		if menu.Name == "Name5" {
			subMenus := menu.GetSubMenus()
			if len(subMenus) != 5 {
				t.Errorf("Should have 5 subMenus for Name5")
			}

			for idx, menu := range subMenus {
				if submenuNames[idx] != menu.Name {
					t.Errorf("#%v menu should be %v, but got %v", idx, submenuNames[idx], menu.Name)
				}
			}
		}

		if menu.Name == "Name1" {
			subMenus := menu.GetSubMenus()
			if len(subMenus) != 5 {
				t.Errorf("Should have 5 subMenus for Name1")
			}

			for idx, menu := range subMenus {
				if submenuNames2[idx] != menu.Name {
					t.Errorf("#%v menu should be %v, but got %v", idx, submenuNames2[idx], menu.Name)
				}
			}
		}
	}
}
