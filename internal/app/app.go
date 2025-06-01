package app

import (
	grpcapp "github.com/dexguitar/gotododbservice/internal/app/grpc"
	"github.com/dexguitar/gotododbservice/internal/service"
	"github.com/dexguitar/gotododbservice/internal/storage/postgres"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(storagePath string) *App {
	storage, err := postgres.New(storagePath)
	if err != nil {
		panic(err)
	}

	dbService := service.New(storage)

	grpcApp := grpcapp.New(dbService)

	return &App{
		GRPCSrv: grpcApp,
	}
}
