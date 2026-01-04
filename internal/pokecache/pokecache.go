package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	cache    map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		cache:    make(map[string]cacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.mutex.Lock()
		for k, v := range c.cache {
			now := time.Now()
			cutoff := now.Add(-c.interval)
			if v.createdAt.Before(cutoff) {
				delete(c.cache, k)
			}
		}
		c.mutex.Unlock()
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.cache[key]
	if ok {
		return entry.val, true
	} else {
		return []byte{}, false
	}
}
