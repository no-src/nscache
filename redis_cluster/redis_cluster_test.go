package redis_cluster

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.RedisClusterConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestRedisClusterCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
}

func TestRedisClusterCache_NewCache_WithNilURL(t *testing.T) {
	_, err := newCache(nil)
	if err == nil {
		t.Errorf("expect get an error, but get nil")
	}
}
