package bigcache

import (
	"errors"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/store"
)

type bigCacheStore struct {
	db         *bigcache.BigCache
	serializer encoding.Serializer
}

func newStore(db *bigcache.BigCache, serializer encoding.Serializer) store.Store {
	return &bigCacheStore{
		db:         db,
		serializer: serializer,
	}
}

func (s *bigCacheStore) Get(k string) *store.Data {
	data, err := s.db.Get(k)
	if err != nil || len(data) == 0 {
		return nil
	}
	var d *store.Data
	if s.serializer.Deserialize(data, &d) != nil {
		return nil
	}
	return d
}

func (s *bigCacheStore) Set(k string, data []byte, expiration time.Duration) error {
	data, err := s.serializer.Serialize(store.NewData(data, expiration))
	if err != nil {
		return err
	}
	return s.db.Set(k, data)
}

func (s *bigCacheStore) Remove(k string) error {
	err := s.db.Delete(k)
	if errors.Is(err, bigcache.ErrEntryNotFound) {
		return nil
	}
	return err
}

func (s *bigCacheStore) Close() error {
	return s.db.Close()
}
