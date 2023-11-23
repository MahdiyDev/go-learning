include .env

GC=go
ENTRY_FILE=main.go

.PHONY: server
server:
	$(GC) run server/$(ENTRY_FILE)

.PHONY: client
client:
	$(GC) run server/$(ENTRY_FILE)

compile:
	$(GC) build -o server.out server/$(ENTRY_FILE)
	$(GC) build -o client.out client/$(ENTRY_FILE)

compile-proto:
	protoc --go_out=./server \
    	--go-grpc_out=./server \
    	server/proto/*.proto
	protoc --go_out=./client \
    	--go-grpc_out=./client \
	client/proto/*.proto

migration-create:
	goose create todo sql

.PHONY: clean
clean:
	rm *.out