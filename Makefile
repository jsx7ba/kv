.PHONY: compile
.PHONY: client
.PHONY: server

PROTOC_GEN_GO := $(GOPATH)/bin/protoc-gen-go

compile: protobuf client server

protobuf: internal/proto/kv.proto
	protoc --go-grpc_out=. --go_out=. internal/proto/kv.proto

bindir:
	mkdir -p ./bin

client: bindir
	go build -o bin/kvclient cmd/client/main.go

server: bindir
	go build -o bin/kvserve cmd/server/main.go

clean:
	rm -fr ./bin
