package testutil

import (
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
	DefaultExpiration = time.Minute
)

// TestCache test the cache with the passed connection string
func TestCache(t *testing.T, conn string, expiration time.Duration) {
	testCases := []struct {
		k string
		v TestStruct
	}{
		{"ts_1", TestStruct{Name: "admin", ID: 1, IsValid: true}},
		{"ts_2", TestStruct{Name: "root", ID: 2, IsValid: false}},
		{"ts_empty", TestStruct{}},
	}

	c, err := nscache.NewCache(conn)
	if err != nil {
		t.Errorf("init cache error, err=%v", err)
		return
	}
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			err = c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("Set: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}
			var actual TestStruct
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("Get: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
			} else if !tc.v.Equal(&actual) {
				t.Errorf("not equal, k=%v, expect:%v, but actual:%v", tc.k, tc.v, actual)
			}
		})
	}
}
