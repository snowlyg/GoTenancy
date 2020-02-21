package amazonpay

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"
	"time"
)

// APIError Amazon Pay API Error definition
type APIError struct {
	XMLName xml.Name `xml:"ErrorResponse"`
	Type    string   `xml:"Error>Type"`
	Code    string   `xml:"Error>Code"`
	Message string   `xml:"Error>Message"`
}

func (apiError APIError) Error() string {
	return apiError.Message
}

// Post post API info
func (amazonPay *AmazonPay) Post(params Params, response interface{}) error {
	if _, ok := params.Get("AWSAccessKeyId"); !ok {
		params.Set("AWSAccessKeyId", amazonPay.Config.AccessKey)
	}

	if _, ok := params.Get("SellerId"); !ok {
		params.Set("SellerId", amazonPay.Config.MerchantID)
	}

	if _, ok := params.Get("SignatureMethod"); !ok {
		params.Set("SignatureMethod", "HmacSHA256")
	}

	if _, ok := params.Get("SignatureVersion"); !ok {
		params.Set("SignatureVersion", "2")
	}

	if _, ok := params.Get("Timestamp"); !ok {
		params.Set("Timestamp", time.Now().UTC().Format("2006-01-02T15:04:05Z"))
	}

	params.Set("Signature", params.Sign())

	if _, ok := params.Get("Version"); !ok {
		params.Set("Version", "2013-01-01")
	}

	URL := url.URL{
		Scheme: "https",
		Host:   amazonPay.Config.Endpoint,
		Path:   path.Join(amazonPay.Config.ModePath, amazonPay.Config.APIVersion),
	}

	resp, err := http.Post(URL.String(), "application/x-www-form-urlencoded", strings.NewReader(amazonPay.buildPostURL(params)))

	var data []byte
	if err == nil {
		defer resp.Body.Close()
		data, err = ioutil.ReadAll(resp.Body)
		if resp.StatusCode == 200 {
			if response != nil {
				err = xml.Unmarshal([]byte(data), response)
			}
		} else {
			var apiError APIError
			if err = xml.Unmarshal([]byte(data), &apiError); err == nil {
				return apiError
			}
			return errors.New("Get error: " + string(data))
		}
	}

	return err
}

// buildPostURL build post URL
func (amazonPay *AmazonPay) buildPostURL(params Params) string {
	apiParams := []string{}

	for key, value := range params {
		if str := fmt.Sprint(value); str != "" {
			apiParams = append(apiParams, key+"="+url.QueryEscape(url.PathEscape(str)))
		}
	}

	sort.Strings(apiParams)
	postURL := strings.Join(apiParams, "&")
	postURL += "&Signature=" + amazonPay.Sign(strings.Join([]string{"POST", amazonPay.Config.Endpoint, fmt.Sprintf("/%v/%v", amazonPay.Config.ModePath, amazonPay.Config.APIVersion), postURL}, "\n"))
	return postURL
}

// Sign sign messages
func (amazonPay *AmazonPay) Sign(message string) string {
	key := []byte(amazonPay.Config.SecretKey)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return url.QueryEscape(base64.StdEncoding.EncodeToString(h.Sum(nil)))
}
