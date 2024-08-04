package rpc

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"kv/internal/gen"
	"kv/pkg/anyval"
)

type Client struct {
	kvc gen.KVClient
}

func New(conn *grpc.ClientConn) *Client {
	return &Client{
		kvc: gen.NewKVClient(conn),
	}
}

func (c *Client) Get(ctx context.Context, key string) (interface{}, error) {
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

func (c *Client) Put(ctx context.Context, key string, val interface{}) error {
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

func (c *Client) Delete(ctx context.Context, key string) error {
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

//func (c *GRPCClient) Watch(ctx context.Context, key string) error {
//
//}
