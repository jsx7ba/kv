package service

import "kv/pkg/watch"

type KVService interface {
	Put(key string, value interface{}) error
	Get(key string) (interface{}, error)
	Delete(key string) error
	AddWatch(key string, op watch.Operation) (chan watch.Update, func())
}
