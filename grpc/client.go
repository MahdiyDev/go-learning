package main

import (
	"context"
	"grpc/chat"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	var conn *grpc.ClientConn

	creds := insecure.NewCredentials()

	conn, err := grpc.Dial(":4040", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Dial error: %v", err)
		return
	}
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from client :(",
	}

	response, err := c.SayHello(context.Background(), &message)

	if err != nil {
		log.Fatalf("SayHello error: %v\n", err)
	}

	log.Printf("Response from server: %v\n", response.Body)
}
