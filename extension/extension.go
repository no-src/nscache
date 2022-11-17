package extension

import (
	"time"

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

func (ext *extension) GetBool(k string) (v bool, ok bool) {
	return getValue[bool](k, ext.c.Get)
}

func (ext *extension) GetUint8(k string) (v uint8, ok bool) {
	return getValue[uint8](k, ext.c.Get)
}

func (ext *extension) GetUint16(k string) (v uint16, ok bool) {
	return getValue[uint16](k, ext.c.Get)
}

func (ext *extension) GetUint32(k string) (v uint32, ok bool) {
	return getValue[uint32](k, ext.c.Get)
}

func (ext *extension) GetUint64(k string) (v uint64, ok bool) {
	return getValue[uint64](k, ext.c.Get)
}

func (ext *extension) GetInt8(k string) (v int8, ok bool) {
	return getValue[int8](k, ext.c.Get)
}

func (ext *extension) GetInt16(k string) (v int16, ok bool) {
	return getValue[int16](k, ext.c.Get)
}

func (ext *extension) GetInt32(k string) (v int32, ok bool) {
	return getValue[int32](k, ext.c.Get)
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

func (ext *extension) GetString(k string) (v string, ok bool) {
	return getValue[string](k, ext.c.Get)
}

func (ext *extension) GetStrings(k string) (v []string, ok bool) {
	return getValue[[]string](k, ext.c.Get)
}

func (ext *extension) GetInt(k string) (v int, ok bool) {
	return getValue[int](k, ext.c.Get)
}

func (ext *extension) GetInts(k string) (v []int, ok bool) {
	return getValue[[]int](k, ext.c.Get)
}

func (ext *extension) GetUint(k string) (v uint, ok bool) {
	return getValue[uint](k, ext.c.Get)
}

func (ext *extension) GetUintptr(k string) (v uintptr, ok bool) {
	return getValue[uintptr](k, ext.c.Get)
}

func (ext *extension) GetByte(k string) (v byte, ok bool) {
	return getValue[byte](k, ext.c.Get)
}

func (ext *extension) GetBytes(k string) (v []byte, ok bool) {
	return getValue[[]byte](k, ext.c.Get)
}

func (ext *extension) GetRune(k string) (v rune, ok bool) {
	return getValue[rune](k, ext.c.Get)
}

func (ext *extension) GetTime(k string) (v time.Time, ok bool) {
	return getValue[time.Time](k, ext.c.Get)
}

func (ext *extension) GetDuration(k string) (v time.Duration, ok bool) {
	return getValue[time.Duration](k, ext.c.Get)
}

func getValue[T any](k string, get func(k string, v any) error) (v T, ok bool) {
	var p *T
	err := get(k, &p)
	if err != nil || p == nil {
		return v, false
	}
	return *p, true
}
