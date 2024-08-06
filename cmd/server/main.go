package main

import (
	"context"
	"golang.org/x/exp/slog"
	"google.golang.org/grpc"
	"kv/cmd/server/rest"
	"kv/cmd/server/rpc"
	"kv/internal/gen"
	"kv/internal/store"
	"kv/internal/store/multilock"
	"kv/internal/store/singlelock"
	"kv/internal/store/watch"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
)

func main() {
	done := make(chan struct{})
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)

	configureLogging()
	kvService := multilock.New(10, watchingKV, multilock.SimpleHashFunc)
	go runGrpc(kvService, done, "127.0.0.1:2000")
	go runHttp(kvService, done, "127.0.0.1:2500")

	select {
	case <-sigChan:
		close(done)
		os.Exit(1)
	}
}

func watchingKV() store.KVStore {
	return watch.New(singlelock.New())
}

func runGrpc(kv store.KVStore, done chan struct{}, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	handlers := rpc.New(kv)
	grpcServer := grpc.NewServer()
	gen.RegisterKVServer(grpcServer, handlers)

	errChan := make(chan error)
	go func() {
		errChan <- grpcServer.Serve(listener)
	}()

	select {
	case <-errChan:
		return
	case <-done:
		grpcServer.GracefulStop()
	}
}

func runHttp(kv store.KVStore, done chan struct{}, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	h := rest.New(kv)
	http.HandleFunc("POST /kv/{key}", h.Put)
	http.HandleFunc("GET /kv/{key}", h.Get)
	http.HandleFunc("DELETE /kv/{key}", h.Delete)

	errChan := make(chan error)
	s := &http.Server{}
	go func() {
		errChan <- s.Serve(listener)
	}()

	select {
	case <-errChan:
		return
	case <-done:
		s.Shutdown(context.Background())
	}
}

func configureLogging() {
	env, exists := os.LookupEnv("KV_LOGLEVEL")
	if !exists {
		return
	}

	level, err := strconv.Atoi(env)
	if err != nil {
		log.Fatal("unable to parse KV_LOGLEVEL", err)
	}
	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.Level(level)})
	slog.SetDefault(slog.New(h))
}
