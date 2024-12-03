package config

import (
	"fmt"
	"github.com/bluele/gcache"
	"short-url-4go/src/infrastrctures"
	"sync"
)

/*func InitCache(cfg Config) gcache.Cache {

	// 创建gcache缓存
	cache := gcache.New(cfg.CacheMaxCap). // 设置缓存最大容量
						Expiration(cfg.CacheLiveTime). // 设置条目过期时间
						LRU().                         // 使用 LRU 策略
						Build()                        // 构建缓存
	return cache
}*/

var (
	gcacheObj  *GCacheHandler
	gcacheOnce sync.Once
)

// IGCacheHandler 是缓存接口
type IGCacheHandler interface {
	InitGCache() // 初始化缓存
}

// GCacheHandler 结构体，封装 gcache.Cache 实例
type GCacheHandler struct {
	Cache *infrastrctures.CacheClient
}

// InitGCache 初始化 gcache 缓存
// CacheMaxItems：缓存的最大容量
// CacheLifetime：缓存项的超时时间
func (g *GCacheHandler) InitGCache() {
	// 初始化缓存，LRU 策略
	gc := gcache.New(EnvVariables.CacheMaxItems).
		LRU().
		Expiration(EnvVariables.CacheLifetime). // 设置缓存过期时间
		Build()

	// 创建 CacheClient 实例并绑定到 GCacheHandler
	g.Cache = &infrastrctures.CacheClient{
		Cache: gc,
	}
	fmt.Printf("GCache Initialized with size: %d and expiration: %v\n", EnvVariables.CacheMaxItems, EnvVariables.CacheLifetime)
}

func Gcache() IGCacheHandler {
	if gcacheObj == nil {
		gcacheOnce.Do(func() {
			gcacheObj = &GCacheHandler{}
		})
	}
	return gcacheObj
}
