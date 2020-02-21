package cache

type CacheStoreInterface interface {
	Get(key string) (string, error)
	Unmarshal(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
}
