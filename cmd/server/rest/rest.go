package rest

import (
	"encoding/json"
	"io"
	"kv/internal/service"
	"log/slog"
	"net/http"
)

type Handlers struct {
	kv service.KVStore
}

func New(kv service.KVStore) *Handlers {
	return &Handlers{
		kv: kv,
	}
}

func (h *Handlers) Put(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bytes, err := readBody(r)
	put := PutRequest{}
	err = json.Unmarshal(bytes, &put)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.kv.Put(key, put.Value)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Get(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	value, err := h.kv.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	doc := GetResponse{
		Key:   key,
		Value: value,
	}

	writeJsonResponse(w, doc)
}

func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	key := r.PathValue("key")
	if len(key) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func readBody(r *http.Request) ([]byte, error) {
	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	return b, err
}

func writeJsonResponse(w http.ResponseWriter, value interface{}) {
	bytes, err := json.Marshal(value)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(bytes)
	if err != nil {
		slog.Error("failed to write json response", err)
	}
}

type PutRequest struct {
	Value interface{} `json:"value"`
}

type GetResponse struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

type GetRequest struct {
	Key string `json:"key"`
}

type DeleteRequest struct {
	Key string `json:"key"`
}
