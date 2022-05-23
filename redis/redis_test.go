package redis

import (
	"testing"
	"time"

	"github.com/no-src/nscache"
)

var (
	connectionString = "redis://127.0.0.1:6379"
	expiration       = time.Minute
)

func TestRedisCache_GetString(t *testing.T) {
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

func TestRedisCache_GetString_NotExpectedValue(t *testing.T) {
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

func TestRedisCache_Get(t *testing.T) {
	testCases := []struct {
		k string
		v Account
	}{
		{"account_1", Account{UserName: "admin", Password: "admin_pwd", IsValid: true}},
		{"account_empty", Account{}},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			err := c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("TestMemoryCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestMemoryCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

func TestRedisCache_Get_Point(t *testing.T) {
	testCases := []struct {
		k string
		v *Account
	}{
		{"account_1", &Account{UserName: "admin", Password: "admin_pwd", IsValid: true}},
		{"account_empty", &Account{}},
		{"account_nil", nil},
	}

	c, _ := nscache.NewCache(connectionString)
	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			err := c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("TestMemoryCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
				return
			}
			var actual *Account
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("TestMemoryCache_Get error k=%v v=%v, err=%s", tc.k, tc.v, err)
			} else if !tc.v.equal(actual) {
				t.Errorf("k=%v v=%v, expect:%v, but actual:%v", tc.k, tc.v, tc.v, actual)
			}
		})
	}
}

type Account struct {
	UserName string
	Password string
	IsValid  bool
}

func (a *Account) equal(account *Account) bool {
	if a == nil || account == nil {
		return a == account
	}
	return a.UserName == account.UserName && a.Password == account.Password && a.IsValid == account.IsValid
}
