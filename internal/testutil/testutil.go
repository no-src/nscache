package testutil

import (
	"testing"
	"time"
)

const (
	// MemoryConnectionString a memory cache driver test connection string
	MemoryConnectionString = "memory:"
	// BuntDBConnectionString a buntdb cache driver test connection string
	BuntDBConnectionString = "buntdb://buntdb.db"
	// BuntDBMemoryConnectionString a buntdb memory cache driver test connection string
	BuntDBMemoryConnectionString = "buntdb://:memory:"
	// EtcdConnectionString a etcd cache driver test connection string
	EtcdConnectionString = "etcd://127.0.0.1:2379?dial_timeout=5s"
	// RedisConnectionString a redis cache driver test connection string
	RedisConnectionString = "redis://127.0.0.1:6379"
	// BoltDBConnectionString a boltdb cache driver test connection string
	BoltDBConnectionString = "boltdb://boltdb.db"
	// FreeCacheConnectionString a freecache driver test connection string
	FreeCacheConnectionString = "freecache://?cache_size=10000000"
	// BigCacheConnectionString a bigcache driver test connection string
	BigCacheConnectionString = "bigcache://?eviction=10m"
	// DefaultExpiration the default expiration time for cache driver tests
	DefaultExpiration = time.Second * 3
	// NoExpiration means never expire
	NoExpiration = 0
)

const (
	TestKey   = "hello"
	TestValue = "world"
)

// TestCache test the cache with the passed connection string
func TestCache(t *testing.T, conn string, expiration time.Duration) {
	testCache(t, conn, expiration)
	testCacheConcurrent(t, conn, expiration)
}
