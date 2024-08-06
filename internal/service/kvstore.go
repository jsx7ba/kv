package service

import "kv/pkg/watch"

// KVStore defines the methods for key value store.
type KVStore interface {
	// Put stores a value in the KV store.
	Put(key string, value interface{}) error

	// Get retrieves a value from the KV store.
	Get(key string) (interface{}, error)

	// Delete removes a key from the KV store.
	Delete(key string) error

	// AddWatch sends updates to values over the returned channel.
	// Call the cancel function when updates are no longer needed.
	AddWatch(key string, op watch.Operation) (chan watch.Update, func())
}
