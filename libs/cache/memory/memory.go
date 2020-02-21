package memory

import (
	"encoding/json"
	"errors"
	"sync"
)

var ErrNotFound = errors.New("not found")

type Memory struct {
	values map[string][]byte
	mutex  *sync.RWMutex
}

func New() *Memory {
	return &Memory{values: map[string][]byte{}, mutex: &sync.RWMutex{}}
}

func (memory *Memory) Get(key string) (string, error) {
	memory.mutex.RLock()
	defer memory.mutex.RUnlock()

	if value, ok := memory.values[key]; ok {
		return string(value), nil
	}
	return "", ErrNotFound
}

func (memory *Memory) Unmarshal(key string, object interface{}) error {
	memory.mutex.RLock()
	defer memory.mutex.RUnlock()

	if value, ok := memory.values[key]; ok {
		return json.Unmarshal(value, object)
	}
	return ErrNotFound
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

func (memory *Memory) Set(key string, value interface{}) error {
	memory.mutex.Lock()
	defer memory.mutex.Unlock()

	memory.values[key] = convertToBytes(value)
	return nil
}

func (memory *Memory) Fetch(key string, fc func() interface{}) (string, error) {
	if str, err := memory.Get(key); err == nil {
		return str, nil
	}
	results := convertToBytes(fc())
	return string(results), memory.Set(key, results)
}

func (memory *Memory) Delete(key string) error {
	memory.mutex.Lock()
	defer memory.mutex.Unlock()

	delete(memory.values, key)
	return nil
}
