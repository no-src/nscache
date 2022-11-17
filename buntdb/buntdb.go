package buntdb

import (
	"errors"
	"net/url"
	"sync"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/extension"
	"github.com/tidwall/buntdb"
)

const (
	driverName = "buntdb"
)

type buntDBCache struct {
	nscache.NSCacheExt

	conn       *url.URL
	serializer encoding.Serializer
	mu         sync.RWMutex
	db         *buntdb.DB
}

func newCache(conn *url.URL) (nscache.NSCache, error) {
	path, err := parseBuntDBConnection(conn)
	if err != nil {
		return nil, err
	}
	db, err := buntdb.Open(path)
	if err != nil {
		return nil, err
	}
	c := &buntDBCache{
		conn:       conn,
		serializer: encoding.DefaultSerializer,
		db:         db,
	}
	c.NSCacheExt = extension.New(c)
	return c, nil
}

func (c *buntDBCache) Get(k string, v any) error {
	c.mu.RLock()
	defer c.mu.RUnlock()
	var data []byte
	err := c.db.View(func(tx *buntdb.Tx) error {
		val, getErr := tx.Get(k)
		if getErr == nil {
			data = []byte(val)
		}
		return getErr
	})
	if err == buntdb.ErrNotFound {
		err = nscache.ErrNil
	}
	if err != nil {
		return err
	}
	return c.serializer.Deserialize(data, &v)
}

func (c *buntDBCache) Set(k string, v any, expiration time.Duration) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	data, err := c.serializer.Serialize(v)
	if err != nil {
		return err
	}
	return c.db.Update(func(tx *buntdb.Tx) error {
		_, _, setErr := tx.Set(k, string(data), &buntdb.SetOptions{Expires: true, TTL: expiration})
		return setErr
	})
}

// parseBuntDBConnection parse the buntdb connection string
func parseBuntDBConnection(u *url.URL) (path string, err error) {
	if u == nil {
		return "", errors.New("invalid buntdb connection string")
	}
	path = u.Host
	return path, nil
}

func init() {
	nscache.Register(driverName, newCache)
}
