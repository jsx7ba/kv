package client

import (
	"context"
	"kv/pkg/watch"
)

// KV defines methods for key value client implementations.
type KV interface {
	Get(ctx context.Context, key string) (interface{}, error)
	Put(ctx context.Context, key string, val interface{}) error
	Delete(ctx context.Context, key string) error
	Watch(ctx context.Context, key string, operation watch.Operation) (chan watch.Update, error)
}
