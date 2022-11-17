package memory

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.MemoryConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestMemoryCache(t *testing.T) {
	testutil.TestCache(t, connectionString, expiration)
}
