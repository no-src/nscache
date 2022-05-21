package memory

import (
	"net/url"
	"sync"

	"github.com/no-src/nscache"
)

const (
	driverName = "memory"
)

type memoryCache struct {
	conn *url.URL
	mu   sync.Mutex
	data map[string]any
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	c := &memoryCache{
		data: make(map[string]any),
		conn: conn,
	}
	return c, nil
}

func (c *memoryCache) Get(k string) any {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.data[k]
}

func (c *memoryCache) GetString(k string) (s string, ok bool) {
	v := c.Get(k)
	if v == nil {
		return "", false
	}
	s, ok = v.(string)
	return s, ok
}

func (c *memoryCache) Set(k string, v any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[k] = v
}

func init() {
	nscache.Register(driverName, newCache)
}
