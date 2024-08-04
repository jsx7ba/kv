package rpc

import (
	"context"
	"kv/internal/gen"
	"kv/internal/service"
	"kv/pkg/anyval"
	"log/slog"
)

type GRPCHandlers struct {
	gen.UnimplementedKVServer
	kv service.KVService
}

func New(service service.KVService) *GRPCHandlers {
	return &GRPCHandlers{
		UnimplementedKVServer: gen.UnimplementedKVServer{},
		kv:                    service,
	}
}

func (s *GRPCHandlers) Put(_ context.Context, r *gen.PutRequest) (*gen.Response, error) {
	response := &gen.Response{Status: gen.Status_ERROR}
	value, err := anyval.Unmarshal(r.GetValue())
	if err != nil {
		slog.Error("put", "key", r.Key, "error", err)
		return response, err
	}

	err = s.kv.Put(r.Key, value)

	if err != nil {
		slog.Error("put", "key", r.Key, "value", value)
		return response, err
	}

	response.Status = gen.Status_OK
	return response, err
}

func (s *GRPCHandlers) Get(_ context.Context, r *gen.GetRequest) (*gen.GetResponse, error) {
	response := &gen.GetResponse{Status: gen.Status_ERROR}
	v, err := s.kv.Get(r.Key)

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

func (s *GRPCHandlers) Delete(_ context.Context, r *gen.DeleteRequest) (*gen.Response, error) {
	err := s.kv.Delete(r.Key)

	status := gen.Status_OK
	if err != nil {
		status = gen.Status_ERROR
	}
	return &gen.Response{Status: status}, nil
}
