package s3

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"GoTenancy/libs/oss"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/credentials/ec2rolecreds"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/ec2metadata"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Client S3 storage
type Client struct {
	*s3.S3
	Config *Config
}

// Config S3 client config
type Config struct {
	AccessID         string
	AccessKey        string
	Region           string
	Bucket           string
	SessionToken     string
	ACL              string
	Endpoint         string
	S3Endpoint       string
	S3ForcePathStyle bool
	CacheControl     string

	Session *session.Session

	RoleARN string
}

func ec2RoleAwsCreds(config *Config) *credentials.Credentials {
	ec2m := ec2metadata.New(session.New(), &aws.Config{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		Endpoint:   aws.String("http://169.254.169.254/latest"),
	})

	return credentials.NewCredentials(&ec2rolecreds.EC2RoleProvider{
		Client: ec2m,
	})
}

// Does this need to be publicly exported?
func EC2RoleAwsConfig(config *Config) *aws.Config {
	return &aws.Config{
		Region:      aws.String(config.Region),
		Credentials: ec2RoleAwsCreds(config),
	}
}

// New initialize S3 storage
func New(config *Config) *Client {
	if config.ACL == "" {
		config.ACL = s3.BucketCannedACLPublicRead
	}

	client := &Client{Config: config}

	if config.RoleARN != "" {
		sess := session.Must(session.NewSession())
		creds := stscreds.NewCredentials(sess, config.RoleARN)

		s3Config := &aws.Config{
			Region:           &config.Region,
			Endpoint:         &config.S3Endpoint,
			S3ForcePathStyle: &config.S3ForcePathStyle,
			Credentials:      creds,
		}

		client.S3 = s3.New(sess, s3Config)
		return client
	}

	s3Config := &aws.Config{
		Region:           &config.Region,
		Endpoint:         &config.S3Endpoint,
		S3ForcePathStyle: &config.S3ForcePathStyle,
	}

	if config.Session != nil {
		client.S3 = s3.New(config.Session, s3Config)
	} else if config.AccessID == "" && config.AccessKey == "" {
		// use aws default Credentials
		// s3Config.Credentials = ec2RoleAwsCreds(config)
		sess := session.Must(session.NewSession())
		client.S3 = s3.New(sess, s3Config)
	} else {
		creds := credentials.NewStaticCredentials(config.AccessID, config.AccessKey, config.SessionToken)
		if _, err := creds.Get(); err == nil {
			s3Config.Credentials = creds
			client.S3 = s3.New(session.New(), s3Config)
		}
	}

	return client
}

// Get receive file with given path
func (client Client) Get(path string) (file *os.File, err error) {
	readCloser, err := client.GetStream(path)

	ext := filepath.Ext(path)
	pattern := fmt.Sprintf("s3*%s", ext)

	if err == nil {
		if file, err = ioutil.TempFile("/tmp", pattern); err == nil {
			defer readCloser.Close()
			_, err = io.Copy(file, readCloser)
			file.Seek(0, 0)
		}
	}

	return file, err
}

// GetStream get file as stream
func (client Client) GetStream(path string) (io.ReadCloser, error) {
	getResponse, err := client.S3.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToRelativePath(path)),
	})

	return getResponse.Body, err
}

// Put store a reader into given path
func (client Client) Put(urlPath string, reader io.Reader) (*oss.Object, error) {
	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	urlPath = client.ToRelativePath(urlPath)
	buffer, err := ioutil.ReadAll(reader)

	fileType := mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	params := &s3.PutObjectInput{
		Bucket:        aws.String(client.Config.Bucket), // required
		Key:           aws.String(urlPath),              // required
		ACL:           aws.String(client.Config.ACL),
		Body:          bytes.NewReader(buffer),
		ContentLength: aws.Int64(int64(len(buffer))),
		ContentType:   aws.String(fileType),
	}
	if client.Config.CacheControl != "" {
		params.CacheControl = aws.String(client.Config.CacheControl)
	}

	_, err = client.S3.PutObject(params)

	now := time.Now()
	return &oss.Object{
		Path:             urlPath,
		Name:             filepath.Base(urlPath),
		LastModified:     &now,
		StorageInterface: client,
	}, err
}

// Delete delete file
func (client Client) Delete(path string) error {
	_, err := client.S3.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(client.Config.Bucket),
		Key:    aws.String(client.ToRelativePath(path)),
	})
	return err
}

// List list all objects under current path
func (client Client) List(path string) ([]*oss.Object, error) {
	var objects []*oss.Object
	var prefix string

	if path != "" {
		prefix = strings.Trim(path, "/") + "/"
	}

	listObjectsResponse, err := client.S3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: aws.String(client.Config.Bucket),
		Prefix: aws.String(prefix),
	})

	if err == nil {
		for _, content := range listObjectsResponse.Contents {
			objects = append(objects, &oss.Object{
				Path:             client.ToRelativePath(*content.Key),
				Name:             filepath.Base(*content.Key),
				LastModified:     content.LastModified,
				StorageInterface: client,
			})
		}
	}

	return objects, err
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (client Client) GetEndpoint() string {
	if client.Config.Endpoint != "" {
		return client.Config.Endpoint
	}

	endpoint := client.S3.Endpoint
	for _, prefix := range []string{"https://", "http://"} {
		endpoint = strings.TrimPrefix(endpoint, prefix)
	}

	return client.Config.Bucket + "." + endpoint
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

// ToRelativePath process path to relative path
func (client Client) ToRelativePath(urlPath string) string {
	if urlRegexp.MatchString(urlPath) {
		if u, err := url.Parse(urlPath); err == nil {
			if client.Config.S3ForcePathStyle { // First part of path will be bucket name
				return strings.TrimPrefix(u.Path, "/"+client.Config.Bucket)
			}
			return u.Path
		}
	}

	if client.Config.S3ForcePathStyle { // First part of path will be bucket name
		return "/" + strings.TrimPrefix(urlPath, "/"+client.Config.Bucket+"/")
	}
	return "/" + strings.TrimPrefix(urlPath, "/")
}

// GetURL get public accessible URL
func (client Client) GetURL(path string) (url string, err error) {
	if client.Endpoint == "" {
		if client.Config.ACL == s3.BucketCannedACLPrivate || client.Config.ACL == s3.BucketCannedACLAuthenticatedRead {
			getResponse, _ := client.S3.GetObjectRequest(&s3.GetObjectInput{
				Bucket: aws.String(client.Config.Bucket),
				Key:    aws.String(client.ToRelativePath(path)),
			})

			return getResponse.Presign(1 * time.Hour)
		}
	}

	return path, nil
}
