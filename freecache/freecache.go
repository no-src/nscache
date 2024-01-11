package freecache

import (
	"errors"
	"net/url"
	"strconv"
	"time"

	"github.com/coocood/freecache"
	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
)

const (
	// DriverName the unique name of the freecache driver for register
	DriverName = "freecache"
)

type freeCache struct {
	nscache.NSCacheExt

	db         *freecache.Cache
	conn       *url.URL
	serializer encoding.Serializer
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	cacheSize, err := parseConnection(conn)
	if err != nil {
		return nil, err
	}
	db := freecache.NewCache(cacheSize)
	c := &freeCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		db:         db,
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func parseConnection(u *url.URL) (cacheSize int, err error) {
	if u == nil {
		return cacheSize, errors.New("invalid freecache connection string")
	}
	sizeStr := u.Query().Get("cache_size")
	cacheSize, err = strconv.Atoi(sizeStr)
	if err != nil {
		return cacheSize, errors.New("invalid cache_size parameter in the freecache connection string")
	}
	return cacheSize, nil
}

func (c *freeCache) Get(k string, v any) error {
	data, err := c.db.Get([]byte(k))
	if errors.Is(err, freecache.ErrNotFound) {
		err = nscache.ErrNil
	}
	if err != nil {
		return err
	}
	return c.serializer.Deserialize(data, &v)
}

func (c *freeCache) Set(k string, v any, expiration time.Duration) error {
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	return c.db.Set([]byte(k), data, int(expiration.Seconds()))
}

func (c *freeCache) Remove(k string) error {
	c.db.Del([]byte(k))
	return nil
}

func (c *freeCache) Close() error {
	return nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
