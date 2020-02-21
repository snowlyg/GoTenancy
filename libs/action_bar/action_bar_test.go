package action_bar

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
	"github.com/fatih/color"

	"github.com/headzoo/surf"
)

var Server *httptest.Server
var CurrentUser *User
var actionBar *ActionBar

type User struct {
	Name string
	Role string
}

func (user *User) DisplayName() string {
	return user.Name
}

type AdminAuth struct {
}

func (AdminAuth) LoginURL(c *admin.Context) string {
	return "/auth/login"
}

func (AdminAuth) LogoutURL(c *admin.Context) string {
	return "/auth/logout"
}

func (AdminAuth) GetCurrentUser(c *admin.Context) qor.CurrentUser {
	if CurrentUser.Role != "admin" {
		return nil
	}
	return CurrentUser
}

// Init
func init() {
	Root, _ := os.Getwd()
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		tmpl, err := template.New("home.tmpl").Funcs(actionBar.FuncMap(w, req)).ParseFiles(Root + "/test/views/home.tmpl")
		if err != nil {
			fmt.Printf("Execute template have error %v\n", err)
		}

		var actions = []ActionInterface{}
		actions = append(actions, HTMLAction{HTML: template.HTML("<a href='/edit_widget'>Edit Widget</a>")})
		actions = append(actions, EditResourceAction{Inline: true, Value: CurrentUser, EditModeOnly: true})
		actions = append(actions, Action{Inline: true, Name: "Edit SEO", Link: "/admin/seo/1/edit", EditModeOnly: true})

		tmpl.Execute(w, struct {
			ActionBar   template.HTML
			CurrentUser *User
		}{
			ActionBar:   actionBar.Actions(actions...).Render(w, req),
			CurrentUser: CurrentUser,
		})
	})
	db := utils.TestDB()
	Server = httptest.NewServer(mux)
	Admin := admin.New(&qor.Config{DB: db})
	Admin.SetAuth(AdminAuth{})
	Admin.MountTo("/admin", mux)
	Admin.AddResource(User{})

	actionBar = New(Admin)
}

func TestAuth(t *testing.T) {
	bow := surf.NewBrowser()
	CurrentUser = &User{Name: "Kaea", Role: ""}
	bow.Open(Server.URL + "/")
	if !strings.Contains(bow.Body(), "LOGIN") {
		t.Errorf(color.RedString("Should Get `LOGIN` link if current user is normal user"))
	}

	CurrentUser = &User{Name: "Admin", Role: "admin"}
	bow.Open(Server.URL + "/")
	if !strings.Contains(bow.Body(), "LOGOUT") {
		t.Errorf(color.RedString("Should Get `LOGOUT` link if current user is admin user"))
	}
}

func TestEditMode(t *testing.T) {
	bow := surf.NewBrowser()

	// Default doesn't checked checkbox
	bow.Open(Server.URL + "/")
	html, _ := bow.Find(".qor-actionbar__module").Html()
	if strings.Contains(html, "checked") {
		t.Errorf(color.RedString("Should don't checked checkbox"))
	}

	// Should checked checkbox after switch mode
	bow.Open(Server.URL + "/admin/action_bar/switch_mode?checked=true")
	html, _ = bow.Find(".qor-actionbar__module").Html()
	if !strings.Contains(html, "checked") {
		t.Errorf(color.RedString("Should checked checkbox"))
	}
}

func TestInilneMode(t *testing.T) {
	bow := surf.NewBrowser()

	// Default should not have inline edit button
	bow.Open(Server.URL + "/")
	if bow.Find(".qor-actionbar-button").Length() != 0 {
		t.Errorf(color.RedString("Should don't have inline edit button"))
	}

	// Should have inline edit button if in edit mode
	bow.Open(Server.URL + "/admin/action_bar/switch_mode?checked=true")
	if bow.Find(".qor-actionbar-button").Length() == 0 {
		t.Errorf(color.RedString("Should have inline edit button"))
	}
}

func TestRegisterAction(t *testing.T) {
	bow := surf.NewBrowser()

	// Default should not have inline edit button
	bow.Open(Server.URL + "/")
	if bow.Find(".qor-actionbar__menu").Length() != 0 {
		t.Errorf(color.RedString("Should don't have additional actions"))
	}

	// Default should not have inline edit button
	actionBar.RegisterAction(&Action{Name: "Admin Dashboard", Link: "/admin"})
	bow.Open(Server.URL + "/")
	if bow.Find(".qor-actionbar__menu").Length() == 0 {
		t.Errorf(color.RedString("Should have additional actions"))
	}
}

func TestInlineActions(t *testing.T) {
	bow := surf.NewBrowser()

	// Default should not have inline edit button
	bow.Open(Server.URL + "/")
	if bow.Find(".qor-inline-actions .qor-action a").Text() != "Edit Widget" {
		t.Errorf(color.RedString("Should only display 'Edit Widget' if haven't turn on inline-edit"))
	}

	bow.Open(Server.URL + "/admin/action_bar/switch_mode?checked=true")
	if bow.Find(".qor-inline-actions .qor-action a").Text() != "Edit WidgetEdit UserEdit SEO" {
		t.Errorf(color.RedString("Should display all actions if turn on inline-edit"))
	}
}
