package main

import (
	"grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":4040")
	if err != nil {
		log.Fatalf("connot listen port 4040 error: %v\n", err)
		return
	}
	log.Printf("Server started on port 4040\n")

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("grpcServer error: %v\n", err)
		return
	}

}
