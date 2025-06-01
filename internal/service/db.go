package service

import (
	"context"
	"fmt"
	"log/slog"

	dbservicev1 "github.com/dexguitar/gotodoprotos/gen/go/dbservice"
)

type DBSrv struct {
	todoProvider TodoProvider
}

type TodoProvider interface {
	GetAllTodos(ctx context.Context) ([]*dbservicev1.Todo, error)
}

// type UserProvider interface {
// 	User(ctx context.Context, email string) (models.User, error)
// 	IsAdmin(ctx context.Context, userID int64) (bool, error)
// 	SetAdmin(ctx context.Context, userID int64) (bool, error)
// }

// type AppProvider interface {
// 	App(ctx context.Context, appID int32) (models.App, error)
// }

// New creates a new DB service.
func New(
	todoProvider TodoProvider,
) *DBSrv {
	return &DBSrv{
		todoProvider: todoProvider,
	}
}

func (s *DBSrv) GetAllTodos(
	ctx context.Context,
) ([]*dbservicev1.Todo, error) {
	const op = "service.GetAllTodos"

	// log := a.log.With(
	// 	slog.String("op", op),
	// 	slog.String("email", email),
	// )

	// log.Info("attempting to login")

	todos, err := s.todoProvider.GetAllTodos(ctx)
	if err != nil {
		slog.Error(err.Error())
		// if errors.Is(err, storage.ErrUserNotFound) {
		// 	a.log.Warn("user not found", sl.Err(err))
		// 	return "", fmt.Errorf("%s: %w", op, ErrInvalidCredentials)
		// }

		// a.log.Error("failed to get user", sl.Err(err))
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}
