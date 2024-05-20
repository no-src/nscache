package client

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.ProxyCacheConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestMemoryCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
}
