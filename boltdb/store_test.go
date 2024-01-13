package boltdb

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/internal/testutil"
	"go.etcd.io/bbolt"
)

func TestBoltDBStore_RemoveDataInANotExistBucket(t *testing.T) {
	db, err := getTestBoltDB()
	if err != nil {
		t.Errorf("open boltdb error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, []byte(defaultBucket), encoding.DefaultSerializer)
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if err != nil {
		t.Errorf("add cache data error => %v", err)
		return
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		return tx.DeleteBucket([]byte(defaultBucket))
	})
	if err != nil {
		t.Errorf("remove bucket error => %v", err)
		return
	}
	err = s.Remove(testutil.TestKey)
	if err != nil {
		t.Errorf("remove cache error => %v", err)
	}
}

func TestBoltDBStore_GetDataReturnDeserializeError(t *testing.T) {
	db, err := getTestBoltDB()
	if err != nil {
		t.Errorf("open boltdb error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, []byte(defaultBucket), encoding.DefaultSerializer)
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if err != nil {
		t.Errorf("add cache data error => %v", err)
		return
	}
	s = newStore(db, []byte(defaultBucket), &testutil.MockErrSerializer{})
	data := s.Get(testutil.TestKey)
	if data != nil {
		t.Errorf("expect to get a nil data, but actual %v", data)
	}
}

func TestBoltDBStore_SetDataReturnSerializeError(t *testing.T) {
	db, err := getTestBoltDB()
	if err != nil {
		t.Errorf("open boltdb error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, []byte(defaultBucket), &testutil.MockErrSerializer{})
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if !errors.Is(err, testutil.ErrMockSerialize) {
		t.Errorf("add cache data expect to get an error %v, but actual %v", testutil.ErrMockSerialize, err)
	}
}

func TestBoltDBStore_WithEmptyBucket(t *testing.T) {
	db, err := getTestBoltDB()
	if err != nil {
		t.Errorf("open boltdb error =>%v", err)
		return
	}
	defer db.Close()

	s := newStore(db, nil, encoding.DefaultSerializer)
	err = s.Set(testutil.TestKey, []byte(testutil.TestValue), time.Second)
	if !errors.Is(err, bbolt.ErrBucketNameRequired) {
		t.Errorf("add cache data expect to get an error %v, but actual %v", bbolt.ErrBucketNameRequired, err)
	}
}

func getTestBoltDB() (*bbolt.DB, error) {
	return bbolt.Open("boltdb_test.db", 0600, nil)
}
