package etcd

import (
	"testing"
	"time"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/internal/testutil"
)

var (
	connectionString = "etcd://127.0.0.1:2379?dial_timeout=5s"
	expiration       = time.Minute
)

func TestEtcdCache_GetString(t *testing.T) {
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

func TestEtcdCache_GetString_NotExpectedValue(t *testing.T) {
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

func TestEtcdCache_Get(t *testing.T) {
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
				t.Errorf("TestEtcdCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *testutil.Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestEtcdCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.Equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

func TestEtcdCache_Get_Point(t *testing.T) {
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
				t.Errorf("TestEtcdCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *testutil.Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestEtcdCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.Equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}
