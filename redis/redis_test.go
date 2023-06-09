package redis

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.RedisConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestRedisCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
}

func TestRedisCache_NewCache_WithNilURL(t *testing.T) {
	_, err := newCache(nil)
	if err == nil {
		t.Errorf("expect get an error, but get nil")
	}
}
