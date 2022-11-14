package nscache

import (
	"errors"
	"testing"

	"github.com/no-src/nscache/internal/testutil"
)

func TestNewCache(t *testing.T) {
	driverName := "testcache"
	conn := driverName + ":"
	Register(driverName, mockFactory)

	testCases := []struct {
		conn  string
		valid bool
	}{
		{conn, true},
		{"invalid", false},
		{"invalid_conn:", false},
		{"invalid+", false},
	}
	for _, tc := range testCases {
		t.Run(tc.conn, func(t *testing.T) {
			_, err := NewCache(tc.conn)
			if tc.valid && err != nil {
				t.Errorf("init cache error, connection string=%s err=%v", tc.conn, err)
				return
			}
			if !tc.valid && err == nil {
				t.Errorf("init cache error, expect to get an error but get nil, connection string=%s err=%v", tc.conn, err)
				return
			}

			if !tc.valid && errors.Is(err, errUnsupportedCacheDriver) {
				t.Errorf("init cache error, get an unsupported error, connection string=%s err=%v", tc.conn, err)
				return
			}
		})
	}
}

func TestNewCache_Unsupported(t *testing.T) {
	testCases := []struct {
		conn string
	}{
		{testutil.MemoryConnectionString},
		{testutil.BuntDBConnectionString},
		{testutil.EtcdConnectionString},
		{testutil.RedisConnectionString},
	}
	for _, tc := range testCases {
		t.Run(tc.conn, func(t *testing.T) {
			_, err := NewCache(tc.conn)
			if err == nil {
				t.Errorf("init cache error, expect to get an error but get nil, connection string=%s err=%v", tc.conn, err)
				return
			}
			if !errors.Is(err, errUnsupportedCacheDriver) {
				t.Errorf("init cache error, expect to get error =>%v, but get error => %v, connection string=%s", errUnsupportedCacheDriver, err, tc.conn)
			}
		})
	}
}
