package rpc

import (
	"context"
	"google.golang.org/protobuf/types/known/anypb"
	"kv/internal/gen"
	"kv/internal/service"
	"kv/pkg/anyval"
	"kv/pkg/watch"
	"log/slog"
)

type Handlers struct {
	gen.UnimplementedKVServer
	kv service.KVService
	// TODO: add done channel
}

func New(service service.KVService) *Handlers {
	return &Handlers{
		UnimplementedKVServer: gen.UnimplementedKVServer{},
		kv:                    service,
	}
}

func (h *Handlers) Put(_ context.Context, r *gen.PutRequest) (*gen.Response, error) {
	response := &gen.Response{Status: gen.Status_ERROR}
	value, err := anyval.Unmarshal(r.GetValue())
	if err != nil {
		slog.Error("put", "key", r.Key, "error", err)
		return response, err
	}

	err = h.kv.Put(r.Key, value)

	if err != nil {
		slog.Error("put", "key", r.Key, "value", value)
		return response, err
	}

	response.Status = gen.Status_OK
	return response, err
}

func (h *Handlers) Get(_ context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	response := &gen.GetResponse{Status: gen.Status_ERROR}
	v, err := h.kv.Get(r.Key)

	if err == nil {
		av, err := anyval.Marshal(v)
		if err == nil {
			response.Value = av
			response.Status = gen.Status_OK
			slog.Debug("get", "key", r.Key, "found", true)
		}
	} else {
		slog.Debug("get", "key", r.Key, "found", false)
	}

	return response, err
}

func (h *Handlers) Delete(_ context.Context, r *gen.DeleteRequest) (*gen.Response, error) {
	err := h.kv.Delete(r.Key)

	status := gen.Status_OK
	if err != nil {
		status = gen.Status_ERROR
	}
	return &gen.Response{Status: status}, nil
}

func (h *Handlers) Watch(r *gen.WatchRequest, server gen.KV_WatchServer) error {
	watchChan, cancelFunc := h.kv.AddWatch(r.Key, watch.Operation(r.GetWatchType()))
	defer cancelFunc()

	var err error
out:
	for {
		select {
		case update := <-watchChan:
			var v *anypb.Any
			if update.Value != nil {
				v, err = anyval.Marshal(update.Value)
				if err != nil {
					break out
				}
			}

			m := gen.WatchResponse{
				WatchType: update.Op.Convert(),
				Key:       update.Key,
				Value:     v,
			}

			err = server.SendMsg(&m)
			if err != nil {
				slog.Error("sending watch response failed", "err", err)
				break out
			}
		}
	}
	return err
}
