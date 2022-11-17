package nscache

import "time"

// NSCacheExt the NSCache extension function collection
type NSCacheExt interface {
	// GetBool get bool cache data by key
	GetBool(k string) (v bool, ok bool)
	// GetUint8 get uint8 cache data by key
	GetUint8(k string) (v uint8, ok bool)
	// GetUint16 get uint16 cache data by key
	GetUint16(k string) (v uint16, ok bool)
	// GetUint32 get uint32 cache data by key
	GetUint32(k string) (v uint32, ok bool)
	// GetUint64 get uint64 cache data by key
	GetUint64(k string) (v uint64, ok bool)
	// GetInt8 get int8 cache data by key
	GetInt8(k string) (v int8, ok bool)
	// GetInt16 get int16 cache data by key
	GetInt16(k string) (v int16, ok bool)
	// GetInt32 get int32 cache data by key
	GetInt32(k string) (v int32, ok bool)
	// GetInt64 get int64 cache data by key
	GetInt64(k string) (v int64, ok bool)
	// GetFloat32 get float32 cache data by key
	GetFloat32(k string) (v float32, ok bool)
	// GetFloat64 get float64 cache data by key
	GetFloat64(k string) (v float64, ok bool)
	// GetString get string cache data by key
	GetString(k string) (v string, ok bool)
	// GetStrings get string list cache data by key
	GetStrings(k string) (v []string, ok bool)
	// GetInt get int cache data by key
	GetInt(k string) (v int, ok bool)
	// GetInts get int list cache data by key
	GetInts(k string) (v []int, ok bool)
	// GetUint get uint cache data by key
	GetUint(k string) (v uint, ok bool)
	// GetUintptr get uintptr cache data by key
	GetUintptr(k string) (v uintptr, ok bool)
	// GetByte get byte cache data by key
	GetByte(k string) (v byte, ok bool)
	// GetBytes get byte list cache data by key
	GetBytes(k string) (v []byte, ok bool)
	// GetRune get rune cache data by key
	GetRune(k string) (v rune, ok bool)
	// GetTime get time.Time cache data by key
	GetTime(k string) (v time.Time, ok bool)
	// GetDuration get time.Duration cache data by key
	GetDuration(k string) (v time.Duration, ok bool)
}
