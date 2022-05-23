package redis

import (
	"testing"

	"github.com/no-src/nscache"
)

func BenchmarkRedisCache_Get(b *testing.B) {
	c, _ := nscache.NewCache(connectionString)
	c.Set("hello", "world", expiration)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v string
		c.Get("hello", &v)
	}
}

func BenchmarkRedisCache_Set(b *testing.B) {
	c, _ := nscache.NewCache(connectionString)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Set("hello", "world", expiration)
	}
}
