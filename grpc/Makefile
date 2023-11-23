compile-proto:
	protoc --go_out=. \
    	--go-grpc_out=. \
    	*.proto

.PHONY: clean
clean:
	rm chat/*