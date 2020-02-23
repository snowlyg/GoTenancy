package widget_test

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
	"GoTenancy/libs/widget"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB
var Widgets *widget.Widgets
var Admin *admin.Admin
var Server *httptest.Server

type bannerArgument struct {
	Title    string
	SubTitle string
}

func init() {
	db = utils.TestDB()
	//db.LogMode(true)
}

// Runner
func setup() {
	if err := db.DropTableIfExists(&widget.QorWidgetSetting{}).Error; err != nil {
		panic(err)
	}
	db.AutoMigrate(&widget.QorWidgetSetting{})
	mux := http.NewServeMux()
	Server = httptest.NewServer(mux)

	Widgets = widget.New(&widget.Config{
		DB: db,
	})
	Widgets.RegisterViewPath("libs/widget/test")

	Admin = admin.New(&qor.Config{DB: db})
	Admin.AddResource(Widgets)
	Admin.MountTo("/admin", mux)

	Widgets.RegisterWidget(&widget.Widget{
		Name:      "Banner",
		Templates: []string{"banner"},
		Setting:   Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			if setting != nil {
				argument := setting.(*bannerArgument)
				context.Options["Title"] = argument.Title
				context.Options["SubTitle"] = argument.SubTitle
			}
			return context
		},
	})

	Widgets.RegisterScope(&widget.Scope{
		Name: "From Google",
		Visible: func(context *widget.Context) bool {
			if request, ok := context.Get("Request"); ok {
				_, ok := request.(*http.Request).URL.Query()["from_google"]
				return ok
			}
			return false
		},
	})

	Widgets.RegisterWidget(&widget.Widget{
		Name:    "NoTemplate",
		Setting: Admin.NewResource(&bannerArgument{}),
		Context: func(context *widget.Context, setting interface{}) *widget.Context {
			context.Body = "<h1>My html</h1>"
			return context
		},
	})
}

func reset() {
	db.DropTable(&widget.QorWidgetSetting{})
	db.AutoMigrate(&widget.QorWidgetSetting{})
	setup()
}

// Test DB's record after call Render
func TestRenderRecord(t *testing.T) {
	reset()
	var count int
	db.Model(&widget.QorWidgetSetting{}).Where(widget.QorWidgetSetting{Name: "HomeBanner", WidgetType: "Banner", Scope: "default", GroupName: "Banner"}).Count(&count)
	if count != 0 {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render Record TestCase: should don't exist widget setting")))
	}

	widgetContext := Widgets.NewContext(&widget.Context{})
	widgetContext.Render("HomeBanner", "Banner")
	db.Model(&widget.QorWidgetSetting{}).Where(widget.QorWidgetSetting{Name: "HomeBanner", WidgetType: "Banner", Scope: "default", GroupName: "Banner"}).Count(&count)
	if count == 0 {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render Record TestCase: should have default widget setting")))
	}

	http.PostForm(Server.URL+"/admin/widgets/HomeBanner",
		url.Values{"_method": {"PUT"},
			"QorResource.Scope":       {"from_google"},
			"QorResource.ActivatedAt": {"2016-07-14 10:10:42.433372925 +0800 CST"},
			"QorResource.Widgets":     {"Banner"},
			"QorResource.Template":    {"banner"},
			"QorResource.Kind":        {"Banner"},
		})
	db.Model(&widget.QorWidgetSetting{}).Where(widget.QorWidgetSetting{Name: "HomeBanner", WidgetType: "Banner", Scope: "from_google"}).Count(&count)
	if count == 0 {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render Record TestCase: should have from_google widget setting")))
	}
}

// Runner
func TestRenderContext(t *testing.T) {
	reset()
	setting := &widget.QorWidgetSetting{}
	db.Where(widget.QorWidgetSetting{Name: "HomeBanner", WidgetType: "Banner", Scope: "default"}).FirstOrInit(setting)
	db.Create(setting)

	html := Widgets.Render("HomeBanner", "Banner")
	if !strings.Contains(string(html), "Hello, \n<h1></h1>\n<h2></h2>\n") {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render TestCase #%d: Failure Result:\n %s\n", 1, html)))
	}

	widgetContext := Widgets.NewContext(&widget.Context{
		Options: map[string]interface{}{"CurrentUser": "Qortex"},
	})
	html = widgetContext.Render("HomeBanner", "Banner")
	if !strings.Contains(string(html), "Hello, Qortex\n<h1></h1>\n<h2></h2>\n") {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render TestCase #%d: Failure Result:\n %s\n", 2, html)))
	}

	db.Where(widget.QorWidgetSetting{Name: "HomeBanner", WidgetType: "Banner"}).FirstOrInit(setting)
	setting.SetSerializableArgumentValue(&bannerArgument{Title: "Title", SubTitle: "SubTitle"})
	err := db.Model(setting).Update(setting).Error
	if err != nil {
		panic(err)
	}

	html = widgetContext.Render("HomeBanner", "Banner")
	if !strings.Contains(string(html), "Hello, Qortex\n<h1>Title</h1>\n<h2>SubTitle</h2>\n") {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render TestCase #%d: Failure Result:\n %s\n", 3, html)))
	}

}

func TestRenderNoTemplate(t *testing.T) {
	reset()

	html := Widgets.Render("abc", "NoTemplate")
	if !strings.Contains(string(html), "<h1>My html</h1>") {
		t.Errorf(color.RedString(fmt.Sprintf("\nWidget Render TestCase #%d: Failure Result:\n %s\n", 5, html)))
	}

}

func TestRegisterFuncMap(t *testing.T) {
	func1 := func() {}
	Widgets.RegisterFuncMap("func1", func1)
	context := Widgets.NewContext(nil)
	if _, ok := context.FuncMaps["func1"]; !ok {
		t.Errorf("func1 should be assigned to context")
	}

	context2 := context.Funcs(template.FuncMap{"func2": func() {}})
	if _, ok := context.FuncMaps["func2"]; !ok {
		t.Errorf("func2 should be assigned to context")
	}
	if _, ok := context2.FuncMaps["func2"]; !ok {
		t.Errorf("func2 should be assigned to context")
	}

	context3 := Widgets.NewContext(nil)
	if _, ok := context3.FuncMaps["func3"]; ok {
		t.Errorf("func3 should not be assigned to other contexts")
	}
}
