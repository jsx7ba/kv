package watch

import (
	"kv/internal/store"
	"kv/pkg/watch"
	"log/slog"
	"sync"
)

type set map[chan watch.Update]struct{}

func (s set) put(c chan watch.Update) {
	s[c] = struct{}{}
}

func (s set) del(c chan watch.Update) {
	delete(s, c)
}

func (s set) empty() bool {
	return len(s) == 0
}

func (s set) forEach(f func(c chan watch.Update)) {
	for k, _ := range s {
		f(k)
	}
}

// KVStoreWatcher wraps a KVStore, and sends changes to channels associated with a specific key and operation.
type KVStoreWatcher struct {
	lock          sync.RWMutex
	service       store.KVStore
	subscriptions map[subscription]set
}

type subscription struct {
	Key string
	Op  watch.Operation
}

func New(service store.KVStore) store.KVStore {
	return &KVStoreWatcher{
		lock:          sync.RWMutex{},
		service:       service,
		subscriptions: make(map[subscription]set),
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
	if err == nil {
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

	updateFunc := func(c chan watch.Update) {
		c <- update
	}

	if watchers, ok := s.subscriptions[sub]; ok {
		watchers.forEach(updateFunc)
	}
}

func (s *KVStoreWatcher) removeWatcher(sub subscription, c chan watch.Update) {
	s.lock.Lock()
	defer s.lock.Unlock()
	slog.Info("removing watch", "sub", sub)
	if watchers, ok := s.subscriptions[sub]; ok {
		watchers.del(c)
		close(c)
		if watchers.empty() {
			delete(s.subscriptions, sub)
		}
	}
}

func (s *KVStoreWatcher) AddWatch(key string, op watch.Operation) (chan watch.Update, func()) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if op != watch.All {
		return s.watchImpl(key, op)
	}

	// make two subscriptions, and combine the channels and cancel func
	updateChan := make(chan watch.Update)
	done := make(chan struct{})
	putChan, putCancel := s.watchImpl(key, watch.Put)
	delChan, delCancel := s.watchImpl(key, watch.Delete)
	go func() {
		for {
			select {
			case <-done:
				return
			case u := <-putChan:
				updateChan <- u
			case u := <-delChan:
				updateChan <- u
			}
		}
	}()

	cancelFunc := func() {
		close(done)
		putCancel()
		delCancel()
	}
	return updateChan, cancelFunc
}

func (s *KVStoreWatcher) watchImpl(key string, op watch.Operation) (chan watch.Update, func()) {
	sub := subscription{
		Key: key,
		Op:  op,
	}
	updateChan := make(chan watch.Update)
	if _, ok := s.subscriptions[sub]; !ok {
		s.subscriptions[sub] = make(set)
	}
	s.subscriptions[sub].put(updateChan)
	return updateChan, func() { s.removeWatcher(sub, updateChan) }
}
