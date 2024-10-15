package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kv/pkg/client"
	"kv/pkg/watch"
	"log"
	"os"
	"strings"
	"time"
)

// ./kvclient put grpc://127.0.0.0:2510 -k asdf -v asdf

var (
	host      = flag.String("h", os.Getenv("KV_HOST"), "[http|grpc]://host:port")
	op        = flag.String("op", "", "[get|put|del|watch]")
	key       = flag.String("k", "", "key name")
	val       = flag.String("v", "", "value")
	watchType = flag.String("t", "", "watch type")
)

const (
	default_http = "127.0.0.1:2500"
	default_grpc = "127.0.0.1:2510"
)

func main() {
	flag.Parse()
	kv, cancel, err := configureTransport(*host)
	checkError(err)
	defer cancel()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *op {
	case "get":
		value, err := kv.Get(ctx, *key)
		checkError(err)
		fmt.Printf("%+v\n", value)
	case "put":
		err := kv.Put(ctx, *key, *val)
		checkError(err)
	case "del":
		err := kv.Delete(ctx, *key)
		checkError(err)
	case "watch":
		watchType, err := watch.OperationFromString(*watchType)
		checkError(err)

		ch, err := kv.Watch(context.Background(), *key, watchType)
		checkError(err)

		for update := range ch {
			log.Printf("%+v", update)
		}
	default:
		fmt.Fprintf(os.Stderr, "error parsing command line: %s\n", *op)
		os.Exit(1)
	}

	os.Exit(0)
}

func checkError(err error) {
	if err == nil {
		return
	}
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}

func configureTransport(hostUrl string) (client.KV, func(), error) {
	if len(hostUrl) == 0 {
		hostUrl = default_grpc
		log.Printf("using %s\n", hostUrl)
	}

	parts := strings.Split(hostUrl, "://")
	if len(parts) != 2 {
		return nil, func() {}, errors.New("invalid transport: " + hostUrl)
	}

	transport := parts[0]
	address := parts[1]

	var kv client.KV
	cancel := func() {}
	var err error

	if transport == "http" {
		kv = client.NewRest(address)
	} else if transport == "grpc" {
		conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		cancel = func() {
			conn.Close()
		}
		kv = client.NewGRPC(conn)
	} else {
		err = errors.New("unknown transport: " + transport)
	}
	return kv, cancel, err

}
