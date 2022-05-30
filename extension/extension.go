package extension

import (
	"github.com/no-src/nscache"
)

type extension struct {
	c nscache.NSCache
}

// New returns an instance of the nscache.NSCacheExt implementation
func New(cache nscache.NSCache) nscache.NSCacheExt {
	return &extension{
		c: cache,
	}
}

func (ext *extension) GetString(k string) (s string, ok bool) {
	var v *string
	err := ext.c.Get(k, &v)
	if err != nil || v == nil {
		return "", false
	}
	return *v, true
}
