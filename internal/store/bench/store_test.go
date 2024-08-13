package bench

import (
	"bytes"
	"fmt"
	"kv/internal/store"
	"kv/internal/store/multilock"
	"kv/internal/store/singlelock"
	"kv/internal/store/syncmap"
	"log"
	"math/rand"
	"sync"
	"testing"
)

var (
	data       [][]string
	iterations = []int{100, 1000, 10_000, 100_000, 1_000_000}
)

func init() {
	size := 100_000
	if data == nil {
		r := rand.New(rand.NewSource(3))
		data = make([][]string, size)
		for i := range size {
			data[i] = make([]string, 2)
			data[i][0] = randomString(r, 5, 124)
			data[i][0] = randomString(r, 100, 100_000)
		}
	}
}

func randomString(r *rand.Rand, min, max uint32) string {
	keyLen := (r.Uint32() % (max - min)) + min
	buffer := bytes.Buffer{}
	for range keyLen {
		b := (r.Uint32() % 93) + 33 // random unsigned int between 33 and 126
		buffer.WriteByte(byte(b))
	}
	return buffer.String()
}

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
