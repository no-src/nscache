package boltdb

import (
	"time"

	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/store"
	"go.etcd.io/bbolt"
)

type boltDBStore struct {
	db         *bbolt.DB
	bucket     []byte
	serializer encoding.Serializer
}

func newStore(db *bbolt.DB, bucket []byte, serializer encoding.Serializer) store.Store {
	return &boltDBStore{
		db:         db,
		bucket:     bucket,
		serializer: serializer,
	}
}

func (s *boltDBStore) Get(k string) *store.Data {
	var data []byte
	s.db.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket(s.bucket)
		if b != nil {
			data = b.Get([]byte(k))
		}
		return nil
	})
	if len(data) == 0 {
		return nil
	}
	var d *store.Data
	if s.serializer.Deserialize(data, &d) != nil {
		return nil
	}
	return d
}

func (s *boltDBStore) Set(k string, data []byte, expiration time.Duration) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists(s.bucket)
		if err != nil {
			return err
		}
		sd, err := s.serializer.Serialize(store.NewData(data, expiration))
		if err != nil {
			return err
		}
		return b.Put([]byte(k), sd)
	})
}

func (s *boltDBStore) Remove(k string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(s.bucket)
		if b == nil {
			return nil
		}
		return b.Delete([]byte(k))
	})
}

func (s *boltDBStore) Close() error {
	return s.db.Close()
}
