package multilock

import (
	"kv/internal/service"
	"kv/pkg/watch"
)

// MultiKVStore Implements KVService and divides the keyspace.  This will prevent one write from locking
// the whole service.
type MultiKVStore struct {
	size   int
	store  []service.KVService
	hasher func(buckets int, key string) int
}

func New(bucketCount int, factory func() service.KVService, hashFunc func(buckets int, key string) int) *MultiKVStore {
	buckets := make([]service.KVService, bucketCount)
	for i := 0; i != bucketCount; i++ {
		buckets[i] = factory()
	}
	return &MultiKVStore{
		size:   bucketCount,
		store:  buckets,
		hasher: hashFunc,
	}
}

func SimpleHashFunc(buckets int, key string) int {
	return int(key[0]) % buckets
}

func (m *MultiKVStore) Put(key string, value interface{}) error {
	idx := m.hasher(m.size, key)
	return m.store[idx].Put(key, value)
}

func (m *MultiKVStore) Get(key string) (interface{}, error) {
	idx := m.hasher(m.size, key)
	return m.store[idx].Get(key)
}

func (m *MultiKVStore) Delete(key string) error {
	idx := m.hasher(m.size, key)
	return m.store[idx].Delete(key)
}

func (m *MultiKVStore) AddWatch(key string, op watch.Operation) (chan watch.Update, func()) {
	idx := m.hasher(m.size, key)
	return m.store[idx].AddWatch(key, op)
}
