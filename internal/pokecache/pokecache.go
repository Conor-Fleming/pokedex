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
	cache map[string]cacheEntry
	mux   *sync.Mutex
}

// use mutex to lock map
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       value,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	if v, ok := c.cache[key]; ok {
		return v.val, true
	}

	return nil, false
}

func (c *Cache) ReapLoop(interval time.Duration) {
	//declare ticker to send on the channel every ten minutes or so
	ticker := time.NewTicker(10 * time.Minute)

	c.mux.Lock()
	defer c.mux.Unlock()
	//remove entries to cache that are older than duration passed to NewCache()
	for range ticker.C {
		for k, v := range c.cache {
			if oldValue(v.createdAt, interval) {
				delete(c.cache, k)
			}
		}
	}
}

func NewCache(interval time.Duration) Cache {
	//Create new Cache
	cache := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}
	//spin off reap go routine
	go cache.ReapLoop(interval)

	return cache
}

func oldValue(t time.Time, duration time.Duration) bool {
	//get length of time since the value was created
	diff := time.Since(t)
	//check if difference is greater than given duration
	return diff > duration
}
