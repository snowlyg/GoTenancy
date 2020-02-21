package yaml_test

import (
	"fmt"
	"net/http"
	"testing"

	"GoTenancy/libs/i18n"
	"GoTenancy/libs/i18n/backends/yaml"
)

var values = map[string][][]string{
	"en": {
		{"hello", "Hello"},
		{"user.name", "User Name"},
		{"user.email", "Email"},
	},
	"de": {
		{"hello", "Hallo"},
		{"user.name", "Benutzername"},
		{"user.email", "E-Mail-Adresse"},
	},
	"zh-CN": {
		{"hello", "你好"},
		{"user.name", "用户名"},
		{"user.email", "邮箱"},
	},
}

func checkTranslations(translations []*i18n.Translation) error {
	for locale, results := range values {
		for _, result := range results {
			var found bool
			for _, translation := range translations {
				if (translation.Locale == locale) && (translation.Key == result[0]) && (translation.Value == result[1]) {
					found = true
				}
			}
			if !found {
				return fmt.Errorf("failed to found translation %v for %v", result[0], locale)
			}
		}
	}
	return nil
}

func TestLoadTranslations(t *testing.T) {
	backend := yaml.New("tests", "tests/subdir")
	if err := checkTranslations(backend.LoadTranslations()); err != nil {
		t.Fatal(err)
	}
}

func TestLoadTranslationsFilesystem(t *testing.T) {
	backend := yaml.NewWithFilesystem(http.Dir("./tests"))
	if err := checkTranslations(backend.LoadTranslations()); err != nil {
		t.Fatal(err)
	}
}

func TestLoadTranslationsWalk(t *testing.T) {
	backend := yaml.NewWithWalk("tests")
	if err := checkTranslations(backend.LoadTranslations()); err != nil {
		t.Fatal(err)
	}
}

var benchmarkResult error

func BenchmarkLoadTranslations(b *testing.B) {
	var backend i18n.Backend
	var err error
	for i := 0; i < b.N; i++ {
		backend = yaml.New("tests", "tests/subdir")
		if err = checkTranslations(backend.LoadTranslations()); err != nil {
			b.Fatal(err)
		}
	}
	benchmarkResult = err
}

var benchmarkResult2 error

func BenchmarkLoadTranslationsWalk(b *testing.B) {
	var backend i18n.Backend
	var err error
	for i := 0; i < b.N; i++ {
		backend = yaml.NewWithWalk("tests")
		if err = checkTranslations(backend.LoadTranslations()); err != nil {
			b.Fatal(err)
		}
	}
	benchmarkResult2 = err
}

var benchmarkResult3 error

func BenchmarkLoadTranslationsFilesystem(b *testing.B) {
	var backend i18n.Backend
	var err error
	for i := 0; i < b.N; i++ {
		backend = yaml.NewWithFilesystem(http.Dir("./tests"))
		if err = checkTranslations(backend.LoadTranslations()); err != nil {
			b.Fatal(err)
		}
	}
	benchmarkResult3 = err
}
