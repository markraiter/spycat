package postgres

import (
	"context"
	"database/sql"
	"fmt"

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
