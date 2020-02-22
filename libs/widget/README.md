# Widget

Widget are defined, customizable, shareable HTML *widgets* that can be used within [QOR Admin](https://GoTenancy/libs/admin) or your own frontend.

## Getting Started

```go
// Initialize widgets container
Widgets := widget.New(&widget.Config{DB: db})

// Define widget's setting struct
type bannerArgument struct {
  Title           string
  Link            string
  BackgroundImage oss.OSS
}

// Register new widget
Widgets.RegisterWidget(&widget.Widget{
  // Widget's Name
  Name:     "Banner",
  // Widget's available templates
  Templates: []string{"banner1", "banner2"},
  // Widget's setting, which is configurable from the place using the widget with inline edit
  Setting:  Admin.NewResource(&bannerArgument{}),
  // InlineEditURL overwrite widget's inline edit's setting URL, usually we don't require do that, Widget will generate setting form based on Widget's setting resource, but in case of you want to fully customize your widget's setting page, you can overwrite it to return setting's URL
  InlineEditURL: func(context *Context) string {
    return "/different-setting-url"
  },
  // Generate a context to render the widget based on the widget's configurations
  Context: func(context *widget.Context, setting interface{}) *widget.Context {
    context.Options["Setting"] = setting
    context.Options["CurrentTime"] = time.Now()
    return context
  },
})


type slideImage struct {
  Title string
  Image oss.OSS
}

type slideShowArgument struct {
  SlideImages []slideImage
}

slideShowResource := Admin.NewResource(&slideShowArgument{})
slideShowResource.AddProcessor(func(value interface{}, metaValues *resource.MetaValues, context *qor.Context) error {
  if slides, ok := value.(*slideShowArgument); ok {
    for _, slide := range slides.SlideImages {
      if slide.Title == "" {
        return errors.New("slide title is blank")
      }
    }
  }
  return nil
})

Widgets.RegisterWidget(&widget.Widget{
  Name:      "SlideShow",
  Templates: []string{"slideshow"},
  Setting:   slideShowResource,
  Context: func(context *widget.Context, setting interface{}) *widget.Context {
    context.Options["Setting"] = setting
    return context
  },
})

// Define Widget Group
Widgets.RegisterWidgetGroup(&widget.WidgetGroup{
  Name: "HomeBanners",
  Widgets: []string{"Banner", "SlideShow"},
})

// Manage widgets from QOR Admin interface
Admin.AddResource(Widgets)
```

### Render Widget From Controller

```go
func Index(request *http.Request, writer http.ResponseWriter) {
  // Generate widget context
  widgetContext := admin.Widgets.NewContext(&widget.Context{
    DB:         DB(request),
    Options:    map[string]interface{}{"Request": request, "CurrentUser": currentUser}, // those options are accessible from widget views
    InlineEdit: true, // enable inline edit mode for widget
  })

  // Render Widget `HomeBanner` based on widget `Banner`'s definition to HTML template
  homeBannerContent := widgetContext.Render("HomeBanner", "Banner")

  // Render Widget `CampaignBanner` based on widget `Banner`'s definition to HTML template
  campaignBannerContent := widgetContext.Render("CampaignBanner", "Banner")

  // Then you could use the banner's HTML content in your templates
}
```

### Templates

Widget's template is Golang HTML template based, and thus quite flexible.

Continuing on the above *widget* example, we configure it to use templates `banner1`, `banner2`. QOR widget will look up templates from path `$APP_ROOT/app/views/widgets` by default, so you need to put your template in the right folder!

Once the template is found, QOR Widget will use the widget's generated context and use it to render the template like how Golang template works.

```go
// app/views/widgets/banner1.tmpl
<div class="banner" style="background:url("{{.Setting.BackgroundImage}}") no-repeat center center">
  <div class="container">
    <div class="row">
      <div class="column column-12">
        <a href="{{.Setting.Link}}" class="button button__primary">{{.Setting.Title}}</a>
        {{.CurrentTime}}
      </div>
    </div>
  </div>
</div>

// app/views/widgets/banner2.tmpl
<div class="banner">
  <div class="column column-9">
    <img src="{{.Setting.BackgroundImage}}" class="banner__logo" />
  </div>

  <div class="column column-3" style="margin-top: 2em;">
    <div><a href="{{.Setting.Link}}" class="button button__primary">{{.Setting.Title}}</a></div>
  </div>
</div>
```

You could also register other paths like:

```go
Widgets.RegisterViewPath("app/views/widgets")
```

### Register Scopes

In some cases, you might want to show a different page according to UTM source/medium/campaign or for A/B testing to increase the conversion rate. Using a Widget's Scope you can register some *scopes*, configure your *widget* for each of them, and when a *scope*'s `Visible` condition is matched the *widget* will use its configuration to render the *widget*.

```go
Widgets.RegisterScope(&widget.Scope{
  Name: "From Google",
  Visible: func(context *widget.Context) bool {
    if request, ok := context.Get("Request"); ok {
      return strings.Contains(request.(*http.Request).Referer(), "google.com")
    }
    return false
  },
})

Widgets.RegisterScope(&widget.Scope{
  Name: "VIP User",
  Visible: func(context *widget.Context) bool {
    if user, ok := context.Get("CurrentUser"); ok {
      return user.(*User).Role == "vip"
    }
    return false
  },
})
```

## Live Widget Demo

[Widget Demo](http://demo.getqor.com)

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
