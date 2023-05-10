package memory

import (
	"net/url"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/cache"
)

const (
	// DriverName the unique name of the memory driver for register
	DriverName = "memory"
)

func newCache(_ *url.URL) (nscache.NSCache, error) {
	return cache.NewCache(newStore())
}

func init() {
	nscache.Register(DriverName, newCache)
}
