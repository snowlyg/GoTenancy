# L10n

L10n gives your [GORM](https://github.com/jinzhu/gorm) models the ability to localize for different Locales. It can be a catalyst for the adaptation of a product, application, or document content to meet the language, cultural, and other requirements of a specific target market.

[![GoDoc](https://godoc.org/GoTenancy/libs/l10n?status.svg)](https://godoc.org/GoTenancy/libs/l10n)

## Usage

L10n utilizes [GORM](https://github.com/jinzhu/gorm) callbacks to handle localization, so you will need to register callbacks first:

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/l10n"
)

func main() {
  db, err := gorm.Open("sqlite3", "demo_db")
  l10n.RegisterCallbacks(&db)
}
```

### Making a Model Localizable

Embed `l10n.Locale` into your model as an anonymous field to enable localization, for example, in a hypothetical project which has a focus on Product management:

```go
type Product struct {
  gorm.Model
  Name string
  Code string
  l10n.Locale
}
```

`l10n.Locale` will add a `language_code` column as a composite primary key with existing primary keys, using [GORM](https://github.com/jinzhu/gorm)'s AutoMigrate to create the field.

The `language_code` column will be used to save a localized model's Locale. If no Locale is set, then the global default Locale (`en-US`) will be used. You can override the global default Locale by setting `l10n.Global`, for example:

```go
l10n.Global = 'zh-CN'
```

### Create localized resources from global product

```go
// Create global product
product := Product{Name: "Global product", Description: "Global product description"}
DB.Create(&product)
product.LanguageCode   // "en-US"

// Create zh-CN product
product.Name = "中文产品"
DB.Set("l10n:locale", "zh-CN").Create(&product)

// Query zh-CN product with primary key 111
DB.Set("l10n:locale", "zh-CN").First(&productCN, 111)
productCN.Name         // "中文产品"
productCN.LanguageCode // "zh"
```

#### Create localized resource directly

By default, only global data allowed to be created, local data have to localized from global one.

If you want to allow user create localized data directly, you can embeded `l10n.LocaleCreatable` for your model/struct, e.g:

```go
type Product struct {
  gorm.Model
  Name string
  Code string
  l10n.LocaleCreatable
}
```

### Keeping localized resources' fields in sync

Add the tag `l10n:"sync"` to the fields that you wish to always sync with the *global* record:

```go
type Product struct {
  gorm.Model
  Name  string
  Code  string `l10n:"sync"`
  l10n.Locale
}
```

Now the localized product's `Code` will be the same as the global product's `Code`. The `Code` is not affected by localized resources, and when the global record changes its `Code` the localized records' `Code` will be synced automatically.

### Query Modes

L10n provides 5 modes for querying.

* global   - find all global records,
* locale   - find localized records,
* reverse  - find global records that haven't been localized,
* unscoped - raw query, won't auto add `locale` conditions when querying,
* default  - find localized record, if not found, return the global one.

You can specify the mode in this way:

```go
dbCN := db.Set("l10n:locale", "zh-CN")

mode := "global"
dbCN.Set("l10n:mode", mode).First(&product, 111)
// SELECT * FROM products WHERE id = 111 AND language_code = 'en-US';

mode := "locale"
db.Set("l10n:mode", mode).First(&product, 111)
// SELECT * FROM products WHERE id = 111 AND language_code = 'zh-CN';
```

## Qor Integration

Although L10n could be used alone, it integrates nicely with [QOR](https://GoTenancy/libs/qor).

[L10n Demo with QOR](http://demo.getqor.com/admin/products)

By default, [QOR](https://GoTenancy/libs/qor) will only allow you to manage the global language. If you have configured [Authentication](http://doc.getqor.com/admin/authentication.html), [QOR Admin](http://GoTenancy/libs/admin) will try to obtain the allowed Locales from the current user.

* Viewable Locales - Locales for which the current user has read permission:

```go
func (user User) ViewableLocales() []string {
  return []string{l10n.Global, "zh-CN", "JP", "EN", "DE"}
}
```

* <a name='editable-locales'></a> Editable Locales - Locales for which the current user has manage (create/update/delete) permission:

```go
func (user User) EditableLocales() []string {
  if user.role == "global_admin" {
    return []string{l10n.Global, "zh-CN", "EN"}
  } else {
    return []string{"zh-CN", "EN"}
  }
}
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
