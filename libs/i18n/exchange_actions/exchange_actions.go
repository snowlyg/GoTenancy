package exchange_actions

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"
	"sort"
	"strings"
	"time"

	"GoTenancy/libs/admin"
	"GoTenancy/libs/i18n"
	"GoTenancy/libs/media/oss"
	"GoTenancy/libs/worker"
)

type ExportTranslationArgument struct {
	Scope string
}

type ImportTranslationArgument struct {
	TranslationsFile oss.OSS
}

// RegisterExchangeJobs register i18n jobs into worker
func RegisterExchangeJobs(I18n *i18n.I18n, Worker *worker.Worker) {
	if I18n.Resource == nil {
		debug.PrintStack()
		fmt.Println("I18n should be registered into `Admin` before register jobs")
		return
	}

	Admin := I18n.Resource.GetAdmin()
	admin.RegisterViewPath("libs/i18n/exchange_actions/views")

	// Export Translations
	exportTranslationResource := Admin.NewResource(&ExportTranslationArgument{})
	exportTranslationResource.Meta(&admin.Meta{Name: "Scope", Type: "select_one", Collection: []string{"All", "Backend", "Frontend"}})

	Worker.RegisterJob(&worker.Job{
		Name:     "Export Translations",
		Group:    "Export/Import Translations From CSV file",
		Resource: exportTranslationResource,
		Handler: func(arg interface{}, qorJob worker.QorJobInterface) (err error) {
			var (
				locales          []string
				translationKeys  []string
				translationsMap  = map[string]bool{}
				filename         = fmt.Sprintf("/downloads/translations.%v.csv", time.Now().UnixNano())
				fullFilename     = filepath.Join("public", filename)
				i18nTranslations = I18n.LoadTranslations()
				scope            = arg.(*ExportTranslationArgument).Scope
			)
			qorJob.AddLog("Exporting translations...")

			// Sort locales
			for locale := range i18nTranslations {
				locales = append(locales, locale)
			}
			sort.Strings(locales)

			// Create download file
			if _, err = os.Stat(filepath.Dir(fullFilename)); os.IsNotExist(err) {
				err = os.MkdirAll(filepath.Dir(fullFilename), os.ModePerm)
			}
			csvfile, err := os.OpenFile(fullFilename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
			defer csvfile.Close()
			if err != nil {
				return err
			}

			writer := csv.NewWriter(csvfile)

			// Append Headers
			writer.Write(append([]string{"Translation Keys"}, locales...))

			// Sort translation keys
			for _, locale := range locales {
				for key := range i18nTranslations[locale] {
					translationsMap[key] = true
				}
			}

			for key := range translationsMap {
				translationKeys = append(translationKeys, key)
			}
			sort.Strings(translationKeys)

			// Write CSV file
			var (
				recordCount         = len(translationKeys)
				perCount            = recordCount/20 + 1
				processedRecordLogs = []string{}
				index               = 0
				progressCount       = 0
			)
			for _, translationKey := range translationKeys {
				// Filter out translation by scope
				index++
				if scope == "Backend" && !strings.HasPrefix(translationKey, "qor_") {
					continue
				}
				if scope == "Frontend" && strings.HasPrefix(translationKey, "qor_") {
					continue
				}
				var translations = []string{translationKey}
				for _, locale := range locales {
					var value string
					if translation := i18nTranslations[locale][translationKey]; translation != nil {
						value = translation.Value
					}
					translations = append(translations, value)
				}
				writer.Write(translations)
				processedRecordLogs = append(processedRecordLogs, fmt.Sprintf("Exported %v\n", strings.Join(translations, ",")))
				if index == perCount {
					qorJob.AddLog(strings.Join(processedRecordLogs, ""))
					processedRecordLogs = []string{}
					progressCount++
					qorJob.SetProgress(uint(float32(progressCount) / float32(20) * 100))
					index = 0
				}
			}
			writer.Flush()

			qorJob.SetProgressText(fmt.Sprintf("<a href='%v'>Download exported translations</a>", filename))
			return
		},
	})

	// Import Translations

	Worker.RegisterJob(&worker.Job{
		Name:     "Import Translations",
		Group:    "Export/Import Translations From CSV file",
		Resource: Admin.NewResource(&ImportTranslationArgument{}),
		Handler: func(arg interface{}, qorJob worker.QorJobInterface) (err error) {
			importTranslationArgument := arg.(*ImportTranslationArgument)
			qorJob.AddLog("Importing translations...")
			if csvfile, err := os.Open(filepath.Join("public", importTranslationArgument.TranslationsFile.URL())); err == nil {
				reader := csv.NewReader(csvfile)
				reader.TrimLeadingSpace = true
				if records, err := reader.ReadAll(); err == nil {
					if len(records) > 1 && len(records[0]) > 1 {
						var (
							recordCount         = len(records) - 1
							perCount            = recordCount/20 + 1
							processedRecordLogs = []string{}
							locales             = records[0][1:]
							index               = 1
						)
						for _, values := range records[1:] {
							logMsg := ""
							for idx, value := range values[1:] {
								if value == "" {
									if values[0] != "" && locales[idx] != "" {
										I18n.DeleteTranslation(&i18n.Translation{
											Key:    values[0],
											Locale: locales[idx],
										})
										logMsg += fmt.Sprintf("%v/%v Deleted %v,%v\n", index, recordCount, locales[idx], values[0])
									}
								} else {
									I18n.SaveTranslation(&i18n.Translation{
										Key:    values[0],
										Locale: locales[idx],
										Value:  value,
									})
									logMsg += fmt.Sprintf("%v/%v Imported %v,%v,%v\n", index, recordCount, locales[idx], values[0], value)
								}
							}
							processedRecordLogs = append(processedRecordLogs, logMsg)
							if len(processedRecordLogs) == perCount {
								qorJob.AddLog(strings.Join(processedRecordLogs, ""))
								processedRecordLogs = []string{}
								qorJob.SetProgress(uint(float32(index) / float32(recordCount+1) * 100))
							}
							index++
						}
						qorJob.AddLog(strings.Join(processedRecordLogs, ""))
					}
				}
				qorJob.AddLog("Imported translations")
			}
			return
		},
	})
}
