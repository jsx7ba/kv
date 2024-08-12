package rest

import (
	"encoding/json"
	"io"
	"kv/internal/store"
	"kv/pkg/rest"
	"kv/pkg/watch"
	"log/slog"
	"net/http"
)

type Handlers struct {
	kv store.KVStore
}

func New(kv store.KVStore) *Handlers {
	return &Handlers{
		kv: kv,
	}
}

func (h *Handlers) Put(w http.ResponseWriter, r *http.Request) {
	bytes, err := readBody(r)
	put := rest.PutRequest{}
	err = json.Unmarshal(bytes, &put)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.kv.Put(put.Key, put.Value)
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

	doc := rest.GetResponse{
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
	err := h.kv.Delete(key)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) Watch(w http.ResponseWriter, r *http.Request) {
	slog.Debug("/watch", "host", r.Host)
	watchReq := watch.WatchRequest{}
	b, err := readBody(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(b, &watchReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	updateChan, cancel := h.kv.AddWatch(watchReq.Key, watchReq.WatchType)
	defer cancel()

	w.Header().Set("Cache-Control", "no-cache")
	writeJsonResponse(w, <-updateChan)
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
		slog.Error("failed to write json response", "err", err)
	}
}
