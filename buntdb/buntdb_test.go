package buntdb

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.BuntDBConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestBuntDBCache(t *testing.T) {
	testutil.TestCache(t, connectionString, expiration)
}
