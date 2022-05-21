package memory

import (
	"testing"

	"github.com/no-src/nscache"
)

func BenchmarkGet(b *testing.B) {
	c, _ := nscache.NewCache(connectionString)
	c.Set("hello", "world")
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Get("hello")
	}
}

func BenchmarkSet(b *testing.B) {
	c, _ := nscache.NewCache(connectionString)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Set("hello", "world")
	}
}
