package JCache

import (
	"JCache/cacheItem"
	"sync"
)

type cache struct {
	mu         sync.Mutex
	lru        *cacheItem.LruCache
	cacheBytes int64
}

func (c *cache) add(key string, value Byteview) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = cacheItem.NewLruCache(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value Byteview, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}
	if v, ok := c.lru.Get(key); ok {
		return v.(Byteview), ok
	}
	return
}
