package buntdb

import (
	"testing"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = testutil.BuntDBConnectionString
	expiration       = testutil.DefaultExpiration
)

func TestBuntDBCache_GetString(t *testing.T) {
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
			c.Set(tc.k, tc.v, expiration)
			actual, ok := c.GetString(tc.k)
			if !ok || actual != tc.v {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

func TestBuntDBCache_GetString_NotExpectedValue(t *testing.T) {
	testCases := []struct {
		k string
		v any
	}{
		{"nil_value", nil},
		{"error_type", time.Now()},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			c.Set(tc.k, tc.v, expiration)
			actual, ok := c.GetString(tc.k)
			if ok && actual == tc.v {
				t.Errorf("k=%v expect get not ok, but get ok", tc.k)
			}
		})
	}
}

func TestBuntDBCache_Get(t *testing.T) {
	testCases := []struct {
		k string
		v testutil.Account
	}{
		{"account_1", testutil.Account{UserName: "admin", Password: "admin_pwd", IsValid: true}},
		{"account_empty", testutil.Account{}},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			err := c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("TestBuntDBCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *testutil.Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestBuntDBCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.Equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

func TestBuntDBCache_Get_Pointer(t *testing.T) {
	testCases := []struct {
		k string
		v *testutil.Account
	}{
		{"account_1", &testutil.Account{UserName: "admin", Password: "admin_pwd", IsValid: true}},
		{"account_empty", &testutil.Account{}},
		{"account_nil", nil},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			err := c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("TestBuntDBCache_Get_Pointer error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *testutil.Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestBuntDBCache_Get_Pointer error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.Equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}
