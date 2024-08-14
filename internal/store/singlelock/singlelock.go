package singlelock

import (
	"errors"
	"fmt"
	"kv/pkg/watch"
	"sync"
)

// KVSingleLockMap uses a single lock for the whole key value store.
type KVSingleLockMap struct {
	m     sync.RWMutex
	store map[string]interface{}
}

func New() *KVSingleLockMap {
	return &KVSingleLockMap{
		m:     sync.RWMutex{},
		store: make(map[string]interface{}),
	}
}

func (kv *KVSingleLockMap) Put(key string, value interface{}) error {
	kv.m.Lock()
	defer kv.m.Unlock()
	kv.store[key] = value
	return nil
}

func (kv *KVSingleLockMap) Get(key string) (interface{}, error) {
	kv.m.RLock()
	defer kv.m.RUnlock()
	v, ok := kv.store[key]

	var err error
	if !ok {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}

	return v, err
}

func (kv *KVSingleLockMap) Delete(key string) error {
	kv.m.Lock()
	defer kv.m.Unlock()
	_, ok := kv.store[key]
	var err error
	if ok {
		delete(kv.store, key)
	} else {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}
	return err
}

func (kv *KVSingleLockMap) AddWatch(_ string, _ watch.Operation) (chan watch.Update, func()) {
	panic("not implemented")
}
