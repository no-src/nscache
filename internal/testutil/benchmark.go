package testutil

import (
	"testing"
	"time"

	"github.com/no-src/nscache"
)

// BenchmarkCacheGet the benchmark test of get cache data
func BenchmarkCacheGet(b *testing.B, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		b.Errorf("init cache error, err=%v", err)
		return
	}
	defer c.Close()

	c.Set(TestKey, TestValue, expiration)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v string
		if err = c.Get(TestKey, &v); err != nil {
			b.Errorf("Get: get data error, err=%v", err)
			return
		}
	}
}

// BenchmarkCacheSet the benchmark test of set cache data
func BenchmarkCacheSet(b *testing.B, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		b.Errorf("init cache error, err=%v", err)
		return
	}
	defer c.Close()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err = c.Set(TestKey, TestValue, expiration); err != nil {
			b.Errorf("Set: set data error, err=%v", err)
			return
		}
	}
}

// BenchmarkCacheRemove the benchmark test of remove cache data
func BenchmarkCacheRemove(b *testing.B, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		b.Errorf("init cache error, err=%v", err)
		return
	}
	defer c.Close()

	c.Set(TestKey, TestValue, expiration)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err = c.Remove(TestKey); err != nil {
			b.Errorf("Remove: remove data error, err=%v", err)
			return
		}
	}
}
