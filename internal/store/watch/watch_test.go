package watch

import (
	"kv/internal/store"
	"kv/internal/store/singlelock"
	"kv/pkg/watch"
	"testing"
)

func makeKVStore() store.KVStore {
	return New(singlelock.New())
}

func TestKVStoreWatcher_Put(t *testing.T) {
	key := "foo"
	value := "bar"

	kv := makeKVStore()
	updateChan, cancel := kv.AddWatch(key, watch.Put)
	defer cancel()

	var actual watch.Update
	done := make(chan struct{})
	go func() {
		actual = <-updateChan
		done <- struct{}{}
	}()

	err := kv.Put(key, value)
	if err != nil {
		t.Error("put failed", err)
	}

	<-done

	if actual.Op != watch.Put {
		t.Errorf("operation should have been [%v], but was [%v]", watch.Put, actual.Op)
	} else if actual.Value != value {
		t.Errorf("expected value [%s] got [%s]", value, actual.Value)
	}
}

func TestKVStoreWatcher_Delete(t *testing.T) {
	key := "foo"
	value := "bar"

	kv := makeKVStore()
	updateChan, cancel := kv.AddWatch(key, watch.Delete)
	defer cancel()

	done := make(chan struct{})
	var actual watch.Update
	go func() {
		actual = <-updateChan
		close(done)
	}()

	err := kv.Put(key, value)
	if err != nil {
		t.Error("put failed", err)
	}

	err = kv.Delete(key)
	if err != nil {
		t.Error("delete failed", err)
	}

	<-done

	if actual.Op != watch.Delete {
		t.Errorf("operation should have been [%v], but was [%v]", watch.Delete, actual.Op)
	} else if actual.Value != nil {
		t.Errorf("expected value [nil] got [%s]", actual.Value)
	}
}

func TestKVStoreWatcher_All(t *testing.T) {
	key := "foo"
	value := "bar"

	kv := makeKVStore()
	updateChan, cancel := kv.AddWatch(key, watch.All)
	defer cancel()

	ready := make(chan struct{})
	var actual watch.Update
	go func() {
		actual = <-updateChan
		ready <- struct{}{}
	}()

	err := kv.Put(key, value)
	if err != nil {
		t.Error("put failed", err)
	}
	<-ready

	if actual.Op != watch.Put {
		t.Errorf("operation should have been [%v], but was [%v]", watch.Put, actual.Op)
	} else if actual.Value != value {
		t.Errorf("expected value [%s] got [%s]", value, actual.Value)
	}

	go func() {
		actual = <-updateChan
		ready <- struct{}{}
	}()

	err = kv.Delete(key)
	if err != nil {
		t.Error("delete failed", err)
	}

	<-ready

	if actual.Op != watch.Delete {
		t.Errorf("operation should have been [%v], but was [%v]", watch.Delete, actual.Op)
	} else if actual.Value != nil {
		t.Errorf("expected value [nil] got [%s]", actual.Value)
	}
}
