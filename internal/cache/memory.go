package cache

import (
	"sync"
	"time"
)

type item struct {
	value      any
	expiration int64
}

type Cache struct {
	data map[string]item
	mu   sync.RWMutex
}

func New() *Cache {
	return &Cache{
		data: make(map[string]item),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.RLock()
	val, ok := c.data[key]
	c.mu.RUnlock()

	if !ok {
		return nil, false
	}

	if time.Now().UnixNano() > val.expiration {
		c.mu.Lock()
		delete(c.data, key)
		c.mu.Unlock()
		return nil, false
	}
	return val.value, ok
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	exp := time.Now().Add(ttl).UnixNano()

	c.data[key] = item{
		value:      value,
		expiration: exp,
	}
}
