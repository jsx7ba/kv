package multilock

import (
	"kv/internal/service"
	"kv/internal/service/singlelock"
	w "kv/internal/service/watch"
	"kv/pkg/watch"
	"testing"
)

func basicKV() service.KVService {
	return singlelock.New()
}

func TestRoundTrip(t *testing.T) {
	data := [][]string{
		{"encouraging", "texture"},
		{"ashamed", "toes"},
		{"abortive", "badge"},
		{"well-to-do", "quiver"},
		{"zonked", "step"},
		{"cute", "idea"},
		{"windy", "meal"},
		{"complete", "snail"},
		{"protective", "year"},
		{"stimulating", "dirt"},
	}

	mkv := New(13, basicKV, SimpleHashFunc)
	for i := range data {
		err := mkv.Put(data[i][0], data[i][1])
		if err != nil {
			t.Error(err)
		}
	}

	for i := range data {
		value, err := mkv.Get(data[i][0])
		if err != nil {
			t.Error(err)
		}
		if data[i][1] != value.(string) {
			t.Errorf("expected [%s], got [%s]", data[i][1], value)
		}
	}
}

func TestDelete(t *testing.T) {
	mkv := New(13, basicKV, SimpleHashFunc)
	key := "a key"
	value := "a value"
	err := mkv.Put(key, value)
	if err != nil {
		t.Error(err)
	}

	err = mkv.Delete(key)
	if err != nil {
		t.Error(err)
	}

	_, err = mkv.Get(key)
	if err == nil {
		t.Error("Get should have errored, but did not")
	}
}

func watchingKV() service.KVService {
	return w.New(basicKV())
}

func TestWatch(t *testing.T) {
	expected := "an expected value"
	done := make(chan struct{})
	mkv := New(13, watchingKV, SimpleHashFunc)
	ch, cancel := mkv.AddWatch("foo", watch.Put)
	defer cancel()
	go func() {
		v := <-ch
		if expected != v.Value.(string) {
			t.Errorf("expected [%s], got[%s]", expected, v.Value)
		}
		close(done)
	}()

	err := mkv.Put("foo", expected)
	if err != nil {
		t.Error(err)
	}
	<-done
}
