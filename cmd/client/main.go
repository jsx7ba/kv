package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"kv/pkg/client/rpc"
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

	conn, err := grpc.NewClient("127.0.0.1:2000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := rpc.New(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	switch *op {
	case "get":
		value, err := client.Get(ctx, *key)
		checkError(err)
		fmt.Printf("%+v\n", value)
	case "put":
		err := client.Put(ctx, *key, *val)
		checkError(err)
	case "del":
		err := client.Delete(ctx, *key)
		checkError(err)
	case "watch":
		watchType, err := watch.OperationFromString(*watchType)
		checkError(err)

		ch, err := client.Watch(context.Background(), *key, watchType)
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
