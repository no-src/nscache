package memory

import (
	"net/url"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
)

const (
	driverName = "memory"
)

type memoryCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	data       map[string]*memoryData
}

type memoryData struct {
	data       []byte
	expireTime time.Time
}

func newMemoryData(data []byte, expiration time.Duration) *memoryData {
	return &memoryData{
		data:       data,
		expireTime: time.Now().Add(expiration),
	}
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	c := &memoryCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		data:       make(map[string]*memoryData),
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *memoryCache) Get(k string, v any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	md := c.data[k]
	if md == nil || md.expireTime.Before(time.Now()) {
		delete(c.data, k)
		return nil
	}
	return c.serializer.Deserialize(md.data, &v)
}

func (c *memoryCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	c.data[k] = newMemoryData(data, expiration)
	return nil
}

func init() {
	nscache.Register(driverName, newCache)
}
