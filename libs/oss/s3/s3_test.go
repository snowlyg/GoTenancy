package s3_test

import (
	"fmt"
	"testing"

	"GoTenancy/libs/oss/s3"
	"GoTenancy/libs/oss/tests"
	awss3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/jinzhu/configor"
)

type Config struct {
	AccessID  string `env:"QOR_AWS_ACCESS_KEY_ID"`
	AccessKey string `env:"QOR_AWS_SECRET_ACCESS_KEY"`
	Region    string `env:"QOR_AWS_REGION"`
	Bucket    string `env:"QOR_AWS_BUCKET"`
	Endpoint  string `env:"QOR_AWS_ENDPOINT"`
}

var (
	client *s3.Client
	config = Config{}
)

func init() {
	configor.Load(&config)

	client = s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, Endpoint: config.Endpoint})
}

func TestAll(t *testing.T) {
	fmt.Println("testing S3 with public ACL")
	tests.TestAll(client, t)

	fmt.Println("testing S3 with private ACL")
	privateClient := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, ACL: awss3.BucketCannedACLPrivate, Endpoint: config.Endpoint})
	tests.TestAll(privateClient, t)

	fmt.Println("testing S3 with AuthenticatedRead ACL")
	authenticatedReadClient := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: config.Bucket, ACL: awss3.BucketCannedACLAuthenticatedRead, Endpoint: config.Endpoint})
	tests.TestAll(authenticatedReadClient, t)
}

func TestToRelativePath(t *testing.T) {
	urlMap := map[string]string{
		"https://mybucket.s3.amazonaws.com/myobject.ext": "/myobject.ext",
		"https://qor-example.com/myobject.ext":           "/myobject.ext",
		"//mybucket.s3.amazonaws.com/myobject.ext":       "/myobject.ext",
		"http://mybucket.s3.amazonaws.com/myobject.ext":  "/myobject.ext",
		"myobject.ext": "/myobject.ext",
	}

	for url, path := range urlMap {
		if client.ToRelativePath(url) != path {
			t.Errorf("%v's relative path should be %v, but got %v", url, path, client.ToRelativePath(url))
		}
	}
}

func TestToRelativePathWithS3ForcePathStyle(t *testing.T) {
	urlMap := map[string]string{
		"https://s3.amazonaws.com/mybucket/myobject.ext": "/myobject.ext",
		"https://qor-example.com/myobject.ext":           "/myobject.ext",
		"//s3.amazonaws.com/mybucket/myobject.ext":       "/myobject.ext",
		"http://s3.amazonaws.com/mybucket/myobject.ext":  "/myobject.ext",
		"/mybucket/myobject.ext":                         "/myobject.ext",
		"myobject.ext":                                   "/myobject.ext",
	}

	client := s3.New(&s3.Config{AccessID: config.AccessID, AccessKey: config.AccessKey, Region: config.Region, Bucket: "mybucket", S3ForcePathStyle: true, Endpoint: config.Endpoint})

	for url, path := range urlMap {
		if client.ToRelativePath(url) != path {
			t.Errorf("%v's relative path should be %v, but got %v", url, path, client.ToRelativePath(url))
		}
	}
}
