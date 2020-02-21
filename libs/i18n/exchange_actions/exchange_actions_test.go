package exchange_actions_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/i18n"
	"GoTenancy/libs/i18n/backends/database"
	"GoTenancy/libs/i18n/exchange_actions"
	"GoTenancy/libs/media"
	"GoTenancy/libs/media/oss"
	"GoTenancy/libs/qor"
	"GoTenancy/libs/qor/test/utils"
	"GoTenancy/libs/worker"
	"github.com/fatih/color"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB
var Worker *worker.Worker
var I18n *i18n.I18n

func init() {
	db = utils.TestDB()
	reset()
}

func reset() {
	db.DropTable(&database.Translation{})
	database.New(db)
	Admin := admin.New(&qor.Config{DB: db})
	Worker = worker.New()
	Admin.AddResource(Worker)
	I18n = i18n.New(database.New(db))
	I18n.SaveTranslation(&i18n.Translation{Key: "qor_admin.title", Value: "title", Locale: "en-US"})
	I18n.SaveTranslation(&i18n.Translation{Key: "qor_admin.subtitle", Value: "subtitle", Locale: "en-US"})
	I18n.SaveTranslation(&i18n.Translation{Key: "qor_admin.description", Value: "description", Locale: "en-US"})
	I18n.SaveTranslation(&i18n.Translation{Key: "header.title", Value: "Header Title", Locale: "en-US"})
	exchange_actions.RegisterExchangeJobs(I18n, Worker)
}

// Test export translations with scope
type testExportWithScopedCase struct {
	Scope            string
	ExpectExportFile string
}

func TestExportTranslations(t *testing.T) {
	reset()
	I18n.SaveTranslation(&i18n.Translation{Key: "header.title", Value: "标题", Locale: "zh-CN"})

	testCases := []*testExportWithScopedCase{
		&testExportWithScopedCase{Scope: "", ExpectExportFile: "export_all.csv"},
		&testExportWithScopedCase{Scope: "All", ExpectExportFile: "export_all.csv"},
		&testExportWithScopedCase{Scope: "Backend", ExpectExportFile: "export_backend.csv"},
		&testExportWithScopedCase{Scope: "Frontend", ExpectExportFile: "export_frontend.csv"},
	}

	for i, testcase := range testCases {
		clearDownloadDir()
		for _, job := range Worker.Jobs {
			if job.Name == "Export Translations" {
				job.Handler(&exchange_actions.ExportTranslationArgument{Scope: testcase.Scope}, job.NewStruct().(worker.QorJobInterface))
				if downloadedFileContent() != loadFixture(testcase.ExpectExportFile) {
					t.Errorf(color.RedString(fmt.Sprintf("\nExchange TestCase #%d: Failure (%s)\n", i+1, "export results are incorrect")))
				} else {
					color.Green(fmt.Sprintf("Export with scope TestCase #%d: Success\n", i+1))
				}
			}
		}
	}
}

// Test import translations
type testImportTranslationsCase struct {
	ImportFileDesc string
	ImportFile     string
	ExpectZhValues map[string]string
}

func TestImportTranslations(t *testing.T) {
	reset()
	testCases := []*testImportTranslationsCase{
		&testImportTranslationsCase{
			ImportFileDesc: "Normal tranlsation file",
			ImportFile:     "import_1.csv",
			ExpectZhValues: map[string]string{"qor_admin.title": "标题", "qor_admin.subtitle": "小标题", "qor_admin.description": "描述", "header.title": "标题"},
		},
		&testImportTranslationsCase{
			ImportFileDesc: "Translation file with missing header.title",
			ImportFile:     "import_2.csv",
			ExpectZhValues: map[string]string{"qor_admin.title": "标题", "qor_admin.subtitle": "小标题", "qor_admin.description": "描述"},
		},
		&testImportTranslationsCase{
			ImportFileDesc: "Translation file with empty column",
			ImportFile:     "import_3.csv",
			ExpectZhValues: map[string]string{"qor_admin.title": "标题", "qor_admin.subtitle": "小标题", "qor_admin.description": "描述", "header.title": "标题"},
		},
	}

	for i, testCase := range testCases {
		for _, job := range Worker.Jobs {
			if job.Name == "Import Translations" {
				job.Handler(&exchange_actions.ImportTranslationArgument{TranslationsFile: oss.OSS{media.Base{Url: "imports/" + testCase.ImportFile}}}, job.NewStruct().(worker.QorJobInterface))
				translations := I18n.LoadTranslations()["zh-CN"]
				if len(translations) == 0 {
					t.Errorf(color.RedString(fmt.Sprintf("\nImport TestCase #%d: Failure (%s)\n", i+1, "Doesn't have Zh translations")))
				}
				for key, translation := range translations {
					if testCase.ExpectZhValues[key] != translation.Value {
						t.Errorf(color.RedString(fmt.Sprintf("\nImport TestCase #%d: Failure (%s)\n", i+1, "Zh translations not match")))
					}
				}
				color.Green(fmt.Sprintf("Import TestCase #%d: Success\n", i+1))
			}
		}
	}
}

// Helper functions
func clearDownloadDir() {
	files, _ := ioutil.ReadDir("./public/downloads")
	for _, f := range files {
		os.Remove("./public/downloads/" + f.Name())
	}
}

func downloadedFileContent() string {
	files, _ := ioutil.ReadDir("./public/downloads")
	for _, f := range files {
		if content, err := ioutil.ReadFile("./public/downloads/" + f.Name()); err == nil {
			return string(content)
		}
	}
	return ""
}

func loadFixture(fileName string) string {
	if content, err := ioutil.ReadFile("./fixtures/" + fileName); err == nil {
		return string(content)
	}
	return ""
}
