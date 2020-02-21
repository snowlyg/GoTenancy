# Qiniu

[Qiniu](https://www.qiniu.com) backend for [QOR OSS](https://GoTenancy/libs/oss)

## Usage

```go
import "GoTenancy/libs/oss/qiniu"

func main() {
  storage := qiniu.New(&qiniu.Config{
    AccessID:  "access_id",
    AccessKey: "access_key",
    Bucket:    "bucket",
    Region:    "huadong",
    Endpoint:  "https://up.qiniup.com",
  })

  // Save a reader interface into storage
  storage.Put("/sample.txt", reader)

  // Get file with path
  storage.Get("/sample.txt")

  // Get object as io.ReadCloser
  storage.GetStream("/sample.txt")

  // Delete file with path
  storage.Delete("/sample.txt")

  // List all objects under path
  storage.List("/")

  // Get Public Accessible URL (useful if current file saved privately)
  storage.GetURL("/sample.txt")
}
```

