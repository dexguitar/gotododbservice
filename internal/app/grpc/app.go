package grpcapp

import (
	"fmt"
	"net"

	dbgrpc "github.com/dexguitar/gotododbservice/internal/grpc"
	"google.golang.org/grpc"
)

type App struct {
	gRPCServer *grpc.Server
}

func New(dbService dbgrpc.DBService) *App {
	gRPCServer := grpc.NewServer()
	dbgrpc.Register(gRPCServer, dbService)

	return &App{
		gRPCServer: gRPCServer,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpcapp.Run"

	// log := a.log.With(
	// 	slog.String("op", op),
	// 	slog.Int("port", a.port),
	// )

	listen, err := net.Listen("tcp", ":44044")
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	// log.Info("grpc server is running", slog.String("address", listen.Addr().String()))

	if err := a.gRPCServer.Serve(listen); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() error {
	const op = "grpcapp.Stop"

	// a.log.With(slog.String("op", op)).Info("stopping grpc server", slog.Int("port", a.port))
	a.gRPCServer.GracefulStop()

	return nil
}
