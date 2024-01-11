package bigcache

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/internal/testutil"
)

func TestBigCacheStore_GetDataReturnDeserializeError(t *testing.T) {
	db, err := getTestBigCache()
	if err != nil {
		t.Errorf("init bigcache error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, encoding.DefaultSerializer)
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
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

func TestBigCacheStore_SetDataReturnSerializeError(t *testing.T) {
	db, err := getTestBigCache()
	if err != nil {
		t.Errorf("init bigcache error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, &testutil.MockErrSerializer{})
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if !errors.Is(err, testutil.ErrMockSerialize) {
		t.Errorf("add cache data expect to get an error %v, but actual %v", testutil.ErrMockSerialize, err)
	}
}

func getTestBigCache() (*bigcache.BigCache, error) {
	return bigcache.New(context.Background(), bigcache.DefaultConfig(time.Minute*10))
}
