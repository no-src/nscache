package fastcache

import (
	"time"

	"github.com/VictoriaMetrics/fastcache"
	"github.com/no-src/nscache/encoding"
	"github.com/no-src/nscache/store"
)

type fastCacheStore struct {
	db         *fastcache.Cache
	serializer encoding.Serializer
}

func newStore(db *fastcache.Cache, serializer encoding.Serializer) store.Store {
	return &fastCacheStore{
		db:         db,
		serializer: serializer,
	}
}

func (s *fastCacheStore) Get(k string) *store.Data {
	data := s.db.Get(nil, []byte(k))
	if len(data) == 0 {
		return nil
	}
	var d *store.Data
	if s.serializer.Deserialize(data, &d) != nil {
		return nil
	}
	return d
}

func (s *fastCacheStore) Set(k string, data []byte, expiration time.Duration) error {
	data, err := s.serializer.Serialize(store.NewData(data, expiration))
	if err != nil {
		return err
	}
	s.db.Set([]byte(k), data)
	return nil
}

func (s *fastCacheStore) Remove(k string) error {
	s.db.Del([]byte(k))
	return nil
}

func (s *fastCacheStore) Close() error {
	return nil
}
