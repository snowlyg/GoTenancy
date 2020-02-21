package responder_test

import (
	"net/http"
	"testing"

	"GoTenancy/libs/responder"
)

func checkRespond(request *http.Request, format string, t *testing.T) {
	responder.With("html", func() {
		if format != "html" {
			t.Errorf("Should call %v, but called html", format)
		}
	}).With("json", func() {
		if format != "json" {
			t.Errorf("Should call %v, but called json", format)
		}
	}).With("xml", func() {
		if format != "xml" {
			t.Errorf("Should call %v, but called xml", format)
		}
	}).Respond(request)
}

func newRequestWithAcceptType(acceptType string) *http.Request {
	request, _ := http.NewRequest("GET", "", nil)
	request.Header.Add("Accept", acceptType)
	return request
}

func TestRespond(t *testing.T) {
	mimeMap := map[string]string{
		"text/html":        "html",
		"application/json": "json",
		"application/xml":  "xml",
	}

	for mimeType, format := range mimeMap {
		checkRespond(newRequestWithAcceptType(mimeType), format, t)
	}
}
