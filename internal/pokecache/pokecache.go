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

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		CacheMu:    sync.RWMutex{},
		CacheEntry: make(map[string]cacheEntry),
		Interval:   interval,
	}
	reapLoop(&cache)
	return &cache
}

func Add(key string, val []byte, cache *Cache) {
	cache.CacheMu.Lock()
	defer cache.CacheMu.Unlock()
	cache.CacheEntry[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func Get(key string, cache *Cache) ([]byte, bool) {
	cache.CacheMu.RLock()
	defer cache.CacheMu.RUnlock()
	val, ok := cache.CacheEntry[key]
	if ok {
		return val.val, ok
	}

	return nil, false
}

func reapLoop(cache *Cache) {
	clock := time.NewTicker(cache.Interval)
	defer clock.Stop()
	go func() {
		for {
			select {
			case <-clock.C:
				cache.CacheMu.Lock()

				currentTime := time.Now()
				for key, val := range cache.CacheEntry {
					if currentTime.Sub(val.createdAt) > cache.Interval {
						delete(cache.CacheEntry, key)
					}
				}

				cache.CacheMu.Unlock()
			}
		}
	}()
}
