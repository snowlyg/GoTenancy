package wildcard_router_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"GoTenancy/libs/wildcard_router"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

var (
	mux    = http.NewServeMux()
	Server = httptest.NewServer(mux)
)

type ModuleBeforeA struct {
}

func (a ModuleBeforeA) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/module_a0" {
		w.Header().Set("Content-Type", "text/moduleBeforeA")
		_, err := w.Write([]byte("Module Before A handled"))
		if err != nil {
			panic("ModuleBeforeA A can't handle")
		}
	} else {
		http.NotFound(w, req)
	}
}

type ModuleA struct {
}

func (a ModuleA) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/module_a0" || req.URL.Path == "/module_a" || req.URL.Path == "/module_ab" {
		w.Header().Set("Content-Type", "text/moduleA")
		_, err := w.Write([]byte("Module A handled"))
		if err != nil {
			panic("Module A can't handle")
		}
	} else {
		http.NotFound(w, req)
	}
}

type ModuleB struct {
}

func (b ModuleB) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path == "/module_b" || req.URL.Path == "/module_ab" {
		w.Header().Set("Content-Type", "text/moduleB")
		_, err := w.Write([]byte("Module B handled"))
		if err != nil {
			panic("Module B can't handle")
		}
	} else {
		http.NotFound(w, req)
	}
}

func init() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/gin")
		c.Writer.Write([]byte("Gin Handle HomePage"))
	})

	wildcardRouter := wildcard_router.New()
	wildcardRouter.MountTo("/", mux)
	wildcardRouter.AddHandler(router)
	wildcardRouter.AddHandler(ModuleBeforeA{})
	wildcardRouter.AddHandler(ModuleA{})
	wildcardRouter.AddHandler(ModuleB{})
	wildcardRouter.NoRoute(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("<h1>Sorry, this page was gone!</h1>"))
	})
}

type WildcardRouterTestCase struct {
	URL               string
	ExpectStatusCode  int
	ExpectHasContent  string
	ExpectContentType string
}

func TestWildcardRouter(t *testing.T) {
	testCases := []WildcardRouterTestCase{
		{URL: "/", ExpectStatusCode: 200, ExpectHasContent: "Gin Handle HomePage", ExpectContentType: "text/gin"},
		{URL: "/module_a", ExpectStatusCode: 200, ExpectHasContent: "Module A handled", ExpectContentType: "text/moduleA"},
		{URL: "/module_b", ExpectStatusCode: 200, ExpectHasContent: "Module B handled", ExpectContentType: "text/moduleB"},
		{URL: "/module_x", ExpectStatusCode: 404, ExpectHasContent: "<h1>Sorry, this page was gone!</h1>", ExpectContentType: "text/html; charset=utf-8"},
		{URL: "/module_a0", ExpectStatusCode: 200, ExpectHasContent: "Module Before A handled", ExpectContentType: "text/moduleBeforeA"},
	}

	for i, testCase := range testCases {
		var hasError bool
		req, _ := http.Get(Server.URL + testCase.URL)
		content, _ := ioutil.ReadAll(req.Body)
		if req.StatusCode != testCase.ExpectStatusCode {
			t.Errorf(color.RedString(fmt.Sprintf("WildcardRouter #%v: HTML expect status code '%v', but got '%v'", i+1, testCase.ExpectStatusCode, req.StatusCode)))
			hasError = true
		}
		if string(content) != testCase.ExpectHasContent {
			t.Errorf(color.RedString(fmt.Sprintf("WildcardRouter #%v: HTML expect have content '%v', but got '%v'", i+1, testCase.ExpectHasContent, string(content))))
			hasError = true
		}
		if req.Header["Content-Type"][0] != testCase.ExpectContentType {
			t.Errorf(color.RedString(fmt.Sprintf("WildcardRouter #%v: Expect response Content-Type is '%v', but got '%v'", i+1, testCase.ExpectContentType, req.Header["Content-Type"][0])))
			hasError = true
		}
		if !hasError {
			fmt.Printf(color.GreenString(fmt.Sprintf("WildcardRouter #%v: Success\n", i+1)))
		}
	}
}
