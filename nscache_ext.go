package nscache

// NSCacheExt the NSCache extension function collection
type NSCacheExt interface {
	// GetString get string cache data by key
	GetString(k string) (v string, ok bool)
	// GetBool get bool cache data by key
	GetBool(k string) (v bool, ok bool)
	// GetByte get byte cache data by key
	GetByte(k string) (v byte, ok bool)
	// GetInt get int cache data by key
	GetInt(k string) (v int, ok bool)
	// GetInt64 get int64 cache data by key
	GetInt64(k string) (v int64, ok bool)
	// GetFloat32 get float32 cache data by key
	GetFloat32(k string) (v float32, ok bool)
	// GetFloat64 get float64 cache data by key
	GetFloat64(k string) (v float64, ok bool)
	// GetComplex64 get complex64 cache data by key
	GetComplex64(k string) (v complex64, ok bool)
	// GetComplex128 get complex128 cache data by key
	GetComplex128(k string) (v complex128, ok bool)
}
