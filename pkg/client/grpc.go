package client

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"kv/internal/gen"
	"kv/pkg/anyval"
	"kv/pkg/watch"
	"log/slog"
)

type GPRCClient struct {
	kvc gen.KVClient
}

func NewGRPC(conn *grpc.ClientConn) *GPRCClient {
	return &GPRCClient{
		kvc: gen.NewKVClient(conn),
	}
}

func (c *GPRCClient) Get(ctx context.Context, key string) (interface{}, error) {
	req := gen.GetRequest{
		Key: key,
	}

	var remoteVal interface{}
	r, err := c.kvc.Get(ctx, &req)
	if err != nil {
		return nil, err
	} else {
		if r.Status != gen.Status_OK {
			err = errors.New(fmt.Sprintf("key [%s] not found\n", key))
		} else {
			remoteVal, _ = anyval.Unmarshal(r.GetValue())
		}
	}

	return remoteVal, err
}

func (c *GPRCClient) Put(ctx context.Context, key string, val interface{}) error {
	anyVal, err := anyval.Marshal(val)
	if err != nil {
		return err
	}

	req := gen.PutRequest{
		Key:   key,
		Value: anyVal,
	}

	r, err := c.kvc.Put(ctx, &req)
	if err != nil {
		return err
	}

	if r.Status != gen.Status_OK {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}

	return err
}

func (c *GPRCClient) Delete(ctx context.Context, key string) error {
	req := gen.DeleteRequest{
		Key: key,
	}

	r, err := c.kvc.Delete(ctx, &req)
	if err != nil {
		return err
	} else if r.Status != gen.Status_OK {
		err = errors.New(fmt.Sprintf("key [%s] not found", key))
	}

	return err
}

func (c *GPRCClient) Watch(ctx context.Context, key string, operation watch.Operation) (chan watch.Update, error) {
	req := gen.WatchRequest{
		Key:       key,
		WatchType: operation.Convert(),
	}

	watchClient, err := c.kvc.Watch(ctx, &req)
	if err != nil {
		return nil, err
	}

	updateChan := make(chan watch.Update)

	go func() {
		response := gen.WatchResponse{}
		for {
			err := watchClient.RecvMsg(&response)
			if err != nil {
				close(updateChan)
				break
			}

			v, err := anyval.Unmarshal(response.Value)
			if err != nil {
				slog.Error("unmarshal value failed", "err", err)
				continue
			}

			updateChan <- watch.Update{
				Op:    watch.OperationFrom(response.WatchType),
				Key:   response.Key,
				Value: v,
			}

			response.Reset()
		}
	}()

	return updateChan, nil
}
