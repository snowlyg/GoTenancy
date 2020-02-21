package banner_editor

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/banner_editor/test/config/bindatafs"
	"GoTenancy/libs/media"
	"GoTenancy/libs/media/media_library"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
	qor_utils "GoTenancy/libs/qor/utils"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
)

var (
	mux                  = http.NewServeMux()
	Server               = httptest.NewServer(mux)
	db                   = utils.TestDB()
	Admin                = admin.New(&qor.Config{DB: db})
	assetManagerResource *admin.Resource
)

type bannerEditorArgument struct {
	gorm.Model
	Value string `gorm:"size:4294967295";`
}

func init() {
	// Migrate database
	if err := db.DropTableIfExists(&QorBannerEditorSetting{}, &bannerEditorArgument{}, &media_library.MediaLibrary{}).Error; err != nil {
		panic(err)
	}
	media.RegisterCallbacks(db)
	db.AutoMigrate(&QorBannerEditorSetting{}, &bannerEditorArgument{}, &media_library.MediaLibrary{})

	// Banner Editor
	type subHeaderSetting struct {
		Text  string
		Color string
	}
	type buttonSetting struct {
		Text  string
		Link  string
		Color string
	}
	subHeaderRes := Admin.NewResource(&subHeaderSetting{})
	subHeaderRes.Meta(&admin.Meta{Name: "Text"})
	subHeaderRes.Meta(&admin.Meta{Name: "Color"})

	buttonRes := Admin.NewResource(&buttonSetting{})
	buttonRes.Meta(&admin.Meta{Name: "Text"})
	buttonRes.Meta(&admin.Meta{Name: "Link"})
	RegisterViewPath("GoTenancy/libs/banner_editor/test/views")

	RegisterElement(&Element{
		Name:     "Sub Header",
		Template: "sub_header",
		Resource: subHeaderRes,
		Context: func(c *admin.Context, r interface{}) interface{} {
			return r.(QorBannerEditorSettingInterface).GetSerializableArgument(r.(QorBannerEditorSettingInterface))
		},
	})
	RegisterElement(&Element{
		Name:     "Button",
		Label:    "Add Button",
		Template: "button",
		Resource: buttonRes,
		Context: func(c *admin.Context, r interface{}) interface{} {
			setting := r.(QorBannerEditorSettingInterface).GetSerializableArgument(r.(QorBannerEditorSettingInterface)).(*buttonSetting)
			setting.Color = "Red"
			return setting
		},
	})

	// Add asset resource
	assetManagerResource = Admin.AddResource(&media_library.MediaLibrary{})
	assetManagerResource.IndexAttrs("Title", "File")

	bannerEditorResource := Admin.AddResource(&bannerEditorArgument{}, &admin.Config{Name: "Banner"})
	bannerEditorResource.Meta(&admin.Meta{Name: "Value", Config: &BannerEditorConfig{
		MediaLibrary: assetManagerResource,
		Platforms: []Platform{
			{
				Name:     "Laptop",
				SafeArea: Size{Width: 1000, Height: 500},
			},
			{
				Name:     "Mobile",
				SafeArea: Size{Width: 600, Height: 300},
			},
		},
	}})

	Admin.MountTo("/admin", mux)
	mux.Handle("/system/", qor_utils.FileServer(http.Dir("public")))

	// Add dummy background image
	image := media_library.MediaLibrary{}
	file, err := os.Open("test/views/images/background.jpg")
	if err != nil {
		panic(err)
	}
	image.File.Scan(file)
	db.Create(&image)

	if os.Getenv("MODE") == "server" {
		db.Create(&bannerEditorArgument{
			Value: `<span id="qor-bannereditor__i9mt1" class="qor-bannereditor__draggable" data-edit-id="1" data-position-left="202" data-position-top="152" style="position: absolute; left: 16.8896%; top: 50.6667%;"><em style="color: #ff0000;">Hello World!</em>
</span>`,
		})
		fmt.Printf("Test Server URL: %v\n", Server.URL+"/admin")
		time.Sleep(time.Second * 3000)
	}
}

func TestGetConfig(t *testing.T) {
	otherBannerEditorResource := Admin.AddResource(&bannerEditorArgument{}, &admin.Config{Name: "other_banner_editor_argument"})
	otherBannerEditorResource.Meta(&admin.Meta{Name: "Value", Config: &BannerEditorConfig{
		Elements:     []string{"Sub Header"},
		MediaLibrary: assetManagerResource,
	}})

	anotherBannerEditorResource := Admin.AddResource(&bannerEditorArgument{}, &admin.Config{Name: "another_banner_editor_argument"})
	anotherBannerEditorResource.Meta(&admin.Meta{Name: "Value", Config: &BannerEditorConfig{
		Elements:     []string{"Button"},
		MediaLibrary: assetManagerResource,
	}})

	testCases := []struct {
		ConfigureName        string
		ElementNameAndLabels [][]string
		Platforms            []string
	}{
		{ConfigureName: "banners", ElementNameAndLabels: [][]string{{"Sub Header", "Sub Header"}, {"Button", "Add Button"}}, Platforms: []string{"Laptop:1000:500", "Mobile:600:300"}},
		{ConfigureName: "other_banner_editor_arguments", ElementNameAndLabels: [][]string{{"Sub Header", "Sub Header"}}, Platforms: []string{"Laptop:0:0", "Mobile:0:0"}},
		{ConfigureName: "another_banner_editor_arguments", ElementNameAndLabels: [][]string{{"Button", "Add Button"}}, Platforms: []string{"Laptop:0:0", "Mobile:0:0"}},
	}
	for _, testCase := range testCases {
		assertConfigIncludeElements(t, testCase.ConfigureName, testCase.ElementNameAndLabels, testCase.Platforms)
	}
}

func TestControllerCRUD(t *testing.T) {
	resp, _ := http.Get(Server.URL + "/admin/qor_banner_editor_settings/new?kind=Sub%20Header")
	assetPageHaveAttributes(t, resp, "Text", "Color")

	resp, _ = http.Get(Server.URL + "/admin/qor_banner_editor_settings/new?kind=Button")
	assetPageHaveAttributes(t, resp, "Text", "Link")

	// Test create setting via HTML request
	resp, _ = http.PostForm(Server.URL+"/admin/qor_banner_editor_settings?kind=Button", url.Values{
		"QorResource.Kind":                  {"Button"},
		"QorResource.SerializableMeta.Text": {"Search by Google"},
		"QorResource.SerializableMeta.Link": {"http://www.google.com"},
	})
	body, _ := ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), "Search by Google")
	assetPageHaveText(t, string(body), "http://www.google.com")

	resp, _ = http.Get(Server.URL + "/admin/qor_banner_editor_settings/1/edit")
	body, _ = ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), "Search by Google")
	assetPageHaveText(t, string(body), "http://www.google.com")

	// Test create setting via JSON request
	resp, _ = http.PostForm(Server.URL+"/admin/qor_banner_editor_settings.json?kind=Button", url.Values{
		"QorResource.Kind":                  {"Button"},
		"QorResource.SerializableMeta.Text": {"Search by Yahoo"},
		"QorResource.SerializableMeta.Link": {"http://www.yahoo.com"},
	})
	body, _ = ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), `{"ID":2,"Template":"<a style='color:Red' href='http://www.yahoo.com'>Search by Yahoo</a>\n"`)

	// Test update setting via JSON request
	resp, _ = http.PostForm(Server.URL+"/admin/qor_banner_editor_settings/2.json?kind=Button", url.Values{
		"_method":                           {"PUT"},
		"QorResource.Kind":                  {"Button"},
		"QorResource.SerializableMeta.Text": {"Search by Bing"},
		"QorResource.SerializableMeta.Link": {"http://www.bing.com"},
	})
	body, _ = ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), `{"ID":2,"Template":"<a style='color:Red' href='http://www.bing.com'>Search by Bing</a>\n"`)

	// Test Customize AssetFS
	SetAssetFS(bindatafs.AssetFS)
	resp, _ = http.PostForm(Server.URL+"/admin/qor_banner_editor_settings.json?kind=Button", url.Values{
		"QorResource.Kind":                  {"Button"},
		"QorResource.SerializableMeta.Text": {"Search by Baidu"},
		"QorResource.SerializableMeta.Link": {"http://www.baidu.com"},
	})
	body, _ = ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), `{"ID":3,"Template":"<a style='color:Red' href='http://www.baidu.com'>Search by Baidu</a>\n"`)
}

func TestMediaLibraryURL(t *testing.T) {
	resp, _ := http.Get(Server.URL + "/admin/media_libraries")
	body, _ := ioutil.ReadAll(resp.Body)
	assetPageHaveText(t, string(body), "/system/media_libraries/1/file.jpg")
}

func TestGetContent(t *testing.T) {
	iphone := "UserAgent: Mozilla/5.0 (iPhone; CPU iPhone OS 10_2_1 like Mac OS X) AppleWebKit/602.4.6 (KHTML, like Gecko) Version/10.0 Mobile/14D27 Safari/602.1"
	ipad := "Mozilla/5.0 (iPad; CPU OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"
	mac := "UserAgent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_4) AppleWebKit/600.7.12 (KHTML, like Gecko) Version/8.0.7 Safari/600.7.12"
	type testCase struct {
		Value         string
		Detector      interface{}
		ExpectedValue string
	}
	testCases := []testCase{
		// Detect by platform string
		{Value: "Laptop Content", Detector: "Laptop", ExpectedValue: "Laptop Content"},
		{Value: `[]`, Detector: "Laptop", ExpectedValue: ""},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}]`, Detector: Laptop, ExpectedValue: "Laptop Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: Laptop, ExpectedValue: "Laptop Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: Mobile, ExpectedValue: "Mobile Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: "Unknown", ExpectedValue: "Laptop Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": ""}]`, Detector: Mobile, ExpectedValue: "Laptop Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "      "}]`, Detector: Mobile, ExpectedValue: "Laptop Content"},
		// Detect by request
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: mac, ExpectedValue: `Laptop Content`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: iphone, ExpectedValue: `Mobile Content`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "Mobile Content"}]`, Detector: ipad, ExpectedValue: `Laptop Content`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}]`, Detector: mac, ExpectedValue: `Laptop Content`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}]`, Detector: iphone, ExpectedValue: `Laptop Content`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": ""}]`, Detector: iphone, ExpectedValue: "Laptop Content"},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}, {"Name": "Mobile", "Value": "      "}]`, Detector: iphone, ExpectedValue: "Laptop Content"},
		// Detect by nil or empty string
		{Value: "Laptop Content", Detector: "", ExpectedValue: "Laptop Content"},
		{Value: "Laptop Content", Detector: nil, ExpectedValue: "Laptop Content"},
		// Unscape content
		{Value: "Laptop Content: %3Cdiv%3Ehello%20%E3%83%A1%E3%83%B3%E3%82%BA%E3%82%92%E3%83%81%E3%82%A7%E3%83%83%E3%82%AF%20%3C%2Fdiv%3E", Detector: "Laptop", ExpectedValue: "Laptop Content: <div>hello メンズをチェック </div>"},
	}
	for i, testcase := range testCases {
		detector := testcase.Detector
		if d, ok := testcase.Detector.(string); ok {
			if strings.HasPrefix(d, "UserAgent") {
				req, _ := http.NewRequest("GET", "http://localhost:30000/user-agent", nil)
				req.Header.Set("User-Agent", d)
				detector = req
			}
		}
		value := GetContent(testcase.Value, detector)
		if value != testcase.ExpectedValue {
			t.Error(color.RedString("TestGetContent #%v: expect value is %v, but got %v", i+1, testcase.ExpectedValue, value))
		} else {
			color.Green("TestGetContent #%v: Success", i+1)
		}
	}
}

func TestFormattedValue(t *testing.T) {
	type testCase struct {
		Value         string
		ExpectedValue string
	}
	testCases := []testCase{
		{Value: "", ExpectedValue: `[]`},
		{Value: "Laptop Content", ExpectedValue: `[{"Name":"Laptop","Value":"Laptop Content"}]`},
		{Value: `[{"Name": "Laptop", "Value": "Laptop Content"}]`, ExpectedValue: `[{"Name": "Laptop", "Value": "Laptop Content"}]`},
	}
	for i, testcase := range testCases {
		value := formattedValue(testcase.Value)
		if value != testcase.ExpectedValue {
			t.Error(color.RedString("TestFormattedValue #%v: expect value is %v, but got %v", i+1, testcase.ExpectedValue, value))
		} else {
			color.Green("TestFormattedValue #%v: Success", i+1)
		}
	}
}

func assetPageHaveText(t *testing.T, body string, text string) {
	if !strings.Contains(body, text) {
		t.Error(color.RedString("PageHaveText: expect page have text %v, but got %v", text, body))
	}
}

func assetPageHaveAttributes(t *testing.T, resp *http.Response, attributes ...string) {
	body, _ := ioutil.ReadAll(resp.Body)
	for _, attr := range attributes {
		if !strings.Contains(string(body), fmt.Sprintf("QorResource.SerializableMeta.%v", attr)) {
			t.Error(color.RedString("PageHaveAttrributes: expect page have attributes %v, but got %v", attr, string(body)))
		}
	}
}

func assertConfigIncludeElements(t *testing.T, resourceName string, elements [][]string, sizes []string) {
	resp, _ := http.Get(fmt.Sprintf("%v/admin/%v/new", Server.URL, resourceName))
	body, _ := ioutil.ReadAll(resp.Body)
	elementDatas := []string{}
	for _, elm := range elements {
		urlParam := strings.Replace(elm[0], " ", "&#43;", -1)
		data := fmt.Sprintf("{\"Label\":\"%v\",\"CreateURL\":\"/admin/qor_banner_editor_settings/new?kind=%v\",\"Icon\":\"\"}", elm[1], urlParam)
		elementDatas = append(elementDatas, data)
	}
	sizeDatas := []string{}
	for _, size := range sizes {
		datas := strings.Split(size, ":")
		data := fmt.Sprintf("{\"Name\":\"%v\",\"Width\":%v,\"Height\":%v}", datas[0], datas[1], datas[2])
		sizeDatas = append(sizeDatas, data)
	}
	elementsStr := strings.Join(elementDatas, ",")
	sizesStr := strings.Join(sizeDatas, ",")
	expectedConfig := fmt.Sprintf("data-configure='{\"Elements\":[%v],\"ExternalStylePath\":null,\"EditURL\":\"/admin/qor_banner_editor_settings/:id/edit\",\"Platforms\":[%v]}'", elementsStr, sizesStr)
	expectedConfig = strings.Replace(expectedConfig, "\"", "&#34;", -1)
	expectedConfig = strings.Replace(expectedConfig, "'", "\"", -1)
	assetPageHaveText(t, string(body), expectedConfig)
}
