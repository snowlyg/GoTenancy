package i18n

import (
	"fmt"
	"testing"
)

type backend struct{}

func (b *backend) LoadTranslations() (translations []*Translation) { return translations }
func (b *backend) SaveTranslation(t *Translation) error            { return nil }
func (b *backend) DeleteTranslation(t *Translation) error          { return nil }

const BIGNUM = 10000

// run TestConcurrent* tests with -race flag would be better

func TestConcurrentReadWrite(t *testing.T) {
	i18n := New(&backend{})
	go func() {
		for i := 0; i < BIGNUM; i++ {
			i18n.AddTranslation(&Translation{Key: fmt.Sprintf("xx-%d", i), Locale: "xx", Value: fmt.Sprint(i)})
		}
	}()
	for i := 0; i < BIGNUM; i++ {
		i18n.T("xx", fmt.Sprintf("xx-%d", i))
	}
}

func TestConcurrentDeleteWrite(t *testing.T) {
	i18n := New(&backend{})
	go func() {
		for i := 0; i < BIGNUM; i++ {
			i18n.AddTranslation(&Translation{Key: fmt.Sprintf("xx-%d", i), Locale: "xx", Value: fmt.Sprint(i)})
		}
	}()
	for i := 0; i < BIGNUM; i++ {
		i18n.DeleteTranslation(&Translation{Key: fmt.Sprintf("xx-%d", i), Locale: "xx", Value: fmt.Sprint(i)})
	}
}

func TestFallbackLocale(t *testing.T) {
	i18n := New(&backend{})
	i18n.AddTranslation(&Translation{Key: "hello-world", Locale: "en-AU", Value: "Hello World"})

	if i18n.Fallbacks("en-AU").T("en-UK", "hello-world") != "Hello World" {
		t.Errorf("Should fallback en-UK to en-US")
	}

	if i18n.T("en-DE", "hello-world") != "hello-world" {
		t.Errorf("Haven't setup any fallback")
	}
}
