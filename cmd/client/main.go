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
	"time"
)

var (
	op        = flag.String("op", "", "[get|put|del|watch]")
	key       = flag.String("k", "", "key name")
	val       = flag.String("v", "", "value")
	watchType = flag.String("t", "", "watch type")
)

func main() {
	flag.Parse()

	kv, cancel, err := configureTransport()
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
		fmt.Fprintf(os.Stderr, "error parsing command line\n")
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

func configureTransport() (client.KV, func(), error) {
	name := os.Getenv("KV_TRANSPORT")
	var kv client.KV
	cancel := func() {}
	var err error

	if name == "rest" {
		kv = client.NewRest("127.0.0.1:2500")
	} else if len(name) == 0 || name == "grpc" {
		conn, err := grpc.NewClient("127.0.0.1:2000", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		cancel = func() {
			conn.Close()
		}
		kv = client.NewGRPC(conn)
	} else {
		err = errors.New("unknown transport: " + name)
	}
	return kv, cancel, err

}
