package todo

import (
	"context"
	"server/lib"
)

type Server struct {
	Conn lib.DatabaseConnection
}

// GetAll implements TodoServiceServer.
func (s *Server) GetAll(ctx context.Context, msg *TodoMessage) (*TodoMessageGetResponse, error) {
	todos, err := ReadAll(s.Conn, msg)
	return todos, err
}

// Create implements TodoServiceServer.
func (s *Server) Create(ctx context.Context, msg *TodoMessage) (*TodoMessage, error) {
	err := Insert(s.Conn, msg)
	return msg, err
}

// mustEmbedUnimplementedTodoServiceServer implements TodoServiceServer.
func (*Server) mustEmbedUnimplementedTodoServiceServer() {
	panic("unimplemented")
}
