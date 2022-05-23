package nscache

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

// NSCache the core interface of the cache
type NSCache interface {
	// Get get cache data by key
	Get(k string, v any) (err error)

	// GetString get string cache data by key
	GetString(k string) (s string, ok bool)

	// Set set new cache data
	Set(k string, v any, expiration time.Duration) error
}

// NewCache get an instance of NSCache by connection string
func NewCache(conn string) (NSCache, error) {
	u, err := url.Parse(conn)
	if err != nil {
		return nil, err
	}
	driverName := strings.ToLower(u.Scheme)
	factory := drivers[driverName]
	if factory == nil {
		return nil, fmt.Errorf("find unsupported cache driver => %s", driverName)
	}
	return factory(u)
}
