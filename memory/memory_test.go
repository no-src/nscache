package memory

import (
	"testing"
	"time"

	"github.com/no-src/nscache"
)

var connectionString = "memory:"

func TestMemoryCache_GetString(t *testing.T) {
	testCases := []struct {
		k string
		v any
	}{
		{"hello", "world"},
		{"empty_key", ""},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			c.Set(tc.k, tc.v)
			actual, ok := c.GetString(tc.k)
			if !ok || actual != tc.v {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

func TestMemoryCache_GetString_NotExpectedValue(t *testing.T) {
	testCases := []struct {
		k string
		v any
	}{
		{"nil_key", nil},
		{"error_type", time.Now()},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			c.Set(tc.k, tc.v)
			_, ok := c.GetString(tc.k)
			if ok {
				t.Errorf("k=%v expect get not ok, but get ok", tc.k)
			}
		})
	}
}
