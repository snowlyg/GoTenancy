package activity

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/l10n"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
	"GoTenancy/libs/roles"
	"github.com/jinzhu/gorm"
	"github.com/theplant/testingutils"
)

var db *gorm.DB
var Admin *admin.Admin

var (
	mux    = http.NewServeMux()
	Server = httptest.NewServer(mux)
)

type Order struct {
	gorm.Model
	Code string
}

type Product struct {
	gorm.Model
	Code string
	l10n.Locale
}

func init() {
	db = utils.TestDB()
	Admin = admin.New(&qor.Config{DB: db})
	orderRes := Admin.AddResource(&Order{})
	orderRes.Permission = roles.Allow(roles.CRUD, roles.Anyone)
	productRes := Admin.AddResource(&Product{})
	productRes.Permission = roles.Allow(roles.CRUD, roles.Anyone)
	Register(orderRes)
	Register(productRes)
	Admin.MountTo("/admin", mux)
	l10n.Global = "Global"
	l10n.RegisterCallbacks(db)
}

func setup() {
	if err := db.DropTableIfExists(&QorActivity{}, &Order{}, &Product{}).Error; err != nil {
		panic(err)
	}
	db.AutoMigrate(&QorActivity{}, &Order{}, &Product{})
	db.Create(&Order{Code: "1000001"})
	db.Create(&Product{Code: "1000001", Locale: l10n.Locale{LanguageCode: "Global"}})
}

func TestActivityWithNotLocalization(t *testing.T) {
	setup()
	http.PostForm(Server.URL+"/admin/orders/1/!qor_activities", url.Values{
		"QorResource.Action":  {"comment on"},
		"QorResource.Content": {"Activity Title"},
		"QorResource.Note":    {"Activity Note"},
	})

	req, _ := http.Get(Server.URL + "/admin/orders/1/!qor_activities.json")
	assertActivityEqual(t, req.Body, QorActivity{
		Action:       "comment on",
		Content:      "Activity Title",
		Note:         "Activity Note",
		ResourceID:   "1",
		ResourceType: "orders",
	})

	http.PostForm(Server.URL+"/admin/orders/1/!qor_activities/1/edit", url.Values{
		"QorResource.ID":   {"1"},
		"QorResource.Note": {"Activity Note Changed"},
	})

	req, _ = http.Get(Server.URL + "/admin/orders/1/!qor_activities.json")
	assertActivityEqual(t, req.Body, QorActivity{
		Action:       "comment on",
		Content:      "Activity Title",
		Note:         "Activity Note Changed",
		ResourceID:   "1",
		ResourceType: "orders",
	})
}

func TestActivityWithLocalization(t *testing.T) {
	setup()
	http.PostForm(Server.URL+"/admin/products/1/!qor_activities", url.Values{
		"QorResource.Action":  {"comment on"},
		"QorResource.Content": {"Activity Title"},
		"QorResource.Note":    {"Activity Note"},
	})

	req, _ := http.Get(Server.URL + "/admin/products/1/!qor_activities.json")
	assertActivityEqual(t, req.Body, QorActivity{
		Action:       "comment on",
		Content:      "Activity Title",
		Note:         "Activity Note",
		ResourceID:   "1::Global",
		ResourceType: "products",
	})

	http.PostForm(Server.URL+"/admin/products/1/!qor_activities/1/edit", url.Values{
		"QorResource.ID":   {"1"},
		"QorResource.Note": {"Activity Note Changed"},
	})

	req, _ = http.Get(Server.URL + "/admin/products/1/!qor_activities.json")
	assertActivityEqual(t, req.Body, QorActivity{
		Action:       "comment on",
		Content:      "Activity Title",
		Note:         "Activity Note Changed",
		ResourceID:   "1::Global",
		ResourceType: "products",
	})
}

func assertActivityEqual(t *testing.T, results io.ReadCloser, act QorActivity) {
	records := []*QorActivity{}
	content, _ := ioutil.ReadAll(results)
	json.Unmarshal(content, &records)
	for _, record := range records {
		record.ID = 0
	}
	diff := testingutils.PrettyJsonDiff(records, []QorActivity{act})
	if len(diff) > 0 {
		t.Errorf(diff)
	}
}
