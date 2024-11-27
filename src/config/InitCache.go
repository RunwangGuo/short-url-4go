package config

import (
	"github.com/bluele/gcache"
)

func InitCache(cfg Config) gcache.Cache {

	// 创建gcache缓存
	cache := gcache.New(cfg.CacheMaxCap). // 设置缓存最大容量
						Expiration(cfg.CacheLiveTime). // 设置条目过期时间
						LRU().                         // 使用 LRU 策略
						Build()                        // 构建缓存
	return cache
}
