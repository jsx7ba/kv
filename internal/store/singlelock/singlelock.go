package singlelock

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/raft"
	"io"
	"kv/internal/store"
	"kv/pkg/watch"
	"maps"
	"sync"
)

type fsmSnapshot struct {
	data map[string]interface{}
}

func (s *fsmSnapshot) Persist(sink raft.SnapshotSink) error {
	persistErr := func() error {
		b, err := json.Marshal(s.data)
		if err != nil {
			return err
		}

		if _, err := sink.Write(b); err != nil {
			return err
		}

		// Close the sink.
		return sink.Close()
	}()

	if persistErr != nil {
		sink.Cancel()
	}

	return persistErr
}

func (s *fsmSnapshot) Release() {}

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

// applyResponse is expected to be returned from The Apply method when an error occurs.
type applyResponse struct {
	index uint64
	err   error
}

func (r applyResponse) Index() uint64 {
	return r.index
}

func (r applyResponse) Response() interface{} {
	return r.err
}

func (kv *KVSingleLockMap) Apply(log *raft.Log) interface{} {
	update := &store.Entry{}
	err := json.Unmarshal(log.Data, &update)

	if err != nil {
		return applyResponse{
			index: log.Index,
			err:   err,
		}
	}

	switch update.Action {
	case store.Put:
		err = kv.Put(update.Key, update.Value)
	case store.Delete:
		err = kv.Delete(update.Key)
	}

	if err != nil {
		return applyResponse{
			index: log.Index,
			err:   err,
		}
	}

	return nil
}

func (kv *KVSingleLockMap) Snapshot() (raft.FSMSnapshot, error) {
	kv.m.Lock()
	defer kv.m.Unlock()
	data := maps.Clone(kv.store)
	return &fsmSnapshot{data: data}, nil
}

func (kv *KVSingleLockMap) Restore(snapshot io.ReadCloser) error {
	kv.m.Lock()
	defer kv.m.Unlock()
	defer snapshot.Close()
	kv.store = make(map[string]interface{})

	b, err := io.ReadAll(snapshot)
	if err != nil {
		return err
	}

	snap := fsmSnapshot{}
	err = json.Unmarshal(b, &snap)
	if err != nil {
		return err
	}

	maps.Copy(kv.store, snap.data)
	return nil
}
