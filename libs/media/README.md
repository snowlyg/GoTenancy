## Media Library

Media is a [Golang](http://golang.org/) library that supports the upload of *files*/*images*/*videos* to a filesystem or cloud storage as well as *linked videos* (i.e. YouTube, Vimeo, etc.). The plugin includes:

- cropping and resizing features for images.
- optional multiple sizes for each media resource.
- Accessibility helpers.

[![GoDoc](https://godoc.org/GoTenancy/libs/media?status.svg)](https://godoc.org/GoTenancy/libs/media)

###### File Types

Media accepts any and every file type, yet it associates certain file types as *images* or *videos* so as to provide helpers supporting those media's specific needs.


    Images: .jpg, .jpeg, .png, .tif, .tiff, .bmp, .gif

    Videos: .mp4, .m4p, .m4v, .m4v, .mov, .mpeg, .webm, .avi, .ogg, .ogv


## Usage

Media depends on [GORM](https://github.com/jinzhu/gorm) models as it is using [GORM](https://github.com/jinzhu/gorm)'s callbacks to handle file processing, so you will need to register callbacks first:

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/media"
)

DB, err = gorm.Open("sqlite3", "demo_db") // [gorm](https://github.com/jinzhu/gorm)

media.RegisterCallbacks(DB)
```

Then add [OSS(Object Storage Service)](https://GoTenancy/libs/oss) to your model:

```go
import (
  "github.com/jinzhu/gorm"
  "GoTenancy/libs/media/oss"
)

type Product struct {
  gorm.Model
  Image oss.OSS
}
```

Last, configure the storage. The default value is `oss.Storage := filesystem.New("public")`. Here we configure S3 as storage.

```go
import (
  // "github.com/oss/filesystem"
  "github.com/oss/s3"
)

oss.Storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: aws.BucketCannedACLPublicRead})
// Default configuration `oss.Storage := filesystem.New("public")`
```

## Operate stored files

The [OSS(Object Storage Service)](https://GoTenancy/libs/oss) provides a pretty simple API to operate files on filesytem or cloud storage

```go
type StorageInterface interface {
  Get(path string) (*os.File, error)
  GetStream(path string) (io.ReadCloser, error)
  Put(path string, reader io.Reader) (*Object, error)
  Delete(path string) error
  List(path string) ([]*Object, error)
  GetEndpoint() string
  GetURL(path string) (string, error)
}
```

So once you finished the setting, you could operate saved files like this:

```go
storage := s3.New(s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: aws.BucketCannedACLPublicRead})
// storage := filesystem.New("public")

// Save a reader interface into storage
storage.Put("/sample.txt", reader)

// Get file with path
storage.Get("/sample.txt")

// Delete file with path
storage.Delete("/sample.txt")

// List all objects under path
storage.List("/")
```

### Predefine common image size

You can implement the `GetSizes` function to predefine image sizes. The size name can be used to fetch image of corresponding size.

```go
import (
  "GoTenancy/libs/media/oss"
  "github.com/jinzhu/gorm"
)

type Product struct {
  gorm.Model
  Image ProductIconImageStorage
}

type ProductIconImageStorage struct{
  oss.OSS
}

func (ProductIconImageStorage) GetSizes() map[string]*media.Size {
  return map[string]*media.Size{
    // Add padding to thumbnail if ratio doesn't match, by default, crop center
    "small":    {Width: 60 * 2, Height: 60 * 2, Padding: true},
    "small@ld": {Width: 60, Height: 60},

    "middle":    {Width: 108 * 2, Height: 108 * 2},
    "middle@ld": {Width: 108, Height: 108},

    "big":    {Width: 144 * 2, Height: 144 * 2},
    "big@ld": {Width: 144, Height: 144},
  }
}

// Get image's url with style
product.Image.URL("small")
product.Image.URL("big@ld")
```

### Set file storage path in the file system

The default size and path for `Image`. The default size is `4294967295` and default path is `{repo_path}/public/system/{{class}}/{{primary_key}}/{{column}}.{{extension}}`.

You can set the path and size manually by adding tag to the field like this:

```go
type Product struct {
  gorm.Model
  Image oss.OSS `sql:"size:4294967295;" media_library:"url:/backend/{{class}}/{{primary_key}}/{{column}}.{{extension}};path:./private"`
}
```

The `media` takes two parameters, `url` and `path`. the `url` set the relative file path and the `path` set the prefix path. So suppose we uploaded a image called `demo.png`. The file will be stored at `{repo_path}/private/backend/products/1/demo.png`.

#### Be careful when using `http.FileServer`

The `http.FileServer` not only serves the files, but also shows the directory contents. This is very dangerous when you do things like

```go
for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
  mux.Handle(fmt.Sprintf("/%s/", path), http.FileServer(http.Dir("public")))
}
```

The files under `public` will be exposed to public(especially the search engine!), Imagine someone upload a illegal or sensitive file to your server. The directory is fully visible to everyone and its indexable by search engines and boom!

To avoid this problem, we made a safer `FileServer` function [here](https://GoTenancy/libs/qor/blob/master/utils/utils.go#L176). This function serves file only. So the previous code now turned into:

```go
for _, path := range []string{"system", "javascripts", "stylesheets", "images"} {
  mux.Handle(fmt.Sprintf("/%s/", path), utils.FileServer(http.Dir("public")))
}
```

## Accessibility helpers

Media Library has some features aimed at helping achieve Accessibile frontends:

- capture of a textual description for *images*, *videos*, and *linked videos* to aid with Accessibility.
- capture of textual transcript for *videos* and *linked videos* to aid with Accessibility.

The values captured are fed into the sub-templates for each media type to be used if/where necessary. For example, an *image*'s HTML output (an `img` tag) manifests the textual description within an `alt` attribute while a video's HTML (an `iframe` tag) manifests the textual description within a `title` attribute.

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
