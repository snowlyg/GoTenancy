# AWS S3

[AWS S3](https://aws.amazon.com/cn/s3/) backend for [QOR OSS](https://GoTenancy/libs/oss)

## Usage

```go
import "GoTenancy/libs/oss/aliyun"

func main() {
  storage := s3.New(s3.Config{
    AccessID: "access_id",
    AccessKey: "access_key",
    Region: "region",
    Bucket: "bucket",
    Endpoint: "cdn.getqor.com",
    ACL: awss3.BucketCannedACLPublicRead,
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


