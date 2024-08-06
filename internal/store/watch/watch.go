package watch

import (
	"kv/internal/store"
	"kv/pkg/watch"
	"log/slog"
	"sync"
)

// KVStoreWatcher wraps a KVStore, and sends changes to channels associated with a specific key and operation.
type KVStoreWatcher struct {
	lock          sync.RWMutex
	service       store.KVStore
	subscriptions map[subscription][]chan watch.Update
}

type subscription struct {
	Key string
	Op  watch.Operation
}

func New(service store.KVStore) store.KVStore {
	return &KVStoreWatcher{
		lock:          sync.RWMutex{},
		service:       service,
		subscriptions: make(map[subscription][]chan watch.Update),
	}
}

func (s *KVStoreWatcher) Put(key string, value interface{}) error {
	err := s.service.Put(key, value)
	if err == nil {
		update := watch.Update{
			Key:   key,
			Op:    watch.Put,
			Value: value,
		}
		go s.updateWatchers(key, update)
	}
	return err
}

func (s *KVStoreWatcher) Get(key string) (interface{}, error) {
	return s.service.Get(key)
}

func (s *KVStoreWatcher) Delete(key string) error {
	err := s.service.Delete(key)
	if err != nil {
		update := watch.Update{
			Key: key,
			Op:  watch.Delete,
		}
		go s.updateWatchers(key, update)
	}
	return err
}

func (s *KVStoreWatcher) updateWatchers(key string, update watch.Update) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	sub := subscription{
		Key: key,
		Op:  update.Op,
	}

	if watchers, ok := s.subscriptions[sub]; ok {
		for _, w := range watchers {
			w <- update
		}
	}
}

func (s *KVStoreWatcher) removeWatcher(sub subscription, index int) {
	s.lock.Lock()
	defer s.lock.Unlock()
	slog.Info("removing watch", "sub", sub)
	if watchers, ok := s.subscriptions[sub]; ok {
		if len(watchers)-1 <= index {
			close(watchers[index])
			watchers = append(watchers[:index], watchers[index+1:]...)
			if len(watchers) == 0 {
				delete(s.subscriptions, sub)
			}
		}
	}
}

// TODO: handle the All operation
func (s *KVStoreWatcher) AddWatch(key string, op watch.Operation) (chan watch.Update, func()) {
	s.lock.Lock()
	defer s.lock.Unlock()

	sub := subscription{
		Key: key,
		Op:  op,
	}

	updateChan := make(chan watch.Update)
	if _, ok := s.subscriptions[sub]; !ok {
		s.subscriptions[sub] = make([]chan watch.Update, 1)
		s.subscriptions[sub][0] = updateChan
	} else {
		s.subscriptions[sub] = append(s.subscriptions[sub], updateChan)
	}

	idx := len(s.subscriptions) - 1
	return updateChan, func() { s.removeWatcher(sub, idx) }
}
