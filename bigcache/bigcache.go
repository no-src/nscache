package bigcache

import (
	"context"
	"errors"
	"net/url"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/no-src/nscache"
	"github.com/no-src/nscache/cache"
	"github.com/no-src/nscache/encoding"
)

const (
	// DriverName the unique name of the bigcache driver for register
	DriverName = "bigcache"
)

func newCache(conn *url.URL) (nscache.NSCache, error) {
	eviction, err := parseConnection(conn)
	if err != nil {
		return nil, err
	}
	db, _ := bigcache.New(context.Background(), bigcache.DefaultConfig(eviction))
	return cache.NewCache(newStore(db, encoding.DefaultSerializer))
}

func parseConnection(u *url.URL) (eviction time.Duration, err error) {
	if u == nil {
		return eviction, errors.New("invalid bigcache connection string")
	}
	evictionStr := u.Query().Get("eviction")
	eviction, err = time.ParseDuration(evictionStr)
	if err != nil {
		return eviction, errors.New("invalid eviction parameter in the bigcache connection string")
	}
	return eviction, nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
