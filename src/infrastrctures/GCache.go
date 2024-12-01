package infrastrctures

import (
	"github.com/bluele/gcache"
	"log"
)

type CacheClient struct {
	Cache gcache.Cache
}

// Get 读取缓存
func (c *CacheClient) Get(key string) (*string, error) {
	value, err := c.Cache.Get(key)
	if err != nil {
		return nil, err
	}
	return value.(*string), nil
}

// Set 写入缓存
func (c *CacheClient) Set(key string, value string) error {
	return c.Cache.Set(key, value)
}

// Remove 批量删除缓存中的键
func (c *CacheClient) Remove(keys []string) error {
	for _, key := range keys {
		err := c.Cache.Remove(key)
		if !err {
			log.Printf("[CacheClient] Remove key: %s error: %v]", key, err)
		}
	}
	return nil
}
