package memcached

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.MemcachedConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestMemcachedCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
}
