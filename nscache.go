package nscache

import (
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// ErrNil get nil data
var ErrNil = errors.New("nscache: nil")

// NSCache the core interface of the cache
type NSCache interface {
	NSCacheExt

	// Get get cache data by key
	Get(k string, v any) error

	// Set set new cache data
	Set(k string, v any, expiration time.Duration) error

	// Remove remove the specified key
	Remove(k string) error

	// Close close the cache component
	Close() error
}

// NewCache get an instance of NSCache by connection string
func NewCache(conn string) (NSCache, error) {
	u, err := url.Parse(conn)
	if err != nil {
		return nil, err
	}
	driverName := strings.ToLower(u.Scheme)
	if len(driverName) == 0 {
		return nil, errInvalidCacheDriverName
	}
	mu.RLock()
	factory := drivers[driverName]
	mu.RUnlock()
	if factory == nil {
		return nil, fmt.Errorf("%w => %s", errUnsupportedCacheDriver, driverName)
	}
	return factory(u)
}
