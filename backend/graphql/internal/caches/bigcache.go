package caches

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

const (
	DefaultCacheExpiration = 10 * time.Minute
	DefaultMaxCacheSizeMB  = 800
)

var _ Cache = (*bigcache.BigCache)(nil)

// NewInMemoryCache creates a new in-memory cache.
// It uses bigcache as a backend.
// It uses default cache expiration and max cache size.
func NewInMemoryCache() *bigcache.BigCache {
	const maxShard = 64
	cfg := bigcache.DefaultConfig(DefaultCacheExpiration)
	cfg.Shards = maxShard
	cfg.HardMaxCacheSize = DefaultMaxCacheSizeMB
	cfg.MaxEntrySize = DefaultMaxEntrySize
	cache, err := bigcache.NewBigCache(cfg)
	if err != nil {
		panic(err)
	}
	return cache
}
