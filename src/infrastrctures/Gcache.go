package infrastrctures

import (
	"github.com/bluele/gcache"
	"log"
)

type CacheClient struct {
	Cache gcache.Cache
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
