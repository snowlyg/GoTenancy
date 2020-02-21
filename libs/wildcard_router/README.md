# WildcardRouter

WildcardRouter handles [*dynamic routes*](https://en.wikipedia.org/wiki/Dynamic_routing).

[![GoDoc](https://godoc.org/GoTenancy/libs/wildcard_router?status.svg)](https://godoc.org/GoTenancy/libs/wildcard_router)

## Usage

Suppose you have a model, `Page`, that handles requests by `ServeHTTP` function based on a URL in a database.

```go
import (
  "GoTenancy/libs/wildcard_router"
)

type PageHandler struct{}

type Page struct {
	URL  string
	Body string
}

// Page's records in database:

// Record1(URL: /page1, Content: "Page1")
// Record2(URL: /page2, Content: "Page2")

func (PageHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var page Page
	if !db.First(&page, "url = ?", req.URL.Path).RecordNotFound() {
		w.Write([]byte(page.Body))
	}
}
```

And you have another model, `FAQ`, which also handles requests by `ServeHTTP` function based on a URL in a database...

```go
type FAQHandler struct{}

type FAQ struct {
	URL      string
	Question string
	Answer   string
}

// FAQ's records in database:

// Record1(URL: /faq1, Question: "FAQ1", Answer: "Answer1")
// Record2(URL: /faq2, Question: "FAQ2", Answer: "Answer2")

func (FAQHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var faq FAQ
	if !db.First(&faq, "url = ?", req.URL.Path).RecordNotFound() {
		w.Write([]byte(fmt.Sprintf("%v: %v", faq.Question, faq.Answer)))
	}
}
```

Those URLs could be anything, with no rule... Let's initialize [WildcardRouter](https://GoTenancy/libs/wildcard_router) and mount it.

```go
func main() {
	mux := http.NewServeMux()

	wildcardRouter := wildcard_router.New()
	wildcardRouter.MountTo("/", mux)
}
```

`AddHandler` to [WildcardRouter](https://GoTenancy/libs/wildcard_router) and any model that implements method `ServeHTTP` can be routed to as a handler.

```go
    wildcardRouter.AddHandler(PageHandler{})
    wildcardRouter.AddHandler(FAQHandler{})
```

If you would like to customize your 404 page, you could set handlefunc to `NoRoute`

```go
  wildcardRouter.NoRoute(func(w http.ResponseWriter, req *http.Request) {
      // You need to set Content-Type to `text/html` if you would like brower recognize as HTML
      w.Header().Set("Content-Type", "text/html; charset=utf-8")
      // You need to set Status as NotFound too if you would like status code is 404
      w.WriteHeader(http.StatusNotFound)
      w.Write([]byte("Sorry, this page was gone!"))
  })
```

The behavior will be:

```go
// Visit "/page1"   will return "Page1"
// Visit "/page2"   will return "Page2"
// Visit "/faq1"    will return "FAQ1: Answer1"
// Visit "/faq2"    will return "FAQ2: Answer2"
// Visit "/unknown" will return "Sorry, this page was gone!" with statu code 404
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
