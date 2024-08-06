package bench

import (
	"fmt"
	"kv/internal/store"
	"kv/internal/store/multilock"
	"kv/internal/store/singlelock"
	"kv/internal/store/syncmap"
	"log"
	"sync"
	"testing"
)

var (
	data = [][]string{
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
		{"plough", "representative"},
		{"brush", "hope"},
		{"oatmeal", "sun"},
		{"jeans", "burn"},
		{"fight", "vest"},
		{"self", "ice"},
		{"disease", "brake"},
		{"friends", "texture"},
		{"bulb", "push"},
		{"thing", "scale"},
		{"cattle", "cakes"},
		{"pickle", "purpose"},
		{"cars", "crime"},
		{"cannon", "class"},
		{"skirt", "sleet"},
		{"pen", "partner"},
		{"rhythm", "root"},
		{"noise", "name"},
		{"pot", "protest"},
		{"surprise", "song"},
	}

	iterations = []int{100, 1000, 10000}
)

func basicKV() store.KVStore {
	return singlelock.New()
}

func syncmapKV() store.KVStore {
	return syncmap.New()
}

func BenchmarkSyncMap(b *testing.B) {
	kv := syncmap.New()
	for _, count := range iterations {
		b.Run(fmt.Sprintf("iter-%d", count), func(b *testing.B) {
			benchmarkDriver(kv, b)
		})
	}
}

func BenchmarkBasicKVService(b *testing.B) {
	kv := basicKV()
	for _, count := range iterations {
		b.Run(fmt.Sprintf("iter-%d", count), func(b *testing.B) {
			benchmarkDriver(kv, b)
		})
	}
}

func BenchmarkMultikeySyncMapKVService(b *testing.B) {
	mkv := multilock.New(16, syncmapKV, multilock.SimpleHashFunc)

	for _, count := range iterations {
		b.Run(fmt.Sprintf("iter-%d", count), func(b *testing.B) {
			benchmarkDriver(mkv, b)
		})
	}
}

func BenchmarkMultikeyService(b *testing.B) {
	mkv := multilock.New(16, basicKV, multilock.SimpleHashFunc)

	for _, count := range iterations {
		b.Run(fmt.Sprintf("iter-%d", count), func(b *testing.B) {
			benchmarkDriver(mkv, b)
		})
	}
}

func benchmarkDriver(kv store.KVStore, b *testing.B) {
	wg := sync.WaitGroup{}
	wg.Add(b.N * 2)

	for i := 0; i != b.N; i++ {
		go concurrentPut(kv, i, &wg)
		go concurrentGet(kv, i, &wg)
	}
	wg.Wait()
}

func concurrentPut(kv store.KVStore, idx int, wg *sync.WaitGroup) {
	di := idx % len(data)
	err := kv.Put(data[di][0], data[di][1])
	if err != nil {
		log.Fatal("put failed", err)
	}
	wg.Done()
}

func concurrentGet(kv store.KVStore, idx int, wg *sync.WaitGroup) {
	di := idx % len(data)
	_, _ = kv.Get(data[di][0])
	wg.Done()
}
