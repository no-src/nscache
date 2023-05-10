package cache_test

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
	_ "github.com/no-src/nscache/memory"
)

func TestCache(t *testing.T) {
	testutil.TestCache(t, testutil.MemoryConnectionString, testutil.DefaultExpiration)
}
