// Package cache s
package cache

import (
	"time"

	"github.com/hashicorp/golang-lru/v2/expirable"
)

// Cache s.
type Cache[K comparable, V any] interface {
	Add(key K, value V) bool
	Get(key K) (value V, ok bool)
	Keys() []K
}

// SetCache s.
func SetCache[K comparable, V any](ttl, size int) Cache[K, V] {
	cache := expirable.NewLRU[K, V](size, nil, time.Second*time.Duration(ttl))
	return cache
}
