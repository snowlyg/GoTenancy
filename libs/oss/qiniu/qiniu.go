package qiniu

import (
	"bytes"
	"context"
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
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

// Client Qiniu storage
type Client struct {
	Config        *Config
	mac           *qbox.Mac
	storageCfg    storage.Config
	bucketManager *storage.BucketManager
	putPolicy     *storage.PutPolicy
}

// Config Qiniu client config
type Config struct {
	AccessID      string
	AccessKey     string
	Region        string
	Bucket        string
	Endpoint      string
	UseHTTPS      bool
	UseCdnDomains bool
	PrivateURL    bool
}

var zonedata = map[string]*storage.Zone{
	"huadong": &storage.ZoneHuadong,
	"huabei":  &storage.ZoneHuabei,
	"huanan":  &storage.ZoneHuanan,
	"beimei":  &storage.ZoneBeimei,
}

func New(config *Config) *Client {

	client := &Client{Config: config, storageCfg: storage.Config{}}

	client.mac = qbox.NewMac(config.AccessID, config.AccessKey)

	if z, ok := zonedata[strings.ToLower(config.Region)]; ok {
		client.storageCfg.Zone = z
	} else {
		panic(fmt.Sprintf("Zone %s is invalid, only support huadong, huabei, huanan, beimei.", config.Region))
	}
	if len(config.Endpoint) == 0 {
		panic("endpoint must be provided.")
	}
	client.storageCfg.UseHTTPS = config.UseHTTPS
	client.storageCfg.UseCdnDomains = config.UseCdnDomains
	client.bucketManager = storage.NewBucketManager(client.mac, &client.storageCfg)

	return client
}

func (client Client) SetPutPolicy(putPolicy *storage.PutPolicy) {
	client.putPolicy = putPolicy
}

// Get receive file with given path
func (client Client) Get(path string) (file *os.File, err error) {
	readCloser, err := client.GetStream(path)

	if file, err = ioutil.TempFile("/tmp", "qiniu"); err == nil {
		defer readCloser.Close()
		_, err = io.Copy(file, readCloser)
		file.Seek(0, 0)
	}

	return file, err
}

// GetStream get file as stream
func (client Client) GetStream(path string) (io.ReadCloser, error) {
	purl, err := client.GetURL(path)
	if err != nil {
		return nil, err
	}

	var res *http.Response
	res, err = http.Get(purl)
	if err == nil && res.StatusCode != http.StatusOK {
		err = fmt.Errorf("file %s not found", path)
	}

	return res.Body, err
}

// Put store a reader into given path
func (client Client) Put(urlPath string, reader io.Reader) (r *oss.Object, err error) {
	if seeker, ok := reader.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}

	urlPath = storageKey(urlPath)
	var buffer []byte
	buffer, err = ioutil.ReadAll(reader)
	if err != nil {
		return
	}

	fileType := mime.TypeByExtension(path.Ext(urlPath))
	if fileType == "" {
		fileType = http.DetectContentType(buffer)
	}

	putPolicy := storage.PutPolicy{
		Scope: fmt.Sprintf("%s:%s", client.Config.Bucket, urlPath),
	}

	if client.putPolicy != nil {
		putPolicy = *client.putPolicy
	}

	upToken := putPolicy.UploadToken(client.mac)

	formUploader := storage.NewFormUploader(&client.storageCfg)
	ret := storage.PutRet{}
	dataLen := int64(len(buffer))

	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	err = formUploader.Put(context.Background(), &ret, upToken, urlPath, bytes.NewReader(buffer), dataLen, &putExtra)
	if err != nil {
		return
	}

	now := time.Now()
	return &oss.Object{
		Path:             ret.Key,
		Name:             filepath.Base(urlPath),
		LastModified:     &now,
		StorageInterface: client,
	}, err
}

// Delete delete file
func (client Client) Delete(path string) error {
	return client.bucketManager.Delete(client.Config.Bucket, storageKey(path))
}

// List list all objects under current path
func (client Client) List(path string) (objects []*oss.Object, err error) {
	var prefix = storageKey(path)
	var listItems []storage.ListItem
	listItems, _, _, _, err = client.bucketManager.ListFiles(
		client.Config.Bucket,
		prefix,
		"",
		"",
		100,
	)

	if err != nil {
		return
	}

	for _, content := range listItems {
		t := time.Unix(content.PutTime, 0)
		objects = append(objects, &oss.Object{
			Path:             "/" + storageKey(content.Key),
			Name:             filepath.Base(content.Key),
			LastModified:     &t,
			StorageInterface: client,
		})
	}

	return
}

// GetEndpoint get endpoint, FileSystem's endpoint is /
func (client Client) GetEndpoint() string {
	return client.Config.Endpoint
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

func storageKey(urlPath string) string {
	if urlRegexp.MatchString(urlPath) {
		if u, err := url.Parse(urlPath); err == nil {
			urlPath = u.Path
		}
	}
	return strings.TrimPrefix(urlPath, "/")
}

// GetURL get public accessible URL
func (client Client) GetURL(path string) (url string, err error) {
	if len(path) == 0 {
		return
	}
	key := storageKey(path)

	if client.Config.PrivateURL {
		deadline := time.Now().Add(time.Second * 3600).Unix()
		url = storage.MakePrivateURL(client.mac, client.Config.Endpoint, key, deadline)
		return
	}

	url = storage.MakePublicURL(client.GetEndpoint(), key)

	return
}
