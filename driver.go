package nscache

import (
	"net/url"
	"strings"

	"github.com/no-src/log"
)

var (
	drivers = make(map[string]CacheFactoryFunc)
)

// CacheFactoryFunc the cache driver factory function
type CacheFactoryFunc func(conn *url.URL) (NSCache, error)

// Register register a new cache driver
func Register(name string, factory CacheFactoryFunc) {
	name = strings.ToLower(name)
	if _, exist := drivers[name]; exist {
		log.Debug("the cache driver [%s] already existed", name)
	}
	drivers[name] = factory
	log.Debug("the cache driver [%s] is registered", name)
}
