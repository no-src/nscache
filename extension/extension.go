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

func (ext *extension) GetString(k string) (v string, ok bool) {
	return getValue[string](k, ext.c.Get)
}

func (ext *extension) GetBool(k string) (v bool, ok bool) {
	return getValue[bool](k, ext.c.Get)
}

func (ext *extension) GetByte(k string) (v byte, ok bool) {
	return getValue[byte](k, ext.c.Get)
}

func (ext *extension) GetInt(k string) (v int, ok bool) {
	return getValue[int](k, ext.c.Get)
}

func (ext *extension) GetInt64(k string) (v int64, ok bool) {
	return getValue[int64](k, ext.c.Get)
}

func (ext *extension) GetFloat32(k string) (v float32, ok bool) {
	return getValue[float32](k, ext.c.Get)
}

func (ext *extension) GetFloat64(k string) (v float64, ok bool) {
	return getValue[float64](k, ext.c.Get)
}

func (ext *extension) GetComplex64(k string) (v complex64, ok bool) {
	return getValue[complex64](k, ext.c.Get)
}

func (ext *extension) GetComplex128(k string) (v complex128, ok bool) {
	return getValue[complex128](k, ext.c.Get)
}

func getValue[T any](k string, get func(k string, v any) error) (v T, ok bool) {
	var p *T
	err := get(k, &p)
	if err != nil || p == nil {
		return v, false
	}
	return *p, true
}
