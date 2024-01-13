package testutil

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/nscache"
)

func testCache(t *testing.T, conn string, expiration time.Duration) {
	testCases := []struct {
		k string
		v testStruct
	}{
		{"ts_1", testStruct{Name: "admin", ID: 1, IsValid: true}},
		{"ts_2", testStruct{Name: "root", ID: 2, IsValid: false}},
		{"ts_empty", testStruct{}},
	}

	c, err := nscache.NewCache(conn)
	if err != nil {
		t.Errorf("init cache error, err=%v", err)
		return
	}
	defer c.Close()

	testCacheReturnError(t, c, expiration)

	for _, tc := range testCases {
		t.Run(tc.k, func(t *testing.T) {
			// remove the key to ensure the key does not exist
			err = c.Remove(tc.k)
			if err != nil {
				t.Errorf("Remove: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}

			var actual *testStruct
			// get data before set
			err = c.Get(tc.k, &actual)
			if !errors.Is(err, nscache.ErrNil) {
				t.Errorf("Get: expect to get error => %v, but get %v, k=%v", nscache.ErrNil, err, tc.k)
				return
			}

			// set data
			err = c.Set(tc.k, tc.v, expiration)
			if err != nil {
				t.Errorf("Set: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}

			time.Sleep(time.Millisecond)

			// get data after set
			err = c.Get(tc.k, &actual)
			if err != nil {
				t.Errorf("Get: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			} else if !tc.v.Equal(actual) {
				t.Errorf("not equal, k=%v, expect:%v, but actual:%v", tc.k, tc.v, actual)
				return
			}

			if expiration > 0 {
				// get data after data is expired
				<-time.After(expiration + time.Second*2)
				err = c.Get(tc.k, &actual)
				if !errors.Is(err, nscache.ErrNil) {
					t.Errorf("Get: expect to get error => %v, but get %v, k=%v", nscache.ErrNil, err, tc.k)
					return
				}
			}

			// set data with long expiration time
			err = c.Set(tc.k, tc.v, expiration*10)
			if err != nil {
				t.Errorf("Set: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}

			// remove the key
			err = c.Remove(tc.k)
			if err != nil {
				t.Errorf("Remove: get an error, k=%v v=%v, err=%v", tc.k, tc.v, err)
				return
			}

			// get data after the key is removed
			err = c.Get(tc.k, &actual)
			if !errors.Is(err, nscache.ErrNil) {
				t.Errorf("Get: expect to get error => %v, but get %v, k=%v", nscache.ErrNil, err, tc.k)
				return
			}
		})
	}
}

func testCacheReturnError(t *testing.T, c nscache.NSCache, expiration time.Duration) {
	ts2 := &testCycleStruct{}
	ts2.Self = ts2
	err := c.Set("unsupported-type", ts2, expiration)
	if err == nil {
		t.Errorf("Set: expect to get an error but get nil")
	}
}
