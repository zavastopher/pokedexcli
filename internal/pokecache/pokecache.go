package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMu     sync.RWMutex
	CacheEntrie map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}
