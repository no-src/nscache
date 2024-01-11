package extension_test

import (
	"testing"
	"time"

	_ "github.com/no-src/nscache/bigcache"
	_ "github.com/no-src/nscache/buntdb"
	_ "github.com/no-src/nscache/etcd"
	_ "github.com/no-src/nscache/freecache"
	_ "github.com/no-src/nscache/memory"
	_ "github.com/no-src/nscache/redis"

	"github.com/no-src/nscache"
	"github.com/no-src/nscache/extension"
	"github.com/no-src/nscache/internal/testutil"
)

func TestExtension(t *testing.T) {
	testCases := []struct {
		conn string
	}{
		{testutil.MemoryConnectionString},
		{testutil.BuntDBConnectionString},
		{testutil.BuntDBMemoryConnectionString},
		{testutil.EtcdConnectionString},
		{testutil.RedisConnectionString},
		{testutil.FreeCacheConnectionString},
		{testutil.BigCacheConnectionString},
	}
	for _, tc := range testCases {
		t.Run(tc.conn, func(t *testing.T) {
			testExtension(t, tc.conn)
		})
	}
}

func testExtension(t *testing.T, conn string) {
	c, err := nscache.NewCache(conn)
	if err != nil {
		t.Errorf("init cache error, err=%v", err)
		return
	}
	expiration := testutil.DefaultExpiration
	ext := extension.New(c)
	var expect any

	expect = true
	c.Set("bool", expect, expiration)
	if v, ok := ext.GetBool("bool"); !ok || v != expect {
		t.Errorf("GetBool: expect to get %v but get %v", expect, v)
	}

	expect = uint8(123)
	c.Set("uint8", expect, expiration)
	if v, ok := ext.GetUint8("uint8"); !ok || v != expect {
		t.Errorf("GetUint8: expect to get %v but get %v", expect, v)
	}

	expect = uint16(123)
	c.Set("uint16", expect, expiration)
	if v, ok := ext.GetUint16("uint16"); !ok || v != expect {
		t.Errorf("GetUint16: expect to get %v but get %v", expect, v)
	}

	expect = uint32(123)
	c.Set("uint32", expect, expiration)
	if v, ok := ext.GetUint32("uint32"); !ok || v != expect {
		t.Errorf("GetUint32: expect to get %v but get %v", expect, v)
	}

	expect = uint64(123)
	c.Set("uint64", expect, expiration)
	if v, ok := ext.GetUint64("uint64"); !ok || v != expect {
		t.Errorf("GetUint64: expect to get %v but get %v", expect, v)
	}

	expect = int8(-123)
	c.Set("int8", expect, expiration)
	if v, ok := ext.GetInt8("int8"); !ok || v != expect {
		t.Errorf("GetInt8: expect to get %v but get %v", expect, v)
	}

	expect = int16(-123)
	c.Set("int16", expect, expiration)
	if v, ok := ext.GetInt16("int16"); !ok || v != expect {
		t.Errorf("GetInt16: expect to get %v but get %v", expect, v)
	}

	expect = int32(-123)
	c.Set("int32", expect, expiration)
	if v, ok := ext.GetInt32("int32"); !ok || v != expect {
		t.Errorf("GetInt32: expect to get %v but get %v", expect, v)
	}

	expect = int64(-123)
	c.Set("int64", expect, expiration)
	if v, ok := ext.GetInt64("int64"); !ok || v != expect {
		t.Errorf("GetInt64: expect to get %v but get %v", expect, v)
	}

	expect = float32(10.01)
	c.Set("float32", expect, expiration)
	if v, ok := ext.GetFloat32("float32"); !ok || v != expect {
		t.Errorf("GetFloat32: expect to get %v but get %v", expect, v)
	}

	expect = float64(-10.01)
	c.Set("float64", expect, expiration)
	if v, ok := ext.GetFloat64("float64"); !ok || v != expect {
		t.Errorf("GetFloat64: expect to get %v but get %v", expect, v)
	}

	expect = "hello"
	c.Set("string", expect, expiration)
	if v, ok := ext.GetString("string"); !ok || v != expect {
		t.Errorf("GetString: expect to get %v but get %v", expect, v)
	}

	expect = []string{"hello", "world"}
	c.Set("strings", expect, expiration)
	if v, ok := ext.GetStrings("strings"); !ok || !equal(v, expect.([]string)) {
		t.Errorf("GetStrings: expect to get %v but get %v", expect, v)
	}

	expect = int(123)
	c.Set("int", expect, expiration)
	if v, ok := ext.GetInt("int"); !ok || v != expect {
		t.Errorf("GetInt: expect to get %v but get %v", expect, v)
	}

	expect = []int{123, 456}
	c.Set("ints", expect, expiration)
	if v, ok := ext.GetInts("ints"); !ok || !equal(v, expect.([]int)) {
		t.Errorf("GetInts: expect to get %v but get %v", expect, v)
	}

	expect = uint(123)
	c.Set("uint", expect, expiration)
	if v, ok := ext.GetUint("uint"); !ok || v != expect {
		t.Errorf("GetUint: expect to get %v but get %v", expect, v)
	}

	expect = uintptr(123)
	c.Set("uintptr", expect, expiration)
	if v, ok := ext.GetUintptr("uintptr"); !ok || v != expect {
		t.Errorf("GetUintptr: expect to get %v but get %v", expect, v)
	}

	expect = byte(123)
	c.Set("byte", expect, expiration)
	if v, ok := ext.GetByte("byte"); !ok || v != expect {
		t.Errorf("GetByte: expect to get %v but get %v", expect, v)
	}

	expect = []byte{10, 20}
	c.Set("bytes", expect, expiration)
	if v, ok := ext.GetBytes("bytes"); !ok || !equal(v, expect.([]byte)) {
		t.Errorf("GetBytes: expect to get %v but get %v", expect, v)
	}

	expect = rune(123)
	c.Set("rune", expect, expiration)
	if v, ok := ext.GetRune("rune"); !ok || v != expect {
		t.Errorf("GetRune: expect to get %v but get %v", expect, v)
	}

	expect, _ = time.Parse(time.RFC3339Nano, time.Now().Format(time.RFC3339Nano))
	c.Set("time", expect, expiration)
	if v, ok := ext.GetTime("time"); !ok || v != expect {
		t.Errorf("GetTime: expect to get %v but get %v", expect, v)
	}

	expect = time.Minute
	c.Set("duration", expect, expiration)
	if v, ok := ext.GetDuration("duration"); !ok || v != expect {
		t.Errorf("GetDuration: expect to get %v but get %v", expect, v)
	}

	// the key does not exist
	if _, ok := ext.GetDuration("not_exist_key"); ok {
		t.Errorf("GetDuration: expect to get ok = false but get ok = %v", ok)
	}
}

func equal[T comparable](x, y []T) bool {
	if len(x) != len(y) {
		return false
	}
	for i := 0; i < len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}
