package oss

import (
	"bytes"
	"io"
	"io/ioutil"
	"strings"

	"GoTenancy/libs/media"
	"GoTenancy/libs/oss"
	"GoTenancy/libs/oss/filesystem"
	"GoTenancy/libs/qor/utils"
)

var (
	// URLTemplate default URL template
	URLTemplate = "/system/{{class}}/{{primary_key}}/{{column}}/{{filename_with_hash}}"
	// Storage the storage used to save medias
	Storage oss.StorageInterface = filesystem.New("public")
	_       media.Media          = &OSS{}
)

// OSS common storage interface
type OSS struct {
	media.Base
}

// DefaultURLTemplateHandler used to generate URL and save into database
var DefaultURLTemplateHandler = func(oss OSS, option *media.Option) (url string) {
	if url = option.Get("URL"); url == "" {
		url = URLTemplate
	}

	url = strings.Join([]string{strings.TrimSuffix(Storage.GetEndpoint(), "/"), strings.TrimPrefix(url, "/")}, "/")
	if strings.HasPrefix(url, "/") {
		return url
	}

	for _, prefix := range []string{"https://", "http://"} {
		url = strings.TrimPrefix(url, prefix)
	}

	// convert `getqor.com/hello` => `//getqor.com/hello`
	return "//" + url
}

// GetURLTemplate URL's template
func (o OSS) GetURLTemplate(option *media.Option) (url string) {
	return DefaultURLTemplateHandler(o, option)
}

// DefaultStoreHandler used to store reader with default Storage
var DefaultStoreHandler = func(oss OSS, path string, option *media.Option, reader io.Reader) error {
	_, err := Storage.Put(path, reader)
	return err
}

// Store save reader's content with path
func (o OSS) Store(path string, option *media.Option, reader io.Reader) error {
	return DefaultStoreHandler(o, path, option, reader)
}

// DefaultRetrieveHandler used to retrieve file
var DefaultRetrieveHandler = func(oss OSS, path string) (media.FileInterface, error) {
	result, err := Storage.GetStream(path)
	if f, ok := result.(media.FileInterface); ok {
		return f, err
	}

	if err == nil {
		buf := []byte{}
		if buf, err = ioutil.ReadAll(result); err == nil {
			result := utils.ClosingReadSeeker{bytes.NewReader(buf)}
			result.Seek(0, 0)
			return result, err
		}
	}
	return nil, err
}

// Retrieve retrieve file content with url
func (o OSS) Retrieve(path string) (media.FileInterface, error) {
	return DefaultRetrieveHandler(o, path)
}

// URL return file's url with given style
func (o OSS) URL(styles ...string) string {
	url := o.Base.URL(styles...)

	newurl, err := Storage.GetURL(url)
	if err != nil || len(newurl) == 0 {
		return url
	}

	return newurl
}

func (o OSS) String() string {
	url := o.Base.URL()

	newurl, err := Storage.GetURL(url)
	if err != nil || len(newurl) == 0 {
		return url
	}

	return newurl
}
