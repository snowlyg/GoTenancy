package admin

import (
	"fmt"
	"html/template"
	"testing"

	"GoTenancy/libs/qor"
	"github.com/fatih/color"
)

type rawTestCase struct {
	HTML         string
	ExpectResult string
}

func TestFuncMaps(t *testing.T) {
	rawTestCases := []rawTestCase{
		{HTML: "<a href='#'>Hello</a>", ExpectResult: "Hello"},
		{HTML: "<a href='http://www.google.com'>Hello</a>", ExpectResult: "<a href=\"http://www.google.com\" rel=\"nofollow\">Hello</a>"},
		{HTML: "<a href='http://www.google.com' data-hint='Hello'>Hello</a>", ExpectResult: "<a href=\"http://www.google.com\" rel=\"nofollow\">Hello</a>"},
	}

	unsafeRawTestCases := []rawTestCase{
		{HTML: "<a href='http://g.cn'>Hello</a>", ExpectResult: "<a href='http://g.cn'>Hello</a>"},
		{HTML: "<a href='#' data-hint='Hello'>Hello</a>", ExpectResult: "<a href='#' data-hint='Hello'>Hello</a>"},
	}

	context := Context{
		Admin: New(&qor.Config{}),
	}
	funcMaps := context.FuncMap()

	for i, testcase := range rawTestCases {
		result := funcMaps["raw"].((func(string) template.HTML))(testcase.HTML)
		var hasError bool
		if result != template.HTML(testcase.ExpectResult) {
			t.Errorf(color.RedString(fmt.Sprintf("Admin FuncMap raw #%v: expect get %v, but got '%v'", i+1, testcase.ExpectResult, result)))
			hasError = true
		}
		if !hasError {
			fmt.Printf(color.GreenString(fmt.Sprintf("Admin FuncMap raw #%v: Success\n", i+1)))
		}
	}

	for i, testcase := range unsafeRawTestCases {
		result := funcMaps["unsafe_raw"].((func(string) template.HTML))(testcase.HTML)
		var hasError bool
		if result != template.HTML(testcase.ExpectResult) {
			t.Errorf(color.RedString(fmt.Sprintf("Admin FuncMap unsafe_raw #%v: expect get %v, but got '%v'", i+1, testcase.ExpectResult, result)))
			hasError = true
		}
		if !hasError {
			fmt.Printf(color.GreenString(fmt.Sprintf("Admin FuncMap unsafe_raw #%v: Success\n", i+1)))
		}
	}
}
