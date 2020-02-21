package tencent

import (
	"GoTenancy/libs/oss/tests"
	"bytes"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestClient_Get(t *testing.T) {

}

var client *Client

func init() {
	client = New(&Config{
		AppID:     "1252882253",
		AccessID:  "AKIDToxukQWBG8nGXcBN8i662nOo12sc5Wjl",
		AccessKey: "40jNrBf5mLiuuiU8HH7lDTXP5at00sbA",
		Bucket:    "tets-1252882253",
		Region:    "ap-shanghai",
		ACL:       "public-read", // private，public-read-write，public-read；默认值：private
		//Endpoint:  config.Public.Endpoint,
	})
}

func TestClient_Put(t *testing.T) {
	f, err := ioutil.ReadFile("/home/owen/Downloads/2.png")
	if err != nil {
		t.Error(err)
		return
	}

	client.Put("test.png", bytes.NewReader(f))
}

func TestClient_Put2(t *testing.T) {
	tests.TestAll(client, t)
}

func TestClient_Delete(t *testing.T) {
	fmt.Println(client.Delete("test.png"))
}
