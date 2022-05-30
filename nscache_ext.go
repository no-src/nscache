package nscache

// NSCacheExt the NSCache extension function collection
type NSCacheExt interface {
	// GetString get string cache data by key
	GetString(k string) (s string, ok bool)
}
