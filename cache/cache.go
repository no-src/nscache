package cache

import (
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/no-src/nscache/store"
)

type cache struct {
	nscache.NSCacheExt

	serializer encoding.Serializer
	mu         sync.RWMutex
	store      store.Store
}

// NewCache create an instance of NSCache with the Store instance
func NewCache(store store.Store) (nscache.NSCache, error) {
	c := &cache{
		serializer: encoding.DefaultSerializer,
		store:      store,
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *cache) Get(k string, v any) error {
	c.mu.RLock()
	md := c.store.Get(k)
	c.mu.RUnlock()
	if md == nil {
		return nscache.ErrNil
	}
	if md.IsExpired() {
		go c.Remove(k)
		return nscache.ErrNil
	}
	return c.serializer.Deserialize(md.Data, &v)
}

func (c *cache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	return c.store.Set(k, data, expiration)
}

func (c *cache) Remove(k string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.store.Remove(k)
}

func (c *cache) Close() error {
	return c.store.Close()
}
