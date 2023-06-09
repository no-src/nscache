package buntdb

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

func BenchmarkBuntDBCache_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, connectionString, expiration)
}

func BenchmarkBuntDBCache_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, connectionString, expiration)
}

func BenchmarkBuntDBCache_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, connectionString, expiration)
}

func BenchmarkBuntDBCache_Memory_Get(b *testing.B) {
	testutil.BenchmarkCacheGet(b, memoryConnectionString, expiration)
}

func BenchmarkBuntDBCache_Memory_Set(b *testing.B) {
	testutil.BenchmarkCacheSet(b, memoryConnectionString, expiration)
}

func BenchmarkBuntDBCache_Memory_Remove(b *testing.B) {
	testutil.BenchmarkCacheRemove(b, memoryConnectionString, expiration)
}
