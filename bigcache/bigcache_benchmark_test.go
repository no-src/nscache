package bigcache

import (
	"testing"
	
	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkBigCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkBigCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkBigCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}
