package testutil

import "time"

const (
	MemoryConnectionString = "memory:"
	BuntDBConnectionString = "buntdb://:memory:"
	EtcdConnectionString   = "etcd://127.0.0.1:2379?dial_timeout=5s"
	RedisConnectionString  = "redis://127.0.0.1:6379"

	DefaultExpiration = time.Minute
)
