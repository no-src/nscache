package buntdb

import (
	"net/url"
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString       = testutil.BuntDBConnectionString
	memoryConnectionString = testutil.BuntDBMemoryConnectionString
	expiration             = testutil.DefaultExpiration
)

func TestBuntDBCache(t *testing.T) {
	testutil.TestCache(t, connectionString, testutil.NoExpiration)
	testutil.TestCache(t, memoryConnectionString, testutil.NoExpiration)
	testutil.TestCache(t, connectionString, expiration)
	testutil.TestCache(t, memoryConnectionString, expiration)
}

func TestBuntDBCache_NewCache_WithNilURL(t *testing.T) {
	_, err := newCache(nil)
	if err == nil {
		t.Errorf("expect get an error, but get nil")
	}
}

func TestBuntDBCache_NewCache_WithInvalidURL(t *testing.T) {
	testCases := []struct {
		conn string
	}{
		{"buntdb:///invalid"},
		{"buntdb://"},
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
