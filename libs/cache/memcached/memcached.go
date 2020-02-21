package memcached

import (
	"encoding/json"

	"github.com/bradfitz/gomemcache/memcache"
)

type Memcached struct {
	Config *Config
	Client *memcache.Client
}

type Config struct {
	NameSpace string
	Hosts     []string
}

func New(config *Config) *Memcached {
	client := memcache.New(config.Hosts...)
	return &Memcached{Config: config, Client: client}
}

func (memcached *Memcached) KeyWithNameSpance(key string) string {
	if memcached.Config.NameSpace != "" {
		key = memcached.Config.NameSpace + ":" + key
	}
	return key
}

func (memcached *Memcached) Get(key string) (string, error) {
	if item, err := memcached.Client.Get(memcached.KeyWithNameSpance(key)); err == nil {
		return string(item.Value), nil
	} else {
		return "", err
	}
}

func (memcached *Memcached) Unmarshal(key string, object interface{}) error {
	item, err := memcached.Client.Get(memcached.KeyWithNameSpance(key))
	if err == nil {
		err = json.Unmarshal(item.Value, object)
	}
	return err
}

func convertToBytes(value interface{}) []byte {
	switch result := value.(type) {
	case string:
		return []byte(result)
	case []byte:
		return result
	default:
		bytes, _ := json.Marshal(value)
		return bytes
	}
}

func (memcached *Memcached) Set(key string, value interface{}) error {
	return memcached.Client.Set(&memcache.Item{Key: memcached.KeyWithNameSpance(key), Value: convertToBytes(value)})
}

func (memcached *Memcached) Fetch(key string, fc func() interface{}) (string, error) {
	if str, err := memcached.Get(key); err == nil {
		return str, nil
	}
	results := convertToBytes(fc())
	return string(results), memcached.Set(key, results)
}

func (memcached *Memcached) Delete(key string) error {
	return memcached.Client.Delete(memcached.KeyWithNameSpance(key))
}
