package freecache

import (
	"testing"
	
	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkFreeCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkFreeCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkFreeCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}
