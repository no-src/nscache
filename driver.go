package nscache

import (
	"errors"
	"net/url"
	"strings"
	"sync"

	"github.com/no-src/log"
)

var (
	drivers = make(map[string]CacheFactoryFunc)
	mu      sync.RWMutex
)

var (
	errCacheDriverFactoryIsNil = errors.New("the cache driver factory can't be nil")
	errUnsupportedCacheDriver  = errors.New("unsupported cache driver")
	errInvalidCacheDriverName  = errors.New("invalid cache driver name")
)

// CacheFactoryFunc the cache driver factory function
type CacheFactoryFunc func(conn *url.URL) (NSCache, error)

// Register register a new cache driver
func Register(name string, factory CacheFactoryFunc) (overwritten bool) {
	if factory == nil {
		panic(errCacheDriverFactoryIsNil)
	}
	name = strings.ToLower(name)
	checked, checkedName := checkDriverName(name)
	if !checked || checkedName != name {
		panic(errInvalidCacheDriverName)
	}
	mu.Lock()
	if _, exist := drivers[name]; exist {
		log.Debug("the cache driver [%s] already existed", name)
		overwritten = true
	}
	drivers[name] = factory
	mu.Unlock()
	log.Debug("the cache driver [%s] is registered", name)
	return overwritten
}

func checkDriverName(name string) (checked bool, checkedName string) {
	u, err := url.Parse(name + ":")
	if err != nil {
		return
	}
	return true, u.Scheme
}
