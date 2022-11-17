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
	testutil.TestCache(t, connectionString, expiration)
}
