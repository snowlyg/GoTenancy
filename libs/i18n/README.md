# I18n

I18n provides internationalization support for your application, it supports 2 kinds of storages(backends), the database and file system.

[![GoDoc](https://godoc.org/GoTenancy/libs/i18n?status.svg)](https://godoc.org/GoTenancy/libs/i18n)

## Usage

Initialize I18n with the storage mode. You can use both storages together, the earlier one has higher priority. So in the example, I18n will look up the translation in database first, then continue finding it in the YAML file if not found.

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/i18n"
  "GoTenancy/libs/i18n/backends/database"
  "GoTenancy/libs/i18n/backends/yaml"
)

func main() {
  db, _ := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")

  I18n := i18n.New(
    database.New(&db), // load translations from the database
    yaml.New(filepath.Join(config.Root, "config/locales")), // load translations from the YAML files in directory `config/locales`
  )

  I18n.T("en-US", "demo.greeting") // Not exist at first
  I18n.T("en-US", "demo.hello") // Exists in the yml file
}
```

Once a database has been set for I18n, all **untranslated** translations inside `I18n.T()` will be loaded into `translations` table in the database when compiling the application. For example, we have an untranslated `I18n.T("en-US", "demo.greeting")` in the example, so I18n will generate this record in the `translations` table after compiling.

| locale | key           | value  |
| ---    | ---           | ---    |
| en-US  | demo.greeting | &nbsp; |

The YAML file format is

```yaml
en-US:
  demo:
    hello: "Hello, world"
```

### Use built-in interface for translation management with [QOR Admin](http://GoTenancy/libs/admin)

I18n has a built-in web interface for translations which is integrated with [QOR Admin](http://GoTenancy/libs/admin).

```go
Admin.AddResource(I18n)
```

Then a page like this will be added to [QOR Admin](http://GoTenancy/libs/admin) interface

Refer the [online demo](http://demo.getqor.com/admin/translations).

### Use with Golang templates

The easy way to use I18n in a template is to define a `t` function and register it as `FuncMap`:

```go
func T(key string, value string, args ...interface{}) string {
  return I18n.Default(value).T("en-US", key, args...)
}

// then use it in the template
{{ t "demo.greet" "Hello, {{$1}}" "John" }} // -> Hello, John
```

### Built-in functions for translations management

I18n has functions to manage translation directly.

```go
// Add Translation
I18n.AddTranslation(&i18n.Translation{Key: "hello-world", Locale: "en-US", Value: "hello world"})

// Update Translation
I18n.SaveTranslation(&i18n.Translation{Key: "hello-world", Locale: "en-US", Value: "Hello World"})

// Delete Translation
I18n.DeleteTranslation(&i18n.Translation{Key: "hello-world", Locale: "en-US", Value: "Hello World"})
```

### Scope and default value

Call Translation with `Scope` or set default value.

```go
// Read Translation with `Scope`
I18n.Scope("home-page").T("zh-CN", "hello-world") // read translation with translation key `home-page.hello-world`

// Read Translation with `Default Value`
I18n.Default("Default Value").T("zh-CN", "non-existing-key") // Will return default value `Default Value`
```

### Fallbacks

I18n has a `Fallbacks` function to register fallbacks. For example, registering `en-GB` as a fallback to `zh-CN`:

```go
i18n := New(&backend{})
i18n.AddTranslation(&Translation{Key: "hello-world", Locale: "en-GB", Value: "Hello World"})

fmt.Print(i18n.Fallbacks("en-GB").T("zh-CN", "hello-world")) // "Hello World"
```

**To set fallback [*Locale*](https://en.wikipedia.org/wiki/Locale_(computer_software)) globally** you can use `I18n.FallbackLocales`. This function accepts a `map[string][]string` as parameter. The key is the fallback *Locale* and the `[]string` is the *Locales* that could fallback to the first *Locale*.

For example, setting `"fr-FR", "de-DE", "zh-CN"` fallback to `en-GB` globally:

```go
I18n.FallbackLocales = map[string][]string{"en-GB": []{"fr-FR", "de-DE", "zh-CN"}}
```

### Interpolation

I18n utilizes a Golang template to parse translations with an interpolation variable.

```go
type User struct {
  Name string
}

I18n.AddTranslation(&i18n.Translation{Key: "hello", Locale: "en-US", Value: "Hello {{.Name}}"})

I18n.T("en-US", "hello", User{Name: "Jinzhu"}) //=> Hello Jinzhu
```

### Pluralization

I18n utilizes [cldr](https://github.com/theplant/cldr) to achieve pluralization, it provides the functions `p`, `zero`, `one`, `two`, `few`, `many`, `other` for this purpose. Please refer to [cldr documentation](https://github.com/theplant/cldr) for more information.

```go
I18n.AddTranslation(&i18n.Translation{Key: "count", Locale: "en-US", Value: "{{p "Count" (one "{{.Count}} item") (other "{{.Count}} items")}}"})
I18n.T("en-US", "count", map[string]int{"Count": 1}) //=> 1 item
```

### Ordered Params

```go
I18n.AddTranslation(&i18n.Translation{Key: "ordered_params", Locale: "en-US", Value: "{{$1}} {{$2}} {{$1}}"})
I18n.T("en-US", "ordered_params", "string1", "string2") //=> string1 string2 string1
```

### Inline Edit

You could manage translations' data with [QOR Admin](http://GoTenancy/libs/admin) interface (UI) after registering it into [QOR Admin](http://GoTenancy/libs/admin), however we warn you that it is usually quite hard (and error prone!) to *translate a translation* without knowing its context...Fortunately, the *Inline Edit* feature of [QOR Admin](http://GoTenancy/libs/admin) was developed to resolve this problem!

*Inline Edit* allows administrators to manage translations from the frontend. Similarly to [integrating with Golang Templates](#integrate-with-golang-templates), you need to register a func map for Golang templates to render *inline editable* translations.

The good thing is we have created a package for you to do this easily, it will generate a `FuncMap`, you just need to use it when parsing your templates:

```go
// `I18n` hold translations backends
// `en-US` current locale
// `true` enable inline edit mode or not, if inline edit not enabled, it works just like the funcmap in section "Integrate with Golang Templates"
inline_edit.FuncMap(I18n, "en-US", true) // => map[string]interface{}{
                                         //     "t": func(string, ...interface{}) template.HTML {
                                         //        // ...
                                         //      },
                                         //    }
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
