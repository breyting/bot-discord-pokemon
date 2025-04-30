package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cacheEntries map[string]CacheEntry
	mu           *sync.RWMutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cacheEntries: map[string]CacheEntry{},
		mu:           &sync.RWMutex{},
	}

	go cache.reapLoop(interval)

	return cache
}

func (cache *Cache) Add(key string, value []byte) {
	cache.mu.Lock()
	cache.cacheEntries[key] = CacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
	cache.mu.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	val, ok := cache.cacheEntries[key]
	return val.val, ok
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		cache.reap(time.Now().UTC(), interval)
	}
}

func (cache *Cache) reap(now time.Time, last time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for k, v := range cache.cacheEntries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(cache.cacheEntries, k)
		}
	}
}
