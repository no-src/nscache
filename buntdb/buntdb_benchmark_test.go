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
