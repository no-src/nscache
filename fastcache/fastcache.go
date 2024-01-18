package fastcache

import (
	"errors"
	"net/url"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/no-src/nscache"
	"github.com/no-src/nscache/cache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/internal/unit"
)

const (
	// DriverName the unique name of the fastcache driver for register
	DriverName = "fastcache"
)

func newCache(conn *url.URL) (nscache.NSCache, error) {
	maxBytes, err := parseConnection(conn)
	if err != nil {
		return nil, err
	}
	return cache.NewCache(newStore(fastcache.New(maxBytes), encoding.DefaultSerializer))
}

func parseConnection(u *url.URL) (maxBytes int, err error) {
	if u == nil {
		return maxBytes, errors.New("invalid fastcache connection string")
	}
	maxBytesStr := u.Query().Get("max_bytes")
	maxBytes, err = unit.ParseBytes(maxBytesStr)
	if err != nil {
		return maxBytes, errors.New("invalid max_bytes parameter in the fastcache connection string")
	}
	return maxBytes, nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
