package memcached

import (
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
)

const (
	// DriverName the unique name of the Memcached driver for register
	DriverName = "memcached"
)

type memcachedCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	client     *memcache.Client
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	servers, err := parseMemcachedConnection(conn)
	if err != nil {
		return nil, err
	}
	c := &memcachedCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		client:     memcache.New(servers...),
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *memcachedCache) Get(k string, v any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, err := c.client.Get(k)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nscache.ErrNil
	}
	if err != nil {
		return err
	}
	return c.serializer.Deserialize(item.Value, &v)
}

func (c *memcachedCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	return c.client.Set(&memcache.Item{Key: k, Value: data, Expiration: int32(expiration.Seconds())})
}

func (c *memcachedCache) Remove(k string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	err := c.client.Delete(k)
	if errors.Is(err, memcache.ErrCacheMiss) {
		return nil
	}
	return err
}

func (c *memcachedCache) Close() error {
	return c.client.Close()
}

// parseMemcachedConnection parse the Memcached connection string
func parseMemcachedConnection(u *url.URL) (servers []string, err error) {
	if u == nil {
		return nil, errors.New("invalid memcached connection string")
	}
	servers = append(servers, u.Host)
	for _, addr := range u.Query()["addr"] {
		servers = append(servers, addr)
	}
	return servers, nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
