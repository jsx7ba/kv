package syncmap

import (
	"errors"
	"fmt"
	"kv/pkg/watch"
	"sync"
)

type KVSyncMap struct {
	store sync.Map
}

func New() *KVSyncMap {
	return &KVSyncMap{
		store: sync.Map{},
	}
}

func (kv *KVSyncMap) Put(key string, value interface{}) error {
	kv.store.Store(key, value)
	return nil
}

func (kv *KVSyncMap) Get(key string) (interface{}, error) {
	var err error
	value, ok := kv.store.Load(key)
	if !ok {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}
	return value, err
}

func (kv *KVSyncMap) Delete(key string) error {
	kv.store.Delete(key)
	return nil
}

func (kv *KVSyncMap) AddWatch(key string, op watch.Operation) (chan watch.Update, func()) {
	panic("watch not implemented")
}
