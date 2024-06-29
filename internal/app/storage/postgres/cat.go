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

func (s *Storage) SaveCat(ctx context.Context, cat *domain.Cat) (int, error) {
	const op = "storage.SaveCat"

	query := "INSERT INTO cats (name, breed, years_of_experience, salary) VALUES ($1, $2, $3, $4) RETURNING id"
	err := s.PostgresDB.QueryRow(query, cat.Name, cat.Breed, cat.YearsOfExperience, cat.Salary).Scan(&cat.ID)
	if err != nil {
		var pgErr *pq.Error

		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrAlreadyExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return cat.ID, nil
}

func (s *Storage) Cat(ctx context.Context, id int) (*domain.Cat, error) {
	const op = "storage.Cat"

	query, err := s.PostgresDB.Prepare("SELECT id, name, breed, years_of_experience, salary FROM cats WHERE id = $1")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	row := query.QueryRowContext(ctx, id)

	cat := &domain.Cat{}
	err = row.Scan(&cat.ID, &cat.Name, &cat.Breed, &cat.YearsOfExperience, &cat.Salary)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return cat, nil
}

func (s *Storage) Cats(ctx context.Context) ([]*domain.Cat, error) {
	const op = "storage.Cats"

	query, err := s.PostgresDB.Prepare("SELECT id, name, breed, years_of_experience, salary FROM cats ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	cats := make([]*domain.Cat, 0)
	for rows.Next() {
		cat := &domain.Cat{}
		err = rows.Scan(&cat.ID, &cat.Name, &cat.Breed, &cat.YearsOfExperience, &cat.Salary)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		cats = append(cats, cat)
	}

	return cats, nil
}

func (s *Storage) UpdateCat(ctx context.Context, cat *domain.Cat) error {
	const op = "storage.UpdateCat"

	query := "UPDATE cats SET name = $1, breed = $2, years_of_experience = $3, salary = $4 WHERE id = $5"
	result, err := s.PostgresDB.Exec(query, cat.Name, cat.Breed, cat.YearsOfExperience, cat.Salary, cat.ID)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
	}

	return nil
}

func (s *Storage) DeleteCat(ctx context.Context, id int) error {
	const op = "storage.DeleteCat"

	query := "DELETE FROM cats WHERE id = $1"
	result, err := s.PostgresDB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
	}

	return nil
}
