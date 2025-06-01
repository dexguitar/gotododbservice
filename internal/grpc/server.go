package grpc

import (
	"context"

	dbservicev1 "github.com/dexguitar/gotododbservice/internal/grpc/gen/go/dbservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DBService interface {
	GetAllTodos(ctx context.Context) ([]*dbservicev1.Todo, error)
}

type serverApi struct {
	dbservicev1.UnimplementedDBServiceServer
	dbService DBService
}

func Register(gRPC *grpc.Server, dbService DBService) {
	dbservicev1.RegisterDBServiceServer(gRPC, &serverApi{dbService: dbService})
}

func (s *serverApi) GetAllTodos(ctx context.Context, req *dbservicev1.GetAllTodosRequest) (*dbservicev1.GetAllTodosResponse, error) {
	todos, err := s.dbService.GetAllTodos(ctx)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal error")
	}

	return &dbservicev1.GetAllTodosResponse{
		Todos: todos,
	}, nil
}
