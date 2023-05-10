package memory

import (
	"time"

	"github.com/no-src/nscache/store"
)

type memoryStore struct {
	data map[string]*store.Data
}

func newStore() store.Store {
	return &memoryStore{
		data: make(map[string]*store.Data),
	}
}

func (s *memoryStore) Get(k string) *store.Data {
	return s.data[k]
}

func (s *memoryStore) Set(k string, data []byte, expiration time.Duration) error {
	s.data[k] = store.NewData(data, expiration)
	return nil
}

func (s *memoryStore) Remove(k string) error {
	delete(s.data, k)
	return nil
}

func (s *memoryStore) Close() error {
	return nil
}
