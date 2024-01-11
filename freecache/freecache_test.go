package freecache

import (
	"net/url"
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.FreeCacheConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestFreeCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
}

func TestFreeCache_NewCache_WithNilURL(t *testing.T) {
	_, err := newCache(nil)
	if err == nil {
		t.Errorf("expect get an error, but get nil")
	}
}

func TestFreeCache_NewCache_WithInvalidURL(t *testing.T) {
	testCases := []struct {
		conn string
	}{
		{"freecache:"},
		{"freecache://"},
		{"freecache://?cache_size=5z"},
	}
	for _, tc := range testCases {
		t.Run(tc.conn, func(t *testing.T) {
			u, err := url.Parse(tc.conn)
			if err != nil {
				t.Errorf("invalid url => %s", tc.conn)
				return
			}
			_, err = newCache(u)
			if err == nil {
				t.Errorf("expect get an error, but get nil")
			}
		})
	}
}
