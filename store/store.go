package store

import "time"

// Store the cache store interface, it must be concurrency safe
type Store interface {
	// Get get cache data
	Get(k string) *Data
	// Set set cache data with expiration time
	Set(k string, data []byte, expiration time.Duration) error
	// Remove remove the cache data
	Remove(k string) error
	// Close close the store component
	Close() error
}
