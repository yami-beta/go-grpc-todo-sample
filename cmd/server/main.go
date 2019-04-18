package main

import (
	"log"
	"net"

	"github.com/yami-beta/go-grpc-todo-sample/server"
	"github.com/yami-beta/go-grpc-todo-sample/todo"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	todo.RegisterTodoServiceServer(s, &server.Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
