package server

import (
	"context"

	"github.com/yami-beta/go-grpc-todo-sample/todo"
)

// Server gRPC に渡すサーバの構造体
type Server struct{}

// ListTodo todo.proto の ListTodo ハンドラ
func (s *Server) ListTodo(ctx context.Context, in *todo.ListTodoRequest) (*todo.ListTodoResponse, error) {
	return &todo.ListTodoResponse{
		Todos: []*todo.Todo{},
	}, nil
}
