package singlelock

import (
	"errors"
	"fmt"
	"kv/pkg/watch"
	"sync"
)

type KVService struct {
	m     sync.RWMutex
	store map[string]interface{}
}

func New() *KVService {
	return &KVService{
		m:     sync.RWMutex{},
		store: make(map[string]interface{}),
	}
}

func (kv *KVService) Put(key string, value interface{}) error {
	kv.m.Lock()
	defer kv.m.Unlock()
	kv.store[key] = value
	return nil
}

func (kv *KVService) Get(key string) (interface{}, error) {
	kv.m.RLock()
	defer kv.m.RUnlock()
	v, ok := kv.store[key]

	var err error
	if !ok {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}

	return v, err
}

func (kv *KVService) Delete(key string) error {
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

func (kv *KVService) AddWatch(_ string, _ watch.Operation) (chan watch.Update, func()) {
	panic("not implemented")
}
