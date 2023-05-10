package boltdb

import (
	"errors"
	"net/url"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/cache"
	"github.com/no-src/nscache/encoding"
	"go.etcd.io/bbolt"
)

const (
	// DriverName the unique name of the boltdb driver for register
	DriverName    = "boltdb"
	defaultBucket = "nscache-default"
)

func newCache(conn *url.URL) (nscache.NSCache, error) {
	path, err := parseConnection(conn)
	if err != nil {
		return nil, err
	}
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	return cache.NewCache(newStore(db, []byte(defaultBucket), encoding.DefaultSerializer))
}

// parseConnection parse the boltdb connection string
func parseConnection(u *url.URL) (path string, err error) {
	if u == nil {
		return "", errors.New("invalid boltdb connection string")
	}
	path = u.Host
	return path, nil
}

func init() {
	nscache.Register(DriverName, newCache)
}
