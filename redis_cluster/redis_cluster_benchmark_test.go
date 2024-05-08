package redis_cluster

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkRedisClusterCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkRedisClusterCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkRedisClusterCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}
