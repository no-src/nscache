package nscache

import (
	"net/url"
	"strings"
	"sync"

	"github.com/no-src/log"
)

var (
	drivers = make(map[string]CacheFactoryFunc)
	mu      sync.RWMutex
)

// CacheFactoryFunc the cache driver factory function
type CacheFactoryFunc func(conn *url.URL) (NSCache, error)

// Register register a new cache driver
func Register(name string, factory CacheFactoryFunc) (overwritten bool) {
	if factory == nil {
		panic("the cache driver factory can't be nil")
	}
	name = strings.ToLower(name)
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
