# Publish2

The Publish2 Plugin is the successor to [Publish](https://github.com/qor/publish). It generalizes publishing using 3 important modules:

* Visible: to flag an object be online/offline
* Schedule: to schedule objects to be online/offline automatically
* Version: to allow an object to have more than one copies and chain them together

You can read the [Introducing Publish2 blog](http://getqor.com/en/blogs/article/title=introducing-publish2) to understand our design idea in detail.

## Usage

First, add Publish2 fields to the model. You can choose the module you need, we provide composability here. Please note that it requires [GORM](https://github.com/jinzhu/gorm) as ORM.

```go
type Product struct {
  ...

  publish2.Version
  publish2.Schedule
  publish2.Visible
}
```

Then, register database callback.

```go
var db *gorm.DB
publish2.RegisterCallbacks(db)
```

See [here](#disable-callbacks) for callback details.

Last, Setup [QOR Admin](http://github.com/qor/admin)

```go
import (
  "github.com/qor/admin"
)

func init() {
  admin.New(&qor.Config{DB: db.DB.Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)})
}
```

Now, start your application, you should see the [Publish2](https://github.com/qor/publish2) UI has been displayed in QOR Admin UI.

## Publish2 UI intro

The [Publish2](https://github.com/qor/publish2) section will be added to the index and new/edit page.

[Demo Site](http://demo.getqor.com/admin/products)

#### How to publish a product immediately
Tick the `Publish ready` option and leave `Schedule Start At` and `Schedule End At` blank.

#### How to schedule to publish a product
Fill the `Schedule Start At` and `Schedule End At` fields.

The [Publish2](https://github.com/qor/publish2) section in index page

#### How to make a new version of a product
Click the `...` icon(C), you can see a "Create new version" button in the popup.

#### How to view all versions of a product
Click the clock icon(B) to toggle all versions panel.

#### Which version of a product will be live if there’re many version
In all versions panel, the one with green circle(A) icon is the live version.

## Advanced usage

### <a name="disable-callbacks"></a> Disable callbacks

Depend on the modules you used, [Publish2](https://github.com/qor/publish2) callback attaches different SQL conditions to your object queries.

This is a SQL sample of select product with language_id is 6. All 3 modules are integrated with `Product`.

```sql
SELECT * FROM `products`  WHERE (language_id = '6') AND ((products.id, `products`.version_priority) IN (SELECT products.id, MAX(`products`.version_priority) FROM `products` WHERE (scheduled_start_at IS NULL OR scheduled_start_at <= '2017-02-13 02:04:09') AND (scheduled_end_at IS NULL OR scheduled_end_at >= '2017-02-13 02:04:09') AND publish_ready = 'true' AND deleted_at IS NULL GROUP BY products.id))) ORDER BY `products`.`id`, `products`.version_priority DESC
```

- Visible: `publish_ready = 'true'`
- Version: `(products.id, `products`.version_priority) IN (SELECT products.id, MAX(`products`.version_priority))`
- Schedule: `(scheduled_start_at IS NULL OR scheduled_start_at <= 'CURRENT_TIME') AND (scheduled_end_at IS NULL OR scheduled_end_at >= 'CURRENT_TIME')`

Sometimes you may need do a pure query without [Publish2](https://github.com/qor/publish2) conditions. You disable callbacks by

```go
db.DB.Set(publish2.VersionMode, publish2.VersionMultipleMode).Set(publish2.VisibleMode, publish2.ModeOff).Set(publish2.ScheduleMode, publish2.ModeOff)
```

The `Set(publish2.VersionMode, publish2.VersionMultipleMode)` means use `VersionMultipleMode` in this query. The default `VersionMode` is single version mode. The difference between single and multiple is the single mode always query the live version from all versions, the multiple mode queries all versions.

### Customize default version name

The default version name is `Default`. To overwrite it, you can use

```go
publish2.DefaultVersionName = "1.0"
```

## Digging deeper: Modules of publish2

### Visible

`Visible` module controls the visibility of the object by `PublishReady`.

```go
type Visible struct {
  PublishReady bool
}
```

### Schedule

`Schedule` module schedules objects to be online/offline automatically by `ScheduledStartAt` and `ScheduledEndAt`.

```go
type Schedule struct {
  ScheduledStartAt *time.Time `gorm:"index"`
  ScheduledEndAt   *time.Time `gorm:"index"`
  ScheduledEventID *uint
}

type ScheduledEvent struct {
  gorm.Model
  Name             string
  ScheduledStartAt *time.Time
  ScheduledEndAt   *time.Time
}
```

If an object has `PublishReady` set to true and current date is inside the range between `ScheduledStartAt` and `ScheduledEndAt`, it will be visible. Otherwise, it is invisible.

Manage the date range one by one is inefficient, so we added `ScheduledEvent` to manage them. Imagine black Friday and Christmas are pre-created in the system, All you need to do is set a price of the product for these holidays.

We have integrated an UI for the `ScheduledEvent` at `QOR Admin > Sidebar > Publishing > Events`.

### Version

`Version` module allow one object to have multiple copies, with `Schedule`, you can schedule different prices of a product for a whole year.

```go
type Version struct {
  VersionName     string `gorm:"primary_key"`
  VersionPriority string `gorm:"index"`
}
```

The `VersionName` will be the primary key of the object, So if you set a new `VersionName` for an object, means you will create a new copy. To set a new version name. We have `obj.SetVersionName("new name")`. When an object has multiple versions the database would looks like:

| id | version_name | name |
| --- | --- | --- |
| 1 | v1 | Product - v1 |
| 1 | v2 | Product - v2 |

The `VersionPriority` represents the priority of current version. The rule when different versions have overlapped schedule range is, the newer the higher. For example, product A has version 1 for the Christmas(12-20 ~ 12-31) and version 2 for the New Year holiday(12-30 ~ 1-3). At 12-31, the version 2 will be the visible one.

To get version work properly, You need define an extra function to your struct called `GetStatus` like this:

```go
func (yourStruct YourStruct) GetStatus() string {
	return yourStruct.Status
}
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
