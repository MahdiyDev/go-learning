package main

import (
	"log"
	"net"

	"server/config"
	"server/lib"
	"server/todo"

	"google.golang.org/grpc"
)

func main() {
	env := config.Env{}
	err := env.Load()
	if err != nil {
		panic(err)
	}

	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("connot listen port 4040 error: %v\n", err)
		return
	}
	log.Printf("Server started on port 4040\n")

	conn := lib.DatabaseConnection{DATABASE_URL: env.DATABASE_URL}
	err = conn.Setup()
	if err != nil {
		panic(err)
	}

	t := todo.Server{Conn: conn}

	grpcServer := grpc.NewServer()

	todo.RegisterTodoServiceServer(grpcServer, &t)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer error: %v\n", err)
		return
	}
}
