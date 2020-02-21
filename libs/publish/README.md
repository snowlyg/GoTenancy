## Publish

The Publish Plugin decouples the timing of data updates in the [QOR Admin](http://GoTenancy/libs/admin) interface from the display of data on the frontend of the website which [QOR Admin](http://GoTenancy/libs/admin) is backing. Effectively, a [GORM](https://github.com/jinzhu/gorm) model can be withheld from frontend display until it is "published".

[![GoDoc](https://godoc.org/GoTenancy/libs/publish?status.svg)](https://godoc.org/GoTenancy/libs/publish)

## Usage

Embed `publish.Status` as an anonymous field in your model to apply the [Publish](https://GoTenancy/libs/publish) feature.

```go
type Product struct {
  Name        string
  Description string
  publish.Status
}
```

```go
db, err = gorm.Open("sqlite3", "demo_db")
publish := publish.New(db)

// Run migration to add `publish_status` column from `publish.Status`
db.AutoMigrate(&Product{})
// Create `products_draft` table
publish.AutoMigrate(&Product{})

draftDB := publish.DraftDB() // Draft resources are saved here
productionDB := publish.ProductionDB() // Published resources saved here
```

With the draft db, all your changes will be saved in `products_draft`, with the `productionDB`, all your changes will be in `products` table and sync back to `products_draft`.

```go
// Publish changes you made in `products_draft` to `products`
publish.Publish(&product)

// Overwrite this product's data in `products_draft` with its data in `products`
publish.Discard(&product)
```

## Publish Event

To avoid a large amount (*n*) of draft events when applying the [Publish](https://GoTenancy/libs/publish) feature to a set of (*n*) records, it is possible to batch the events into a single, high-level event which represents the set of (*n*) events. To do this, use the `PublishEvent` feature.

For example, when sorting products, say you have changed 100 products' positions but don't want to show 100 products as draft nor have to go through the horrible process of publishing 100 products, one at a time. Instead, you can show a single, high-level event such as `Changed products' sorting` in the draft page. After publishing this single event, all of the associated position changes will be published.


```go
// Register Publish Event
publish.RegisterEvent("changed_product_sorting", changedSortingPublishEvent{})

// Publish Event definition
type changedSortingPublishEvent struct {
  Table       string
  PrimaryKeys []string
}

func (e changedSortingPublishEvent) Publish(db *gorm.DB, event publish.PublishEventInterface) (err error) {
  if event, ok := event.(*publish.PublishEvent); ok {
    if err = json.Unmarshal([]byte(event.Argument), &e); err == nil {
      // e.PrimaryKeys => get primary keys
      // sync your changes made for above primary keys to production
    }
  }
}

func (e changedSortingPublishEvent) Discard(db *gorm.DB, event publish.PublishEventInterface) (err error) {
  // discard your changes made in draft
}

func init() {
  // change an `changed_product_sorting` event, including primary key 1, 2, 3
  db.Create(&publish.PublishEvent{
    Name: "changed_product_sorting",
    Argument: "[1,2,3]",
  })
}
```

## Store all changes in draft table by default

If you want all changes made to be stored in the draft table by default, initialize [QOR Admin](http://GoTenancy/libs/admin) with the [Publish](https://GoTenancy/libs/publish) value's draft DB. If you then want to manage those draft data, add the [Publish](https://GoTenancy/libs/publish) value as resource to [QOR Admin](../chapter2/setup.md):

```go
Admin := admin.New(&qor.Config{DB: Publish.DraftDB()})
Admin.AddResource(Publish)
```

[Publish Demo: http://demo.getqor.com/admin/publish](http://demo.getqor.com/admin/publish)
## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
