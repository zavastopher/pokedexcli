package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	CacheMu    sync.RWMutex
	CacheEntry map[string]cacheEntry
	Interval   time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	return Cache{
		CacheMu:    sync.RWMutex{},
		CacheEntry: make(map[string]cacheEntry),
		Interval:   interval,
	}
}
func Add(key string, val []byte, cache *Cache) {
	cache.CacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func Get(key string, cache *Cache) ([]byte, bool) {

	val, ok := cache.CacheEntry[key]
	if ok {
		return val.val, ok
	}

	return nil, ok
}

func ReapLoop(cache *Cache) {
	clock := time.NewTicker(cache.Interval)

}
