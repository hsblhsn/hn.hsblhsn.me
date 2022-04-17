package caches

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

const (
	DefaultCacheExpiration = 10 * time.Minute
)

var _ Cache = (*bigcache.BigCache)(nil)

// NewInMemoryCache creates a new in-memory cache.
// It uses bigcache as a backend.
// It uses default cache expiration and max cache size.
func NewInMemoryCache() *bigcache.BigCache {
	// these constants are here because they are bigcache specific.
	const (
		MaxShard              = 64
		DefaultMaxCacheSizeMB = 800
	)
	cfg := bigcache.DefaultConfig(DefaultCacheExpiration)
	cfg.Shards = MaxShard
	cfg.HardMaxCacheSize = DefaultMaxCacheSizeMB
	cache, err := bigcache.NewBigCache(cfg)
	if err != nil {
		panic(err)
	}
	return cache
}
