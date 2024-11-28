package config

import (
	"fmt"
	"github.com/bluele/gcache"
	"sync"
)

//func InitCache(cfg Config) gcache.Cache {
//
//	// 创建gcache缓存
//	cache := gcache.New(cfg.CacheMaxCap). // 设置缓存最大容量
//						Expiration(cfg.CacheLiveTime). // 设置条目过期时间
//						LRU().                         // 使用 LRU 策略
//						Build()                        // 构建缓存
//	return cache
//}

var (
	gcacheObj  *GCacheHandler
	gcacheOnce sync.Once
)

// IGCacheHandler 是缓存接口
type IGCacheHandler interface {
	InitGCache() // 初始化缓存
	//GetCache() gcache.Cache                    // 获取缓存实例
}

// GCacheHandler 结构体，封装 gcache.Cache 实例
type GCacheHandler struct {
	Cache gcache.Cache
}

// InitGCache 初始化 gcache 缓存
// size：缓存的最大容量
// expire：缓存项的超时时间
func (g *GCacheHandler) InitGCache() {
	// 初始化缓存，LRU 策略
	g.Cache = gcache.New(EnvVariables.CacheMaxItems).
		LRU().
		Expiration(EnvVariables.CacheLifetime). // 设置缓存过期时间
		Build()
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
