package memory

import (
	"net/url"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
)

const (
	driverName = "memory"
)

type memoryCache struct {
	conn       *url.URL
	mu         sync.Mutex
	data       map[string]*memoryData
	serializer encoding.Serializer
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
		data:       make(map[string]*memoryData),
		serializer: encoding.DefaultSerializer,
	}
	return c, nil
}

func (c *memoryCache) Get(k string, v any) (err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	md := c.data[k]
	if md == nil || md.expireTime.Before(time.Now()) {
		delete(c.data, k)
		return nil
	}
	err = c.serializer.Deserialize(md.data, &v)
	return err
}

func (c *memoryCache) GetString(k string) (s string, ok bool) {
	var v any
	err := c.Get(k, &v)
	if err != nil || v == nil {
		return "", false
	}
	s, ok = v.(string)
	return s, ok
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
