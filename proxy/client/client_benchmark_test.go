package client

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkMemoryCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkMemoryCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkMemoryCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}
