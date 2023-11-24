- ### Install necessary instruments and tools
    ### docker postgres image with name - postgres-container or local postgres

        docker pull postgres:14.3
        docker run -d --name postgres-container -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postres postgres
    ### docker redis image with name - redis-container or local redis
        
        docker pull redis:latest
        docker run --name redis-container -d redis
        docker run -it --network some-network --rm redis redis-cli -h redis-container
    ### install sqlc
        # MacOs
        brew install sqlc

        # Ubuntu
        sudo snap install sqlc

        # Docker
        docker pull sqlc/sqlc

        - Run sqlc using docker run:
        docker run --rm -v $(pwd):/src -w /src sqlc/sqlc generate
    ### install protobuffer
        - protoc                    v4.23.4
        - protoc-gen-go             v1.31.0
        - protoc-gen-go-grpc        v1.2.0
        - protoc-gen-openapiv2
        - protoc-gen-grpc-gateway

    ### install gitsecrets
        # MacOs
        brew install git-secret

        # Ubuntu
        sudo apt-get update && sudo apt-get install git-secret
    ### install goose migration tool
        https://github.com/pressly/goose#install

        # MacOs
        brew install goose

        # Go
        go install github.com/pressly/goose/v3/cmd/goose@latest

        # Linux
        https://pressly.github.io/goose/installation/
- Read through documentation of used tools
    - Read through 
        - Docs: https://grpc-ecosystem.github.io/grpc-gateway/
        - Example: https://github.com/grpc-ecosystem/grpc-gateway/tree/main/examples/internal
    - Read about integration testing:
        - Docs: https://dev.to/kliukovkin/integration-tests-with-go-and-testcontainers-6o5
        - Example: https://github.com/kliukovkin/integration-test-go/tree/main
    - Read about mocking the interface with testify/mock
        - Docs: https://github.com/stretchr/testify#mock-package
        - Example: https://go.dev/play/p/IrpjrkJ1WG1
    - Read through the doc for sqlc - https://sqlc.dev/
    - Read about - https://sobolevn.me/git-secret/
    - Learn the queue package that works with redis - https://github.com/hibiken/asynq
    - Read about goose migrate tool - https://github.com/pressly/goose
    - Read the docs of dbdiagram.io - https://dbdiagram.io/docs/
    - Read through available commands in Makefile
- Add your first contribution
    - create one rpc service named TestService
    - Add test-endpoint and generate api gateway
    - Implement a service
    - add a log when request is fired from swagger
    - write a dummy query for integration dummy table and generate with sqlc
    - write a test in service layer that mocks the single method of db Query interface
