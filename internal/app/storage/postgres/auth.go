package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/domain"
)

func (s *Storage) SaveUser(ctx context.Context, user *domain.User) (int, error) {
	const op = "storage.SaveUser"

	query := "INSERT INTO users (username, password, email) VALUES ($1, $2, $3) RETURNING id"
	err := s.PostgresDB.QueryRow(query, user.Username, user.Password, user.Email).Scan(&user.ID)
	if err != nil {
		var pgErr *pq.Error

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return user.ID, nil
}

func (s *Storage) User(ctx context.Context, email string) (*domain.User, error) {
	const op = "storage.UserByEmail"

	query, err := s.PostgresDB.Prepare("SELECT id, username, password, email FROM users WHERE email = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := query.QueryRowContext(ctx, email)

	user := &domain.User{}
	err = row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}

		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return user, nil
}
