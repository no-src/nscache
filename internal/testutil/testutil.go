package testutil

import "time"

const (
	// MemoryConnectionString a memory cache driver test connection string
	MemoryConnectionString = "memory:"
	// BuntDBConnectionString a buntdb cache driver test connection string
	BuntDBConnectionString = "buntdb://:memory:"
	// EtcdConnectionString a etcd cache driver test connection string
	EtcdConnectionString = "etcd://127.0.0.1:2379?dial_timeout=5s"
	// RedisConnectionString a redis cache driver test connection string
	RedisConnectionString = "redis://127.0.0.1:6379"
	// DefaultExpiration the default expiration time for cache driver tests
	DefaultExpiration = time.Minute
)
