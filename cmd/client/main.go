package main

import (
	"context"
	"log"
	"time"

	"github.com/yami-beta/go-grpc-todo-sample/todo"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	c := todo.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.ListTodo(ctx, &todo.ListTodoRequest{})
	if err != nil {
		log.Fatalf("Could not get todos: %v", err)
	}
	log.Printf("Todos: %s", r.Todos)
}
