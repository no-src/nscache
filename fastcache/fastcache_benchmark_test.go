package fastcache

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkFastCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkFastCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkFastCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}
