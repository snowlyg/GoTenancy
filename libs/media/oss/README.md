# Media Library OSS

Use [OSS](https://GoTenancy/libs/oss) as backend to store medias

# Usage

```go
import (
	"GoTenancy/libs/media/oss"
	"GoTenancy/libs/oss/filesystem"
	"GoTenancy/libs/oss/s3"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
)

type Product struct {
	gorm.Model
	Image oss.OSS
}

func init() {
  // OSS's default storage is directory `public`, change it to S3
	oss.Storage = s3.New(&s3.Config{AccessID: "access_id", AccessKey: "access_key", Region: "region", Bucket: "bucket", Endpoint: "cdn.getqor.com", ACL: awss3.BucketCannedACLPublicRead})

  // or change directory to `download`
	oss.Storage = filesystem.New("download")
}
```

# Advanced Usage

```go
// change URL template
oss.URLTemplate = "/system/{{class}}/{{primary_key}}/{{column}}/{{filename_with_hash}}"

// change default URL handler
oss.DefaultURLTemplateHandler = func(option *media_library.Option) (url string) {
  // ...
}

// change default save handler
oss.DefaultStoreHandler = func(path string, option *media_library.Option, reader io.Reader) error {
  // ...
}

// change default retrieve handler
oss.DefaultRetrieveHandler = func(path string) (*os.File, error) {
	// ...
}

// By overwritting default store, retrieve handler, you could do some advanced tasks, like use private mode when store sensitive data to S3, public read mode for other files
```

## License

Released under the [MIT License](http://opensource.org/licenses/MIT).
