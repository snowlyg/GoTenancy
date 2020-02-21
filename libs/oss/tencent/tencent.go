package tencent

import (
	"GoTenancy/libs/oss"
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

var _ oss.StorageInterface = (*Client)(nil)

type Config struct {
	AppID     string
	AccessID  string
	AccessKey string
	Region    string
	Bucket    string
	ACL       string
	CORS      string
	Endpoint  string
}

type Client struct {
	Config *Config
	Client *http.Client
}

func New(conf *Config) *Client {
	return &Client{conf, &http.Client{}}
}

func (client Client) getUrl() string {
	return fmt.Sprintf("http://%s.cos.%s.myqcloud.com/", client.Config.Bucket, client.Config.Region)
}

func (client Client) Get(path string) (file *os.File, err error) {
	readCloser, err := client.GetStream(path)
	if err == nil {
		if file, err = ioutil.TempFile("/tmp", "tencent"); err == nil {
			defer readCloser.Close()
			_, err = io.Copy(file, readCloser)
			file.Seek(0, 0)
		}
	}
	return file, err
}

var urlRegexp = regexp.MustCompile(`(https?:)?//((\w+).)+(\w+)/`)

// ToRelativePath process path to relative path
func (client Client) ToRelativePath(urlPath string) string {
	if urlRegexp.MatchString(urlPath) {
		if u, err := url.Parse(urlPath); err == nil {
			return strings.TrimPrefix(u.Path, "/")
		}
	}

	return strings.TrimPrefix(urlPath, "/")
}

func (client Client) GetStream(path string) (io.ReadCloser, error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", client.getUrl(), client.ToRelativePath(path)))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("get file fail")
	}
	return resp.Body, nil
}

func (client Client) Put(path string, body io.Reader) (*oss.Object, error) {
	if seeker, ok := body.(io.ReadSeeker); ok {
		seeker.Seek(0, 0)
	}
	switch body.(type) {
	case *bytes.Buffer, *bytes.Reader, *strings.Reader:
	default:
		if body != nil {
			b, err := ioutil.ReadAll(body)
			if err != nil {
				return nil, err
			}
			body = bytes.NewReader(b)
		}
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("%s%s", client.getUrl(), client.ToRelativePath(path)), body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Host", client.GetEndpoint())
	req.Header.Set("Authorization", client.authorization(req))
	result, err := client.Client.Do(req)
	if err != nil {
		return nil, err
	}
	if result.StatusCode != http.StatusOK {
		d, err := ioutil.ReadAll(ioutil.NopCloser(result.Body))
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(d))
	}
	now := time.Now()
	return &oss.Object{
		Path:             path,
		Name:             filepath.Base(path),
		LastModified:     &now,
		StorageInterface: client,
	}, nil
}

func (client Client) Delete(path string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s%s", client.getUrl(), client.ToRelativePath(path)), nil)
	if err != nil {
		return err
	}
	req.Header.Set("Host", client.GetEndpoint())
	req.Header.Set("Authorization", client.authorization(req))
	result, err := client.Client.Do(req)
	if err != nil {
		return err
	}
	if result.StatusCode != http.StatusOK && result.StatusCode != http.StatusNoContent {
		d, err := ioutil.ReadAll(ioutil.NopCloser(result.Body))
		if err != nil {
			return err
		}
		return errors.New(string(d))
	}
	return nil
}

//todo not found api
func (client Client) List(path string) ([]*oss.Object, error) {
	var objects []*oss.Object

	results, err := client.Get(path)

	if err == nil {
		objects = append(objects, &oss.Object{
			Path: client.ToRelativePath(path),
			Name: results.Name(),
			//LastModified:     &obj.LastModified,
			StorageInterface: client,
		})

	}
	return objects, err
}

func (client Client) GetEndpoint() string {
	if client.Config.Endpoint != "" {
		return client.Config.Endpoint
	}
	return fmt.Sprintf("%s.cos.%s.myqcloud.com", client.Config.Bucket, client.Config.Region)
}

func (client Client) GetURL(path string) (string, error) {
	return fmt.Sprintf("%s%s", client.getUrl(), client.ToRelativePath(path)), nil
}

func (client Client) authorization(req *http.Request) string {
	signTime := getSignTime()
	signature := getSignature(client.Config.AccessKey, req, signTime)
	authStr := fmt.Sprintf("q-sign-algorithm=sha1&q-ak=%s&q-sign-time=%s&q-key-time=%s&q-header-list=%s&q-url-param-list=%s&q-signature=%s",
		client.Config.AccessID, signTime, signTime, getHeadKeys(req.Header), getParamsKeys(req.URL.RawQuery), signature)

	return authStr
}
