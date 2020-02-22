# SEO

The SEO library allows for the management and injection of dynamic data into HTML tags for the purpose of Search Engine Optimisation. Using the QOR Admin interface, an administrator can easily manage the content of an HTML page's title, description, and meta tags.

[![GoDoc](https://godoc.org/GoTenancy/libs/seo?status.svg)](https://godoc.org/GoTenancy/libs/seo)

## Definition

```go
// The `QorSeoSetting` struct is a normal GORM-backend model, need to run migration before using it
db.AutoMigrate(&seo.QorSeoSetting{})

// SeoGlobalSetting used to generate `Site-wide Settings` part
type SeoGlobalSetting struct {
    SiteName string
}

SeoCollection = seo.New()

// Configure `Site-wide Settings`
SeoCollection.RegisterGlobalVaribles(&SeoGlobalSetting{SiteName: "ASICS"})

// Configure SEO storage model, you could customize it by embed seo.QorSeoSetting to your custom model
SeoCollection.SettingResource = Admin.AddResource(&seo.QorSeoSetting{}, &admin.Config{Name: "SEO", Invisible: true})

// Configure `Page Metadata Defaults`
SeoCollection.RegisterSeo(&seo.SEO{
    Name:     "Default Page",
})

SeoCollection.RegisterSeo(&seo.SEO{
    Name:     "Category Page",
    // Defined what Varibles could be using in title, description and keywords
    Varibles: []string{"CategoryName"},
    // Generated a mapping to replace the Variable, e.g. title: 'Qor - {{CategoryName}}", will be dislayed as 'Qor - Clothing'
    Context: func(objects ...interface{}) map[string]string {
        values := make(map[string]string)
        if len(objects) > 0 {
            category := objects[0].(Category)
            values["CategoryName"] = category.Name
        }
        return values
    },
})
```

## Usage

```go
qorContext := &qor.Context{DB: db}

// render default meta tags
SeoCollection.Render(qorContext, "Default Page")

// render cateogry pages' meta tags
var category Category
db.First(&category, "code = ?", "clothing")
SeoCollection.Render(qorContext, "Category Page", category)

```

## Structured Data

```go
// micro search
seo.MicroSearch{
  URL:    "http://demo.getqor.com",
  Target: "http://demo.getqor.com/search?q=",
}.Render()

// micro contact
seo.MicroContact{
  URL:         "http://demo.getqor.com",
  Telephone:   "080-0012-3232",
  ContactType: "Customer Service",
}.Render()

// micro product
seo.MicroProduct{
  Name: "Kenmore White 17 Microwave",
  Image: "http://getqor.com/source/images/qor-logo.png",
  Description: "0.7 cubic feet countertop microwave. Has six preset cooking categories and convenience features like Add-A-Minute and Child Lock."
  BrandName: "ThePlant",
  SKU: "L1212",
  PriceCurrency: "USD",
  Price: 100,
  SellerName: "ThePlant",
}.Render()
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
