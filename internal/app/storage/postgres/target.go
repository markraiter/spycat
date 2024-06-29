package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/markraiter/spycat/internal/app/storage"
	"github.com/markraiter/spycat/internal/domain"
)

func (s *Storage) SaveTarget(ctx context.Context, tx *sql.Tx, target *domain.Target) error {
	const op = "storage.SaveTarget"
	query := "INSERT INTO targets (mission_id, name, country, notes, completed) VALUES ($1, $2, $3, $4, $5)"

	_, err := tx.ExecContext(ctx, query, target.MissionID, target.Name, target.Country, target.Notes, target.Completed)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) Targets(ctx context.Context) ([]*domain.Target, error) {
	const op = "storage.Targets"

	query, err := s.PostgresDB.Prepare("SELECT id, mission_id, name, country, notes, completed FROM targets ORDER BY created_at DESC")
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	rows, err := query.QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	targets := make([]*domain.Target, 0)
	for rows.Next() {
		t := &domain.Target{}
		err = rows.Scan(&t.ID, &t.MissionID, &t.Name, &t.Country, &t.Notes, &t.Completed)
		if err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}

		targets = append(targets, t)
	}

	return targets, nil
}

func (s *Storage) TargetCompleted(ctx context.Context, id int) error {
	const op = "storage.TargetCompleted"

	query := "UPDATE targets SET completed = true WHERE id = $1"
	result, err := s.PostgresDB.ExecContext(ctx, query, id)
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

func (s *Storage) TargetByID(ctx context.Context, id int) (*domain.Target, error) {
	const op = "storage.TargetByID"

	query := "SELECT id, mission_id, name, country, notes, completed FROM targets WHERE id = $1"
	row := s.PostgresDB.QueryRowContext(ctx, query, id)

	t := &domain.Target{}
	err := row.Scan(&t.ID, &t.MissionID, &t.Name, &t.Country, &t.Notes, &t.Completed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return t, nil
}

func (s *Storage) AddTargetToMission(ctx context.Context, missionID, targetID int) error {
	const op = "storage.AddTargetToMission"

	tx, err := s.PostgresDB.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	mission, err := s.MissionByID(ctx, missionID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			tx.Rollback()
			return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	_, err = s.TargetByID(ctx, targetID)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			tx.Rollback()
			return fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	if mission.Completed {
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, storage.ErrMissionCompleted)
	}

	query := "UPDATE targets SET mission_id = $1 WHERE id = $2"
	_, err = tx.ExecContext(ctx, query, missionID, targetID)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("%s: %w", op, err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
