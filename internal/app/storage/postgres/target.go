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
