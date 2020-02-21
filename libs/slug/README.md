# Slug

Slug provides an easy way to create a [pretty URL](https://en.wikipedia.org/wiki/Semantic_URL#Slug), a.k.a. [Search Engine Friendly (SEF) URL](https://en.wikipedia.org/wiki/Semantic_URL#Slug), for your model.

[![GoDoc](https://godoc.org/GoTenancy/libs/slug?status.svg)](https://godoc.org/GoTenancy/libs/slug)

## Usage

Use `slug.Slug` as your field type with the same name as the benefactor field, from which the slug's value should be dynamically derived, and prepended with `WithSlug`, for example:

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/slug"
)

type User struct {
  gorm.Model
  Name            string
  NameWithSlug    slug.Slug
}
```

Then in the [QOR Admin](https://GoTenancy/libs/admin) interface, you will see a slug field.

## License

Released under the MIT License.
