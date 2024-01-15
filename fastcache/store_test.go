package fastcache

import (
	"errors"
	"testing"
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/internal/testutil"
)

func TestFastCacheStore_GetDataReturnDeserializeError(t *testing.T) {
	db := getTestFastCache()
	s := newStore(db, encoding.DefaultSerializer)
	err := s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if err != nil {
		t.Errorf("add cache data error => %v", err)
		return
	}
	s = newStore(db, &testutil.MockErrSerializer{})
	data := s.Get(testutil.TestKey)
	if data != nil {
		t.Errorf("expect to get a nil data, but actual %v", data)
	}
}

func TestFastCacheStore_SetDataReturnSerializeError(t *testing.T) {
	db := getTestFastCache()
	s := newStore(db, &testutil.MockErrSerializer{})
	err := s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if !errors.Is(err, testutil.ErrMockSerialize) {
		t.Errorf("add cache data expect to get an error %v, but actual %v", testutil.ErrMockSerialize, err)
	}
}

func getTestFastCache() *fastcache.Cache {
	return fastcache.New(10000000)
}
