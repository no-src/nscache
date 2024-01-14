package memory

import (
	"sync"
	"time"

	"github.com/no-src/nscache/store"
)

type memoryStore struct {
	data map[string]*store.Data
	mu   sync.RWMutex
}

func newStore() store.Store {
	return &memoryStore{
		data: make(map[string]*store.Data),
	}
}

func (s *memoryStore) Get(k string) *store.Data {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.data[k]
}

func (s *memoryStore) Set(k string, data []byte, expiration time.Duration) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[k] = store.NewData(data, expiration)
	return nil
}

func (s *memoryStore) Remove(k string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.data, k)
	return nil
}

func (s *memoryStore) Close() error {
	return nil
}
