package postgres

import (
	"context"
	"fmt"

	dbservicev1 "github.com/dexguitar/gotodoprotos/gen/go/dbservice"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Storage struct {
	db *sqlx.DB
}

func New(storagePath string) (*Storage, error) {
	const op = "storage.postgres.New"

	db, err := sqlx.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Storage{db: db}, nil
}

func (s *Storage) GetAllTodos(ctx context.Context) ([]*dbservicev1.Todo, error) {
	const op = "storage.postgres.GetAllTodos"

	todos := make([]*dbservicev1.Todo, 0)
	q := "SELECT * FROM todos"
	rows, err := s.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	for rows.Next() {
		var todo dbservicev1.Todo
		err := rows.Scan(&todo.Id, &todo.Title, &todo.Content, &todo.Done)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		todos = append(todos, &todo)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return todos, nil
}
