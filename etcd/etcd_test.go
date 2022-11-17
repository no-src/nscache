package etcd

import (
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.EtcdConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestEtcdCache(t *testing.T) {
	testutil.TestCache(t, connectionString, expiration)
}
