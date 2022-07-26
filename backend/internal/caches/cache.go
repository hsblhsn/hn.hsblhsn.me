package caches

// Cache is a minimal interface for caching expensive results.
type Cache interface {
	Get(key string) ([]byte, error)
	Set(key string, val []byte) error
}
