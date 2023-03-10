package cache

import (
	"encoding/json"
	"os"
	"path/filepath"

	cache "github.com/gregjones/httpcache/diskcache"
)

type Cache struct {
	client *cache.Cache
}

type CacheInterface interface {
	Set(key string, data interface{}) error
	Get(key string) (interface{}, bool)
}

// Initializes a new cache
func NewCache() CacheInterface {
	homeDir, _ := os.UserHomeDir()
	newCache := &Cache{
		client: cache.New(filepath.Join(homeDir, ".gogitignore")),
	}

	return newCache
}

func (c *Cache) Set(key string, data interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.client.Set(key, b)
	return nil
}

func (c *Cache) Get(key string) (interface{}, bool) {
	data, found := c.client.Get(key)
	if !found {
		return nil, false
	}
	var resp interface{}
	_ = json.Unmarshal(data, &resp)
	return resp, true
}
