package main

import (
	"client/todo"
	"context"
	"fmt"
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

	c := todo.NewTodoServiceClient(conn)

	var input string
	for input != "exit" {
		fmt.Print("To create new todo type: add\nTo get all todos type: get\nTo exit program type exit\n")
		fmt.Print("> ")
		fmt.Scanf("%s", &input)

		switch input {
		case "add":
			message := todo.CreateTodo()
			response, err := c.Create(context.Background(), message)
			if err != nil {
				log.Fatalf("Create error: %v\n", err)
			}
			log.Printf("Response from server: %v\n", response)

		case "get":
			getResponse, err := c.GetAll(context.Background(), &todo.TodoMessage{})

			if err != nil {
				log.Fatalf("Create error: %v\n", err)
			}

			for i := 0; i < len(getResponse.Todos); i++ {
				fmt.Println(getResponse.Todos[i])
			}
		}
	}
}
