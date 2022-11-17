package testutil

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/nscache"
)

const (
	// MemoryConnectionString a memory cache driver test connection string
	MemoryConnectionString = "memory:"
	// BuntDBConnectionString a buntdb cache driver test connection string
	BuntDBConnectionString = "buntdb://:memory:"
	// EtcdConnectionString a etcd cache driver test connection string
	EtcdConnectionString = "etcd://127.0.0.1:2379?dial_timeout=5s"
	// RedisConnectionString a redis cache driver test connection string
	RedisConnectionString = "redis://127.0.0.1:6379"
	// DefaultExpiration the default expiration time for cache driver tests
	DefaultExpiration = time.Second * 3
)

// TestCache test the cache with the passed connection string
func TestCache(t *testing.T, conn string, expiration time.Duration) {
	testCases := []struct {
		k string
		v testStruct
	}{
		{"ts_1", testStruct{Name: "admin", ID: 1, IsValid: true}},
		{"ts_2", testStruct{Name: "root", ID: 2, IsValid: false}},
		{"ts_empty", testStruct{}},
	}

	c, err := nscache.NewCache(conn)
	if err != nil {
		t.Errorf("init cache error, err=%v", err)
		return
	}

	testCacheReturnError(t, c, expiration)

	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			var actual *testStruct
			// get data before set
			err = c.Get(tc.k, &actual)
			if err == nil {
				t.Errorf("Get: expect to get error => %v, but get nil, k=%v", nscache.ErrNil, tc.k)
				return
			} else if !errors.Is(err, nscache.ErrNil) {
				t.Errorf("Get: expect to get error => %v, but get %v, k=%v", nscache.ErrNil, err, tc.k)
				return
			}

			// set data
			err = c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("Set: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}

			// get data after set
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("Get: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			} else if !tc.v.Equal(actual) {
				t.Errorf("not equal, k=%v, expect:%v, but actual:%v", tc.k, tc.v, actual)
				return
			}

			// get data after data is expired
			<-time.After(expiration + time.Second*2)
			err = c.Get(tc.k, &actual)
			if err == nil {
				t.Errorf("Get: expect to get error => %v, but get nil, k=%v", nscache.ErrNil, tc.k)
				return
			} else if !errors.Is(err, nscache.ErrNil) {
				t.Errorf("Get: expect to get error => %v, but get %v, k=%v", nscache.ErrNil, err, tc.k)
				return
			}
		})
	}
}

func testCacheReturnError(t *testing.T, c nscache.NSCache, expiration time.Duration) {
	ts2 := &testCycleStruct{}
	ts2.Self = ts2
	err := c.Set("unsupported-type", ts2, expiration)
	if err == nil {
		t.Errorf("Set: expect to get an error but get nil")
	}
}

// BenchmarkCacheGet the benchmark test of get cache data
func BenchmarkCacheGet(b *testing.B, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		b.Errorf("init cache error, err=%v", err)
		return
	}
	c.Set("hello", "world", expiration)
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var v string
		c.Get("hello", &v)
	}
}

// BenchmarkCacheSet the benchmark test of set cache data
func BenchmarkCacheSet(b *testing.B, conn string, expiration time.Duration) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		b.Errorf("init cache error, err=%v", err)
		return
	}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		c.Set("hello", "world", expiration)
	}
}
