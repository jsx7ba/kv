package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kv/cmd/server/rest"
	"kv/cmd/server/rpc"
	"kv/internal/gen"
	"kv/internal/store"
	"kv/internal/store/singlelock"
	"kv/internal/store/watch"
	"log"
	"log/slog"
	"net"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
)

const (
	default_http = "0.0.0.0:2500"
	default_grpc = "0.0.0.0:2510"
)

func main() {

	config := parseFlags()
	err := configureLogging(config.LogLevel)
	checkError(err)
	run(config)
}

func run(config KVConfig) {
	done := make(chan struct{})
	errorChan := make(chan error)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	//kvService := multilock.New(10, watchingKV, multilock.SimpleHashFunc)
	kvService := singlelock.New()
	r, err := bootstrapRaft(config, kvService)
	checkError(err)
	defer r.Shutdown()

	go runGrpc(kvService, done, errorChan, config.GrpcAddr)
	go runHttp(kvService, done, errorChan, config.HttpAddr)

	select {
	case <-errorChan:
		close(done)
		os.Exit(1)
	case <-sigChan:
		close(done)
		os.Exit(1)
	}
}

func watchingKV() store.KVStore {
	return watch.New(singlelock.New())
}

func runGrpc(kv store.KVStore, done chan struct{}, errorChan chan error, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	handlers := rpc.New(kv)
	grpcServer := grpc.NewServer()
	gen.RegisterKVServer(grpcServer, handlers)

	go func() {
		errorChan <- grpcServer.Serve(listener)
	}()

	select {
	case <-errorChan:
		return
	case <-done:
		grpcServer.GracefulStop()
	}
}

func runHttp(kv store.KVStore, done chan struct{}, errorChan chan error, address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer listener.Close()

	h := rest.New(kv)
	http.HandleFunc("POST /kv/", h.Put)
	http.HandleFunc("GET /kv/{key}", h.Get)
	http.HandleFunc("DELETE /kv/{key}", h.Delete)
	http.HandleFunc("POST /watch", h.Watch)

	s := &http.Server{}
	go func() {
		errorChan <- s.Serve(listener)
	}()

	select {
	case <-errorChan:
		return
	case <-done:
		s.Shutdown(context.Background())
	}
}

func configureLogging(logLevel string) error {
	var level slog.Level
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		return err
	}

	h := slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: level})
	slog.SetDefault(slog.New(h))
	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func parseFlags() KVConfig {

	cfg := KVConfig{}

	serveCmd := flag.NewFlagSet("serve", flag.ExitOnError)
	serveCmd.StringVar(&cfg.ClusterAddr, "cluster", "", "The listen address for cluster connections: address:port")
	serveCmd.StringVar(&cfg.NodeId, "id", "", "The raft node id.")
	serveCmd.StringVar(&cfg.HttpAddr, "http", default_http, "The listen address for http client connections: address:port")
	serveCmd.StringVar(&cfg.GrpcAddr, "grpc", default_grpc, "The listen address for grpc client connections: address:port")
	serveCmd.StringVar(&cfg.LogLevel, "loglevel", "INFO", "[DEBUG|INFO|WARN|ERROR]")
	serveCmd.StringVar(&cfg.LogDir, "logfile", "/tmp/kvserve", "Full path to the log directory")

	joinCmd := flag.NewFlagSet("join", flag.ExitOnError)
	joinCmd.StringVar(&cfg.ClusterAddr, "cluster", "", "The listen address for cluster connections: address:port")
	joinCmd.StringVar(&cfg.NodeId, "id", "", "The raft node id.")
	joinCmd.StringVar(&cfg.HttpAddr, "http", default_http, "The listen address for http client connections: address:port")
	joinCmd.StringVar(&cfg.GrpcAddr, "grpc", default_grpc, "The listen address for grpc client connections: address:port")
	joinCmd.StringVar(&cfg.LogLevel, "loglevel", "INFO", "[DEBUG|INFO|WARN|ERROR]")
	joinCmd.StringVar(&cfg.LogDir, "logfile", "/tmp/kvserve", "Full path to the log directory")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "for help try: ./kvserve help\n")
		os.Exit(1)
	}

	var err error
	switch os.Args[1] {
	case "serve":
		cfg.Mode = BootStrap
		err = serveCmd.Parse(os.Args[2:])
	case "join":
		cfg.Mode = Join
		err = joinCmd.Parse(os.Args[2:])
	default:
		flag.Usage()
		serveCmd.Usage()
		joinCmd.Usage()
		os.Exit(1)
	}
	checkError(err)

	slog.Info("", "grpc listen", cfg.GrpcAddr)
	slog.Info("", "http listen", cfg.HttpAddr)

	return cfg
}

func bootstrapRaft(config KVConfig, fsm raft.FSM) (*raft.Raft, error) {
	raftConfig := raft.DefaultConfig()
	raftConfig.LocalID = raft.ServerID(config.NodeId)
	raftConfig.LogLevel = convertLogLevels(config.LogLevel)

	raftTransport := transport.New(raft.ServerAddress(config.ClusterAddr), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})

	nodeDir := filepath.Join(config.LogDir, config.NodeId)

	_, err := os.Stat(nodeDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(nodeDir, 755)
	}
	if err != nil {
		return nil, fmt.Errorf("Unable to create log directory [%s]: %v", config.LogDir, err)
	}

	logFile := filepath.Join(nodeDir, "logs.dat")
	ldb, err := boltdb.NewBoltStore(logFile)
	if err != nil {
		return nil, fmt.Errorf("boltdb.NewBoltStore(%q): %v", logFile, err)
	}

	stableFile := filepath.Join(nodeDir, "stable.dat")
	sdb, err := boltdb.NewBoltStore(stableFile)
	if err != nil {
		return nil, fmt.Errorf("boltdb.NewBoltStore(%q): %v", stableFile, err)
	}

	fss, err := raft.NewFileSnapshotStore(nodeDir, 3, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf(`raft.NewFileSnapshotStore(%q, ...): %v`, nodeDir, err)
	}

	r, err := raft.NewRaft(raftConfig, fsm, ldb, sdb, fss, raftTransport.Transport())
	if err != nil {
		return nil, fmt.Errorf(`raft.NewRaft(%q, ...): %v`, config.LogDir, err)
	}

	if config.Mode == BootStrap {
		servers := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(config.NodeId),
					Address:  raft.ServerAddress(config.ClusterAddr),
				},
			},
		}

		future := r.BootstrapCluster(servers)
		if err := future.Error(); err != nil {
			return nil, err
		}
	}

	return r, nil
}

func convertLogLevels(kvLevel string) string {
	switch kvLevel {
	case "DEBUG":
		return "debug"
	case "INFO":
		return "info"
	case "WARN":
		return "warn"
	case "ERROR":
		return "error"
	default:
		panic("unknown log level " + kvLevel)
	}
}
